package xerr

// 成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const ServerCommonError uint32 = 100001
const RequestParamError uint32 = 100002
const TokenExpireError uint32 = 100003
const TokenGenerateError uint32 = 100004
const DbError uint32 = 100005
const DbUpdateAffectedZeroError uint32 = 100006
const UserExistMobile uint32 = 100007
const PasswordError uint32 = 100008
const ErrWxPayment = 100009
const ErrWxPrePayment = 100010

// 用户模块
const UserHasBeenRegister uint32 = 200001
const GenerateTokenFailed uint32 = 200002
const ErrGenerateTokenError uint32 = 200003
const UserNotExitsError uint32 = 200004
