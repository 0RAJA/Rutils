package email

import (
	"fmt"
	"testing"
	"time"

	"github.com/0RAJA/Rutils/pkg/times"
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
		UserName: "1296643805@qq.com",
		Password: "yfbxrbthozfnifjd",
		From:     "1296643805@qq.com",
	})
	err := defailtMailer.SendMail( // 短信通知
		[]string{"1296643805@qq.com"},
		fmt.Sprintf("异常抛出，发生时间: %s,%d", times.GetNowDateTimeStr(), time.Now().Unix()),
		fmt.Sprintf("错误信息: %v", "NO"),
	)
	if err != nil {
		fmt.Println(err)
	}
}
