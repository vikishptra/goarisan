package findgruparisanbyidowner

import (
	"context"
)

type findgruparisanbyidOwnerInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &findgruparisanbyidOwnerInteractor{
		outport: outputPort,
	}
}

func (r *findgruparisanbyidOwnerInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	resultGrupArisan, err := r.outport.Getfindgruparisanbyiduser(ctx, "003de5b6-6304-4603-90d0-dd2736217d93")
	if err != nil {
		return nil, err
	}

	if err := req.Grup.ValidateTokenUserGrupArisan(req.Grup.ID_Owner, req.JwtToken); err != nil {
		return nil, err
	}

	res.Item = resultGrupArisan

	return res, nil
}
