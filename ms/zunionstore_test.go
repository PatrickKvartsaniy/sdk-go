package ms_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZunionstoreEmptyKeys(t *testing.T) {
	k, _ := kuzzle.NewKuzzle(&internal.MockedConnection{}, nil)

	_, err := k.MemoryStorage.Zunionstore("foo", []string{}, nil)

	assert.NotNil(t, err)
	assert.Equal(t, "[400] Ms.Zunionstore: please provide at least one key", fmt.Sprint(err))
}

func TestZunionstoreError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: &types.KuzzleError{Message: "Unit test error"}}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := k.MemoryStorage.Zunionstore("foo", []string{"bar", "rab"}, nil)

	assert.NotNil(t, err)
}

func TestZunionstore(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zunionstore", parsedQuery.Action)

			r, _ := json.Marshal(2)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := k.MemoryStorage.Zunionstore("foo", []string{"bar", "rab"}, nil)

	assert.Equal(t, 2, res)
}

func TestZunionstoreWithOptions(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "ms", parsedQuery.Controller)
			assert.Equal(t, "zunionstore", parsedQuery.Action)
			assert.Equal(t, "sum", options.Aggregate())
			assert.Equal(t, []int{1, 2}, options.Weights())

			r, _ := json.Marshal(2)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetAggregate("sum")
	qo.SetWeights([]int{1, 2})
	res, _ := k.MemoryStorage.Zunionstore("foo", []string{"bar", "rab"}, qo)

	assert.Equal(t, 2, res)
}

func ExampleMs_Zunionstore() {
	c := websocket.NewWebSocket("localhost:7512", nil)
	k, _ := kuzzle.NewKuzzle(c, nil)
	qo := types.NewQueryOptions()

	qo.SetAggregate("sum")
	qo.SetWeights([]int{1, 2})
	res, err := k.MemoryStorage.Zunionstore("foo", []string{"bar", "rab"}, qo)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
