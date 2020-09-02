package bstStore

import (
	"github.com/milQA/int-slice-crud-bst/internal/api/store"
	"github.com/milQA/int-slice-crud-bst/pkg/bst"
	"go.uber.org/zap"
)

type Store struct {
	logger   *zap.Logger
	intSlice *IntSliceRepository
}

func NewStore(log *zap.Logger) store.Store {
	logger := log.With(
		zap.String(
			"module", "bst_store",
		),
	)

	return &Store{
		logger: logger,
	}
}

func (s *Store) IntSlice() store.IntSliceRepository {
	if s.intSlice != nil {
		return s.intSlice
	}

	s.intSlice = &IntSliceRepository{
		store: s,
		bst:   bst.NewBST(s.logger.Sugar().Infof),
	}

	return s.intSlice
}
