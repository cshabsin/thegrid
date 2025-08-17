<html>
<head>
    <title>{{.Title}}</title>
    <link rel="stylesheet" href="/firebase/authui/auth.css">
    <script>
        var firebaseConfig = {{.FirebaseConfig}};
    </script>
</head>
<body>
    {{template "auth_ui" .}}
    <div id="content">
        {{template "body" .}}
    </div>
</body>
</html>
