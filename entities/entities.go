package entities

import "time"

type User struct {
	ID    int
	NAMA   string
	LINKTUGAS string
	WAKTUKUMPUL time.Time
}