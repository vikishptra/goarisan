package entity

import (
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/model/vo"
)

type User struct {
	ID       vo.UserID `bson:"_id" json:"id"`
	Created  time.Time `bson:"created" json:"created"`
	Name     string    `bson:"name" json:"name" form:"name" binding:"required"`
	Password string    `form:"password" json:"password" binding:"required"`
	Money    int64     `bson:"money" json:"money"`
}

type DataUserDetailGrupArisan struct {
	ID               vo.UserID          `bson:"_id" json:"id"`
	Name             string             `bson:"name" json:"name" form:"name" binding:"required"`
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
	Password     string    `form:"password" json:"password" binding:"required"`
	Money        int64     `json:"-"`
}

func (r *User) ValidateUserCreate(req UserCreateRequest) error {

	//hashpassword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	r.Password = string(hashedPassword)

	//hapusspace
	r.Name = html.EscapeString(strings.TrimSpace(req.Name))

	return nil

}
func (r *User) ValidateTokenUser(IDUser, jwtToken vo.UserID) error {

	if IDUser != vo.UserID(jwtToken) {
		return errorenum.HayoMauNgapain
	}

	return nil
}

func NewUser(req UserCreateRequest) (*User, error) {

	id, err := vo.NewUserID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj User
	obj.ID = id
	obj.Created = req.Now
	obj.Name = req.Name
	obj.Money = 0

	return &obj, nil
}

type UserUpdateRequest struct {
	ID   vo.UserID `uri:"id"`
	Name string    `json:"name"`
	Jwt  vo.UserID `json:"token"`
}

func (r *User) Update(req UserUpdateRequest) error {
	r.Name = req.Name
	if strings.TrimSpace(req.Name) == "" {
		return errorenum.MessageNotEmpty
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
