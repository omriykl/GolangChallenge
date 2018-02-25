package structs

import (
	"time"
)

type User struct {
	Id int
	FirstName string
	LastName string
	CreatedDate time.Time
}