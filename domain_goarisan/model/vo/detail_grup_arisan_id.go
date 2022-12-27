package vo

import (
	"fmt"
	"time"
)

type DetailGrupArisanID string

func NewDetailGrupArisanID(randomStringID string, now time.Time) (DetailGrupArisanID, error) {
	var obj = DetailGrupArisanID(fmt.Sprintf("DGA-%s", randomStringID))
	return obj, nil
}

func (r DetailGrupArisanID) String() string {
	return string(r)
}
