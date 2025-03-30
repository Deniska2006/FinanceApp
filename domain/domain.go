package domain

type Cost struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Price       int    `db:"price"`
	CreatedTime string `db:"createdtime"`
}
