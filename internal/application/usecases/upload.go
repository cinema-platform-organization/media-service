package usecases

import (
	"context"
	"fmt"

	"github.com/cinema-platform-organization/media-service/internal/application/dto"
	"github.com/cinema-platform-organization/media-service/internal/infrastructure/images"
	"github.com/cinema-platform-organization/media-service/internal/infrastructure/storage"
)

type UploadUseCase struct {
	storage   storage.Storage
	processor images.Processor
}

func NewUploadUseCase(s storage.Storage, p images.Processor) *UploadUseCase {
	return &UploadUseCase{
		storage:   s,
		processor: p,
	}
}

func (u *UploadUseCase) Execute(
	ctx context.Context,
	input dto.UploadMediaRequest,
) (*dto.UploadMediaResponse, error) {
	key := fmt.Sprintf("%s/%s", input.Folder, input.FileName)

	if err := u.storage.UploadStream(
		ctx,
		key,
		input.Reader,
		input.ContentType,
	); err != nil {
		return nil, err
	}

	return &dto.UploadMediaResponse{
		Key: key,
	}, nil
}
