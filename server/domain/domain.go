package domain

// User доменная сущность
type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
	Gender    string
	Interests string
	City      string
	Email     string
	Password  string
}

type Post struct {
	ID      int
	Content string
	UserID  int
}
