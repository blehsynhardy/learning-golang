package main

import (
	"fmt"
	"html/template"
	"os"
)

type EmailData struct {
	RecipientName string
	SenderName    string
	Subject       string
	Body          string
	Items         []string
	UnreadCount   int
}

func main() {

	fmt.Println("===========TEMPLATE START HERE ===========")

	emailTemplate := `Dear {{.RecipientName}},
	{{.SenderName}} has sent you an email with the subject "{{.Subject}}".

	{{.Body}}

{{ if .Items}}}

	related items:
	{{ range .Items}}
	 	- {{.}}
	{{ end }}
{{ end }}
 
{{if gt .UnreadCount 0}} {
	You have {{.UnreadCount}} unread emails.
}
{{ else }}

	You have no unread emails.
{{ end }}

Thanks and regards,
{{.SenderName}}
`

	tmpl, err := template.New("email").Parse(emailTemplate)

	if err != nil {
		fmt.Print(err.Error())

		os.Exit(1)
	}

	sampleData := EmailData{
		RecipientName: "John Doe",
		SenderName:    "Jane Smith",
		Subject:       "Meeting Reminder",
		Body:          "Just a reminder about our meeting scheduled for tomorrow at 10 AM.",
		Items:         []string{"Agenda", "Project Update", "Budget Review"},
		UnreadCount:   0,
	}

	err = tmpl.Execute(os.Stdout, sampleData)
	if err != nil {
		fmt.Print(err.Error())

		os.Exit(1)
	}

	fmt.Println("===========TEMPLATE END HERE ===========")

}
