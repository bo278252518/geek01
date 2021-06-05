package types

type Test struct {
	ID      uint
	Message string
}

func (Test) TableName() string {
	return "test"
}
