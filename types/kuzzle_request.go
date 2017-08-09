package types

type KuzzleRequest struct {
	RequestId  string       `json:"requestId,omitempty"`
	Controller string       `json:"controller,omitempty"`
	Action     string       `json:"action,omitempty"`
	Index      string       `json:"index,omitempty"`
	Collection string       `json:"collection,omitempty"`
	Body       interface{}  `json:"body"`
	Id         string       `json:"_id,omitempty"`
	From       int          `json:"from"`
	Size       int          `json:"size"`
	Scroll     string       `json:"scroll,omitempty"`
	ScrollId   string       `json:"scrollId,omitempty"`
	Strategy   string       `json:"strategy,omitempty"`
	ExpiresIn  int          `json:"expiresIn"`
	Volatile   VolatileData `json:"volatile"`
	Scope      string       `json:"scope"`
	State      string       `json:"state"`
	User       string       `json:"user"`
}

type SubscribeQuery struct {
	Scope string      `json:"scope"`
	State string      `json:"state"`
	User  string      `json:"user"`
	Body  interface{} `json:"body"`
}

type VolatileData map[string]interface{}

type UserCredentials map[string]interface{}

type UserData struct {
	ProfileIds  []string               `json:"profileIds"`
	Content     map[string]interface{} `json:"content"`
	Credentials map[string]interface{} `json:"credentials"`
}

type Policy struct {
	RoleId string `json:"roleId"`
}

type Policies struct {
	Policies []Policy `json:"policies"`
}

type SearchFilters struct {
	Query        interface{}          `json:"query,omitempty"`
	Sort         []map[string]string  `json:"sort,omitempty"`
	Aggregations interface{}          `json:"aggregations,omitempty"`
	SearchAfter  []interface{}        `json:"search_after,omitempty"`
}
