package auth

import (
	"syscall/js"
)

type Config struct {
	APIKey            string `json:"apiKey"`
	AuthDomain        string `json:"authDomain"`
	ProjectID         string `json:"projectId"`
	StorageBucket     string `json:"storageBucket"`
	MessagingSenderID string `json:"messagingSenderId"`
	AppID             string `json:"appId"`
}

func (c *Config) toJS() js.Value {
	return js.ValueOf(map[string]any{
		"apiKey":            c.APIKey,
		"authDomain":        c.AuthDomain,
		"projectId":         c.ProjectID,
		"storageBucket":     c.StorageBucket,
		"messagingSenderId": c.MessagingSenderID,
		"appId":             c.AppID,
	})
}

func InitializeApp(config *Config) {
	js.Global().Get("firebaseAuth").Call("initializeApp", config.toJS())
}

func SignIn() {
	js.Global().Get("firebaseAuth").Call("signIn")
}

func SignOut() {
	js.Global().Get("firebaseAuth").Call("signOut")
}

func OnAuthStateChanged(callback func(user js.Value)) {
	js.Global().Get("firebaseAuth").Call("onAuthStateChanged", js.FuncOf(func(this js.Value, args []js.Value) any {
		callback(args[0])
		return nil
	}))
}
