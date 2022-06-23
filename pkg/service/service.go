package service

type Service struct {
}

func (s *Service) GetAccessList() []string {
	return []string{"ModuloA", "ModuloB"}
}
