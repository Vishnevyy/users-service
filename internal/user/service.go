package user

type Service struct{ repo *Repository }

func NewService(r *Repository) *Service { return &Service{repo: r} }

func (s *Service) Create(email string) (*User, error) {
	u := &User{Email: email}
	return u, s.repo.Create(u)
}
func (s *Service) Get(id uint32) (*User, error) { return s.repo.GetByID(id) }
func (s *Service) List() ([]User, error)        { return s.repo.List() }
func (s *Service) Update(id uint32, email string) (*User, error) {
	u := &User{ID: id, Email: email}
	return u, s.repo.Update(u)
}
func (s *Service) Delete(id uint32) error { return s.repo.Delete(id) }
