package models

type MailForm struct {
	To      string `json:"to"`
	ToName  string `json:"to_name"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
