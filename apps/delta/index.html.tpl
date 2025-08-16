<html>
    <head>
        <meta charset="utf-8"/>
        <script>
            var firebaseConfig = {{.FirebaseConfig}};
        </script>
        <script src="/static/wasm_exec.js"></script>
        <script src="../../firebase/auth/bundle.js"></script>
        <script>
            const go = new Go();
            WebAssembly.instantiateStreaming(fetch("delta.wasm"), go.importObject).then((result) => {
                go.run(result.instance);
            });
        </script>
        <title>Delta</title>
        <link rel="stylesheet" href="ui.css">
        <link rel="stylesheet" href="delta.css">
</head>
    <body>
        <div id="game-board">
        </div>
        <div id="auth-container">
            <div id="logged-out-view">
                <button id="login-button">Login</button>
            </div>
            <div id="logged-in-view" style="display: none;">
                <span id="user-name"></span>
                <button id="logout-button">Logout</button>
            </div>
        </div>
    </body>
</html>