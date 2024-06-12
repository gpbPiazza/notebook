package linkedlist

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	type testCase struct {
		name  string
		input int
		want  *Node
	}

	tests := []testCase{
		{
			name:  "should return new pointer to new list with val int",
			input: 10,
			want:  &Node{Val: 10},
		},
		{
			name: "should return empty head where input is empty",
			want: &Node{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.input)
			if !reflect.DeepEqual(*tt.want, *got) {
				t.Errorf("%s - got: %v want: %v", tt.name, *got, *tt.want)
			}
		})
	}

}

func Test_SliceToList(t *testing.T) {
	type testCase struct {
		name  string
		input []int
		want  *Node
	}

	tests := []testCase{
		{
			name:  "should return list poiting in the same order as the slice",
			input: []int{10, 2, 3},
			want:  &Node{Val: 10, Next: &Node{Val: 2, Next: &Node{Val: 3}}},
		},
		{
			name:  "should return list poiting to only on node",
			input: []int{10},
			want:  &Node{Val: 10},
		},
		{
			name:  "should return nil if input is nil",
			input: nil,
			want:  nil,
		},
		{
			name:  "should return nil if input has zero len",
			input: []int{},
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SliceToList(tt.input)
			if got != nil && tt.want != nil && !reflect.DeepEqual(*tt.want, *got) {
				t.Errorf("%s - got: %v want: %v", tt.name, *got, *tt.want)
			}
		})
	}

}

func Test_FindMiddleNode(t *testing.T) {
	type testCase struct {
		name      string
		want      *Node
		buildList func() *Node
	}

	cases := []testCase{
		{
			name: "should return the first node when list only have one node",
			want: &Node{Val: 10},
			buildList: func() *Node {
				return &Node{Val: 10}
			},
		},
		{
			name: "should return empty Node if the list is empty",
			want: &Node{},
			buildList: func() *Node {
				return &Node{}
			},
		},
		{
			name: "should return the middle node (index 1) in the list of 3 elements",
			want: &Node{Val: 2, Next: &Node{Val: 3}},
			buildList: func() *Node {
				return &Node{Val: 10, Next: &Node{Val: 2, Next: &Node{Val: 3}}}
			},
		},
		{
			name: "should return the middle node (index 3) in the list of 6 elements",
			want: &Node{
				Val: 4,
				Next: &Node{
					Val: 5,
					Next: &Node{
						Val: 6,
					},
				},
			},
			buildList: func() *Node {
				return &Node{
					Val: 1,
					Next: &Node{
						Val: 2,
						Next: &Node{
							Val: 3,
							Next: &Node{
								Val: 4,
								Next: &Node{
									Val: 5,
									Next: &Node{
										Val: 6,
									},
								},
							},
						},
					},
				}
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			middleNode := tt.buildList().FindMiddleNode()
			if middleNode != nil && tt.want != nil && !reflect.DeepEqual(*tt.want, *middleNode) {
				t.Errorf("%s - want: %v - got: %v", tt.name, *tt.want, *middleNode)
			}
		})
	}
}

func Test_DeleteMiddleNode(t *testing.T) {
	type testCase struct {
		name      string
		want      *Node
		buildList func() *Node
	}

	cases := []testCase{
		{
			name: "should return first node when has two elements",
			want: &Node{Val: 1},
			buildList: func() *Node {
				return &Node{Val: 1, Next: &Node{Val: 2}}
			},
		},
		{
			name: "should return the list without the node with val 7",
			// [1,3,4,1,2,6]
			want: &Node{Val: 1, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 6}}}}}},
			// [1,3,4,7,1,2,6]
			buildList: func() *Node {
				return &Node{Val: 1, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: &Node{Val: 7, Next: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 6}}}}}}}
			},
		},
		{
			name: "should return the list without the node with val 3",
			// [1,2,4]
			want: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 4}}},
			// [1,2,3,4]
			buildList: func() *Node {
				return &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4}}}}
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.buildList().DeleteMiddleNode()

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("%s - want: %v got: %v", tt.name, tt.want, got)
			}
		})
	}
}

func Test_Add(t *testing.T) {
	type testCase struct {
		name          string
		want          *Node
		elementsToAdd []int
	}

	cases := []testCase{
		{
			name:          "should return first node when has one element",
			want:          &Node{Val: 1},
			elementsToAdd: []int{1},
		},
		{
			name:          "should return new list if the same amount of nodes",
			want:          &Node{Val: 1, Next: &Node{Val: 3, Next: &Node{Val: 3, Next: &Node{Val: 7, Next: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 6}}}}}}},
			elementsToAdd: []int{1, 3, 3, 7, 1, 2, 6},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			list := New(tt.elementsToAdd[0])

			for i := 1; i < len(tt.elementsToAdd); i++ {
				list.Add(tt.elementsToAdd[i])
			}

			if !reflect.DeepEqual(tt.want, list) {
				t.Errorf("%s - want: %v got: %v", tt.name, tt.want, list)
			}
		})
	}
}

func Test_SortByOddIndexes(t *testing.T) {
	type testCase struct {
		name      string
		buildList func() *Node
		want      *Node
	}

	cases := []testCase{
		{
			name: "sort by odd indexex 01",
			buildList: func() *Node {
				return SliceToList([]int{1, 2, 3, 4, 5})
			},
			want: &Node{Val: 1, Next: &Node{Val: 3, Next: &Node{Val: 5, Next: &Node{Val: 2, Next: &Node{Val: 4}}}}},
		},
		{
			name: "sort by odd indexes 02",
			buildList: func() *Node {
				return SliceToList([]int{30, 2, 3, 4, 5, 1, 3, 4})
			},
			want: &Node{Val: 30, Next: &Node{Val: 3, Next: &Node{Val: 5, Next: &Node{Val: 3,
				Next: &Node{Val: 2, Next: &Node{Val: 4, Next: &Node{Val: 1, Next: &Node{Val: 4}}}}}}}},
		},
		{
			name: "sort by odd indexes 03",
			buildList: func() *Node {
				return SliceToList([]int{1, 2})
			},
			want: &Node{Val: 1, Next: &Node{Val: 2}},
		},
		{
			name: "sort by odd indexes 04",
			buildList: func() *Node {
				return SliceToList([]int{1})
			},
			want: &Node{Val: 1},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.buildList()

			sortedList := list.SortByOddIndexes()

			if !reflect.DeepEqual(sortedList, tt.want) {
				t.Errorf("%s - wat: %v got: %v", tt.name, tt.want, sortedList)
			}
		})
	}
}

func Test_Reverse(t *testing.T) {
	type testCase struct {
		name      string
		buildList func() *Node
		want      *Node
	}

	cases := []testCase{
		{
			name: "should revert list order 01",
			buildList: func() *Node {
				return SliceToList([]int{1, 2, 3, 4, 5})
			},
			want: &Node{Val: 5, Next: &Node{Val: 4, Next: &Node{Val: 3, Next: &Node{Val: 2, Next: &Node{Val: 1}}}}},
		},
		{
			name: "should revert list order 02",
			buildList: func() *Node {
				return SliceToList([]int{30, 2, 3, 4, 5, 1, 3, 4})
			},
			want: &Node{Val: 4, Next: &Node{Val: 3, Next: &Node{Val: 1, Next: &Node{Val: 5,
				Next: &Node{Val: 4, Next: &Node{Val: 3, Next: &Node{Val: 2, Next: &Node{Val: 30}}}}}}}},
		},
		{
			name: "should revert list order 03",
			buildList: func() *Node {
				return SliceToList([]int{1, 2})
			},
			want: &Node{Val: 2, Next: &Node{Val: 1}},
		},
		{
			name: "should revert list order 04",
			buildList: func() *Node {
				return SliceToList([]int{1})
			},
			want: &Node{Val: 1},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.buildList()

			reversed := list.Reverse()

			if !reflect.DeepEqual(reversed, tt.want) {
				t.Errorf("%s - wat: %v got: %v", tt.name, tt.want, list)
			}
		})
	}
}

func TestNode_Print(t *testing.T) {
	t.Run("print  list", func(t *testing.T) {
		list := SliceToList([]int{1, 11, 2, 22, 3, 33, 4, 44, 5, 55})
		list.Print()

		emptyList := SliceToList([]int{})
		emptyList.Print()
	})
}

func TestNode_LastNode(t *testing.T) {
	t.Run("return 6", func(t *testing.T) {
		list := SliceToList([]int{1, 2, 3, 4, 5, 6})

		lastNode := list.LastNode()

		assert.Equal(t, 6, lastNode.Val)
	})

	t.Run("return nil in empty list", func(t *testing.T) {
		list := SliceToList([]int{})
		empty := list.LastNode()
		assert.Nil(t, empty)
	})

	t.Run("return first element when list only have one node", func(t *testing.T) {
		list := SliceToList([]int{1})
		lastNode := list.LastNode()
		assert.Equal(t, 1, lastNode.Val)
	})
}

func TestNode_SelfDelete(t *testing.T) {
	tests := []struct {
		name      string
		buildList func() (*Node, *Node)
		want      *Node
	}{
		{
			name: "should delete node 3 and all list are OK",
			buildList: func() (*Node, *Node) {
				list := SliceToList([]int{1, 2, 3, 4, 5})
				return list, list.Next.Next
			},
			want: &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 4, Next: &Node{Val: 5}}}},
		},
		{
			name: "should delete node 20 and all list are OK",
			buildList: func() (*Node, *Node) {
				list := SliceToList([]int{12, 20, 2, 33, 44})
				return list, list.Next
			},
			want: &Node{Val: 12, Next: &Node{Val: 2, Next: &Node{Val: 33, Next: &Node{Val: 44}}}},
		},
		{
			name: "should should delete node 12",
			buildList: func() (*Node, *Node) {
				list := SliceToList([]int{12, 22})
				return list, list
			},
			want: &Node{Val: 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entirelist, nodeToBeDeleted := tt.buildList()

			nodeToBeDeleted.SelfDelete()

			if !reflect.DeepEqual(entirelist, tt.want) {
				t.Errorf("%v wanted %v - got %v", tt.name, tt.want, entirelist)
			}
		})
	}
}
