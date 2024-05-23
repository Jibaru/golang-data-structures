package structures_test

import (
	"testing"

	"github.com/Jibaru/golang-data-structures/structures"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := structures.NewPriorityQueue[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	if pq == nil {
		t.Fatal("NewPriorityQueue should not return nil")
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	pq := structures.NewPriorityQueue[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	nums := []int{5, 3, 7, 1, 4, 6, 8}
	for _, num := range nums {
		pq.Push(num)
	}
}

func TestPriorityQueue_Pop_Lowest(t *testing.T) {
	pq := structures.NewPriorityQueue[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	nums := []int{5, 3, 7, 1, 4, 6, 8}
	for _, num := range nums {
		pq.Push(num)
	}

	sortedNums := []int{1, 3, 4, 5, 6, 7, 8}

	for _, num := range sortedNums {
		top, err := pq.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if top != num {
			t.Fatalf("expected %d but got %d", num, top)
		}
	}
}

func TestPriorityQueue_Pop_Highest(t *testing.T) {
	pq := structures.NewPriorityQueue[int](func(a, b int) int {
		if a < b {
			return 1
		}
		if a > b {
			return -1
		}
		return 0
	})
	nums := []int{5, 3, 7, 1, 4, 6, 8}
	for _, num := range nums {
		pq.Push(num)
	}

	sortedNums := []int{8, 7, 6, 5, 4, 3, 1}

	for _, num := range sortedNums {
		top, err := pq.Pop()
		if err != nil {
			t.Fatal(err)
		}
		if top != num {
			t.Fatalf("expected %d but got %d", num, top)
		}
	}
}

func TestPriorityQueue_Pop_Empy(t *testing.T) {
	pq := structures.NewPriorityQueue[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	_, err := pq.Pop()
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestPriorityQueue_Top(t *testing.T) {
	pq := structures.NewPriorityQueue[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	nums := []int{5, 3, 7, 1, 4, 6, 8}
	for _, num := range nums {
		pq.Push(num)
	}

	sortedNums := []int{1, 3, 4, 5, 6, 7, 8}

	for _, num := range sortedNums {
		top, err := pq.Top()
		if err != nil {
			t.Fatal(err)
		}
		if top != num {
			t.Fatalf("expected %d but got %d", num, top)
		}
		pq.Pop()
	}
}

func TestPriorityQueue_Top_Empty(t *testing.T) {
	pq := structures.NewPriorityQueue[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	_, err := pq.Top()
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestPriorityQueue_Size(t *testing.T) {
	pq := structures.NewPriorityQueue[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	nums := []int{5, 3, 7, 1, 4, 6, 8}
	for i, num := range nums {
		if pq.Size() != int64(i) {
			t.Fatalf("expected size to be %d but got %d", i, pq.Size())
		}
		pq.Push(num)
	}
	if pq.Size() != int64(len(nums)) {
		t.Fatalf("expected size to be %d but got %d", len(nums), pq.Size())
	}
}

func TestPriorityQueue_Size_Empy(t *testing.T) {
	pq := structures.NewPriorityQueue[int](func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})
	if pq.Size() != 0 {
		t.Fatalf("expected size to be 0 but got %d", pq.Size())
	}
}
