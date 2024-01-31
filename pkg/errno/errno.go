package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode    = 0
	ServiceErrCode = iota + 10000
	ParamErrCode
	AuthorizationFailedErrCode

	UserAlreadyExistErrCode
	UserIsNotExistErrCode
	BalanceNotEnoughErrCode

	RecordAlreadyExistErrCode
	RecordNotExistErrCode

	TransactionDoneErrCode
	AmountNotMatchErrCode
	PayIdNotMatchErrCode

	SdProcessErrCode
)

const (
	SuccessMsg               = "Success"
	ServerErrMsg             = "Service is unable to start successfully"
	ParamErrMsg              = "Wrong Parameter has been given"
	UserIsNotExistErrMsg     = "user is not exist"
	PasswordIsNotVerifiedMsg = "username or password not verified"
)

type ErrNo struct {
	ErrCode int32
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(code int32, msg string) ErrNo {
	return ErrNo{code, msg}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

var (
	Success                = NewErrNo(SuccessCode, SuccessMsg)
	ServiceErr             = NewErrNo(ServiceErrCode, ServerErrMsg)
	ParamErr               = NewErrNo(ParamErrCode, ParamErrMsg)
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
	UserIsNotExistErr      = NewErrNo(UserIsNotExistErrCode, UserIsNotExistErrMsg)
	PasswordIsNotVerified  = NewErrNo(AuthorizationFailedErrCode, PasswordIsNotVerifiedMsg)
	BalanceNotEnoughErr    = NewErrNo(BalanceNotEnoughErrCode, "User balance not enough")
	RecordNotExistErr      = NewErrNo(RecordNotExistErrCode, "Record does not exist")

	TransactionDoneErr = NewErrNo(TransactionDoneErrCode, "Transaction already done")
	AmountNotMatchErr  = NewErrNo(AmountNotMatchErrCode, "Transaction amount does not match")
	PayIdNotMatchErr   = NewErrNo(PayIdNotMatchErrCode, "Transaction pay_id does not match")
	SdProcessErr       = NewErrNo(SdProcessErrCode, "Sd process Failed")
)

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
