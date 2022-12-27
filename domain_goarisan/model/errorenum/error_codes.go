package errorenum

import "vikishptra/shared/model/apperror"

const (
	SomethingError  apperror.ErrorType = "ER0000 something error"
	DataNotFound    apperror.ErrorType = "ER0001 data tidak di temukan"
	MessageNotEmpty apperror.ErrorType = "ER0002 pesan tidak boleh error"
	MoneyMin        apperror.ErrorType = "ER0003 uang anda tidak boleh 0 minimal 10000"
	UserStrapped    apperror.ErrorType = "ER0004 uang anda harus lebih besar dari rules money"
)
