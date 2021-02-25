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
	"ads-node-module/internal/message"
	"ads-node-module/internal/message/version"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/amenzhinsky/iothub/iotdevice"
	iotmqtt "github.com/amenzhinsky/iothub/iotdevice/transport/mqtt"
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

	c, err := iotdevice.NewModuleFromEnvironment(iotmqtt.NewModuleTransport(), true)
	if err != nil {
		log.Fatal(err)
	}

	if err = c.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}
	counter := 0
	for {
		message := &message.Message{
			Timestamp: time.Now().Unix(),
			Payload: message.Payload{
				Counter: counter,
			},
		}
		counter++
		j, err := json.Marshal(message)
		if err != nil {
			log.Fatal(err)
		}
		if err = c.SendEvent(context.Background(), j); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Duration(updateIntervalMs) * time.Millisecond)
		fmt.Printf("Sending value: %d\n", counter)
	}
}
