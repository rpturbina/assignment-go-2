package order

type Order struct {
	ID     uint64 `json:"id" gorm:"column:id;primaryKey"`
	Item   string `json:"item" gorm:"column:item"`
	UserID uint64 `json:"user_id"`
}

type User struct {
	ID        uint64  `json:"id" gorm:"column:id;primaryKey"`
	FirstName string  `json:"first_name" gorm:"column:first_name"`
	LastName  string  `json:"last_name" gorm:"column:last_name"`
	Email     string  `json:"email" gorm:"column:email"`
	Orders    []Order `json:"orders"`
}
