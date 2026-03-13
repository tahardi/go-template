package live_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tahardi/go-template/internal"
)

func TestToDoImpl_Foo(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// given
		todo := &internal.ToDoImpl{}

		// when
		err := todo.Foo()

		// then
		assert.NoError(t, err)
	})
}
