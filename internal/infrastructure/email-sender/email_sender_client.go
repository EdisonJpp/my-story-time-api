package email_sender

import (
	"github.com/aws/aws-sdk-go/aws"
	credential "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gofiber/fiber/v2/log"
	configDomain "my-story-time-api/internal/domain/config"
)

func EmailSenderClient(config *configDomain.Config) *ses.SES {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Aws.Region),
		Credentials: credential.NewStaticCredentials(config.Aws.AccessKey, config.Aws.SecretAccessKey, ""),
	})

	if err != nil {
		log.Fatal("Error creating AWS session:", err)
	}

	sesClient := ses.New(sess)

	return sesClient
}
