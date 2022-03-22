package provider

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_speedTestNet_MeasureDownloadSpeed(t *testing.T) {
	speedtestnet, err := newSpeedTestNet()
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	speed, err := speedtestnet.MeasureDownloadSpeed(ctx)
	require.NoError(t, err)
	require.NotZero(t, speed)
}

func Test_speedTestNet_MeasureUploadSpeed(t *testing.T) {
	speedtestnet, err := newSpeedTestNet()
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	speed, err := speedtestnet.MeasureUploadSpeed(ctx)
	require.NoError(t, err)
	require.NotZero(t, speed)
}

func Benchmark_speedTestNet_MeasureDownloadSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		provider, err := newSpeedTestNet()
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

func Benchmark_speedTestNet_MeasureUploadSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		provider, err := newSpeedTestNet()
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
