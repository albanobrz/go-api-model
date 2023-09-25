package services

import (
	"context"
	"math/rand"
	"strconv"
	"time"
)

type TestService interface {
	GetNumber(ctx context.Context, max string) (int, error)
}

type DefaultTestService struct {
}

func (s *DefaultTestService) GetNumber(ctx context.Context, max string) (int, error) {
	rand.Seed(time.Now().UnixNano())

	number, err := strconv.Atoi(max)
	if err != nil {
		return 0, err
	}

	randomNumber := rand.Intn(number)
	return randomNumber, nil
}
