// Copyright 2015-2017 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth_test

import (
	"encoding/json"
	"testing"

	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
)

func TestCredentialsExistEmptyStrategy(t *testing.T) {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	_, err := k.Auth.CredentialsExist("", nil)
	assert.NotNil(t, err)
}

func TestCredentialsExistQueryError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "credentialsExist", request.Action)

			return &types.KuzzleResponse{Error: types.KuzzleError{Message: "error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	_, err := k.Auth.CredentialsExist("local", nil)
	assert.NotNil(t, err)
}

func TestCredentialsExists(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			request := types.KuzzleRequest{}
			json.Unmarshal(query, &request)
			assert.Equal(t, "auth", request.Controller)
			assert.Equal(t, "credentialsExist", request.Action)

			ret, _ := json.Marshal(true)
			return &types.KuzzleResponse{Result: ret}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	res, _ := k.Auth.CredentialsExist("local", nil)
	assert.Equal(t, true, res)
}

func ExampleAdminExists() {
	c := websocket.NewWebSocket("localhost", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	k.Connect()
	res, _ := k.Auth.CredentialsExist("local", nil)
	println(res)
}
