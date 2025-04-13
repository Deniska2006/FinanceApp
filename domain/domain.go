package domain

type Cost struct {
	ID          int     `db:"id"`
	Uid 		int  	`db:"uid"`
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
	Uid int64   `db:"uid"`
}

type Account struct{
	Id int64 `db:"id,omitempty"`
	Mail string `db:"mail,omitempty"`
	Password string `db:"password,omitempty"`
	HashedPassword string `db:"hashed_password"` 
}
