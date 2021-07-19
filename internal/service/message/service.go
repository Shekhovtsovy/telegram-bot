package message

// Service is an interface which provides methods for message service work
type Service interface {
	SaveMessage() error
}

type repository interface {
	InsertMessage() error
}

type service struct {
	repository repository
}

// Save message
func (s *service) SaveMessage() error {
	if err := s.repository.InsertMessage(); err != nil {
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
