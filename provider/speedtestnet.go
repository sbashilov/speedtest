package provider

import (
	"context"

	"github.com/friendsofgo/errors"
	"go.jonnrb.io/speedtest/geo"
	"go.jonnrb.io/speedtest/speedtestdotnet"
	"go.jonnrb.io/speedtest/units"
)

// Implementation of speedtest.Provider with speedtest.net
type speedTestNet struct {
	client *speedtestdotnet.Client
	config speedtestdotnet.Config
	server speedtestdotnet.Server
}

func newSpeedTestNet() (*speedTestNet, error) {
	ctx := context.Background()
	client := &speedtestdotnet.Client{}
	config, err := client.Config(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get speedtestnet config")
	}
	servers, err := client.LoadAllServers(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load speedtestnet servers")
	}
	if len(servers) == 0 {
		return nil, errors.New("no servers found to test with speedtestnet")
	}
	return &speedTestNet{
		client: client,
		config: config,
		server: getNearestServer(servers, config.Coordinates),
	}, nil
}

func (s *speedTestNet) MeasureDownloadSpeed(ctx context.Context) (float64, error) {
	stream := make(chan units.BytesPerSecond)
	go func() {
		for range stream {
		}
	}()
	bytesPerSec, err := s.server.ProbeDownloadSpeed(ctx, s.client, stream)
	if err != nil {
		return 0, errors.Wrap(err, "failed to test speedtestnet download speed")
	}
	return float64(bytesPerSec / units.MBps), nil
}

func (s *speedTestNet) MeasureUploadSpeed(ctx context.Context) (float64, error) {
	stream := make(chan units.BytesPerSecond)
	go func() {
		for range stream {
		}
	}()
	bytesPerSec, err := s.server.ProbeUploadSpeed(ctx, s.client, stream)
	if err != nil {
		return 0, errors.Wrap(err, "failed to test speedtestnet upload speed")
	}
	return float64(bytesPerSec / units.MBps), nil
}

func (s *speedTestNet) GetName() string {
	return "speedtestnet"
}

func getNearestServer(servers []speedtestdotnet.Server, coords geo.Coordinates) speedtestdotnet.Server {
	minDist := coords.DistanceTo(servers[0].Coordinates)
	minDistServer := servers[0]
	for i := 1; i < len(servers); i++ {
		dist := coords.DistanceTo(servers[i].Coordinates)
		if dist > minDist {
			continue
		}
		minDist = dist
		minDistServer = servers[i]
	}
	return minDistServer
}
