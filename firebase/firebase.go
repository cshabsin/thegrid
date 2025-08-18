package firebase

import "syscall/js"

type Config struct {
	APIKey            string `json:"apiKey"`
	AuthDomain        string `json:"authDomain"`
	ProjectID         string `json:"projectId"`
	StorageBucket     string `json:"storageBucket"`
	MessagingSenderID string `json:"messagingSenderId"`
	AppID             string `json:"appId"`
}

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
