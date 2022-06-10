package helpers

import (
	"encoding/json"
	"log"
	"time"
)

// CustomDateFormat supports yyyy-mm-dd time format
type CustomDateFormat time.Time

var _ json.Unmarshaler = &CustomDateFormat{}

// UnmarshalJSON converts a string to yyyy-mm-dd time format
func (cdf *CustomDateFormat) UnmarshalJSON(bs []byte) error {
	var s string
	log.Print(bs)
	if err := json.Unmarshal(bs, &s); err != nil {
		log.Fatal(err.Error())
	}

	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*cdf = CustomDateFormat(t)
	return nil
}
