package linkedlist

import (
	"testing"

	"github.com/gpbPiazza/notebook/algo_data/stdout"
	"github.com/stretchr/testify/assert"
)

func Test_DLL_Print(t *testing.T) {
	dll := NewDoubleLinkedList(NewDLLNode(10), NewDLLNode(9))

	dll.InsertAtHead(11)
	dll.InsertAtHead(12)
	dll.InsertAtHead(13)
	dll.InsertAtHead(14)

	t.Run("test ReversePrint", func(t *testing.T) {
		stdOutStr := stdout.String((func() { dll.ReversePrint() }))
		assert.Equal(t, "14\n13\n12\n11\n10\n9\n", stdOutStr)
	})

	t.Run("test Print", func(t *testing.T) {
		stdOutStr := stdout.String((func() { dll.Print() }))
		assert.Equal(t, "9\n10\n11\n12\n13\n14\n", stdOutStr)
	})
}
