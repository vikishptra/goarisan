package vo

import (
	"time"
)

type UserID string

func NewUserID(randomStringID string, now time.Time) (UserID, error) {
	var obj = UserID(randomStringID)
	return obj, nil
}

func (r UserID) String() string {
	return string(r)
}
