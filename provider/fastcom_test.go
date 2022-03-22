package provider

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_fastCom_MeasureDownloadSpeed(t *testing.T) {
	fastcom, err := newFastCom()
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	speed, err := fastcom.MeasureDownloadSpeed(ctx)
	require.NoError(t, err)
	require.NotZero(t, speed)
}

func Test_fastCom_MeasureUploadSpeed(t *testing.T) {
	fastcom, err := newFastCom()
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	speed, err := fastcom.MeasureUploadSpeed(ctx)
	require.NoError(t, err)
	require.NotZero(t, speed)
}

func Benchmark_fastCom_MeasureDownloadSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		provider, err := newFastCom()
		if err != nil {
			b.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		_, err = provider.MeasureDownloadSpeed(ctx)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_fastCom_MeasureUploadSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		provider, err := newFastCom()
		if err != nil {
			b.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		_, err = provider.MeasureUploadSpeed(ctx)
		if err != nil {
			b.Fatal(err)
		}
	}
}
