package mail

import (
	"gopkg.in/gomail.v2"
     "strconv"
)

func SendEmail(mailTo []string,subject string,body string) error {

	mailConn := map[string]string{
		"user": "fengjunhua@yeah.net",
		"pass":"120225fjhFJH!",
		"host":"",
		"port":"465",
	}
	port,_ := strconv.Atoi(mailConn["port"])

	m := gomail.NewMessage()
	m.SetHeader("To",mailTo...)
	m.SetBody("Subject",subject)
	m.SetBody("text/html",body)
	d := gomail.NewDialer(mailConn["host"],port,mailConn["user"],mailConn["pass"])

	err := d.DialAndSend(m)

	return err


}

func Send()  {
	mailTo := []string{
		"2424503808@qq.com",
	}
	subject := "hello"
	body := "hello 你好!"
	SendEmail(mailTo,subject,body)
}