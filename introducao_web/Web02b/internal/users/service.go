package users

type Service interface  {
	GetAll() ([]User, err)
	CreateUser(name, email, password string) (User, err)
}

