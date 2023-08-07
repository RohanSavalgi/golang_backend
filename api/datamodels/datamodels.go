package datamodels

type Row struct {
	ID int `json:"ID"`
	Name string `json:"Name"`
	Email string `json:"Email"`
}

type Operations interface {
	CreateRow(Row) error
	GetAll() ([]Row, error)
	UpdateRow(int, Row) error
	DeleteRow(int) error
}