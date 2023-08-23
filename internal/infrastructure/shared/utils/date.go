package utils

import "time"

func Date(rawTimezone string) (*time.Time, error) {
	var timezone string

	if rawTimezone == "" {
		timezone = "America/Santo_Domingo"
	} else {
		timezone = rawTimezone
	}

	loc, err := time.LoadLocation(timezone)

	if err != nil {
		return nil, err
	}

	currentTime := time.Now().In(loc)

	return &currentTime, nil
}
