package bst

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrAddressNotFound = errors.New("Address not found")
)

type Note struct {
	address     *int
	value       interface{}
	parent      *Note
	rightBranch *Note
	leftBranch  *Note
}

type BST struct {
	mu     sync.Mutex
	Logger *PrintfLogger
	root   *Note
}

func NewBST(printf Printf) *BST {

	bst := BST{
		root: &Note{
			address:     nil,
			value:       nil,
			rightBranch: nil,
			leftBranch:  nil,
		},
		Logger: NewPrintfLogger(printf),
	}

	bst.root.parent = &Note{
		rightBranch: bst.root,
	}

	return &bst
}

func (n *Note) findValue(address int) (interface{}, error) {
	if n.address == nil {
		return nil, ErrAddressNotFound
	}

	if *n.address == address {
		return n.value, nil
	}

	if *n.address < address {
		return n.rightBranch.findValue(address)
	}

	if *n.address > address {
		return n.leftBranch.findValue(address)
	}
	return nil, nil
}

func (c *BST) Find(address int) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Logger.beforeFind(address)
	start := time.Now()
	value, err := c.root.findValue(address)
	c.Logger.afterFind(address, value, time.Since(start), err)
	return value, err
}

func (n *Note) insertValue(address int, value interface{}) error {

	if n.address == nil {
		n.address = &address
		n.value = value
		n.rightBranch = &Note{
			address:     nil,
			value:       nil,
			rightBranch: nil,
			leftBranch:  nil,
			parent:      n,
		}
		n.leftBranch = &Note{
			address:     nil,
			value:       nil,
			rightBranch: nil,
			leftBranch:  nil,
			parent:      n,
		}

		return nil
	}

	if *n.address == address {
		n.value = value
		return nil
	}

	if *n.address < address {
		return n.rightBranch.insertValue(address, value)
	}

	if *n.address > address {
		return n.leftBranch.insertValue(address, value)
	}
	return nil
}

func (c *BST) Insert(address int, value interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Logger.beforeInsert(address, value)
	start := time.Now()
	err := c.root.insertValue(address, value)
	c.Logger.afterInsert(address, value, time.Since(start), err)
	return err
}

func (c *BST) Delete(address int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Logger.beforeDelete(address)
	start := time.Now()
	err := c.root.deleteValue(address)
	c.Logger.afterDelete(address, time.Since(start), err)
	return err
}

func (n *Note) deleteValue(address int) error {
	if n.address == nil {
		return ErrAddressNotFound
	}

	if *n.address < address {
		if n.rightBranch == nil {
			return ErrAddressNotFound
		}
		return n.rightBranch.deleteValue(address)
	}

	if *n.address > address {
		if n.leftBranch == nil {
			return ErrAddressNotFound
		}
		return n.leftBranch.deleteValue(address)
	}

	if *n.address == address {
		if (n.leftBranch == nil || n.leftBranch.address == nil) &&
			(n.rightBranch == nil || n.rightBranch.address == nil) {
			n.address = nil
			n.value = nil
			n.leftBranch = &Note{
				address:     nil,
				value:       nil,
				rightBranch: nil,
				leftBranch:  nil,
				parent:      n,
			}
			n.rightBranch = &Note{
				address:     nil,
				value:       nil,
				rightBranch: nil,
				leftBranch:  nil,
				parent:      n,
			}
			return nil
		}

		if (n.leftBranch != nil && n.leftBranch.address != nil) &&
			(n.rightBranch == nil || n.rightBranch.address == nil) {
			*n = *n.leftBranch
			return nil
		}

		if (n.leftBranch == nil || n.leftBranch.address == nil) &&
			(n.rightBranch != nil && n.rightBranch.address != nil) {
			*n = *n.rightBranch
			return nil
		}

		if (n.leftBranch != nil && n.leftBranch.address != nil) &&
			(n.rightBranch != nil && n.rightBranch.address != nil) {

			if n.rightBranch.leftBranch == nil || n.rightBranch.leftBranch.address == nil {
				n.address = n.rightBranch.address
				n.value = n.rightBranch.value
				n.rightBranch = n.rightBranch.rightBranch
				return nil
			} else {
				min := n.rightBranch.findMin()
				n.address = min.address
				n.value = min.value
				return n.rightBranch.deleteValue(*min.address)
			}
		}
	}

	return nil
}

func (n *Note) findMin() *Note {
	if n.leftBranch == nil || n.leftBranch.address == nil {
		if n.rightBranch == nil || n.rightBranch.address == nil {
			return n
		}
		return n.rightBranch.findMin()
	}

	return n.leftBranch.findMin()
}
