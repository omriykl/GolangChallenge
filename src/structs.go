//this file will include all the data structures

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