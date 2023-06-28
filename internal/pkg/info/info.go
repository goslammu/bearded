package info

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"sync"

	"gopkg.in/yaml.v2"
)

const (
	dumpFileName     = "dump.yaml"
	addressDelimiter = "."
)

var (
	ErrInvalidformat = errors.New("invalid address format")
)

type Info struct {
	mu sync.Mutex

	root *storage
}

type storage struct {
	Name     string           `yaml:"—"`
	Content  map[int]string   `yaml:"——"`
	Branches map[int]*storage `yaml:"———"`
}

func New(restore bool) (*Info, error) {
	newInfo := &Info{
		root: &storage{
			Name:     "ROOT",
			Branches: map[int]*storage{},
			Content:  map[int]string{},
		},
	}

	if restore {
		if err := newInfo.restore(); err != nil {
			return nil, err
		}
	}

	return newInfo, nil
}

func (slf *Info) List() (string, error) {
	data, err := yaml.Marshal(slf.root)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (slf *Info) dump() error {
	data, err := yaml.Marshal(slf.root)
	if err != nil {
		return err
	}

	return os.WriteFile(dumpFileName, data, os.ModePerm)
}

func (slf *Info) restore() error {
	data, err := os.ReadFile(dumpFileName)
	if err != nil {
		return os.WriteFile(dumpFileName, nil, os.ModePerm)
	}

	return yaml.Unmarshal(data, slf.root)
}

func (slf *Info) branchByAddress(address string) (*storage, error) {
	branchNumbers, err := parseAddress(address)
	if err != nil {
		return nil, err
	}

	branch := slf.root
	var ok bool

	for _, v := range branchNumbers {
		if branch, ok = branch.Branches[v]; !ok {
			return nil, ErrBranchDoesntExist
		}
	}

	return branch, nil
}

func parseAddress(input string) (output []int, err error) {
	trimmed := strings.TrimFunc(
		input,
		func(r rune) bool {
			if r == '.' || r == '0' {
				return true
			}
			return false
		})

	if trimmed == "" {
		return
	}

	splittedAddress := strings.Split(trimmed, addressDelimiter)

	var n int

	for _, v := range splittedAddress {
		n, err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		output = append(output, n)
	}

	return
}
