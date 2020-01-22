package main

import (
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

func storeInInflux(minRtt time.Duration, avgRtt time.Duration, maxRtt time.Duration) {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "mydb",
		Precision: "s",
	})

	fields := map[string]interface{}{
		"min": minRtt.Seconds(),
		"avg": avgRtt.Seconds(),
		"max": maxRtt.Seconds(),
	}

	pt, err := client.NewPoint("ping_rtt", nil, fields, time.Now().UTC())
	if err != nil {
		panic(err)
	}

	bp.AddPoint(pt)
	c.Write(bp)
}