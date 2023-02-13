package runnotificationmidtrans

import (
	"context"
)

type runnotificationmidtransInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runnotificationmidtransInteractor{
		outport: outputPort,
	}
}

func (r *runnotificationmidtransInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	resultObj, err := r.outport.PushNotificationMidtrans(ctx, req.OrderID, req.TrcIDMidtrans)
	if err != nil {
		return nil, err
	}
	if err := r.outport.SavePayment(ctx, resultObj); err != nil {
		return nil, err
	}
	res.Message = "OK"

	return res, nil
}
