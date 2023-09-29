package service

type TickService struct {
	ratesRepository RatesRepository
}

func NewTickService(ratesRepository RatesRepository) *TickService {
	return &TickService{
		ratesRepository: ratesRepository,
	}
}

func (s *TickService) GetTickById(id int64) (Tick, error) {
	tick, err := s.ratesRepository.GetTickById(id)
	if err != nil {
		return Tick{}, err
	}
	return tick, nil
}
