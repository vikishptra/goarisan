package vo

import (
	"fmt"
	"time"
)

type TransaksiID string

func NewTransaksiID(randomStringID string, now time.Time) (TransaksiID, error) {
	var obj = TransaksiID(fmt.Sprintf("TRC-%s", randomStringID))
	return obj, nil
}

func (r TransaksiID) String() string {
	return string(r)
}
