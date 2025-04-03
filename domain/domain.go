package domain

type Cost struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Price       float64 `db:"price"`
	CreatedTime string  `db:"createdtime"`
}
