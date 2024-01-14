package tree

type Node[T Entity] struct {
	Data     *T        `json:"data"`
	Children []Node[T] `json:"children"`
}

type Entity interface {
	GetId() string
	SetId(id string)
}
