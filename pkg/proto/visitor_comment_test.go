package proto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var commentVisitor = &CommentVisitor{}

func TestCommentVisitor_CanVisit(t *testing.T) {
	l := &Line{Comment: "Test Comment", Token: InlineCommentPrefix}
	assert.True(t, commentVisitor.CanVisit(l))

	l = &Line{Comment: "Test Comment", Token: "/*"}
	assert.True(t, commentVisitor.CanVisit(l))
}

func TestCommentVisitor_Visit(t *testing.T) {

}
