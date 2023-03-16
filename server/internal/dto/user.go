package dto

type UserRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Interests string `json:"interests"`
	City      string `json:"city"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Interests string `json:"interests"`
	City      string `json:"city"`
	Email     string `json:"email"`
	Token     string `json:"token,omitempty"`
}

type SecurityUser struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type FriendRequest struct {
	ID int `json:"friendID"`
}
