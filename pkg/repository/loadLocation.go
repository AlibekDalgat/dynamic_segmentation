package repository

import "time"

func loadLocation(timeZone string) (*time.Location, error) {
	loc, err := time.LoadLocation("Europe/Moscow")
	return loc, err
}
