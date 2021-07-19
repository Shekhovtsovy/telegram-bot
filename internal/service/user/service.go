package message

// Service is an interface which provides methods for user service work
type Service interface {
	SaveUser() error
}

type repository interface {
	InsertUser() error
}

type service struct {
	repository repository
}

// Save user
func (s *service) SaveUser() error {
	if err := s.repository.InsertUser(); err != nil {
		return err
	}
	return nil
}

// NewService return a new message service
func NewService(r repository) Service {
	return &service{
		repository: r,
	}
}
