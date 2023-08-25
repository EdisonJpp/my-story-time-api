package email_sender

type EmailSenderRepository interface {
	Send(from string, to []string, subject string, body string) error
}
