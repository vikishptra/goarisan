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

	if err := r.outport.DeleteUserDetailGrupArisan(ctx, req.IDUserDetailGrup); err != nil {
		return nil, err
	}

	res.Message = "delete user grup success"

	return res, nil
}
