package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Truncate delete every Documents from the provided Collection.
func (dc Collection) Truncate(options types.QueryOptions) (*types.AckResponse, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "truncate",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch
	ack := &types.AckResponse{}

	if res.Error != nil {
		return ack, res.Error
	}

	json.Unmarshal(res.Result, ack)

	return ack, nil
}
