package datamodels

// swagger:model User 
type Row struct {
	// ID of User
	// in: int64
	ID int `json:"ID"`
	// Name of User
	// in: string
	Name string `json:"Name"`
	// Email of User
	// in: string
	Email string `json:"Email"`
}

type Operations interface {
	CreateRow(Row) error
	GetAll() ([]Row, error)
	UpdateRow(int, Row) error
	DeleteRow(int) error
}