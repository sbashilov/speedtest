package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sbashilov/speedtest"
	"github.com/sbashilov/speedtest/provider"
)

func main() {
	speedtestnet, err := provider.NewProvider(provider.SpeedTestNetProvider)
	if err != nil {
		log.Fatal(err)
	}
	dwnSpeed, updSpeed, err := speedtest.MeasureSpeed(context.Background(), speedtestnet, time.Second*1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Speedtestnet dwn: %f, upd: %f\n", dwnSpeed, updSpeed)
	fastcom, err := provider.NewProvider(provider.FastComProvider)
	if err != nil {
		log.Fatal(err)
	}
	dwnSpeed, updSpeed, err = speedtest.MeasureSpeed(context.Background(), fastcom, time.Second*1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Fastcom dwn: %f, upd: %f", dwnSpeed, updSpeed)
}
