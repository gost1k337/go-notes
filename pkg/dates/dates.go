package dates

import (
	"fmt"
	"strings"
	"time"
)

func BytesToTimeSlice(remindsSlice []byte) ([]time.Time, error) {
	timeString := string(remindsSlice)
	timeString = strings.Trim(timeString, "{}")
	timeStrings := strings.Split(timeString, ",")

	var reminds []time.Time

	for _, str := range timeStrings {
		t, err := time.Parse("2006-01-02 15:04:05-07", strings.Trim(str, `"`))
		if err != nil {
			return nil, fmt.Errorf("parse: %w", err)
		} else {
			reminds = append(reminds, t)
		}
	}
	return reminds, nil
}
