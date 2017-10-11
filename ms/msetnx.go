package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// MsetNx sets the provided keys to their respective values, only if they do not exist.
// If a key exists, then the whole operation is aborted and no key is set.
func (ms Ms) MsetNx(entries []*types.MSKeyValue, options types.QueryOptions) (int, error) {
	if len(entries) == 0 {
		return 0, types.NewError("Ms.MsetNx: please provide at least one key/value entry")
	}

	result := make(chan *types.KuzzleResponse)

	type body struct {
		Entries []*types.MSKeyValue `json:"entries"`
	}

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "msetnx",
		Body:       &body{Entries: entries},
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
