package domain

type Cost struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Price       float64 `db:"price"`
	CreatedTime string  `db:"createdtime"`
}
type Data struct {
	Name  string  `db:"name"`
	Price float64 `db:"price"`
}
type Category struct {
	ID   int64  `db:"id,omitempty"` 
	Name string `db:"name"`
}
