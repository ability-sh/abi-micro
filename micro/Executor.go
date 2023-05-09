package micro

type Executor interface {
	Scheme() *Scheme
	Exec(ctx Context, name string, data interface{}) (interface{}, error)
}

type SchemeObject struct {
	Fields []*SchemeField `json:"fields"`
}

type SchemeField struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Required bool   `json:"required"`
}

type SchemeItem struct {
	Name   string        `json:"name"`
	Task   *SchemeObject `json:"task"`
	Result *SchemeObject `json:"result"`
}

type Scheme struct {
	Name    string                   `json:"name"`
	Items   []*SchemeItem            `json:"items"`
	Objects map[string]*SchemeObject `json:"objects"`
}
