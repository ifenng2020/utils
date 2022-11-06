package algorithm

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "hello",
			arr:  []int{6, 3, 1, 7, 2, 9},
			want: []int{1, 2, 3, 6, 7, 9},
		},
		{
			name: "world",
			arr:  []int{8, 3, 1, 7, 2, 9},
			want: []int{1, 2, 3, 7, 8, 9},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			param := tC.arr
			SortInsert(tC.arr)
			if !reflect.DeepEqual(tC.arr, tC.want) {
				t.Errorf("SortInsert(%v) expect: %v and got %v", param, tC.want, tC.arr)
			}
		})
	}
}
