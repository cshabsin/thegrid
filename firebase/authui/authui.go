package authui

import (
	"fmt"

	"github.com/cshabsin/thegrid/firebase"
	"github.com/cshabsin/thegrid/firebase/auth"
	"github.com/cshabsin/thegrid/js"
)

func Initialize(config *firebase.Config) {
	auth.InitializeApp(config)

	doc := js.Document()

	loginButton := doc.GetElementByID("login-button")
	loginButton.AddEventListener("click", func(_ js.DOMElement, _ js.DOMEvent) {
		auth.SignIn()
	})

	logoutButton := doc.GetElementByID("logout-button")
	logoutButton.AddEventListener("click", func(_ js.DOMElement, _ js.DOMEvent) {
		auth.SignOut()
	})

	loggedOutView := doc.GetElementByID("logged-out-view")
	loggedInView := doc.GetElementByID("logged-in-view")
	userName := doc.GetElementByID("user-name")

	auth.OnAuthStateChanged(func(user auth.User) {
		if !user.IsNull() {
			fmt.Println("User is signed in:", user.DisplayName())
			userName.SetInnerHTML(user.DisplayName())
			loggedOutView.Style().SetProperty("display", "none")
			loggedInView.Style().SetProperty("display", "block")
		} else {
			fmt.Println("User is signed out")
			loggedOutView.Style().SetProperty("display", "block")
			loggedInView.Style().SetProperty("display", "none")
		}
	})
}
