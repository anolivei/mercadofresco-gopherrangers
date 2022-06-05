package products

type Service interface {
	Store(prod Product) (Product, error)
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	UpdatePut(prod Product, id int) (Product, error)
	//UpdatePatch(prod Product) (Product, error)
	//Delete(id int) (error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Store(prod Product) (Product, error) {
	lastId, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	lastId++
	product, err := s.repository.Store(prod, lastId)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetById(id int) (Product, error) {
	ps, err := s.repository.GetById(id)
	if err != nil {
		return Product{}, err
	}
	return ps, nil
}

func (s *service) UpdatePut(prod Product, id int) (Product, error) {
	product, err := s.repository.UpdatePut(prod, id)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

/*
func (s *service) UpdatePatch(prod Product) (Product, error) {

}

func (s *service) Delete(id int) (error) {

}
*/
