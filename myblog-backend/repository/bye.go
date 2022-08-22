package repository

import "context"

type IByeRepository interface {
	Bye(ctx context.Context) map[string]interface{}
}

type ByeRepository struct{}

func (h *ByeRepository) Bye(ctx context.Context) map[string]interface{} {
	return map[string]interface{}{
		"message": "bye",
	}
}
