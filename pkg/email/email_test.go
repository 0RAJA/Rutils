package email

import (
	"Rutils/pkg/times"
	"fmt"
	"testing"
	"time"
)

/*
Host: smtp.qq.com
  Port: 465
  UserName: 1647193241@qq.com
  Password: hrefwwxzvxgbehfc
  IsSSL: true
  From: 1647193241@qq.com
  To:
    - 1647193241@qq.com
*/
func TestEmail_SendMail(t *testing.T) {
	defailtMailer := NewEmail(&SMTPInfo{
		Host:     "smtp.qq.com",
		Port:     465,
		IsSSL:    true,
		UserName: "1647193241@qq.com",
		Password: "hrefwwxzvxgbehfc",
		From:     "1647193241@qq.com",
	})
	err := defailtMailer.SendMail( //短信通知
		[]string{"1647193241@qq.com"},
		fmt.Sprintf("异常抛出，发生时间: %s,%d", times.GetNowDateTimeStr(), time.Now().Unix()),
		fmt.Sprintf("错误信息: %v", "NO"),
	)
	if err != nil {
		fmt.Println(err)
	}
}
