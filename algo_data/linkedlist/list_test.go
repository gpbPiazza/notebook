package linkedlist

import (
	"reflect"
	"testing"
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
