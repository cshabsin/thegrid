{{define "auth_ui"}}
<div id="auth-container">
    <div id="logged-out-view">
        <button id="login-button">Login</button>
    </div>
    <div id="logged-in-view" style="display: none;">
        <span id="user-name"></span>
        <button id="logout-button">Logout</button>
    </div>
</div>
{{end}}
