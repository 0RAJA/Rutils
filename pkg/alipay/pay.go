package alipay

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/smartwalle/alipay/v3"
)

type Client struct {
	client    *alipay.Client
	notifyURL string
	returnURL string
}

// Config 初始化配置文件
type Config struct {
	KAppID               string // 应用ID
	KPrivateKey          string // 应用私钥
	IsProduction         bool   // 是否是正式环境
	AppPublicCertPath    string // app公钥证书路径
	AliPayRootCertPath   string // alipay根证书路径
	AliPayPublicCertPath string // alipay公钥证书路径
	NotifyURL            string // 异步通知地址
	ReturnURL            string // 支付后回调链接地址
}

// Init 客户端初始化
func Init(config Config) *Client {
	var err error
	var aliClient *alipay.Client
	doThat := func(f func() error) {
		if err = f(); err != nil {
			panic(err)
		}
	}
	doThat(func() error {
		aliClient, err = alipay.New(config.KAppID, config.KPrivateKey, config.IsProduction)
		return err
	})
	doThat(func() error { return aliClient.LoadAppPublicCertFromFile(config.AppPublicCertPath) })
	doThat(func() error { return aliClient.LoadAliPayRootCertFromFile(config.AliPayRootCertPath) })
	doThat(func() error { return aliClient.LoadAliPayPublicCertFromFile(config.AliPayPublicCertPath) })
	return &Client{client: aliClient, notifyURL: config.NotifyURL, returnURL: config.ReturnURL}
}

type Order struct {
	ID          string      // 订单ID
	Subject     string      // 订单标题
	TotalAmount float32     // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	Code        ProductCode // 销售产品码，与支付宝签约的产品码名称
}

type ProductCode string

const (
	AppPay       ProductCode = "QUICK_MSECURITY_PAY"    // app支付
	PhoneWebPay  ProductCode = "QUICK_WAP_WAY"          // 手机网站支付
	LaptopWebPay ProductCode = "FAST_INSTANT_TRADE_PAY" // 电脑网站支付
)

var (
	ErrOrderAmountOver = errors.New("订单金额超限")
	ErrVerifySign      = errors.New("异步通知验证签名未通过")
)

// Pay 订单支付请求，返回支付界面链接及可能出现的错误
func (client *Client) Pay(order Order) (payUrl string, err error) {
	if order.TotalAmount < 0.01 || order.TotalAmount > 100000000 {
		return "", ErrOrderAmountOver
	}
	var p = alipay.TradePagePay{}
	p.NotifyURL = client.notifyURL
	p.ReturnURL = client.returnURL
	p.Subject = order.Subject
	p.OutTradeNo = order.ID
	p.TotalAmount = strconv.FormatFloat(float64(order.TotalAmount), 'f', 2, 32)
	p.ProductCode = string(order.Code)
	pay, err := client.client.TradePagePay(p)
	if err != nil {
		return "", err
	}
	return pay.String(), nil
}

// VerifyForm 校验form表单并返回对应订单ID(注意: callback为get,notify为post)
func (client *Client) VerifyForm(form url.Values) (orderID string, err error) {
	ok, err := client.client.VerifySign(form)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", ErrVerifySign
	}
	orderID = form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = orderID
	rsp, err := client.client.TradeQuery(p)
	if err != nil {
		return "", fmt.Errorf("异步通知验证订单 %s 信息发生错误: %s", orderID, err.Error())
	}
	if !rsp.IsSuccess() {
		return "", fmt.Errorf("异步通知验证订单 %s 信息发生错误: %s-%s", orderID, rsp.Content.Msg, rsp.Content.SubMsg)
	}
	return orderID, nil
}
