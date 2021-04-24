package dao

type data struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Dao interface {
	GetData(id string) (*data, error)
}
