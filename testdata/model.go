package testdata

import "time"

type UnknownType string

type Person struct {
	IgnoredFirstName string `json:""`

	FirstName string `json:"" validate:"required,min=2,max=255"`
	LastName  string `json:"" validate:"required"`
	Email     string `json:"" validate:"required,email,max=255"`

	DateOfBirth time.Time `json:"" validate:required`

	SecretIdentity UnknownType `json:"" validate:required` // this still passes despite tag being malformed because of missing quotes

	Registered bool `json:"" validate:"required"`

	Age   int    `json:"" validate:"required,min=18,max=69"`
	Age1  int8   `json:"" validate:"required,min=18,max=69"`
	Age2  int16  `json:"" validate:"required,min=18,max=69"`
	Age3  int32  `json:"" validate:"required,min=18,max=69"`
	Age4  int64  `json:"" validate:"required,min=18,max=69"`
	Age5  uint   `json:"" validate:"required,min=18,max=69"`
	Age6  uint8  `json:"" validate:"required,min=18,max=69"`
	Age7  uint16 `json:"" validate:"required,min=18,max=69"`
	Age8  uint32 `json:"" validate:"required,min=18,max=69"`
	Age9  uint64 `json:"" validate:"required,min=18,max=69"`
	Age10 uint8  `json:"" validate:"required,min=18,max=69"`

	Salary3 float32 `json:"" validate:"required,min=18,max=69"`
	Salary4 float64 `json:"" validate:"required,min=18,max=69"`
}
