package info

import (
	"errors"
	"fmt"
)

const contentMaxLen = 128

var (
	ErrContentAlreadyExists = errors.New("content already exists")
	ErrContentDoesntExist   = errors.New("content doesn't exist")
	ErrTooLargeContent      = fmt.Errorf("too large content (max %d)", contentMaxLen)
)

func (slf *Info) CreateContent(address, content string) error {
	slf.mu.Lock()
	defer slf.mu.Unlock()

	if len(content) > contentMaxLen {
		return ErrTooLargeContent
	}

	branch, err := slf.branchByAddress(address)
	if err != nil {
		return err
	}

	branch.Content[len(branch.Content)+1] = content

	return slf.dump()
}

func (slf *Info) UpdateContent(address string, contentNum int, newContent string) error {
	slf.mu.Lock()
	defer slf.mu.Unlock()

	branch, err := slf.branchByAddress(address)
	if err != nil {
		return err
	}

	if _, ok := branch.Content[contentNum]; !ok {
		return ErrContentDoesntExist
	}

	if len(newContent) > contentMaxLen {
		return ErrTooLargeContent
	}

	branch.Content[contentNum] = newContent

	return slf.dump()
}

func (slf *Info) DeleteContent(address string, contentNum int) error {
	slf.mu.Lock()
	defer slf.mu.Unlock()

	branch, err := slf.branchByAddress(address)
	if err != nil {
		return err
	}

	if _, ok := branch.Content[contentNum]; !ok {
		return ErrContentDoesntExist
	}

	for i := contentNum; i < len(branch.Content); i++ {
		branch.Content[i] = branch.Content[i+1]
	}

	delete(branch.Content, len(branch.Content))

	return slf.dump()
}
