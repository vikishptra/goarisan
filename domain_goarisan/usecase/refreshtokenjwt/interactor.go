package refreshtokenjwt

import (
	"context"

	"vikishptra/domain_goarisan/model/vo"
)

type refreshtokenjwtInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &refreshtokenjwtInteractor{
		outport: outputPort,
	}
}

func (r *refreshtokenjwtInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	token, err := r.outport.RunRefreshTokenJwt(ctx, vo.UserID(req.IDUser), req.Token)
	if err != nil {
		return nil, err
	}
	res.AccessToken = token
	return res, nil
}
