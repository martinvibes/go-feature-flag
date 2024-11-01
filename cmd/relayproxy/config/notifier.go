package config

import "fmt"

type NotifierConf struct {
	Kind            NotifierKind        `mapstructure:"kind" koanf:"kind"`
	SlackWebhookURL string              `mapstructure:"slackWebhookUrl" koanf:"slackWebhookUrl"`
	MicrosoftTeamsWebhookURL string              `mapstructure:"microsoftteamsWebhookUrl" koanf:"microsoftteamsWebhookUrl"`
	EndpointURL     string              `mapstructure:"endpointUrl" koanf:"endpointUrl"`
	Secret          string              `mapstructure:"secret" koanf:"secret"`
	Meta            map[string]string   `mapstructure:"meta" koanf:"meta"`
	Headers         map[string][]string `mapstructure:"headers" koanf:"headers"`
}

func (c *NotifierConf) IsValid() error {
	if err := c.Kind.IsValid(); err != nil {
		return err
	}
	if c.Kind == SlackNotifier && c.SlackWebhookURL == "" {
		return fmt.Errorf("invalid notifier: no \"slackWebhookUrl\" property found for kind \"%s\"", c.Kind)
	}
	if c.Kind == MicrosoftTeamsNotifier && c.MicrosoftTeamsWebhookURL == "" {
		return fmt.Errorf("invalid notifier: no \"microsoftteamsWebhookUrl\" property found for kind \"%s\"", c.Kind)
	}
	if c.Kind == WebhookNotifier && c.EndpointURL == "" {
		return fmt.Errorf("invalid notifier: no \"endpointUrl\" property found for kind \"%s\"", c.Kind)
	}
	return nil
}

type NotifierKind string

const (
	SlackNotifier   NotifierKind = "slack"
	MicrosoftTeamsNotifier   NotifierKind = "microsoftteams"
	WebhookNotifier NotifierKind = "webhook"
)

// IsValid is checking if the value is part of the enum
func (r NotifierKind) IsValid() error {
	switch r {
	case SlackNotifier, MicrosoftTeamsNotifier, WebhookNotifier:
		return nil
	}
	return fmt.Errorf("invalid notifier: kind \"%s\" is not supported", r)
}
