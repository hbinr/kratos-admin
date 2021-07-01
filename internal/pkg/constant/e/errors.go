package e

import "errors"

var (
	ErrNotFound  = errors.New(CodeNotFound.Msg())
	ErrInvalidID = errors.New(CodeInvalidID.Msg())

	ErrUserNotLogin = errors.New(CodeNeedLogin.Msg())
	ErrUserHasExist = errors.New(CodeUserHasExist.Msg())
	ErrEmailExist   = errors.New(CodeEmailExist.Msg())
)
