package errorenum

import "vikishptra/shared/model/apperror"

const (
	SomethingError                                       apperror.ErrorType = "ER0000 something error"
	DataNotFound                                         apperror.ErrorType = "ER0001 data tidak di temukan"
	MessageNotEmpty                                      apperror.ErrorType = "ER0002 pesan tidak boleh kosong"
	MoneyMin                                             apperror.ErrorType = "ER0003 uang anda tidak boleh 0 minimal 10000"
	UserStrapped                                         apperror.ErrorType = "ER0004 uang anda harus lebih besar dari rules money"
	UserAlreadyJoin                                      apperror.ErrorType = "ER0005 user already join"
	AndaBukanAdmin                                       apperror.ErrorType = "ER0006 anda bukan owner"
	AnggotaGrupAndaMasihAdaYangBelumSetoran              apperror.ErrorType = "ER0007 anggota grup anda masih ada yang belum setoran"
	UsernameAtauPasswordAndaSalah                        apperror.ErrorType = "ER0008 username atau password anda salah"
	UsernameTidakDiTemukan                               apperror.ErrorType = "ER0009 username tidak di temukan"
	HayoMauNgapain                                       apperror.ErrorType = "ER0010 hayo mau ngapain"
	DataUserNotFound                                     apperror.ErrorType = "ER0011 data user not found"
	DataGrupNotFound                                     apperror.ErrorType = "ER0012 data grup not found"
	UndianDuplicated                                     apperror.ErrorType = "ER0013 undian duplicated"
	MoneyTidakBoleh0                                     apperror.ErrorType = "ER0014 money tidak boleh0"
	MoneyAndaTidakBolehKurangDariUpdateMoney             apperror.ErrorType = "ER0015 money anda tidak boleh kurang dari update money"
	GabisaAksesBro                                       apperror.ErrorType = "ER0016 gabisa akses bro belum login"
	RulesMoneyTidakBolehKurangDari0                      apperror.ErrorType = "ER0017 rules money tidak boleh kurang dari 0"
	GabolehLebihDariRulesMoney                           apperror.ErrorType = "ER0018 gaboleh lebih dari rules money"
	NoundianTersebutSudahAda                             apperror.ErrorType = "ER0019 noundian tersebut sudah ada"
	UserStrappedAtauLebihDariRulesMoney                  apperror.ErrorType = "ER0020 uang anda harus pas dengan rules money"
	DataArisanAndaSudahBerakhir                          apperror.ErrorType = "ER0021 data arisan anda sudah berakhir"
	DataUserMasihAdaSaldoArisan                          apperror.ErrorType = "ER0022 data user masih ada saldo arisan"
	AndaAdalahAdmin                                      apperror.ErrorType = "ER0023 anda adalah owner"
	EmailIsNotValid                                      apperror.ErrorType = "ER0024 email is not valid"
	EmailLengthIsTooLong                                 apperror.ErrorType = "ER0025 email length is too long"
	CouldNotFindEmail                                    apperror.ErrorType = "ER0026 could not find email"
	UsernameAndaSudahDigunakan                           apperror.ErrorType = "ER0027 username anda sudah digunakan"
	EmailAndaSudahDigunakan                              apperror.ErrorType = "ER0028 email anda sudah digunakan"
	EmailBelumDiKonfirmasi                               apperror.ErrorType = "ER0029 email belum di konfirmasi"
	TokenExpired                                         apperror.ErrorType = "ER0030 token expired"
	SepertinyaAdaYangSalahDariAnda                       apperror.ErrorType = "ER0031 sepertinya ada yang salah dari anda"
	EmailSudahDiKonfirmasi                               apperror.ErrorType = "ER0032 email sudah di konfirmasi"
	KataSandiHarusBerisiSetidaknyaSatuHurufKecil         apperror.ErrorType = "ER0033 kata sandi harus berisi setidaknya satu huruf kecil"
	KataSandiHarusBerisiSetidaknyaSatuHurufBesar         apperror.ErrorType = "ER0034 kata sandi harus berisi setidaknya satu huruf besar"
	KataSandiHarusBerisiSetidaknyaSatuAngka              apperror.ErrorType = "ER0035 kata sandi harus berisi setidaknya satu angka"
	KataSandiHarusBerisiSetidaknyaSatuSpesialKarakter    apperror.ErrorType = "ER0036 kata sandi harus berisi setidaknya satu spesial karakter"
	PanjangKataSandiHarusMinimal8KarakterDanKurangDari60 apperror.ErrorType = "ER0037 panjang kata sandi harus minimal 8 karakter dan kurang dari  60"
	KataSandiTidakBolehMemilikiSpasi                     apperror.ErrorType = "ER0038 kata sandi tidak boleh memiliki spasi"
	KonfirmasiEmailAndaSudahKadaluawarsa                 apperror.ErrorType = "ER0039 konfirmasi email anda sudah kadaluarwarsa"
	NamaTidakBolehLebihDari253                           apperror.ErrorType = "ER0040 nama tidak boleh lebih dari 253"
	TooManyRequests                                      apperror.ErrorType = "ER0041 too many requests"
	PasswordTidakSama                                    apperror.ErrorType = "ER0042 password tidak sama"
	EmailAndaTidakTerdaftarPergiUntukDaftarAkunAnda      apperror.ErrorType = "ER0043 email anda tidak terdaftar pergi untuk daftar akun anda"
)
