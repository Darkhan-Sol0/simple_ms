package service

type (
	Storage interface {
		AddUser() (string, error)
		GetListUser() (map[string]string, error)
	}

	ServiceYmpl struct {
		Storage Storage
	}

	Service interface {
		Hello() string
		GetList() map[string]string
	}
)

func NewService(storage Storage) Service {
	return &ServiceYmpl{
		Storage: storage,
	}
}

func (s *ServiceYmpl) Hello() string {
	a, _ := s.Storage.AddUser()
	return a
}

func (s *ServiceYmpl) GetList() map[string]string {
	a, _ := s.Storage.GetListUser()
	return a
}
