package service

import (
	"fmt"

	"github.com/milQA/int-slice-crud-bst/internal/api/store"
)

type IntSliceService struct {
	store store.Store
}

func NewIntSliceService(store store.Store) *IntSliceService {

	return &IntSliceService{
		store: store,
	}
}

func (s *IntSliceService) Insert(vals ...int) error {

	for i, val := range vals {
		if err := s.store.IntSlice().Save(val); err != nil {
			return fmt.Errorf("cannot insert vals = %v: %w", vals[i:], err)
		}
	}

	return nil
}

func (s *IntSliceService) Delete(val int) error {

	if err := s.store.IntSlice().Delete(val); err != nil {
		return fmt.Errorf("cannot delete val = %v: %w", val, err)
	}

	return nil
}

func (s *IntSliceService) Search(val int) (int, error) {

	answer, err := s.store.IntSlice().Search(val)
	if err != nil {
		return 0, fmt.Errorf("cannot search val = %v: %w", val, err)
	}

	return answer, nil
}
