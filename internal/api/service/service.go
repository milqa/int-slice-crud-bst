package service

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/milQA/int-slice-crud-bst/internal/api/store"
	"go.uber.org/zap"
)

type (
	Service struct {
		intSlice *IntSliceService
		store    store.Store
		logger   *zap.Logger
	}

	IntSliceService struct {
		service *Service
	}
)

func NewService(log *zap.Logger, store store.Store, opts ...Option) (*Service, error) {
	logger := log.With(
		zap.String(
			"module", "service",
		),
	)

	service := &Service{
		logger: logger,
		store:  store,
	}

	for _, opt := range opts {
		if err := opt(service); err != nil {
			return nil, fmt.Errorf("cannot update service with options: %w", err)
		}
	}

	return service, nil
}

type Option func(*Service) error

func WithUpdateIntSliceByFile(filePath string) Option {
	return func(service *Service) error {
		if err := service.IntSlice().UpdateByFile(filePath); err != nil {
			return fmt.Errorf("cannot upload int slice by file: %w", err)
		}
		return nil
	}
}

func (s *Service) IntSlice() *IntSliceService {
	if s.intSlice != nil {
		return s.intSlice
	}

	s.intSlice = &IntSliceService{
		service: s,
	}

	return s.intSlice
}

func (s *IntSliceService) UpdateByFile(filePath string) error {

	vals := make([]int, 0)

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("cannot open file: %w", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&vals); err != nil {
		return fmt.Errorf("cannot unmarshal file: %w", err)
	}

	for _, val := range vals {
		if err := s.service.store.IntSlice().Save(val); err != nil {
			return fmt.Errorf("cannot save val = %d: %w", val, err)
		}
	}

	return nil
}

func (s *IntSliceService) Insert(val int) error {

	if err := s.service.store.IntSlice().Save(val); err != nil {
		return fmt.Errorf("cannot insert val = %v: %w", val, err)
	}

	return nil
}

func (s *IntSliceService) Delete(val int) error {

	if err := s.service.store.IntSlice().Delete(val); err != nil {
		return fmt.Errorf("cannot delete val = %v: %w", val, err)
	}

	return nil
}

func (s *IntSliceService) Search(val int) (int, error) {

	answer, err := s.service.store.IntSlice().Search(val)
	if err != nil {
		return 0, fmt.Errorf("cannot search val = %v: %w", val, err)
	}

	return answer, nil
}
