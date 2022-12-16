package proto

import "errors"

var AttributeNotFound = errors.New("attribute not found")

type Named interface {
	GetID() string
	GetFQN() string
	GetName() string
	ToMermaid() string
}
