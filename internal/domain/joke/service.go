package joke

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s *Service) GetRandom() (Model, error) {
	return s.repo.Random()
}

func (s *Service) Query(query string) (ModelQuery, error) {
	return s.repo.Query(query)
}
