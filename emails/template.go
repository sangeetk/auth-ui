package emails

import (
	"fmt"
	"io/ioutil"
	"log"
	"text/template"
)

const (
	// Activation mail
	Activation = "activation"
	// Welcome mail
	Welcome = "welcome"
	// PasswordReset mail
	PasswordReset = "password-reset"
	// ResetConfirmation mail
	ResetConfirmation = "reset-confirmation"
)

var funcMap = template.FuncMap{}

// Emails templates
var Emails map[string]*template.Template
var base = template.New("Email").Funcs(funcMap)

// MustFindTemplate loads the template file into memory
func MustFindTemplate(dir, file string) string {
	filepath := fmt.Sprintf("%s/%s", dir, file)
	content, _ := ioutil.ReadFile(filepath)
	return fmt.Sprintf("%s", content)
}

func init() {
	var err error
	Emails = make(map[string]*template.Template)

	activationMail := MustFindTemplate("emails", "account-activation.tpl")
	if Emails[Activation], err = template.Must(base.Clone()).Parse(activationMail); err != nil {
		log.Fatal("Activation Mail:", err)
	}

	resetConfirmation := MustFindTemplate("emails", "reset-confirmation.tpl")
	if Emails[ResetConfirmation], err = template.Must(base.Clone()).Parse(resetConfirmation); err != nil {
		log.Fatal("ResetConfirmation Mail:", err)
	}

	passwordReset := MustFindTemplate("emails", "password-reset.tpl")
	if Emails[PasswordReset], err = template.Must(base.Clone()).Parse(passwordReset); err != nil {
		log.Fatal("PasswordReset Mail:", err)
	}

	welcomeMail := MustFindTemplate("emails", "welcome.tpl")
	if Emails[Welcome], err = template.Must(base.Clone()).Parse(welcomeMail); err != nil {
		log.Fatal("Verification Mail:", err)
	}
}
