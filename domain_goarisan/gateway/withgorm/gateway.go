package withgorm

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/config"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/token"
	"vikishptra/shared/util"
)

type Gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	Db      *gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *Gateway {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// dbUser := os.Getenv("MYSQLUSER")
	// dbPassword := os.Getenv("MYSQLPASSWORD")
	// dbHost := os.Getenv("MYSQLHOST")
	// dbPort := os.Getenv("MYSQLPORT")
	// database := os.Getenv("MYSQLDATABASE")

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, database)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	Db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	err = Db.AutoMigrate(entity.User{}, entity.Gruparisan{}, entity.DetailGrupArisan{}).Error
	Db.Model(entity.Gruparisan{}).AddForeignKey("id_owner", "users(id)", "CASCADE", "CASCADE")
	Db.Model(entity.DetailGrupArisan{}).AddForeignKey("id_detail_grup", "gruparisans(id)", "CASCADE", "CASCADE")
	Db.Model(entity.DetailGrupArisan{}).AddForeignKey("id_user", "users(id)", "CASCADE", "CASCADE")
	Db.Model(entity.User{}).AddIndex("idx_email", "email")
	Db.Model(entity.User{}).AddIndex("idx_name", "name")
	if err != nil {
		panic(err)
	}
	return &Gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		Db:      Db,
	}
}

func (r *Gateway) SaveUser(ctx context.Context, obj *entity.User) error {
	r.log.Info(ctx, "called")

	if err := r.Db.Save(obj).Error; err != nil {
		return err
	}
	return nil
}

func (r *Gateway) FindUserByID(ctx context.Context, UserID vo.UserID) (*entity.User, error) {
	r.log.Info(ctx, "called")
	var user entity.User
	if err := r.Db.First(&user, "id = ?", UserID); err.RecordNotFound() {
		return nil, errorenum.DataNotFound
	}
	return &user, nil
}

func (r *Gateway) SaveGrupArisan(ctx context.Context, obj *entity.Gruparisan) error {

	r.log.Info(ctx, "called")
	if err := r.Db.Save(obj).Error; err != nil {
		return err
	}

	return nil
}

func (r *Gateway) SaveDetailGrupArisan(ctx context.Context, obj *entity.DetailGrupArisan) error {
	r.log.Info(ctx, "called")

	if err := r.Db.Save(obj).Error; err != nil {
		return err
	}

	return nil
}

func (r *Gateway) FindGrupArisanAndUserById(ctx context.Context, GrupArisanId vo.GruparisanID, UserID vo.UserID) (*entity.Gruparisan, error) {
	var gruparisan entity.Gruparisan
	var user entity.DetailGrupArisan
	if err := r.Db.First(&gruparisan, "id = ?", GrupArisanId); err.RecordNotFound() {
		return nil, errorenum.DataNotFound
	}
	if err := r.Db.First(&user, "id_user = ?", UserID).First(&user, "id_detail_grup = ?", GrupArisanId); !err.RecordNotFound() {
		return nil, errorenum.UserAlreadyJoin
	}

	return &gruparisan, nil
}

func (r *Gateway) FindUndianArisanUser(ctx context.Context, IDGrup vo.GruparisanID) ([]map[string]any, error) {
	var detailGrupArisan entity.DetailGrupArisan
	var MaxStatus, MoneyUser int64
	var result []map[string]any

	//generate nilai random
	if err := r.Db.Table("detail_grup_arisans").Select("id_user,name, no_undian").Joins("INNER JOIN users ON users.id = detail_grup_arisans.id_user").Where("status_user_putaran_arisan = 0 AND id_detail_grup = ?", IDGrup).Order("RAND()").Find(&detailGrupArisan); err.RecordNotFound() {
		return nil, errorenum.DataArisanAndaSudahBerakhir
	}
	//menjumlahkan data uang dari grup arisan kalo 0 maka gabisa ngocok arisan
	MoneyUser, err := r.FindSumMoneyByIDGrup(ctx, IDGrup, MoneyUser)
	if err != nil {
		return nil, err
	}

	//temuin user id untuk dapatin data users yang di undi
	users, _ := r.FindUserByID(ctx, detailGrupArisan.ID_User)

	//pindahin ke result dengan tipe map
	result = append(result, map[string]any{"id_user": detailGrupArisan.ID_User, "no_undian": detailGrupArisan.No_undian, "name": users.Name})

	//mengambil data max dari status user by id grup
	if err := r.Db.Model(&entity.DetailGrupArisan{}).Select("MAX(status_user_putaran_arisan)").Where("id_detail_grup = ?", IDGrup).Row().Scan(&MaxStatus); err != nil {
		return nil, errorenum.SomethingError
	}
	MaxStatus = MaxStatus + 1
	//update stats dengan status_user_arisan
	if err := r.Db.Model(entity.DetailGrupArisan{}).Where("id_user = ? AND id_detail_grup = ?", detailGrupArisan.ID_User, IDGrup).Update("status_user_putaran_arisan", MaxStatus); err.Error != nil {
		return nil, err.Error
	}

	//update money
	MoneyUser = users.Money + MoneyUser
	if err := r.Db.Model(entity.User{}).Where("id = ?", detailGrupArisan.ID_User).Update("money", MoneyUser); err.Error != nil {
		return nil, errorenum.SomethingError
	}

	if err := r.Db.Model(entity.DetailGrupArisan{}).Where("id_detail_grup = ?", IDGrup).Update("money", 0); err.Error != nil {
		return nil, errorenum.SomethingError
	}

	return result, nil
}

func (r *Gateway) FindOneGrupByOwner(ctx context.Context, IDUser vo.UserID, IDGrup vo.GruparisanID) (*entity.Gruparisan, error) {
	var grup entity.Gruparisan
	if err := r.Db.Where("id_owner = ? AND id = ?", IDUser, IDGrup).Find(&grup); err.RecordNotFound() {
		return nil, errorenum.AndaBukanAdmin
	}

	return &grup, nil

}

func (r *Gateway) FindSumMoneyByIDGrup(ctx context.Context, IDGrup vo.GruparisanID, MoneyUser int64) (int64, error) {

	if err := r.Db.Model(&entity.DetailGrupArisan{}).Select("SUM(money)").Where("id_detail_grup = ?", IDGrup).Row().Scan(&MoneyUser); err != nil {
		return 0, errorenum.SomethingError
	}

	grupObj, err := r.FindGrupArisanyIdGrup(ctx, IDGrup)
	if err != nil {
		return 0, err
	}

	CountUser, err := r.CountDetailGrupByID(ctx, IDGrup)
	if err != nil {
		return 0, err
	}

	test := (CountUser * grupObj.RulesMoney) - 1

	if MoneyUser < test {
		return 0, errorenum.AnggotaGrupAndaMasihAdaYangBelumSetoran
	}
	return MoneyUser, nil
}

func (r *Gateway) CountDetailGrupByID(ctx context.Context, IDGrup vo.GruparisanID) (int64, error) {
	var CountUser int64
	if err := r.Db.Model(&entity.DetailGrupArisan{}).Select("COUNT(status_user_putaran_arisan)").Where("id_detail_grup = ?", IDGrup).Row().Scan(&CountUser); err != nil {
		return 0, errorenum.SomethingError
	}

	return CountUser, nil

}

func (r *Gateway) FindUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	if err := r.Db.Model(&user).Where("name = ?", username).Take(&user); !err.RecordNotFound() {
		return nil, errorenum.UsernameAndaSudahDigunakan
	}
	return &user, nil
}

func (r *Gateway) FindEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.Db.Model(&user).Where("email = ?", email).Take(&user); !err.RecordNotFound() {
		return nil, errorenum.EmailAndaSudahDigunakan
	}
	return &user, nil
}
func (r *Gateway) FindEmailConfirmUser(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.Db.Model(&user).Where("email = ? AND is_active = 0", email).Take(&user); err.RecordNotFound() {
		return nil, errorenum.SepertinyaAdaYangSalahDariAnda
	}
	return &user, nil
}
func (r *Gateway) RunLogin(ctx context.Context, email, password string) (string, string, *entity.User, error) {
	var user entity.User
	var UserPassword *entity.User
	if err := r.Db.First(&user, "email = ?", email); err.RecordNotFound() {
		return "", "", nil, errorenum.UsernameAtauPasswordAndaSalah
	}

	err := UserPassword.VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", "", nil, errorenum.UsernameAtauPasswordAndaSalah
	}
	//dapat token
	access_token, err := token.GenerateToken(user.ID)
	refreshToken, _ := token.RefreshToken(user.ID)
	if err != nil {
		return "", "", nil, err
	}

	if err := r.Db.Model(entity.User{}).Where("id = ?", user.ID).Update("refresh_token", refreshToken); err.Error != nil {
		return "", "", nil, errorenum.SomethingError
	}

	return refreshToken, access_token, &user, nil
}

func (r *Gateway) RunLogout(ctx context.Context, user vo.UserID) error {

	if err := r.Db.Model(entity.User{}).Where("id = ?", user).Update("refresh_token", ""); err.Error != nil {
		return errorenum.SomethingError
	}

	return nil
}

func (r *Gateway) Getfindgrupbyidowner(ctx context.Context, IDUser vo.UserID) ([]any, error) {
	var user entity.DataUserDetailGrupArisan
	var detail_grup_arisans []entity.DetailGrupArisan
	var testt []any
	if err := r.Db.Table("users AS u").Select("u.name, u.id, u.money").Where("u.id = ?", IDUser).Find(&user); err.RecordNotFound() {
		return nil, errorenum.DataGrupNotFound
	}

	if err := r.Db.Table("detail_grup_arisans AS d").Select("*").Where("d.id_user = ?", IDUser).Find(&detail_grup_arisans); err.RecordNotFound() {
		return nil, errorenum.DataGrupNotFound
	}

	user.DetailGrupArisan = detail_grup_arisans
	testt = append(testt, user)

	return testt, nil
}

func (r *Gateway) Findoneuserdetailgruparisans(ctx context.Context, IDGrup vo.GruparisanID, IDUser vo.UserID) (*entity.DetailGrupArisan, error) {
	var grup entity.DetailGrupArisan

	if err := r.Db.Where("id_detail_grup = ? AND  id_user = ? AND money = 0", IDGrup, IDUser).Find(&grup); err.RecordNotFound() {
		return nil, errorenum.DataNotFound
	}

	return &grup, nil
}

func (r *Gateway) FindGrupArisanyIdGrup(ctx context.Context, GrupArisanId vo.GruparisanID) (*entity.Gruparisan, error) {
	var gruparisan entity.Gruparisan

	if err := r.Db.Where("id = ?", GrupArisanId).Find(&gruparisan); err.RecordNotFound() {
		return nil, errorenum.DataGrupNotFound
	}

	return &gruparisan, nil
}
func (r *Gateway) Getfindgruparisanbyiduser(ctx context.Context, IDUser vo.UserID) ([]any, error) {
	var user entity.DataUserGrupArisan
	var gruparisan []entity.Gruparisan
	var test []any

	resultUser, err := r.FindUserByID(ctx, IDUser)
	if err != nil {
		return nil, err
	}

	if err := r.Db.Where("id_owner = ?", resultUser.ID).Find(&gruparisan); err.RecordNotFound() {
		return nil, errorenum.DataUserNotFound
	}
	user.GrupArisan = gruparisan
	user.ID = resultUser.ID
	user.Name = resultUser.Name
	user.Money = resultUser.Money
	test = append(test, user)

	return test, nil

}

func (r *Gateway) DeleteUserDetailGrupArisan(ctx context.Context, IDUser vo.UserID, IDGrup vo.GruparisanID, IDOwner vo.UserID) error {
	var detailGrupObj entity.DetailGrupArisan

	//verifikasi dulu ownernya
	if _, err := r.FindOneGrupByOwner(ctx, IDOwner, IDGrup); err != nil {
		return err
	}
	//verifikasi dulu usernya
	if err := r.Db.Where("id_user = ? AND id_detail_grup = ?", IDUser, IDGrup).Find(&detailGrupObj); err.RecordNotFound() {
		return errorenum.DataUserNotFound
	}

	//proses delete grup
	if err := r.Db.Where("id_user = ? AND money = 0 AND id_detail_grup = ? ", IDUser, IDGrup).Find(&detailGrupObj).Delete(&detailGrupObj); err.RecordNotFound() {
		return errorenum.DataUserMasihAdaSaldoArisan
	}

	return nil
}

func (r *Gateway) RunUpdateOwnerGrup(ctx context.Context, IDUser vo.UserID, IDGrup vo.GruparisanID, IDOwner vo.UserID) (*entity.Gruparisan, error) {

	grupArisan, err := r.FindGrupArisanyIdGrup(ctx, IDGrup)
	if err != nil {
		return nil, err
	}

	if _, err := r.FindOneGrupByOwner(ctx, IDOwner, IDGrup); err != nil {
		return nil, err
	}
	//verifikasi dulu usernya
	if err := r.Db.Model(&entity.DetailGrupArisan{}).Where("id_detail_grup = ? AND id_user = ?", IDGrup, IDUser); err.RecordNotFound() {
		return nil, errorenum.DataUserNotFound
	}
	return grupArisan, nil
}

func (r *Gateway) RunRefreshTokenJwt(ctx context.Context, IDuser vo.UserID, tokens string) (string, error) {
	var user entity.User
	if err := r.Db.First(&user, "id = ?", IDuser); err.RecordNotFound() {
		return "", errorenum.GabisaAksesBro
	}
	tokenss := user.RefreshToken
	if IDuser == "" || tokenss != tokens {
		return "", errorenum.GabisaAksesBro
	}
	if err := r.Db.Model(entity.User{}).Where("id = ?", IDuser); err.RecordNotFound() {
		return "", errorenum.HayoMauNgapain
	}
	token, err := token.GenerateToken(IDuser)
	if err != nil {
		return "", err
	}
	return token, nil

}

func (r *Gateway) RunVerifyEmail(ctx context.Context, id, code string) error {
	var user entity.User
	verification_code := util.Encode(code)
	if err := r.Db.First(&user, "id = ? AND verification_code = ?", id, verification_code); err.RecordNotFound() {
		return errorenum.SepertinyaAdaYangSalahDariAnda
	}
	if user.IsActive {
		return errorenum.EmailSudahDiKonfirmasi
	}
	currentTime := time.Now()
	then := user.Created.Add(time.Duration(24) * time.Hour)
	if currentTime.After(then) {
		return errorenum.KonfirmasiEmailAndaSudahKadaluawarsa
	}
	if err := r.Db.Model(entity.User{}).Where("id = ?", id).Update("is_active", 1); err.RecordNotFound() {
		return errorenum.SomethingError
	}

	return nil

}
