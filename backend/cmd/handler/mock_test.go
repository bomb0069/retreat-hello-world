package handler

import "say/internal/platform/repository"

type MockRepository struct {
}

func (mock MockRepository) GetFirstMessage() repository.Say {
	return repository.Say{Message: "Hello World"}
}
