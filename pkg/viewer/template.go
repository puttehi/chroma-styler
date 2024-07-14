package viewer

const Template = `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>Chroma style viewer</title>
        <!--Page styles-->
        <link href="css/style.css" rel="stylesheet">
        <!--User code styles-->
        <link href="css/syntax.css" rel="stylesheet">
    </head>
    <body>
        {{- range .CodeBlocks -}}
        <div>
            {{- . -}}
        </div>
        {{- end -}}
    </body>
</html>
`
