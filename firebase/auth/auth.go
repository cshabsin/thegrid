package auth

import (
	syscalljs "syscall/js"
)

type Config struct {
	APIKey            string `json:"apiKey"`
	AuthDomain        string `json:"authDomain"`
	ProjectID         string `json:"projectId"`
	StorageBucket     string `json:"storageBucket"`
	MessagingSenderID string `json:"messagingSenderId"`
	AppID             string `json:"appId"`
}

func (c *Config) toJS() syscalljs.Value {
	return syscalljs.ValueOf(map[string]any{
		"apiKey":            c.APIKey,
		"authDomain":        c.AuthDomain,
		"projectId":         c.ProjectID,
		"storageBucket":     c.StorageBucket,
		"messagingSenderId": c.MessagingSenderID,
		"appId":             c.AppID,
	})
}

func InitializeApp(config *Config) {
	syscalljs.Global().Get("firebaseAuth").Call("initializeApp", config.toJS())
}

func SignIn() {
	syscalljs.Global().Get("firebaseAuth").Call("signIn")
}

func SignOut() {
	syscalljs.Global().Get("firebaseAuth").Call("signOut")
}

func OnAuthStateChanged(callback func(user User)) {
	syscalljs.Global().Get("firebaseAuth").Call("onAuthStateChanged", syscalljs.FuncOf(func(this syscalljs.Value, args []syscalljs.Value) any {
		callback(User{Value: args[0]})
		return nil
	}))
}

// User wraps the JavaScript user object from Firebase auth.
type User struct {
	syscalljs.Value
}

// IsNull returns true if the user is null.
func (u User) IsNull() bool {
	return u.Value.IsNull()
}

// DisplayName returns the user's display name.
func (u User) DisplayName() string {
	return u.Get("displayName").String()
}