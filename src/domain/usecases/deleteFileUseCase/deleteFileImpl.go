package deleteFileUseCase

import "os"

type DeleteFileUseCase struct{}

func NewDeleteFileUseCase() *DeleteFileUseCase {
	return &DeleteFileUseCase{}
}

func (d *DeleteFileUseCase) Run(input Input) {
	os.Remove(input.Path)
}
