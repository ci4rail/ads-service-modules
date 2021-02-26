/*
Copyright © 2021 Ci4Rail GmbH

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"ads-node-module/internal/message/version"
	"context"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/amenzhinsky/iothub/iotdevice"
	iotmqtt "github.com/amenzhinsky/iothub/iotdevice/transport/mqtt"

	nats "github.com/nats-io/nats.go"
)

const (
	defaultUpdateIntervalMs int = 1000
)

func main() {
	log.Printf("ads-node-module version: %s\n", version.Version)
	updateIntervalMs := defaultUpdateIntervalMs
	if i := os.Getenv("UPDATE_INTERVAL_MS"); i != "" {
		interval, err := strconv.Atoi(i)
		updateIntervalMs = interval
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Info: using update interval in %d milliseconds\n", updateIntervalMs)
	} else {
		log.Printf("Info: env UPDATE_INTERVAL_MS. Using default %d\n", updateIntervalMs)
	}

	natsServer := "nats"
	natsServerEnv := os.Getenv("NATS_SERVER")
	if len(natsServerEnv) > 0 {
		natsServer = natsServerEnv
	}

	opts := []nats.Option{nats.Name("ads-node-module")}
	opts = setupConnOptions(opts)

	c, err := iotdevice.NewModuleFromEnvironment(iotmqtt.NewModuleTransport(), true)
	if err != nil {
		log.Fatal(err)
	}

	if err = c.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	nc, err := nats.Connect(natsServer, opts...)
	if err != nil {
		log.Fatal(err)
	}

	_, err = nc.Subscribe("service.location", func(msg *nats.Msg) {
		if err = c.SendEvent(context.Background(), msg.Data); err != nil {
			log.Fatal(err)
		}
	})

	if err != nil {
		log.Fatal(err)
	}

	nc.Flush()
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	runtime.Goexit()
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
