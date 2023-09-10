package service

import "os/exec"

type fileDiffService struct {
}

func NewFileDiffService() *fileDiffService {
	ps := new(fileDiffService)
	return ps
}

func (ps fileDiffService) CreateFileDiffStr() (string, error) {
	diff, err := exec.Command("git", "diff", "--staged").Output()

	return string(diff), err
}
