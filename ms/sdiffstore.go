package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// SdiffStore computes the difference between the set of unique values stored at key
// and the other provided sets, and stores the result in the key stored at destination.
// If the destination key already exists, it is overwritten.
func (ms *Ms) Sdiffstore(key string, sets []string, destination string, options types.QueryOptions) (int, error) {
	if len(sets) == 0 {
		return 0, types.NewError("Ms.Sdiffstore: please provide at least one set to compare", 400)
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Keys        []string `json:"keys"`
		Destination string   `json:"destination"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "sdiffstore",
		Id:         key,
		Body:       &body{Keys: sets, Destination: destination},
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return 0, res.Error
	}

	var returnedResult int
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
