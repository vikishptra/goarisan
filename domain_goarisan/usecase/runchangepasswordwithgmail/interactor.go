package runchangepasswordwithgmail

import (
	"context"
	"time"

	"vikishptra/domain_goarisan/model/entity"
	"vikishptra/domain_goarisan/model/errorenum"
)

type runChangePasswordWithGmailInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runChangePasswordWithGmailInteractor{
		outport: outputPort,
	}
}

func (r *runChangePasswordWithGmailInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	if err := entity.CheckEmailValid(req.Email); err != nil {
		return nil, err
	}
	if err := entity.CheckEmailDomain(req.Email); err != nil {
		return nil, err
	}
	userObj, err := r.outport.FindEmailUser(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if !userObj.IsActive {
		return nil, errorenum.EmailBelumDiKonfirmasi
	}
	go entity.ChangePasswordWithEmail(userObj)
	userObj.Created = time.Now()
	if err := r.outport.SaveUser(ctx, userObj); err != nil {
		return nil, err
	}
	res.Message = "ok success mohon ke email anda untuk mengubah password anda yang sudah di kirimkan"
	return res, nil
}
