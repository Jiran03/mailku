package domain

type Service interface {
	InsertData(domain Agency) (agencyObj Agency, err error)
	UpdateData(id string, domain Agency) (agencyObj Agency, err error)
	GetAllData() (agencyObj []Agency, err error)
	GetByID(id string) (agencyObj Agency, err error)
	DeleteData(id string) (err error)
}

type Repository interface {
	Create(domain Agency) (agencyObj Agency, err error)
	Update(domain Agency) (agencyObj Agency, err error)
	Get() (agencyObj []Agency, err error)
	GetByID(id string) (domain Agency, err error)
	Delete(id string) (err error)
}
