package errorenum

import "vikishptra/shared/model/apperror"

const (
	SomethingError                           apperror.ErrorType = "ER0000 something error"
	DataNotFound                             apperror.ErrorType = "ER0001 data tidak di temukan"
	MessageNotEmpty                          apperror.ErrorType = "ER0002 pesan tidak boleh kosong"
	MoneyMin                                 apperror.ErrorType = "ER0003 uang anda tidak boleh 0 minimal 10000"
	UserStrapped                             apperror.ErrorType = "ER0004 uang anda harus lebih besar dari rules money"
	UserAlreadyJoin                          apperror.ErrorType = "ER0005 user already join"
	AndaBukanAdmin                           apperror.ErrorType = "ER0006 anda bukan admin"
	AnggotaGrupAndaMasihAdaYangBelumSetoran  apperror.ErrorType = "ER0007 anggota grup anda masih ada yang belum setoran"
	UsernameAtauPasswordAndaSalah            apperror.ErrorType = "ER0008 username atau password anda salah"
	UsernameTidakDiTemukan                   apperror.ErrorType = "ER0009 username tidak di temukan"
	HayoMauNgapain                           apperror.ErrorType = "ER0010 hayo mau ngapain"
	DataUserNotFound                         apperror.ErrorType = "ER0011 data user not found"
	DataGrupNotFound                         apperror.ErrorType = "ER0012 data grup not found"
	UndianDuplicated                         apperror.ErrorType = "ER0013 undian duplicated"
	MoneyTidakBoleh0                         apperror.ErrorType = "ER0014 money tidak boleh0"
	MoneyAndaTidakBolehKurangDariUpdateMoney apperror.ErrorType = "ER0015 money anda tidak boleh kurang dari update money"
	GabisaAksesBro                           apperror.ErrorType = "ER0016 gabisa akses bro"
	RulesMoneyTidakBolehKurangDari0          apperror.ErrorType = "ER0017 rules money tidak boleh kurang dari 0"
	GabolehLebihDariRulesMoney               apperror.ErrorType = "ER0018 gaboleh lebih dari rules money"
	NoundianTersebutSudahAda                 apperror.ErrorType = "ER0019 noundian tersebut sudah ada"
	UserStrappedAtauLebihDariRulesMoney      apperror.ErrorType = "ER0020 uang anda harus pas dengan rules money"
	DataArisanAndaSudahBerakhir              apperror.ErrorType = "ER0021 data arisan anda sudah berakhir"
)
