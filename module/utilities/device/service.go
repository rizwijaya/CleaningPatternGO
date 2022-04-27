package device

type Service interface {
	GetAllDevice() ([]Device, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllDevice() ([]Device, error) {
	device, err := s.repository.GetAllDevice()
	if err != nil {
		return device, err
	}

	return device, nil
}
