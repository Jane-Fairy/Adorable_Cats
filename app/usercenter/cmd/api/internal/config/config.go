package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	WxMiniConf WxMiniConf
	WxPayConf  WxPayConf
	JwtAuth    struct {
		AccessSecret string
		AccessExpire int64
	}
	UserCenterRpcConf zrpc.RpcClientConf
}

// wechat config
type WxMiniConf struct {
	AppId  string `json:"AppId"`  //微信小程序appId（非公众号）
	Secret string `json:"Secret"` //微信小程序secret（非公众号）
}

// wechat payment config
type WxPayConf struct {
	MchId      string `json:"MchId"`      //微信商户id
	SerialNo   string `json:"SerialNo"`   //商户证书的证书序列号
	APIv3Key   string `json:"APIv3Key"`   //apiV3Key，商户平台获取
	PrivateKey string `json:"PrivateKey"` //privateKey：私钥 apiclient_key.pem 读取后的内容
	NotifyUrl  string `json:"NotifyUrl"`  //支付通知回调服务端地址
}
