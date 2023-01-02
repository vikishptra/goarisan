package repository

import (
	"context"

	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/domain_goarisan/model/vo"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
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
}

type FindUndianArisanUserRepo interface {
	FindUndianArisanUser(ctx context.Context, IDgrup vo.GruparisanID) ([]map[string]any, error)
}

type FindOneGrupByOwnerRepo interface {
	FindOneGrupByOwner(ctx context.Context, IDUser vo.UserID, IDGrup vo.GruparisanID) error
}

type RunLoginRepo interface {
	RunLogin(ctx context.Context, username, password string) (string, *entity.User, error)
}

type RunLogoutRepo interface {
	DeleteAuth(s []string, r string) []string
}
