package email

import (
	"fmt"
	"github.com/0RAJA/Rutils/pkg/times"
	"testing"
	"time"
)

/*
Host: smtp.qq.com
  Port: 465
  UserName: XXX@qq.com
  Password:
  IsSSL: true
  From: XXX@qq.com
  To:
    - XXX@qq.com
*/
func TestEmail_SendMail(t *testing.T) {
	defailtMailer := NewEmail(&SMTPInfo{
		Host:     "smtp.qq.com",
		Port:     465,
		IsSSL:    true,
		UserName: "XXX@qq.com",
		Password: "",
		From:     "XXX@qq.com",
	})
	err := defailtMailer.SendMail( //短信通知
		[]string{"XXX@qq.com"},
		fmt.Sprintf("异常抛出，发生时间: %s,%d", times.GetNowDateTimeStr(), time.Now().Unix()),
		fmt.Sprintf("错误信息: %v", "NO"),
	)
	if err != nil {
		fmt.Println(err)
	}
}
