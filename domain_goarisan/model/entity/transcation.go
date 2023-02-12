package entity

import (
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"

	"vikishptra/domain_goarisan/model/vo"
)

type Transcation struct {
	ID               vo.TransaksiID `bson:"_id" json:"id"`
	IDUser           vo.UserID      `bson:"id_user" json:"id_user"`
	MoneyUser        int64          `bson:"money_user" json:"money_user"`
	ResponseMidtrans string         `bson:"response_midtrans" json:"response_midtrans"`
	Updated          time.Time      `bson:"updated" json:"updated"`
	Created          time.Time      `bson:"created" json:"created"`
}

type TranscationCreateRequest struct {
	RandomString        string                       `json:"-"`
	IDuser              vo.UserID                    `json:"id_user"`
	PaymentType         coreapi.CoreapiPaymentType   `json:"payment_type"`
	BankTransferDetails *coreapi.BankTransferDetails `json:"bank_transfer"`
	TransactionDetails  midtrans.TransactionDetails  `json:"transaction_details"`
	Now                 time.Time                    `json:"-"`
}
type BcaResponse struct {
	StatusCode        string             `json:"status_code"`
	StatusMessage     string             `json:"status_message"`
	TransactionID     string             `json:"transaction_id"`
	OrderID           string             `json:"order_id"`
	GrossAmount       string             `json:"gross_amount"`
	Currency          string             `json:"currency"`
	PaymentType       string             `json:"payment_type"`
	TransactionTime   string             `json:"transaction_time"`
	TransactionStatus string             `json:"transaction_status"`
	VaNumbers         []coreapi.VANumber `json:"va_numbers"`
	FraudStatus       string             `json:"fraud_status"`
}

func BCA(res coreapi.ChargeResponse) []any {
	var BCAResponse BcaResponse
	BCAResponse.StatusCode = res.StatusCode
	BCAResponse.StatusMessage = res.StatusMessage
	BCAResponse.GrossAmount = res.GrossAmount
	BCAResponse.TransactionID = res.TransactionID
	BCAResponse.OrderID = res.OrderID
	BCAResponse.GrossAmount = res.GrossAmount
	BCAResponse.Currency = res.Currency
	BCAResponse.PaymentType = res.PaymentType
	BCAResponse.TransactionTime = res.TransactionTime
	BCAResponse.TransactionStatus = res.TransactionStatus
	BCAResponse.FraudStatus = res.FraudStatus
	BCAResponse.VaNumbers = res.VaNumbers

	var resultMidtrans []any

	resultMidtrans = append(resultMidtrans, BCAResponse)

	return resultMidtrans
}

func NewTranscation(req TranscationCreateRequest) (*Transcation, error) {

	id, err := vo.NewTransaksiID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	var obj Transcation
	obj.ID = id
	obj.MoneyUser = req.TransactionDetails.GrossAmt
	obj.IDUser = req.IDuser
	obj.Updated = req.Now
	obj.Created = req.Now

	return &obj, nil
}

type TranscationUpdateRequest struct {
	// add field to update here ...
}

func (r *Transcation) Update(req TranscationUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}
