package e

// ResCode .
type ResCode int

const (
	CodeSuccess       ResCode = 200
	CodeInvalidParams ResCode = 400
	CodeNotFound      ResCode = 404
	CodeInternalError ResCode = 500

	CodeConvDataErr       ResCode = 10000
	CodeValidateParamsErr ResCode = 10001
	CodeInvalidToken      ResCode = 10002
	CodeNeedLogin         ResCode = 10003
	CodeInvalidID         ResCode = 10004

	CodeUserHasExist            ResCode = 20002
	CodeEmailExist              ResCode = 20003
	CodeWrongPassword           ResCode = 20004
	CodeWrongUserNameOrPassword ResCode = 20005
)
