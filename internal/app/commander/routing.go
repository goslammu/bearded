package commander

import (
	"fmt"
	"strconv"
	"strings"
)

func (slf *commander) cmdCreateBranch(cmd []string) error {
	if len(cmd) < 3 {
		return ErrInvalidCommand
	}

	branchName := ""
	for i := 2; i < len(cmd); i++ {
		branchName += cmd[i] + " "
	}

	return slf.i.CreateBranch(cmd[1], strings.TrimSpace(branchName))
}

func (slf *commander) cmdDeleteBranch(cmd []string) error {
	if len(cmd) < 3 {
		return ErrInvalidCommand
	}

	n, err := strconv.Atoi(cmd[2])
	if err != nil {
		return err
	}

	return slf.i.DeleteBranch(cmd[1], n)
}

func (slf *commander) cmdCreateContent(cmd []string) error {
	if len(cmd) < 3 {
		return ErrInvalidCommand
	}

	content := ""
	for i := 2; i < len(cmd); i++ {
		content += cmd[i] + " "
	}

	return slf.i.CreateContent(cmd[1], strings.TrimSpace(content))
}

func (slf *commander) cmdDeleteContent(cmd []string) error {
	if len(cmd) < 3 {
		return ErrInvalidCommand
	}

	n, err := strconv.Atoi(cmd[2])
	if err != nil {
		return err
	}

	return slf.i.DeleteContent(cmd[1], n)
}

func (slf *commander) cmdShow() error {
	sh, err := slf.i.List()
	if err != nil {
		return err
	}

	fmt.Print("CURRENT INFO STATE\n\n", sh, "\n")
	return nil
}
