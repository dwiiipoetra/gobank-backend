package interfaces

type User struct {
	// gorm.Model
	ID       uint   `gorm:"primary key;autoIncrement" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Account struct {
	// gorm.Model
	ID      uint   `gorm:"primary key;autoIncrement" json:"id"`
	Type    string `json:"account_type"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
	UserID  uint   `json:"user_id"`
}

type ResponseAccount struct {
	ID      uint
	Name    string
	Balance int
}

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}

type Validation struct {
	Value string
	Valid string
}
