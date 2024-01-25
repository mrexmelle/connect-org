package tree

import (
	"fmt"
	"strings"

	"github.com/mrexmelle/connect-org/internal/localerror"
)

type Class[T Entity] struct {
	Root          *Node[T]
	PathSeparator string
}

func New[T Entity](pathSeparator string) *Class[T] {
	return &Class[T]{
		Root: &Node[T]{
			Data:     new(T),
			Children: []Node[T]{},
		},
		PathSeparator: pathSeparator,
	}
}

func (c *Class[T]) AssignEntity(
	path string,
	data *T,
) error {
	return c.assignEntityIntoNode(path, data, c.Root)
}

func (c *Class[T]) assignEntityIntoNode(
	path string,
	data *T,
	node *Node[T],
) error {
	lineage := strings.Split(path, c.PathSeparator)
	if len(lineage) == 0 {
		return localerror.ErrBadHierarchy
	} else if len(lineage) == 1 {
		node.Data = data
		node.Children = []Node[T]{}
		return nil
	}

	reducedHierarchy := lineage[1]
	if len(lineage) > 2 {
		for i := 2; i < len(lineage); i++ {
			reducedHierarchy += fmt.Sprintf(".%s", lineage[i])
		}
	}
	i := 0
	for i = 0; i < len(node.Children); i++ {
		if (*node.Children[i].Data).GetId() == lineage[1] {
			c.assignEntityIntoNode(reducedHierarchy, data, &node.Children[i])
			return nil
		}
	}
	if i == len(node.Children) {
		newEntity := new(T)
		(*newEntity).SetId(lineage[1])
		node.Children = append(
			node.Children,
			Node[T]{
				Data:     newEntity,
				Children: []Node[T]{},
			},
		)
		c.assignEntityIntoNode(
			reducedHierarchy,
			data,
			&node.Children[len(node.Children)-1],
		)
	}

	return nil
}
