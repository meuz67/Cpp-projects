package service

import "time"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
func (s *Service) Daysuntil() int64 {
	targetDate := time.Date(2027, time.January, 1, 0, 0, 0, 0, time.UTC)
	daysUntil := time.Until(targetDate)
	daysLeft := int64(daysUntil.Hours() / 24)
	return daysLeft
}
