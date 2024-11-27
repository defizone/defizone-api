package model

import "time"

type License struct {
	Key         string
	Token       string
	OwnerId     string
	DeviceId    string
	ExpiredDate time.Time
	Duration    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
