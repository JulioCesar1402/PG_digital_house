package users

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var listOfUsers []User
var lastId int

type Repository interface {
	GetAll() ([]User, err)
	CreateUser(id int, name, email, password string) (User, err)
	LastId() (int, err)
}

func GetAll() ([]User, err) {
	return listOfUsers, nil
}

func CreateUser(id int, name, email, password string) (User, err) {
	user := User{id, name, email, password}
	listOfUsers = append(listOfUsers, user)
	lastId = user.Id
	return user, nil
}

func LastId() (int, err) {
	return lastId, nil
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
