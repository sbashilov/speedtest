package speedtest

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sbashilov/speedtest/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var zeroSpeed = float64(0)

func TestMeasureSpeedSuccess(t *testing.T) {
	downSpeed := float64(10)
	upSpeed := float64(11)
	ctx := context.Background()
	successMock := mocks.Provider{}
	successMock.On("GetName").Return("provider")
	successMock.On("MeasureDownloadSpeed", mock.Anything).Return(downSpeed, nil)
	successMock.On("MeasureUploadSpeed", mock.Anything).Return(upSpeed, nil)
	dwn, upl, err := MeasureSpeed(ctx, &successMock, time.Second)
	require.NoError(t, err)
	require.Equal(t, downSpeed, dwn)
	require.Equal(t, upSpeed, upl)
}

func TestMeasureSpeedUploadSpeedErr(t *testing.T) {
	ctx := context.Background()
	successMock := mocks.Provider{}
	successMock.On("GetName").Return("provider")
	successMock.On("MeasureDownloadSpeed", mock.Anything).Return(float64(10), nil)
	successMock.On("MeasureUploadSpeed", mock.Anything).Return(zeroSpeed, errors.New("failed to measure download speed"))
	dwn, upl, err := MeasureSpeed(ctx, &successMock, time.Second)
	require.Error(t, err)
	require.Equal(t, zeroSpeed, dwn)
	require.Equal(t, zeroSpeed, upl)
}

func TestMeasureSpeedDownloadSpeedErr(t *testing.T) {
	ctx := context.Background()
	successMock := mocks.Provider{}
	successMock.On("GetName").Return("provider")
	successMock.On("MeasureDownloadSpeed", mock.Anything).Return(zeroSpeed, errors.New("failed to measure download speed"))
	dwn, upl, err := MeasureSpeed(ctx, &successMock, time.Second)
	require.Error(t, err)
	require.Equal(t, zeroSpeed, dwn)
	require.Equal(t, zeroSpeed, upl)
}
