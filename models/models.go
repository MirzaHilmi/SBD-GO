package models

type Tabler interface {
	TableName() string
}

func (Address) TableName() string {
	return "address"
}
