package provider

import (
	"github.com/friendsofgo/errors"
	"github.com/sbashilov/speedtest"
)

const numberOfUrlsToCheck = 1

type Type int

const (
	SpeedTestNetProvider Type = iota
	FastComProvider
)

var ErrUnknownProviderType = errors.New("unknown provider type")

// NewProvider returns new speed test provider by passed type
// If passed unknown provider returns ErrUnknownProviderType
func NewProvider(providerType Type) (speedtest.Provider, error) {
	switch providerType {
	case SpeedTestNetProvider:
		return newSpeedTestNet()
	case FastComProvider:
		return newFastCom()
	default:
		return nil, ErrUnknownProviderType
	}
}
