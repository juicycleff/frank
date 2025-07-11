package sms

import (
	"strings"

	"github.com/xraph/frank/config"
	"github.com/xraph/frank/pkg/logging"
)

// SenderFactory creates an SMS sender based on configuration
func SenderFactory(cfg *config.SMSConfig, logger logging.Logger) Provider {
	provider := strings.ToLower(cfg.Provider)

	switch provider {
	case "twilio":
		return NewTwilioProvider(cfg, logger)
	case "sendgrid":
		return NewSendgridProvider(cfg, logger)
	case "aws", "sns":
		return NewAWSProvider(cfg, logger)
	case "messagebird":
		return NewMessageBirdProvider(cfg, logger)
	case "vonage", "nexmo":
		return NewVonageProvider(cfg, logger)
	case "plivo":
		return NewPlivoProvider(cfg, logger)
	case "clicksend":
		return NewClickSendProvider(cfg, logger)
	case "sinch":
		return NewSinchProvider(cfg, logger)
	case "infobip":
		return NewInfobipProvider(cfg, logger)
	case "telnyx":
		return NewTelnyxProvider(cfg, logger)
	case "test", "mock":
		return NewMockProvider(cfg, logger)
	default:
		logger.Warn("Unknown SMS provider, defaulting to no-op sender", logging.String("provider", provider))
		return NewNoOpSender(logger)
	}
}
