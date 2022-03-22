package provider

import (
	"context"
	"time"

	"github.com/friendsofgo/errors"
	"go.jonnrb.io/speedtest/fastdotcom"
	"go.jonnrb.io/speedtest/units"
)

// Implementation of speedtest.Provider with fast.com
type fastCom struct {
	m      *fastdotcom.Manifest
	client *fastdotcom.Client
}

func newFastCom() (*fastCom, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	manifest, err := fastdotcom.GetManifest(ctx, numberOfUrlsToCheck)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get fastcom manifest")
	}
	return &fastCom{
		m:      manifest,
		client: &fastdotcom.Client{},
	}, nil
}

func (f *fastCom) MeasureDownloadSpeed(ctx context.Context) (float64, error) {
	stream := make(chan units.BytesPerSecond)
	go func() {
		for range stream {
		}
	}()
	bytesPerSec, err := f.m.ProbeDownloadSpeed(ctx, f.client, stream)
	if err != nil {
		return 0, errors.Wrap(err, "failed to test fastcom download speed")
	}
	return float64(bytesPerSec / units.MBps), nil
}

func (f *fastCom) MeasureUploadSpeed(ctx context.Context) (float64, error) {
	stream := make(chan units.BytesPerSecond)
	go func() {
		for range stream {
		}
	}()
	bytesPerSec, err := f.m.ProbeUploadSpeed(ctx, f.client, stream)
	if err != nil {
		return 0, errors.Wrap(err, "failed to test fastcom upload speed")
	}
	return float64(bytesPerSec / units.MBps), nil
}

func (f *fastCom) GetName() string {
	return "fastcom"
}
