package repository

import (
	"context"

	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/domain_goarisan/model/vo"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
	FindUsername(ctx context.Context, username string) (*entity.User, error)
	FindEmail(ctx context.Context, email string) (*entity.User, error)
}

type FindUserByIDRepo interface {
	FindUserByID(ctx context.Context, UserID vo.UserID) (*entity.User, error)
}

type SaveGrupArisanRepo interface {
	SaveGrupArisan(ctx context.Context, obj *entity.Gruparisan) error
}

type SaveDetailGrupArisanRepo interface {
	SaveDetailGrupArisan(ctx context.Context, obj *entity.DetailGrupArisan) error
}

type FindGrupArisanByIdRepo interface {
	FindGrupArisanAndUserById(ctx context.Context, someID vo.GruparisanID, userID vo.UserID) (*entity.Gruparisan, error)
	FindGrupArisanyIdGrup(ctx context.Context, GrupArisanId vo.GruparisanID) (*entity.Gruparisan, error)
}

type FindUndianArisanUserRepo interface {
	FindUndianArisanUser(ctx context.Context, IDgrup vo.GruparisanID) ([]map[string]any, error)
}

type FindOneGrupByOwnerRepo interface {
	FindOneGrupByOwner(ctx context.Context, IDUser vo.UserID, IDGrup vo.GruparisanID) (*entity.Gruparisan, error)
}

type RunLoginRepo interface {
	RunLogin(ctx context.Context, email, password string) (string, string, *entity.User, error)
	RunLogout(ctx context.Context, user vo.UserID) error
	RunRefreshTokenJwt(ctx context.Context, IDuser vo.UserID) (string, error)
}

type GetfindgrupbyidownerRepo interface {
	Getfindgrupbyidowner(ctx context.Context, IDUser vo.UserID) ([]any, error)
	Getfindgruparisanbyiduser(ctx context.Context, IDUser vo.UserID) ([]any, error)
}

type FindoneuserdetailgruparisansRepo interface {
	Findoneuserdetailgruparisans(ctx context.Context, IDGrup vo.GruparisanID, IDUser vo.UserID) (*entity.DetailGrupArisan, error)
}

type DeleteDetailGrupArisanRepo interface {
	DeleteUserDetailGrupArisan(ctx context.Context, IDUser vo.UserID, IDGrup vo.GruparisanID, IDOwner vo.UserID) error
}

type RunUpdateOwnerGrupRepo interface {
	RunUpdateOwnerGrup(ctx context.Context, IDUser vo.UserID, IDGrup vo.GruparisanID, IDOwner vo.UserID) (*entity.Gruparisan, error)
}
