package withgorm

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/config"
	"vikishptra/shared/infrastructure/logger"
)

type Gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	Db      *gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *Gateway {
	Db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/arisan?charset=utf8&parseTime=True")

	if err != nil {
		panic(err)
	}
	err = Db.AutoMigrate(entity.User{}, entity.Gruparisan{}, entity.DetailGrupArisan{}).Error
	Db.Model(entity.Gruparisan{}).AddForeignKey("id_owner", "users(id)", "CASCADE", "CASCADE")                   // Foreign key need to define manually
	Db.Model(entity.DetailGrupArisan{}).AddForeignKey("id_detail_grup", "gruparisans(id)", "CASCADE", "CASCADE") // Foreign key need to define manually
	Db.Model(entity.DetailGrupArisan{}).AddForeignKey("id_user", "users(id)", "CASCADE", "CASCADE")              // Foreign key need to define manually

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

	var result []map[string]any
	//generate nilai random
	if err := r.Db.Table("detail_grup_arisans").Select("id_user,name, no_undian").Joins("INNER JOIN users ON users.id = detail_grup_arisans.id_user").Where("status_user_putaran_arisan = 0 AND id_detail_grup = ?", IDGrup).Order("RAND()").Find(&detailGrupArisan); err.RecordNotFound() {
		return nil, errorenum.DataNotFound
	}
	//temuin user id untuk dapatin data users yang di undi
	users, _ := r.FindUserByID(ctx, detailGrupArisan.ID_User)

	//pindahin ke result dengan tipe map
	result = append(result, map[string]any{"id_user": detailGrupArisan.ID_User, "no_undian": detailGrupArisan.No_undian, "name": users.Name})

	var MaxStatus int64
	if err := r.Db.Model(&entity.DetailGrupArisan{}).Select("MAX(status_user_putaran_arisan)").Where("id_detail_grup = ?", IDGrup).Row().Scan(&MaxStatus); err != nil {
		return nil, errorenum.SomethingError
	}
	MaxStatus = MaxStatus + 1
	//update stats dengan status_user_arisan
	if err := r.Db.Model(entity.DetailGrupArisan{}).Where("id_user = ? AND id_detail_grup = ?", detailGrupArisan.ID_User, IDGrup).Update("status_user_putaran_arisan", MaxStatus); err.Error != nil {
		return nil, err.Error
	}

	//select uang money dari money by grup
	var MoneyUser int64
	if err := r.Db.Model(entity.DetailGrupArisan{}).Select("SUM(money)").Where("id_detail_grup = ?", IDGrup).Row().Scan(&MoneyUser); err != nil {
		return nil, errorenum.SomethingError
	}
	//update money
	users.Money = users.Money + MoneyUser
	if err := r.Db.Model(entity.User{}).Where("id = ?", detailGrupArisan.ID_User).Update("money", users.Money); err.Error != nil {
		return nil, err.Error
	}

	if err := r.Db.Model(entity.DetailGrupArisan{}).Where("id_detail_grup = ?", IDGrup).Update("money", 0); err.Error != nil {
		return nil, err.Error
	}

	return result, nil
}

func (r *Gateway) FindOneGrupByOwner(ctx context.Context, IDUser vo.UserID, IDGrup vo.GruparisanID) error {

	if err := r.Db.Model(entity.Gruparisan{}).Where("id = ? AND id_owner = ? ", IDGrup, IDUser); err.Error != nil {
		return errorenum.AndaBukanAdmin
	}

	return nil

}
