package sendemailconfirm

import (
	"context"
	"os"

	"github.com/thanhpk/randstr"

	"vikishptra/domain_goarisan/model/entity"
)

type sendemailconfirmInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &sendemailconfirmInteractor{
		outport: outputPort,
	}
}

func (r *sendemailconfirmInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {
	res := &InportResponse{}
	code := randstr.String(4)
	if err := entity.CheckEmailValid(req.Email); err != nil {
		return nil, err
	}
	if err := entity.CheckEmailDomain(req.Email); err != nil {
		return nil, err
	}
	userObj, err := r.outport.FindEmailConfirmUser(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	emailData := entity.EmailData{
		URL:       os.Getenv("DOMAIN_EMAIL") + "/verifyemail/?code=" + code + "&id=" + userObj.ID.String(),
		FirstName: userObj.Name,
		Subject:   "Verifikasi code kamu!",
	}
	entity.SendEmailConfirmUser(code, userObj)
	if err := r.outport.SaveUser(ctx, userObj); err != nil {
		return nil, err
	}
	go entity.SendEmail(userObj, req.Email, &emailData)

	res.Message = "ok success mohon ke email anda untuk verifikasi akun anda yang sudah di kirimkan"

	return res, nil
}
