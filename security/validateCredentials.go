package security

import (
	"encoding/json"

	"github.com/kuzzleio/sdk-go/types"
)

// ValidateCredentials validates credentials of the specified strategy for the given user.
func (s *Security) ValidateCredentials(strategy string, kuid string, body json.RawMessage, options types.QueryOptions) (bool, error) {
	if strategy == "" || kuid == "" {
		return false, types.NewError("Security.ValidateCredentials: strategy and kuid are required", 400)
	}

	result := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "security",
		Action:     "validateCredentials",
		Body:       body,
		Strategy:   strategy,
		Id:         kuid,
	}
	go s.Kuzzle.Query(query, options, result)

	res := <-result

	if res.Error != nil {
		return false, res.Error
	}

	var hasCredentials bool
	json.Unmarshal(res.Result, &hasCredentials)

	return true, nil
}
