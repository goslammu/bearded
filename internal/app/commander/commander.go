package commander

import (
	"bearded/internal/pkg/info"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	ErrInvalidCommand = errors.New("invalid command format")
	ErrUnknownCommand = errors.New("unknown command")
	ErrExit           = errors.New("exit by command command")
)

type commander struct {
	i *info.Info
}

func New(restore bool) (*commander, error) {
	i, err := info.New(restore)
	if err != nil {
		return nil, err
	}

	return &commander{i}, nil
}

func (slf *commander) Run() {
	var err error
	scn := bufio.NewScanner(os.Stdin)

	slf.cmdShow()
	fmt.Println("ENTER COMMANDS FOR INFO EDITING...")

	for {
		fmt.Print(">>")
		scn.Scan()
		cmd := strings.Split(scn.Text(), " ")

		err = slf.routeCommand(cmd)
		if err != nil {
			if err == ErrExit {
				return
			}

			log.Println(err)
			continue
		}
	}
}

func (slf *commander) routeCommand(cmd []string) (err error) {
	if len(cmd) < 1 {
		return ErrInvalidCommand
	}

	switch cmd[0] {
	case "+b":
		return slf.cmdCreateBranch(cmd)
	case "-b":
		return slf.cmdDeleteBranch(cmd)
	case "+c":
		return slf.cmdCreateContent(cmd)
	case "-c":
		return slf.cmdDeleteContent(cmd)
	case "sh":
		return slf.cmdShow()
	case "reset":
		slf.i, err = info.New(false)
		if err != nil {
			return err
		}

		return nil
	case "exit":
		return ErrExit
	default:
		return ErrUnknownCommand
	}
}
