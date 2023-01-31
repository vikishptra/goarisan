package entity

import (
	"errors"
	"html"
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
	Name             string    `bson:"name" json:"name" form:"name" binding:"required"`
	Email            string    `bson:"email" json:"email" form:"email" binding:"required"`
	IsActive         bool      `bson:"is_active" json:"is_active" form:"is_active"`
	VerificationCode string
	RefreshToken     string
	Password         string `form:"password" json:"password" binding:"required"`
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
	Name       string       `bson:"name" json:"name" form:"name" binding:"required"`
	Money      int64        `bson:"money" json:"money"`
	GrupArisan []Gruparisan `json:"grup"`
}

type UserCreateRequest struct {
	RandomString string    `json:"id"`
	Now          time.Time `json:"time"`
	Name         string    `bson:"name" json:"name" form:"name" binding:"required"`
	Email        string    `bson:"email" json:"email" form:"email" binding:"required"`
	Password     string    `form:"password" json:"password" binding:"required"`
	Money        int64     `json:"-"`
}

func (r *User) ValidateUserCreate(req UserCreateRequest) error {
	if err := ValidateName(req.Name); err != nil {
		return err
	}
	//hashpassword
	if err := checkEmailValid(req.Email); err != nil {
		return err
	}
	if err := checkEmailDomain(req.Email); err != nil {
		return err
	}
	return nil

}

func checkEmailValid(email string) error {
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

func checkEmailDomain(email string) error {
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
	obj.Created = req.Now
	obj.Name = req.Name
	obj.Money = 0
	obj.VerificationCode = verification_code

	emailData := EmailData{
		URL:       os.Getenv("ORIGIN") + "/verifyemail/?code=" + code + "&id=" + string(obj.ID),
		FirstName: obj.Name,
		Subject:   "Verifikasi code kamu!",
	}

	go SendEmail(&obj, req.Email, &emailData)

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

func (r *User) HashPassword(req UserCreateRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	r.Password = string(hashedPassword)

	//hapusspace
	r.Name = html.EscapeString(strings.TrimSpace(req.Name))
	return nil
}
