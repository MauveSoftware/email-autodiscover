package template

type Args struct {
	Domain            string // This one is populated from the request.
	Email             string // This one is populated from the request.
	EmailProvider     string `yaml:"email-provider"`
	ImapHost          string `yaml:"imap-host"`
	ImapPort          int    `yaml:"imap-port"`
	SmtpHost          string `yaml:"smtp-host"`
	SmtpPort          int    `yaml:"smtp-port"`
	PayloadIdentifier string `yaml:"payload-identifier"`
}
