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
