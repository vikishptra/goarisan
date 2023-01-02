package errorenum

import "vikishptra/shared/model/apperror"

const (
	SomethingError                          apperror.ErrorType = "ER0000 something error"
	DataNotFound                            apperror.ErrorType = "ER0001 data tidak di temukan"
	MessageNotEmpty                         apperror.ErrorType = "ER0002 pesan tidak boleh error"
	MoneyMin                                apperror.ErrorType = "ER0003 uang anda tidak boleh 0 minimal 10000"
	UserStrapped                            apperror.ErrorType = "ER0004 uang anda harus lebih besar dari rules money"
	UserAlreadyJoin                         apperror.ErrorType = "ER0005 user already join"
	AndaBukanAdmin                          apperror.ErrorType = "ER0006 anda bukan admin"
	AnggotaGrupAndaMasihAdaYangBelumSetoran apperror.ErrorType = "ER0007 anggota grup anda masih ada yang belum setoran"
	UsernameAtauPasswordAndaSalah           apperror.ErrorType = "ER0008 username atau password anda salah"
	UsernameTidakDiTemukan                  apperror.ErrorType = "ER0009 username tidak di temukan"
	HayoMauNgapain                          apperror.ErrorType = "ER0010 hayo mau ngapain"
)
