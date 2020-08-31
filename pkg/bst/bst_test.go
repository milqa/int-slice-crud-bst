package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	branchInsertTests = []struct {
		in1 int
		in2 interface{}
		err error
	}{
		{1, "test1", nil},
		{3, "test3", nil},
		{15, "test15", nil},
		{10, "test10", nil},
		{8, "test8", nil},
		{9, "test9", nil},
		{10, "test10new", nil},
		{-1, "test-1", nil},
		{99, "test99", nil},
		{1000, "test1000", nil},
		{200, "test200", nil},
		{1200, "test1200", nil},
		{150, "test150", nil},
		{300, "300", nil},
		{250, "250", nil},
		{400, "400", nil},
		{1500, "1500", nil},
		{1600, "1600", nil},
		{1700, "1700", nil},
		{1450, "1450", nil},
		{1750, "1750", nil},
		{1350, "1350", nil},
		{1850, "1850", nil},
		{1950, "1950", nil},
		{1959, "1959", nil},
	}

	branchFindTests = []struct {
		in  int
		out interface{}
		err error
	}{
		{9, "test9", nil},
		{29, nil, ErrAddressNotFound},
		{10, "test10new", nil},
		{8, "test8", nil},
		{1, "test1", nil},
		{33, nil, ErrAddressNotFound},
		{3, "test3", nil},
		{15, "test15", nil},
		{-1, "test-1", nil},
	}

	branchDeleteTests = []struct {
		in  int
		err error
	}{
		{1750, nil},
		{1500, nil},
		{9, nil},
		{29, ErrAddressNotFound},
		{1, nil},
		{33, ErrAddressNotFound},
		{-1, nil},
		{200, nil},
		{250, nil},
		{300, nil},
		{400, nil},
		{150, nil},
		{1000, nil},
		{1200, nil},
		{99, nil},
	}

	branchFindAfterDeleteTests = []struct {
		in  int
		err error
	}{
		{9, ErrAddressNotFound},
		{29, ErrAddressNotFound},
		{10, nil},
		{8, nil},
		{1, ErrAddressNotFound},
		{33, ErrAddressNotFound},
		{3, nil},
		{-1, ErrAddressNotFound},
		{200, ErrAddressNotFound},
		{250, ErrAddressNotFound},
		{300, ErrAddressNotFound},
		{400, ErrAddressNotFound},
		{150, ErrAddressNotFound},
		{1200, ErrAddressNotFound},
		{99, ErrAddressNotFound},
		{1000, ErrAddressNotFound},
	}
)

func TestNew(t *testing.T) {
	got := NewBST()
	assert.NotNil(t, got)
}

func TestInsert(t *testing.T) {
	got := NewBST()
	for _, tt := range branchInsertTests {
		err := got.Insert(tt.in1, tt.in2)
		assert.NoError(t, err)
	}
}

func TestFind(t *testing.T) {
	got := NewBST()
	for _, tt := range branchInsertTests {
		err := got.Insert(tt.in1, tt.in2)
		assert.NoError(t, err)
	}

	for _, tt := range branchFindTests {
		b, e := got.Find(tt.in)
		assert.Equal(t, tt.out, b)
		assert.Equal(t, tt.err, e)
	}
}

func TestDelete(t *testing.T) {
	got := NewBST()
	for _, tt := range branchInsertTests {
		err := got.Insert(tt.in1, tt.in2)
		assert.NoError(t, err)
	}

	for _, tt := range branchDeleteTests {
		err := got.Delete(tt.in)
		assert.Equal(t, tt.err, err)
	}

	for _, tt := range branchFindAfterDeleteTests {
		_, err := got.Find(tt.in)
		assert.Equal(t, tt.err, err, "input = %v", tt.in)
	}
}
