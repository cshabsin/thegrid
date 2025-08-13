package auth

import (
	syscalljs "syscall/js"

	"github.com/cshabsin/thegrid/js"
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

func OnAuthStateChanged(callback func(user js.User)) {
	syscalljs.Global().Get("firebaseAuth").Call("onAuthStateChanged", syscalljs.FuncOf(func(this syscalljs.Value, args []syscalljs.Value) any {
		callback(js.User{Value: args[0]})
		return nil
	}))
}
