package users

type User struct {
	Id          int64  `Id`
	FirstName   string `first_name`
	LastName    string `last_name`
	Email       string `email`
	DateCreated string `date_created`
}
