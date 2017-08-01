package ms

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/types"
)

/*
  Inspects the low-level properties of a key.
*/
func (ms Ms) Object(key string, subcommand string, options types.QueryOptions) (string, error) {
	if key == "" {
		return "", errors.New("Ms.Object: key required")
	}
	if subcommand != "refcount" && subcommand != "encoding" && subcommand != "idletime" {
		return "", errors.New("Ms.Object: subcommand required, possible values: refcount|encoding|idletime")
	}

	result := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "ms",
		Action:     "object",
		Id:         key,
		Subcommand: subcommand,
	}
	go ms.Kuzzle.Query(query, nil, result)

	res := <-result

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}
	var returnedResult string
	json.Unmarshal(res.Result, &returnedResult)

	return returnedResult, nil
}
