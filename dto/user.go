package dto

type UserRequestform struct {
	Name      string  `json:"name" binding:"required"`
	Email     *string `json:"email" binding:"required"`
	Age       uint8   `json:"age" binding:"required"`
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Username  string  `json:"username" binding:"required"`
	Password  string  `json:"password" binding:"required"`
}
