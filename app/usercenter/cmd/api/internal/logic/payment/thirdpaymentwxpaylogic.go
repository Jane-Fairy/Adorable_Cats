package payment

import (
	"catroom/app/usercenter/cmd/api/internal/svc"
	"catroom/app/usercenter/cmd/api/internal/types"
	"catroom/app/usercenter/model"
	"catroom/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/zeromicro/go-zero/core/logx"
)

type ThirdPaymentwxPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThirdPaymentwxPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdPaymentwxPayLogic {
	return &ThirdPaymentwxPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdPaymentwxPayLogic) ThirdPaymentWxPay(req *types.ThirdPaymentWxPayReq) (resp *types.ThirdPaymentWxPayResp, err error) {
	// todo: add your logic here and delete this line
	var totalPrice int64   // Total amount paid for current order(cent)
	var description string // Current Payment Description.

	//payment platform
	switch req.ServiceType {
	case model.ThirdPaymentTypeOrderCatroom: //cat internal payment
		price, catDescription, err := l.getPaymentPriceDescription(req.OrderSn)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.ErrWxPayment), "Get price description err : %v requset: %+v", err, req)
		}
		totalPrice = price
		description = catDescription
	default:
		return nil, errors.Wrapf(xerr.NewErrMsg("Payment for this business type is not supported"), "Payment for this business type is not supported req: %+v", req)
	}

	wechatPrepayResqonse, err := l.createWxPrePayOrder(req.ServiceType, req.OrderSn, totalPrice, description)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ErrWxPrePayment), "create pre payment err : %v requset: %+v", err, req)
	}

	return &types.ThirdPaymentWxPayResp{
		Appid:     l.svcCtx.Config.WxMiniConf.AppId,
		NonceStr:  *wechatPrepayResqonse.NonceStr,
		PaySign:   *wechatPrepayResqonse.PaySign,
		Package:   *wechatPrepayResqonse.Package,
		Timestamp: *wechatPrepayResqonse.TimeStamp,
		SignType:  *wechatPrepayResqonse.SignType,
	}, nil
}

// Obtain the price and description information of the current order
func (l *ThirdPaymentwxPayLogic) getPaymentPriceDescription(orderSn string) (int64, string, error) {

	return 1, "nil", nil
}

// Generate wx pre-payment
func (l *ThirdPaymentwxPayLogic) createWxPrePayOrder(serviceType, orderSn string, totalPrice int64, description string) (*jsapi.PrepayWithRequestPaymentResponse, error) {

	return nil, nil
}
