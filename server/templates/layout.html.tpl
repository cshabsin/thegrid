<html>
<head>
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/firebase/authui/auth.css">
    <script>
        var firebaseConfig = {{.FirebaseConfig}};
    </script>
    {{template "head" .}}
    <script src="/static/wasm_exec.js"></script>
    <script src="/firebase/auth/bundle.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("{{.Title}}.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        });
    </script>
</head>
<body>
    {{template "auth_ui" .}}
    <div id="content">
        {{template "body" .}}
    </div>
</body>
</html>