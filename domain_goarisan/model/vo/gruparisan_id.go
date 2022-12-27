package vo

import (
	"fmt"
	"time"
)

type GruparisanID string

func NewGruparisanID(randomStringID string, now time.Time) (GruparisanID, error) {
	var obj = GruparisanID(fmt.Sprintf("GRP-%s", randomStringID))
	return obj, nil
}

func (r GruparisanID) String() string {
	return string(r)
}
