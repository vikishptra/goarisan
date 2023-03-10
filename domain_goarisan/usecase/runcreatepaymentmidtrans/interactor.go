package runcreatepaymentmidtrans

import (
	"context"

	"vikishptra/domain_goarisan/model/entity"
)

type runcreatepaymentmidtransInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runcreatepaymentmidtransInteractor{
		outport: outputPort,
	}
}

func (r *runcreatepaymentmidtransInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	resultObj, err := entity.NewTranscation(req.TranscationCreateRequest)
	if err != nil {
		return nil, err
	}
	resultMidtrans, err := r.outport.ChargeCoreApiBankTransfer(ctx, resultObj, req.TranscationCreateRequest)
	if err != nil {
		return nil, err
	}
	if err := r.outport.SavePayment(ctx, resultObj); err != nil {
		return nil, err
	}

	res.Item = resultMidtrans

	return res, nil
}
