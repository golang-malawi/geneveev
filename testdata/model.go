package testdata

import "time"

type Person struct {
	FirstName string `validate:"required,min:2,max:255"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email,max:255"`

	DateOfBirth time.Time `validate:required`

	Registered bool `validate:"required"`

	Age   int    `validate:"required,min:18,max:69"`
	Age1  int8   `validate:"required,min:18,max:69"`
	Age2  int16  `validate:"required,min:18,max:69"`
	Age3  int32  `validate:"required,min:18,max:69"`
	Age4  int64  `validate:"required,min:18,max:69"`
	Age5  uint   `validate:"required,min:18,max:69"`
	Age6  uint8  `validate:"required,min:18,max:69"`
	Age7  uint16 `validate:"required,min:18,max:69"`
	Age8  uint32 `validate:"required,min:18,max:69"`
	Age9  uint64 `validate:"required,min:18,max:69"`
	Age10 uint8  `validate:"required,min:18,max:69"`

	Salary3 float32 `validate:"required,min:18,max:69"`
	Salary4 float64 `validate:"required,min:18,max:69"`
}
