package mail

import (
	"backend-code/config"
	"log"
	"net/smtp"
)

func SendEmailOk(email string, password string) {
	send(""+
		"Hi "+email+",\n\n"+
		"Your email succesfully registered with Password : "+password+". Please login and change your password for security Reason"+
		"\n\nThanks"+"\nFflix Admin", email)
}

func SendEmailError(email string) {
	send(""+
		"Hi "+email+",\n\n"+
		"Your email has been registered. Thank You."+

		"", email)
}

func send(body string, reciever string) {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	from := baseConfig.Mail.From
	pass := baseConfig.Mail.Password
	to := reciever

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Registered Status\n\n" +
		body

	err := smtp.SendMail(baseConfig.Mail.Smtp+":"+baseConfig.Mail.Port,
		smtp.PlainAuth("", from, pass, baseConfig.Mail.Smtp),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

}
