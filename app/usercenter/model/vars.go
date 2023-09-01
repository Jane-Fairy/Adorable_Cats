package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

var UserAuthTypeSystem string = "System"   //平台内部
var WxMiNiProgram string = "WxMiNiProgram" //微信小程序

// 支付业务类型
var ThirdPaymentTypeOrderCatroom string = "catroom" //猫咖支付

// 支付方式
var ThirdPaymentWeChatPay = "WECHAT_PAY" //微信支付

// 支付状态
var PaymentFAIL int64 = -1   //支付失败
var PaymenWait int64 = 0     //待支付
var PaymentSuccess int64 = 1 //支付成功
var PaymentRefund int64 = 2  //已退款
