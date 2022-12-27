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

type gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	Db      *gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {
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
	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		Db:      Db,
	}
}

func (r *gateway) SaveUser(ctx context.Context, obj *entity.User) error {
	r.log.Info(ctx, "called")

	if err := r.Db.Save(obj).Error; err != nil {
		return err
	}
	return nil
}

func (r *gateway) FindUserByID(ctx context.Context, UserID vo.UserID) (*entity.User, error) {
	r.log.Info(ctx, "called")
	var user entity.User
	if err := r.Db.First(&user, "id = ?", UserID); err.RecordNotFound() {
		return nil, errorenum.DataNotFound
	}
	return &user, nil
}

func (r *gateway) SaveGrupArisan(ctx context.Context, obj *entity.Gruparisan) error {

	r.log.Info(ctx, "called")
	if err := r.Db.Save(obj).Error; err != nil {
		return err
	}

	return nil
}

func (r *gateway) SaveDetailGrupArisan(ctx context.Context, obj *entity.DetailGrupArisan) error {
	r.log.Info(ctx, "called")

	if err := r.Db.Save(obj).Error; err != nil {
		return err
	}

	return nil
}
