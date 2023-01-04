package findgrupbyidowner

import (
	"context"

	"vikishptra/shared/util"
)

type findgrupbyidownerInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &findgrupbyidownerInteractor{
		outport: outputPort,
	}
}

func (r *findgrupbyidownerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObj, err := r.outport.FindUserByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	if err := userObj.ValidateTokenUser(req.UserID, req.JwtToken); err != nil {
		return nil, err
	}

	resultObj, err := r.outport.Getfindgrupbyidowner(ctx, req.UserID)

	if err != nil {
		return nil, err
	}

	res.Item = util.ToSliceAny(resultObj)
	return res, nil
}
