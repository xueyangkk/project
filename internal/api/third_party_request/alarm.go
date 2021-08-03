package third_party_request

import (
	"exams-api/configs"
	"exams-api/pkg/httpclient"
	"exams-api/pkg/mail"

	"github.com/pkg/errors"
)

// 实现 AlarmObject 告警
var _ httpclient.AlarmObject = (*AlarmEmail)(nil)

type AlarmEmail struct{}

// 邮件告警方式
func (a *AlarmEmail) Send(subject, body string) error {
	cfg := configs.Get().Mail
	if cfg.Host == "" || cfg.Port == 0 || cfg.User == "" || cfg.Pass == "" || cfg.To == "" {
		return errors.New("mail config error")
	}

	options := &mail.Options{
		MailHost: cfg.Host,
		MailPort: cfg.Port,
		MailUser: cfg.User,
		MailPass: cfg.Pass,
		MailTo:   cfg.To,
		Subject:  subject,
		Body:     body,
	}

	return mail.Send(options)
}
