package bstStore

import (
	"errors"
	"fmt"

	"github.com/milQA/int-slice-crud-bst/internal/api/store"
	"github.com/milQA/int-slice-crud-bst/pkg/bst"
)

type IntSliceRepository struct {
	store *Store
	bst   *bst.BST
}

func (r *IntSliceRepository) Save(val int) error {
	if err := r.bst.Insert(val, nil); err != nil {
		return fmt.Errorf("cannot insert val = %v: %w", val, err)
	}
	return nil
}

func (r *IntSliceRepository) Delete(val int) error {
	if err := r.bst.Delete(val); err != nil {
		if errors.Is(err, bst.ErrAddressNotFound) {
			return store.ErrRecordNotFound
		}
		return fmt.Errorf("cannot delete val = %v: %w", val, err)
	}
	return nil
}

func (r *IntSliceRepository) Search(val int) (int, error) {
	_, err := r.bst.Find(val)
	if err != nil {
		if errors.Is(err, bst.ErrAddressNotFound) {
			return 0, store.ErrRecordNotFound
		}
		return 0, fmt.Errorf("cannot find val = %v: %w", val, err)
	}

	return 0, nil
}
