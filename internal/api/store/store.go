package store

import "errors"

type Store interface {
	IntSlice() IntSliceRepository
}

type IntSliceRepository interface {
	Save(val int) error
	Search(val int) (int, error)
	Delete(val int) error
}

var (
	ErrRecordNotFound = errors.New("record not found")
)
