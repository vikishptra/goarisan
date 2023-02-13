package entity

import (
	"errors"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/util"
)

type User struct {
	ID               vo.UserID `bson:"_id" json:"id"`
	Created          time.Time `bson:"created" json:"created"`
	Name             string    `bson:"name" json:"name" form:"name"`
	Email            string    `bson:"email" json:"email" form:"email" `
	IsActive         bool      `bson:"is_active" json:"is_active" form:"is_active"`
	VerificationCode string
	RefreshToken     string
	Password         string `form:"password" json:"password"`
	Money            int64  `bson:"money" json:"money"`
}

type DataUserDetailGrupArisan struct {
	ID               vo.UserID          `bson:"_id" json:"id"`
	Name             string             `bson:"name" json:"name" `
	Money            int64              `bson:"money" json:"money"`
	DetailGrupArisan []DetailGrupArisan `json:"detail_grup"`
}
type DataUserGrupArisan struct {
	ID         vo.UserID    `bson:"_id" json:"id"`
	Name       string       `bson:"name" json:"name" form:"name"`
	Money      int64        `bson:"money" json:"money"`
	GrupArisan []Gruparisan `json:"grup"`
}

type UserCreateRequest struct {
	RandomString    string    `json:"id"`
	Now             time.Time `json:"time"`
	Name            string    `bson:"name" json:"name" form:"name"`
	Email           string    `bson:"email" json:"email" form:"email"`
	Password        string    `form:"password" json:"password"`
	ConfirmPassword string    `form:"confirm_password" json:"confirm_password"`
	Money           int64     `json:"-"`
}

func (r *User) ValidateUserCreate(req UserCreateRequest) error {
	if strings.TrimSpace(req.Email) == "" || strings.TrimSpace(req.Password) == "" {
		return errorenum.SepertinyaAdaYangSalahDariAnda
	} else if req.ConfirmPassword != req.Password {
		return errorenum.PasswordTidakSama
	}
	if err := ValidateName(req.Name); err != nil {
		return err
	}
	//hashpassword
	if err := CheckEmailValid(req.Email); err != nil {
		return err
	}
	if err := CheckEmailDomain(req.Email); err != nil {
		return err
	}
	return nil

}

func CheckEmailValid(email string) error {
	emailRegex, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		return errorenum.EmailIsNotValid
	}
	rg := emailRegex.MatchString(email)
	if !rg {
		return errorenum.EmailIsNotValid
	}
	if len(email) > 253 {
		return errorenum.EmailLengthIsTooLong
	}
	return nil
}

func CheckEmailDomain(email string) error {
	i := strings.Index(email, "@")
	host := email[i+1:]
	// func LookupMX(name string) ([]*MX, error)
	_, err := net.LookupMX(host)
	if err != nil {
		err = errorenum.EmailIsNotValid
		return err
	}
	return nil
}

func (r *User) ValidateTokenUser(IDUser, jwtToken vo.UserID) error {

	if IDUser != vo.UserID(jwtToken) {
		return errorenum.HayoMauNgapain
	}
	return nil
}

func (r *User) ValidateVerifyEmail(VerifyEmail bool) error {

	if !VerifyEmail {
		return errorenum.EmailBelumDiKonfirmasi
	}
	return nil
}

func NewUser(req UserCreateRequest) (*User, error) {

	code := randstr.String(4)

	verification_code := util.Encode(code)
	id, err := vo.NewUserID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj User
	obj.ID = id
	obj.Email = req.Email
	obj.Name = req.Name
	obj.Money = 0
	obj.VerificationCode = verification_code
	file := "verifyemail.html"
	temp := "domain_goarisan/templates/email"
	emailData := EmailData{
		URL:       os.Getenv("DOMAIN_EMAIL") + "/verifyemail/?code=" + code + "&id=" + string(obj.ID),
		FirstName: obj.Name,
		Subject:   "Verifikasi code kamu!",
	}
	obj.Created = req.Now

	go SendEmail(&obj, req.Email, &emailData, file, temp)

	return &obj, nil
}

type UserUpdateRequest struct {
	ID   vo.UserID `uri:"id"`
	Name string    `json:"name"`
	Jwt  vo.UserID `json:"token"`
}

func (r *User) Update(req UserUpdateRequest) error {
	r.Name = req.Name
	if err := ValidateName(req.Name); err != nil {
		return err
	}
	if req.ID != vo.UserID(req.Jwt) {
		return errorenum.HayoMauNgapain
	}

	return nil
}

func (r *User) VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r *User) UpdateMoney(money int64) error {
	if money < 0 {
		return errorenum.MoneyAndaTidakBolehKurangDariUpdateMoney
	}

	r.Money = r.Money + money

	return nil
}

func (r *User) CheckPasswordCriteria(password string) error {

	var err error
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			pswdLowercase = true
		case unicode.IsUpper(char):
			pswdUppercase = true
			err = errors.New("Pa")
		case unicode.IsNumber(char):
			pswdNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	if 7 < len(password) && len(password) < 60 {
		pswdLength = true
	}
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces {
		switch false {
		case pswdLowercase:
			err = errorenum.KataSandiHarusBerisiSetidaknyaSatuHurufKecil
		case pswdUppercase:
			err = errorenum.KataSandiHarusBerisiSetidaknyaSatuHurufBesar
		case pswdNumber:
			err = errorenum.KataSandiHarusBerisiSetidaknyaSatuAngka
		case pswdSpecial:
			err = errorenum.KataSandiHarusBerisiSetidaknyaSatuSpesialKarakter
		case pswdLength:
			err = errorenum.PanjangKataSandiHarusMinimal8KarakterDanKurangDari60
		case pswdNoSpaces:
			err = errorenum.KataSandiTidakBolehMemilikiSpasi
		}
		return err
	}
	return nil
}
func ValidateName(nama string) error {
	if strings.TrimSpace(nama) == "" {
		return errorenum.MessageNotEmpty
	} else if len(nama) > 253 {
		return errorenum.NamaTidakBolehLebihDari253
	}
	return nil
}

func (r *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	r.Password = string(hashedPassword)

	//hapusspace
	// r.Name = html.EscapeString(strings.TrimSpace(req.Name))
	return nil
}

func SendEmailConfirmUser(code string, obj *User) {
	verification_code := util.Encode(code)
	obj.VerificationCode = verification_code
}

func ChangePasswordWithEmail(user *User) {
	code := randstr.String(4)
	file := "verifypassword.html"
	temp := "domain_goarisan/templates/password"
	verification_code := util.Encode(code)
	user.VerificationCode = verification_code
	emailData := EmailData{
		URL:       os.Getenv("DOMAIN_EMAIL") + "/change/password?code=" + code + "&id=" + string(user.ID),
		FirstName: user.Name,
		Subject:   "Password Baru Kamu!",
	}

	go SendEmail(user, user.Email, &emailData, file, temp)
}

func (r *User) NewPasswordWithEmail(password, confirmPassword string) error {

	if password != confirmPassword {
		return errorenum.PasswordTidakSama
	}
	r.Password = password

	return nil
}
