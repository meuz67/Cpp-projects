package service

import "time"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
func (s *Service) DaysLeft() int64 {
	targetDate := time.Date(2027, time.January, 1, 0, 0, 0, 0, time.UTC)
	daysLeft := time.Until(targetDate)
	return int64(daysLeft.Hours() / 24)
}
