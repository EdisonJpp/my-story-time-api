package email_sender

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"log"
	emailSenderDomain "my-story-time-api/internal/domain/email-sender"
)

type emailSenderRepository struct {
	sesClient *ses.SES
}

func NewEmailSenderRepository(sesClient *ses.SES) emailSenderDomain.EmailSenderRepository {
	return &emailSenderRepository{sesClient}
}

func (e *emailSenderRepository) Send(from string, to []string, subject string, body string) error {
	var destinations []*string

	for _, des := range to {
		destinations = append(destinations, aws.String(des))
	}

	emailInput := &ses.SendEmailInput{
		Source: aws.String(from),
		Destination: &ses.Destination{
			ToAddresses: destinations,
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(body),
				},
			},
		},
	}

	_, err := e.sesClient.SendEmail(emailInput)
	if err != nil {
		log.Fatal("Error sending email:", err)
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}
