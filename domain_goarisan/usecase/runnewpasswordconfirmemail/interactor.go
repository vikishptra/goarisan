package runnewpasswordconfirmemail

import (
	"context"
	"time"
)

type runnewpasswordconfirmemailInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runnewpasswordconfirmemailInteractor{
		outport: outputPort,
	}
}

func (r *runnewpasswordconfirmemailInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	userObj, err := r.outport.RunNewPasswordWithEmail(ctx, req.Id, req.Code)
	if err != nil {
		return nil, err
	}
	if err := userObj.NewPasswordWithEmail(req.Password, req.ConfirmPassword); err != nil {
		return nil, err
	}
	if err := userObj.CheckPasswordCriteria(req.Password); err != nil {
		return nil, err
	}
	if err := userObj.HashPassword(req.Password); err != nil {
		return nil, err
	}
	userObj.Created = time.Now()
	userObj.VerificationCode = ""
	if err := r.outport.SaveUser(ctx, userObj); err != nil {
		return nil, err
	}
	res.Message = "ok berhasil mengubah password anda"
	return res, nil
}
