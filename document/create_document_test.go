package collection_test

import (
	"encoding/json"
	"fmt"
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/internal"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDocumentError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{Error: types.NewError("Unit test error")}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	_, err := collection.NewCollection(k, "collection", "index").CreateDocument("id", &collection.Document{Content: []byte(`{"title":"yolo"}`)}, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Unit test error", err.(*types.KuzzleError).Message)
}

func TestCreateDocumentWrongOptionError(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			return &types.KuzzleResponse{}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)
	newCollection := collection.NewCollection(k, "collection", "index")
	opts := types.NewQueryOptions()
	opts.SetIfExist("unknown")

	_, err := newCollection.CreateDocument("id", &collection.Document{Content: []byte(`{"title":"yolo"}`)}, opts)
	assert.Equal(t, "[400] Invalid value for the 'ifExist' option: 'unknown'", fmt.Sprint(err))
}

func TestCreateDocument(t *testing.T) {
	id := "myId"

	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "document", parsedQuery.Controller)
			assert.Equal(t, "create", parsedQuery.Action)
			assert.Equal(t, "index", parsedQuery.Index)
			assert.Equal(t, "collection", parsedQuery.Collection)
			assert.Equal(t, id, parsedQuery.Id)

			body := make(map[string]interface{}, 0)
			body["title"] = "yolo"

			assert.Equal(t, body, parsedQuery.Body)

			res := collection.Document{Id: id, Content: []byte(`{"title":"yolo"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	res, _ := collection.NewCollection(k, "collection", "index").CreateDocument(id, &collection.Document{Content: []byte(`{"title":"yolo"}`)}, nil)
	assert.Equal(t, id, res.Id)
}

func TestCreateDocumentReplace(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "createOrReplace", parsedQuery.Action)

			res := collection.Document{Id: "id", Content: []byte(`{"Title":"yolo"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	newCollection := collection.NewCollection(k, "collection", "index")
	opts := types.NewQueryOptions()
	opts.SetIfExist("replace")
	newCollection.CreateDocument("id", &collection.Document{Content: []byte(`{"Title":"yolo"}`)}, opts)
}

func TestCreateDocumentCreate(t *testing.T) {
	c := &internal.MockedConnection{
		MockSend: func(query []byte, options types.QueryOptions) *types.KuzzleResponse {
			parsedQuery := &types.KuzzleRequest{}
			json.Unmarshal(query, parsedQuery)

			assert.Equal(t, "create", parsedQuery.Action)

			res := collection.Document{Id: "id", Content: []byte(`{"Title":"yolo"}`)}
			r, _ := json.Marshal(res)
			return &types.KuzzleResponse{Result: r}
		},
	}
	k, _ := kuzzle.NewKuzzle(c, nil)

	newCollection := collection.NewCollection(k, "collection", "index")
	opts := types.NewQueryOptions()
	opts.SetIfExist("error")

	newCollection.CreateDocument("id", &collection.Document{Content: []byte(`{"Title":"yolo"}`)}, opts)
}

func ExampleCollection_CreateDocument() {
	c := &internal.MockedConnection{}
	k, _ := kuzzle.NewKuzzle(c, nil)
	id := "myId"

	res, err := collection.NewCollection(k, "collection", "index").CreateDocument(id, &collection.Document{Content: []byte(`{"title":"yolo"}`)}, nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.Id, res.Content)
}
