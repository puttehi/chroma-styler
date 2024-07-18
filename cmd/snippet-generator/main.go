package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/alecthomas/chroma/v2/lexers"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func dumpGenerateHelper() {
	fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("  cwd = %s\n", cwd)
	fmt.Printf("  os.Args = %#v\n", os.Args)

	for _, ev := range []string{"GOARCH", "GOOS", "GOFILE", "GOLINE", "GOPACKAGE", "DOLLAR"} {
		fmt.Println("  ", ev, "=", os.Getenv(ev))
	}
}

type Folder struct {
	Extension string
}

type YAMLTestInfo struct {
	Folder
}

func readTestInfo(ymlPath string) YAMLTestInfo {
	content, err := os.ReadFile(ymlPath)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(content))

	var testInfo YAMLTestInfo
	err = yaml.Unmarshal(content, &testInfo)
	if err != nil {
		panic(err)
	}

	return testInfo
}

type LanguageData struct {
	testInfo YAMLTestInfo
	files    []string // code files
	language string   // "adb"
}

func (ti YAMLTestInfo) compileLanguageData(langDir string) LanguageData {
	var sampleFiles []string
	filepath.WalkDir(langDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}

		if !strings.HasSuffix(path, ti.Extension) {
			return nil // Next
		}

		sampleFiles = append(sampleFiles, path)

		return nil
	})

	return LanguageData{
		testInfo: ti,
		files:    sampleFiles,
		language: filepath.Base(langDir),
	}
}

type TemplateSample struct {
	FunctionName string
	Filename     string
	LexerName    string // Chroma lexer name (or file extension)
	Code         string
}

func (ld LanguageData) toSample() TemplateSample {
	content, err := os.ReadFile(ld.files[0])
	if err != nil {
		panic(err)
	}

	titlecased := cases.Title(language.Und, cases.NoLower).String(ld.language)
	functionName := strings.ReplaceAll(titlecased, "-", "_")

	return TemplateSample{
		FunctionName: functionName,
		Filename:     fmt.Sprintf("%s.go", ld.language),
		LexerName:    ld.testInfo.Extension,
		Code:         string(content),
	}
}

func unsupported(sample TemplateSample) bool {
	lexer := lexers.Get(sample.LexerName)
	if lexer == nil {
		// Unsupported, no lexer found
		return true
	}

	if strings.Contains(sample.Code, "`") {
		// Unsupported, templating goes crazy
		return true
	}

	var unsupportedFiles = []string{
		"piet.go",
		"agda.go",
	}

	// Unsupported, Chroma goes crazy
	return slices.Contains(unsupportedFiles, sample.Filename)
}

func generateSample(relPath string, testInfoPath string) *TemplateSample {
	// Read language config file
	testInfo := readTestInfo(testInfoPath)

	// Find language files and name
	languageData := testInfo.compileLanguageData(relPath)

	// Read (first found) to a sample
	sample := languageData.toSample()

	if unsupported(sample) {
		fmt.Printf("Skipping %s as it's not supported by the templating implementation or is in binary.\n", sample.Filename)
		return nil
	}

	return &sample
}

func generateSamples(relPath string) []TemplateSample {
	abs, err := filepath.Abs(relPath)
	if err != nil {
		panic(err)
	}

	var samples []TemplateSample

	// Let's go with one for now

	filepath.WalkDir(abs, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}

		if !d.IsDir() {
			return nil // Traverse until getting to a directory
		}

		testInfoPath := filepath.Join(path, "testinfo.yml")
		if _, err := os.Stat(testInfoPath); errors.Is(err, os.ErrNotExist) {
			return nil // Skip until we traverse deep enough to find a path that contains code samples
		} else {
			sample := generateSample(path, testInfoPath)
			if sample == nil {
				return nil
			} else {
				samples = append(samples, *sample)
				return fs.SkipDir // Found, continue
			}
		}
	})

	return samples
}

func writeGoFiles(samples []TemplateSample) {
	funcs := make(map[string]any)

	funcs["Backtick"] = func() string {
		return "`"
	}

	tpl, err := template.New("code_sample").Funcs(funcs).Parse(`package code_samples

import "github.com/alecthomas/chroma/v2/lexers"

func {{.FunctionName}}() Sample {
	return Sample{
		Lexer: lexers.Get("{{.LexerName}}"),
		Code: {{Backtick}}{{.Code}}{{Backtick}},
	}
}`)

	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for _, sample := range samples {
		path := filepath.Join(cwd, sample.Filename)
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		err = tpl.Execute(file, sample)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Wrote %s\n", path)
	}
}

func writeCollectionFile(samples []TemplateSample) {

	tpl, err := template.New("code_sample_collection").Parse(`package code_samples

    var CodeSamples = []Sample{
    {{- range . -}}
    {{.FunctionName}}(),
    {{ end }}
}`)

	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(cwd, "all_code_samples.go")
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = tpl.Execute(file, samples)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %s\n", path)
}

func main() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "help":
			dumpGenerateHelper()
			break
		default:
			samples := generateSamples(os.Args[1])
			writeGoFiles(samples)
			writeCollectionFile(samples)
			break
		}
	}

	return
}
