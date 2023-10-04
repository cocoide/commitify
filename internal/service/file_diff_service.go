package service

import "os/exec"

type fileDiffService struct {
}

func NewFileDiffService() fileDiffService {
	fds := fileDiffService{}
	return fds
}

func (fds *fileDiffService) createFileDiffStr() (string, error) {
	diff, err := exec.Command("git", "diff", "--staged").Output()

	return string(diff), err
}
