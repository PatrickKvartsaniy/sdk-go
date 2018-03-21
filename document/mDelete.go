package document

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

func (d *Document) MDelete(index string, collection string, ids []string, options types.QueryOptions) ([]string, error) {
	if index == "" {
		return nil, types.NewError("Document.MDelete: index required", 400)
	}

	if collection == "" {
		return nil, types.NewError("Document.MDelete: collection required", 400)
	}

	if len(ids) == 0 {
		return nil, types.NewError("Document.MDelete: ids filled array required", 400)
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Index:      index,
		Collection: collection,
		Controller: "document",
		Action:     "mDelete",
		Body:       ids,
	}

	go d.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return nil, res.Error
	}

	var mDeleted []string
	json.Unmarshal(res.Result, &mDeleted)

	return mDeleted, nil
}
