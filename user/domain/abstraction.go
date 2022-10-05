package domain

type Service interface {
	CreateToken(username, password string) (token string, userObj User, err error)
	InsertData(domain User) (userObj User, err error)
	UpdateData(id string, domain User) (userObj User, err error)
	GetAllData() (userObj []User, err error)
	GetByID(id string) (userObj User, err error)
	GetByUsername(username string) (userObj User, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Create(domain User) (userObj User, err error)
	Update(domain User) (userObj User, err error)
	Get() (userObj []User, err error)
	GetByID(id string) (domain User, err error)
	GetByUsername(username string) (userObj User, err error)
	Delete(id int) (err error)
}
