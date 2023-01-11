package deletedetailgrupbyowner

import (
	"context"
)

type deletedetailgrupbyownerInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &deletedetailgrupbyownerInteractor{
		outport: outputPort,
	}
}

func (r *deletedetailgrupbyownerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	if err := req.DetailGrupArisan.ValidateTokenUser(req.IDOwner, req.IDOwner); err != nil {
		return nil, err
	}

	if err := r.outport.DeleteUserDetailGrupArisan(ctx, req.IDUser, req.IDGrup, req.IDOwner); err != nil {
		return nil, err
	}

	res.Message = "delete user grup success"

	return res, nil
}
