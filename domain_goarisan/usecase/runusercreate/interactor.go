package runusercreate

import (
	"context"

	"vikishptra/domain_goarisan/model/entity"
)

type runUserCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runUserCreateInteractor{
		outport: outputPort,
	}
}

func (r *runUserCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}
	userObj, err := r.outport.FindEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if err := userObj.CheckPasswordCriteria(req.Password); err != nil {
		return nil, err
	}
	if err := userObj.ValidateUserCreate(req.UserCreateRequest); err != nil {
		return nil, err
	}
	userObj, err = entity.NewUser(req.UserCreateRequest)
	if err != nil {
		return nil, err
	}
	if err := userObj.HashPassword(req.UserCreateRequest); err != nil {
		return nil, err
	}
	if err := r.outport.SaveUser(ctx, userObj); err != nil {
		return nil, err
	}

	// message := []any{
	// 	"ok success create user",
	// }
	res.Message = "ok success mohon ke email anda untuk verifikasi akun anda yang sudah di kirimkan"

	return res, nil
}
