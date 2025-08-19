package firebase

import (
	"syscall/js"
)

// Config holds the Firebase config.
type Config struct {
	APIKey            string `json:"apiKey"`
	AuthDomain        string `json:"authDomain"`
	ProjectID         string `json:"projectId"`
	StorageBucket     string `json:"storageBucket"`
	MessagingSenderID string `json:"messagingSenderId"`
	AppID             string `json:"appId"`
}

// ToJS converts the config to a js.Value.
func (c *Config) ToJS() js.Value {
	return js.ValueOf(map[string]any{
		"apiKey":            c.APIKey,
		"authDomain":        c.AuthDomain,
		"projectId":         c.ProjectID,
		"storageBucket":     c.StorageBucket,
		"messagingSenderId": c.MessagingSenderID,
		"appId":             c.AppID,
	})
}

// InitializeApp initializes the Firebase app and returns the app object.
func InitializeApp(config *Config) js.Value {
	return js.Global().Get("firebase").Call("initializeApp", config.ToJS())
}
