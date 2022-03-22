package speedtest

import (
	"context"
	"log"
	"time"

	"github.com/friendsofgo/errors"
)

//go:generate mockery --name=Provider --case=underscore

// Provider interface of speed test provider
type Provider interface {
	GetName() string
	MeasureDownloadSpeed(ctx context.Context) (float64, error)
	MeasureUploadSpeed(ctx context.Context) (float64, error)
}

// MeasureSpeed measures speed with passed provider and while passed timeout not expires
func MeasureSpeed(ctx context.Context, provider Provider, timeout time.Duration) (float64, float64, error) {
	downloadCtx, downloadCancel := context.WithTimeout(ctx, timeout)
	defer downloadCancel()
	log.Printf("Start measuring with %s provider", provider.GetName())
	download, err := provider.MeasureDownloadSpeed(downloadCtx)
	if err != nil {
		return 0, 0, errors.Wrap(err, "download measure failed")
	}
	uploadCtx, uploadCancel := context.WithTimeout(ctx, timeout)
	defer uploadCancel()
	upload, err := provider.MeasureUploadSpeed(uploadCtx)
	if err != nil {
		return 0, 0, errors.Wrap(err, "upload measure failed")
	}
	return download, upload, nil
}
