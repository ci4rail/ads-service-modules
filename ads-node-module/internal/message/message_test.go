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

package message

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayloadCreation(t *testing.T) {
	assert := assert.New(t)
	message := &Message{
		Timestamp: 234,
		Payload: Payload{
			Counter: 1,
		},
	}
	j, err := json.Marshal(message)
	assert.Nil(err)
	unmarshalled := &Message{}
	err = json.Unmarshal(j, &unmarshalled)
	assert.Nil(err)
	assert.Equal(unmarshalled.Timestamp, int64(234))
	assert.Equal(unmarshalled.Payload.Counter, 1)
}
