package tools

import "gopkg.in/gomail.v2"

//发送邮件
func SendEmail(from string, to []string, title string, content string, host string, port int, password string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	//m.SetAddressHeader("Cc", "904801074@qq.com", "Tools")
	m.SetHeader("Subject", title)
	m.SetBody("text/html", content)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(host, port, from, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		//panic(err)

		return err
	}

	return nil
}
