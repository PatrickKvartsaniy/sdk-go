package ms

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// Scan iterates incrementally the set of keys in the database using a cursor.
// An iteration starts when the cursor is set to 0.
// To get the next page of results, simply re-send the identical request
// with the updated cursor position provided in the result set.
// The scan terminates when the next position cursor returned by the server is 0.
func (ms *Ms) Scan(cursor int, options types.QueryOptions) (*types.MSScanResponse, error) {
	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "ms",
		Action:     "scan",
		Cursor:     cursor,
	}

	if options != nil {
		if options.Count() != 0 {
			query.Count = options.Count()
		}

		if options.Match() != "" {
			query.Match = options.Match()
		}
	}

	go ms.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return nil, res.Error
	}

	var scanResponse = &types.MSScanResponse{}
	json.Unmarshal(res.Result, scanResponse)

	return scanResponse, nil
}
