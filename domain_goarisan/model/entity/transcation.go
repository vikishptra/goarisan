package entity

import (
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"

	"vikishptra/domain_goarisan/model/vo"
)

type Transcation struct {
	ID              vo.TransaksiID `bson:"_id" json:"id"`
	TRC_Midtrans    string         `bson:"trc_midtrans" json:"trc_midtrans"`
	IDUser          vo.UserID      `bson:"id_user" json:"id_user"`
	MoneyUser       int64          `bson:"money_user" json:"money_user"`
	StatusTransaksi string         `bson:"status_transaksi" json:"status_transaksi"`
	Updated         time.Time      `bson:"updated" json:"updated"`
	Created         time.Time      `bson:"created" json:"created"`
}

type TranscationCreateRequest struct {
	RandomString        string                       `json:"-"`
	IDuser              vo.UserID                    `json:"id_user"`
	PaymentType         coreapi.CoreapiPaymentType   `json:"payment_type"`
	BankTransferDetails *coreapi.BankTransferDetails `json:"bank_transfer"`
	TransactionDetails  midtrans.TransactionDetails  `json:"transaction_details"`
	Now                 time.Time                    `json:"-"`
}
type VAResponse struct {
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

type PermataResponse struct {
	StatusCode        string `json:"status_code"`
	StatusMessage     string `json:"status_message"`
	TransactionID     string `json:"transaction_id"`
	OrderID           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	Currency          string `json:"currency"`
	PaymentType       string `json:"payment_type"`
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	PermataVaNumber   string `json:"permata_va_number"`
	FraudStatus       string `json:"fraud_status"`
}

func VABANK(res coreapi.ChargeResponse) []any {
	var VAResponse VAResponse
	VAResponse.StatusCode = res.StatusCode
	VAResponse.StatusMessage = res.StatusMessage
	VAResponse.GrossAmount = res.GrossAmount
	VAResponse.TransactionID = res.TransactionID
	VAResponse.OrderID = res.OrderID
	VAResponse.GrossAmount = res.GrossAmount
	VAResponse.Currency = res.Currency
	VAResponse.PaymentType = res.PaymentType
	VAResponse.TransactionTime = res.TransactionTime
	VAResponse.TransactionStatus = res.TransactionStatus
	VAResponse.FraudStatus = res.FraudStatus
	VAResponse.VaNumbers = res.VaNumbers

	var resultMidtrans []any

	resultMidtrans = append(resultMidtrans, VAResponse)

	return resultMidtrans
}

func PERMATA(res coreapi.ChargeResponse) []any {
	var PERMATAResponse PermataResponse
	PERMATAResponse.StatusCode = res.StatusCode
	PERMATAResponse.StatusMessage = res.StatusMessage
	PERMATAResponse.GrossAmount = res.GrossAmount
	PERMATAResponse.TransactionID = res.TransactionID
	PERMATAResponse.OrderID = res.OrderID
	PERMATAResponse.GrossAmount = res.GrossAmount
	PERMATAResponse.Currency = res.Currency
	PERMATAResponse.PaymentType = res.PaymentType
	PERMATAResponse.TransactionTime = res.TransactionTime
	PERMATAResponse.TransactionStatus = res.TransactionStatus
	PERMATAResponse.FraudStatus = res.FraudStatus
	PERMATAResponse.PermataVaNumber = res.PermataVaNumber

	var resultMidtrans []any

	resultMidtrans = append(resultMidtrans, PERMATAResponse)

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
