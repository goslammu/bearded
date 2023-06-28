package info

import (
	"errors"
)

var (
	ErrBranchDoesntExist = errors.New("branch doesn't exist")
)

func (slf *Info) CreateBranch(address, branchName string) error {
	slf.mu.Lock()
	defer slf.mu.Unlock()

	branch, err := slf.branchByAddress(address)
	if err != nil {
		return err
	}

	newInfo := &storage{
		Name:     branchName,
		Branches: map[int]*storage{},
		Content:  map[int]string{},
	}

	branch.Branches[len(branch.Branches)+1] = newInfo

	return slf.dump()
}

func (slf *Info) DeleteBranch(address string, branchNum int) error {
	slf.mu.Lock()
	defer slf.mu.Unlock()

	branch, err := slf.branchByAddress(address)
	if err != nil {
		return err
	}

	if _, ok := branch.Branches[branchNum]; !ok {
		return ErrBranchDoesntExist
	}

	for i := branchNum; i < len(branch.Branches); i++ {
		branch.Branches[i] = branch.Branches[i+1]
	}

	delete(branch.Branches, len(branch.Branches))

	return slf.dump()
}
