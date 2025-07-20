package domain

type User struct {
	ID        int64
	FirstName string
	LastName  string
	UserName  string
}

type Character struct {
	ID     int64
	Name   string
	UserID int64
}
