package email_sender

import "go.uber.org/fx"

var DIContainer = fx.Options(
	fx.Provide(EmailSenderClient, NewEmailSenderRepository),
)
