package collection

import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

// GetSpecifications retrieves the current specifications of the collection.
func (dc Collection) GetSpecifications(options types.QueryOptions) (*types.KuzzleSpecificationsResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "getSpecifications",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	specification := &types.KuzzleSpecificationsResult{}

	if res.Error != nil {
		return specification, res.Error
	}

	json.Unmarshal(res.Result, specification)

	return specification, nil
}

// SearchSpecifications searches specifications across indexes/collections according to the provided filters.
func (dc Collection) SearchSpecifications(filters interface{}, options types.QueryOptions) (*types.KuzzleSpecificationSearchResult, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "collection",
		Action:     "searchSpecifications",
		Body: struct {
			Query interface{} `json:"query"`
		}{Query: filters},
	}

	if options != nil {
		query.From = options.GetFrom()
		query.Size = options.GetSize()
		scroll := options.GetScroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	specifications := &types.KuzzleSpecificationSearchResult{}

	if res.Error != nil {
		return specifications, res.Error
	}

	json.Unmarshal(res.Result, specifications)

	return specifications, nil
}

// ScrollSpecifications retrieves next result of a specification search with scroll query.
func (dc Collection) ScrollSpecifications(scrollId string, options types.QueryOptions) (*types.KuzzleSpecificationSearchResult, error) {
	specifications := &types.KuzzleSpecificationSearchResult{}

	if scrollId == "" {
		return specifications, types.NewError("Collection.ScrollSpecifications: scroll id required")
	}

	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Controller: "collection",
		Action:     "scrollSpecifications",
		ScrollId:   scrollId,
	}

	if options != nil {
		scroll := options.GetScroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return specifications, res.Error
	}

	json.Unmarshal(res.Result, specifications)

	return specifications, nil
}

// ValidateSpecifications validates the provided specifications.
func (dc Collection) ValidateSpecifications(specifications *types.KuzzleValidation, options types.QueryOptions) (*types.ValidResponse, error) {
	ch := make(chan *types.KuzzleResponse)

	specificationsData := types.KuzzleSpecifications{
		dc.index: {
			dc.collection: specifications,
		},
	}

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "validateSpecifications",
		Body:       specificationsData,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch
	response := &types.ValidResponse{}

	if res.Error != nil {
		return response, res.Error
	}

	json.Unmarshal(res.Result, response)

	return response, nil
}

// UpdateSpecifications updates the current specifications of this collection.
func (dc Collection) UpdateSpecifications(specifications *types.KuzzleValidation, options types.QueryOptions) (*types.KuzzleSpecifications, error) {
	ch := make(chan *types.KuzzleResponse)

	specificationsData := &types.KuzzleSpecifications{
		dc.index: {
			dc.collection: specifications,
		},
	}

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "updateSpecifications",
		Body:       specificationsData,
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch
	specification := &types.KuzzleSpecifications{}

	if res.Error != nil {
		return specification, res.Error
	}

	json.Unmarshal(res.Result, specification)

	return specification, nil
}

// DeleteSpecifications deletes the current specifications of this collection.
func (dc Collection) DeleteSpecifications(options types.QueryOptions) (*types.AckResponse, error) {
	ch := make(chan *types.KuzzleResponse)

	query := &types.KuzzleRequest{
		Collection: dc.collection,
		Index:      dc.index,
		Controller: "collection",
		Action:     "deleteSpecifications",
	}
	go dc.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error != nil {
		return &types.AckResponse{Acknowledged: false}, res.Error
	}

	response := &types.AckResponse{}
	json.Unmarshal(res.Result, response)

	return response, nil
}
