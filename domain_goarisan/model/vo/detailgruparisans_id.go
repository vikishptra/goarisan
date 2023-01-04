package vo

import (
	"fmt"
	"time"
)

type DetailgruparisansID string

func NewDetailgruparisansID(randomStringID string, now time.Time) (DetailgruparisansID, error) {
	var obj = DetailgruparisansID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))
	return obj, nil
}

func (r DetailgruparisansID) String() string {
	return string(r)
}
