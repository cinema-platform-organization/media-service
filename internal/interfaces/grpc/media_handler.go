package grpc

import (
	"bytes"
	"context"
	"io"

	pb "github.com/cinema-platform-organization/contracts/gen/go/media"
	"github.com/cinema-platform-organization/media-service/internal/application/dto"
	"github.com/cinema-platform-organization/media-service/internal/application/usecases"
)

type MediaHandler struct {
	pb.UnimplementedMediaServiceServer
	uploadUC *usecases.UploadUseCase
	getUC    *usecases.GetUseCase
	deleteUC *usecases.DeleteUseCase
}

func NewMediaHandler(
	u *usecases.UploadUseCase,
	g *usecases.GetUseCase,
	d *usecases.DeleteUseCase,
) *MediaHandler {
	return &MediaHandler{
		uploadUC: u,
		getUC:    g,
		deleteUC: d,
	}
}

func (h *MediaHandler) Upload(
	ctx context.Context,
	req *pb.UploadRequest,
) (*pb.UploadResponse, error) {
	res, err := h.uploadUC.Execute(ctx, dto.UploadMediaRequest{
		FileName:    req.FileName,
		Folder:      req.Folder,
		ContentType: req.ContentType,
		Reader:      bytes.NewReader(req.Data),
		Size:        int64(len(req.Data)),
	})

	if err != nil {
		return nil, err
	}

	return &pb.UploadResponse{
		Key: res.Key,
	}, nil
}

func (h *MediaHandler) Delete(
	ctx context.Context,
	req *pb.DeleteRequest,
) (*pb.DeleteResponse, error) {
	res, err := h.deleteUC.Execute(ctx, dto.DeleteMediaRequest{
		Key: req.Key,
	})

	if err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{
		Ok: res.OK,
	}, nil
}

func (h *MediaHandler) Get(
	ctx context.Context,
	req *pb.GetRequest,
) (*pb.GetResponse, error) {
	res, err := h.getUC.Execute(ctx, dto.GetMediaRequest{
		Key: req.Key,
	})

	if err != nil {
		return nil, err
	}
	defer res.Reader.Close()

	data, err := io.ReadAll(res.Reader)
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{
		Data:        data,
		ContentType: res.ContentType,
	}, nil
}