package auth

import (
	"syscall/js"

	"github.com/cshabsin/thegrid/firebase"
)

func InitializeApp(config *firebase.Config) {
	js.Global().Get("firebaseAuth").Call("initializeApp", config.ToJS())
}

func SignIn() {
	js.Global().Get("firebaseAuth").Call("signIn")
}

func SignOut() {
	js.Global().Get("firebaseAuth").Call("signOut")
}

func OnAuthStateChanged(callback func(user User)) {
	js.Global().Get("firebaseAuth").Call("onAuthStateChanged", js.FuncOf(func(this js.Value, args []js.Value) any {
		callback(User{Value: args[0]})
		return nil
	}))
}

// User wraps the JavaScript user object from Firebase auth.
type User struct {
	js.Value
}

// IsNull returns true if the user is null.
func (u User) IsNull() bool {
	return u.Value.IsNull()
}

// DisplayName returns the user's display name.
func (u User) DisplayName() string {
	return u.Get("displayName").String()
}
