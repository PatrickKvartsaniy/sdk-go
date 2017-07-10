package core

import (
  "github.com/kuzzleio/sdk-go/wrappers"
  "github.com/kuzzleio/sdk-go/types"
  "encoding/json"
  "github.com/satori/go.uuid"
  "sync"
  "github.com/kuzzleio/sdk-go/state"
)

type IKuzzle interface {
  Query(types.KuzzleRequest, chan<- types.KuzzleResponse, *types.Options)
}

type Kuzzle struct {
  Host   string
  socket *wrappers.WebSocket
  State  *int

  wasConnected bool
  lastUrl      string
  message      chan []byte
  mu           *sync.Mutex
  jwtToken     string
}

// Kuzzle constructor
func NewKuzzle(host string, options *types.Options) (*Kuzzle, error) {
  var err error

  if options == nil {
    options = types.DefaultOptions()
  }

  k := &Kuzzle{
    Host:   host,
    mu:     &sync.Mutex{},
    socket: wrappers.NewWebSocket(options),
  }
  k.State = &k.socket.State
  if options.Connect == types.Auto {
    err = k.Connect()
  }

  return k, err
}

// Adds a listener to a Kuzzle global event. When an event is fired, listeners are called in the order of their insertion.
func (k *Kuzzle) AddListener(event int, channel chan<- interface{}) {
  k.socket.AddListener(event, channel)
}

// Connects to a Kuzzle instance using the provided host and port.
func (k *Kuzzle) Connect() error {
  if !k.isValidState() {
    return nil
  }

  wasConnected, err := k.socket.Connect(k.Host)
  if err == nil {
    if k.lastUrl != k.Host {
      k.wasConnected = false
      k.lastUrl = k.Host
    }

    if wasConnected {
      if k.jwtToken != "" {
        // todo avoid import cycle (kuzzle)
        //go func() {
        //	res, err := kuzzle.CheckToken(k, k.jwtToken)
        //
        //	if err != nil {
        //		k.jwtToken = ""
        //		k.emitEvent(event.JwtTokenExpired, nil)
        //		k.Reconnect()
        //		return
        //	}
        //
        //	if !res.Valid {
        //		k.jwtToken = ""
        //		k.emitEvent(event.JwtTokenExpired, nil)
        //	}
        //	k.Reconnect()
        //}()
      }
    }
    return nil
  }

  return err
}

func (k Kuzzle) Reconnect() {
  // todo auto resubscribe

  //todo auto replay
}

// Instantiates a new Collection object.
func (k *Kuzzle) Collection(collection, index string) *Collection {
  return NewCollection(k, collection, index)
}

// This is a low-level method, exposed to allow advanced SDK users to bypass high-level methods.
func (k *Kuzzle) Query(query types.KuzzleRequest, responseChannel chan<- types.KuzzleResponse, options *types.Options) {
  requestId := uuid.NewV4().String()

  query.RequestId = requestId

  type body struct{}
  if query.Body == nil {
    query.Body = &body{}
  }

  json, err := json.Marshal(query)
  if err != nil {
    responseChannel <- types.KuzzleResponse{Error: types.MessageError{Message: err.Error()}}
    return
  }

  err = k.socket.Send(json, options, responseChannel, requestId)
  if err != nil {
    responseChannel <- types.KuzzleResponse{Error: types.MessageError{Message: err.Error()}}
    return
  }
}

func (k *Kuzzle) Disconnect() error {
  err := k.socket.Close()

  if err != nil {
    return err
  }
  k.wasConnected = false

  return nil
}

func (k Kuzzle) isValidState() bool {
  switch k.socket.State {
  case state.Initializing, state.Ready, state.Disconnected, state.Error, state.Offline:
    return true
  }
  return false
}
