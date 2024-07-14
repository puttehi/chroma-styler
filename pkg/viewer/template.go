package viewer

const Template = `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>Chroma style viewer</title>
        <!--User code styles-->
        <link href="syntax.css" rel="stylesheet">
    </head>
    <body style="background-color: #1D1D3D">
        {{- range .CodeBlocks -}}
        <div>
            {{- . -}}
        </div>
        {{- end -}}
    </body>
</html>
`
