package REmail

import (
	"github.com/go-gomail/gomail"
	"log"
)

const (
	LengthOfStr = 6
)

type EmailInfo struct {
	ServerHost string // ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerPort int    // ServerPort 邮箱服务器端口，如腾讯企业邮箱为465

	FromEmail  string // FromEmail　发件人邮箱地址
	FromPasswd string //发件人邮箱密码（注意，这里是明文形式)

	Recipient []string //收件人邮箱地址
	CC        []string //抄送
}

var emailMessage *gomail.Message

const (
	ServerHost = "smtp.qq.com"
	ServerPort = 465
	FromEmail  = ""                 //发件人邮箱地址 TODO:填邮箱
	FromPasswd = "twqunycchmxecijg" //TODO:填密码
)

// SendEmail 发送验证码并返回验证码
func SendEmail(email string) string {
	relist := make([]string, 0)
	relist = append(relist, email)

	info := &EmailInfo{
		ServerHost,
		ServerPort,
		FromEmail,
		FromPasswd,
		relist,
		nil,
	}
	message := RandStringRunes(LengthOfStr)
	sendEmail("验证码", "<h1>您的验证码是:"+message+"</h1>", info)
	return message
}

func sendEmail(subject, body string, emailInfo *EmailInfo) {
	if len(emailInfo.Recipient) == 0 {
		log.Print("收件人列表为空")
		return
	}

	emailMessage = gomail.NewMessage()
	//设置收件人
	emailMessage.SetHeader("To", emailInfo.Recipient...)
	//设置抄送列表
	if len(emailInfo.CC) != 0 {
		emailMessage.SetHeader("Cc", emailInfo.CC...)
	}
	// 第三个参数为发件人别名，如"dcj"，可以为空（此时则为邮箱名称）
	emailMessage.SetAddressHeader("From", emailInfo.FromEmail, "dcj")

	//主题
	emailMessage.SetHeader("Subject", subject)

	//正文
	emailMessage.SetBody("text/html", body)

	d := gomail.NewPlainDialer(emailInfo.ServerHost, emailInfo.ServerPort,
		emailInfo.FromEmail, emailInfo.FromPasswd)
	err := d.DialAndSend(emailMessage)
	if err != nil {
		log.Println("发送邮件失败： ", err)
	} else {
		log.Println("已成功发送邮件到指定邮箱")
	}
}
