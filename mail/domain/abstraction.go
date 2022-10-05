package domain

type Service interface {
	InsertData(domain Mail) (mailObj Mail, err error)
	UpdateData(id string, domain Mail) (mailObj Mail, err error)
	GetAllData() (mailObj []Mail, err error)
	GetByID(id string) (mailObj Mail, err error)
	// GetBymailname(mailname string) (mailObj Mail, err error)
	DeleteData(id string) (err error)
}

type Repository interface {
	Create(domain Mail) (mailObj Mail, err error)
	Update(domain Mail) (mailObj Mail, err error)
	Get() (mailObj []Mail, err error)
	GetByID(id string) (domain Mail, err error)
	// GetBymailname(mailname string) (mailObj Mail, err error)
	Delete(id string) (err error)
}
