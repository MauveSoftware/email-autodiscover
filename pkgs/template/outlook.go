package template

import (
	"bytes"
	"text/template"
)

const outlookTemplate = `<?xml version="1.0" encoding="utf-8" ?>
<Autodiscover xmlns="http://schemas.microsoft.com/exchange/autodiscover/responseschema/2006">
	<Response xmlns="https://schemas.microsoft.com/exchange/autodiscover/outlook/responseschema/2006a">
		<User>
			<DisplayName>{{.Email}}</DisplayName>
		</User>

		<Account>
			<AccountType>email</AccountType>
			<Action>settings</Action>

			<Protocol>
				<Type>IMAP</Type>
				<Server>{{.ImapHost}}</Server>
				<Port>{{.ImapPort}}</Port>
				<DomainRequired>off</DomainRequired>
				<LoginName>{{.Email}}</LoginName>
				<SPA>off</SPA>
				<SSL>on</SSL>
				<AuthRequired>on</AuthRequired>
			</Protocol>
		</Account>

		<Account>
			<AccountType>email</AccountType>
			<Action>settings</Action>
			<Protocol>
				<Type>SMTP</Type>
				<Server>{{.SmtpHost}}</Server>
				<Port>{{.SmtpPort}}</Port>
				<DomainRequired>off</DomainRequired>
				<LoginName>{{.Email}}</LoginName>
				<SPA>off</SPA>
				<SSL>on</SSL>
				<AuthRequired>on</AuthRequired>
			</Protocol>
		</Account>
	</Response>
</Autodiscover>
`

func OutlookMail(args *Args) (string, error) {
	reportTmpl, err := template.New("report").Parse(outlookTemplate)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	err = reportTmpl.Execute(&b, args)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
