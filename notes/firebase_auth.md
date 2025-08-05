# Firebase Auth Design Notes

## Core Concepts

- **Client-Side Library:** The Firebase Auth client SDK is a JavaScript library. It runs in the user's browser to manage authentication state, handle sign-in popups, and refresh tokens. The Go *Admin* SDK is for backend servers, not for client-side login.

- **The Bridge:** Go code compiled to Wasm runs in a sandbox and communicates with the JavaScript environment via the `syscall/js` package.

- **Making Symbols Available:** We can't import the JavaScript library directly into Go. Instead, we load the Firebase JS SDK in `index.html` and create simple JavaScript wrapper functions (e.g., `signIn()`, `signOut()`, `onAuthStateChanged(callback)`). These wrappers are attached to the global `window` object, making them callable from Go.

- **Calling from Go:** From Go, we use the `js` library to get a reference to the global `window` object and call the JavaScript wrapper functions. For asynchronous operations like auth state changes, we pass a Go function (wrapped with `js.FuncOf`) to the JavaScript `onAuthStateChanged` wrapper. The JavaScript code then invokes the Go callback with the user data.

## Secure Configuration Management

- **Client-Side Config is Public:** The Firebase config object (`apiKey`, `authDomain`, etc.) is not a secret. Security comes from properly configured Firebase Security Rules and Firebase App Check, not from hiding the config.

- **Untracked Config File:** To manage the configuration securely and avoid committing it to version control, we create an untracked `config.go` file. This file is added to `.gitignore`. The configuration is loaded from this file at build time and passed to the JavaScript side.
