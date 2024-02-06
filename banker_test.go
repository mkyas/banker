package banker

import "testing"

func TestBankers0(t *testing.T) {
	allocation := [][]int{[]int{0, 1, 0},
		[]int{2, 0, 0},
		[]int{3, 0, 2},
		[]int{2, 1, 1},
		[]int{0, 0, 2}}

	need := [][]int{[]int{7, 4, 3},
		[]int{1, 2, 2},
		[]int{6, 0, 0},
		[]int{0, 1, 1},
		[]int{4, 3, 1}}

	available := []int{3, 3, 2}

	if ok, schedule := Safe(available, allocation, need); !ok {
		t.Error("Safe: ", schedule)
	}
	if ok := Bankers(1, []int{1, 0, 2}, available, allocation, need); !ok {
		t.Error("Bankers(1, {1, 0, 2})")
	}
}

func TestBankers1(t *testing.T) {
	allocation := [][]int{[]int{0, 1, 0},
		[]int{3, 0, 2},
		[]int{3, 0, 2},
		[]int{2, 1, 1},
		[]int{0, 0, 2}}

	need := [][]int{[]int{7, 4, 3},
		[]int{0, 2, 0},
		[]int{6, 0, 0},
		[]int{0, 1, 1},
		[]int{4, 3, 1}}

	available := []int{2, 3, 0}

	if ok := Bankers(4, []int{3, 3, 0}, available, allocation, need); ok {
		t.Error("Bankers(4, {3, 3, 0})")
	}
	if ok := Bankers(1, []int{0, 2, 0}, available, allocation, need); !ok {
		t.Error("Bankers(1, {0, 2, 0})")
	}
}

func TestBankers2(t *testing.T) {
	allocation := [][]int{[]int{1, 0, 2, 1, 1},
		[]int{2, 0, 1, 1, 0},
		[]int{1, 1, 0, 1, 0},
		[]int{1, 1, 1, 1, 0}}

	need := [][]int{[]int{0, 1, 0, 2, 0},
		[]int{0, 2, 1, 0, 0},
		[]int{1, 0, 3, 0, 0},
		[]int{0, 0, 1, 1, 1}}

	if ok, schedule := Safe([]int{0, 0, 0, 1, 1}, allocation, need); ok {
		t.Error("Failed x = 0: ", schedule)
	}
	if ok, schedule := Safe([]int{0, 0, 1, 1, 1}, allocation, need); !ok {
		t.Error("Failed x = 1: ", schedule)
	}
}

func TestBankers3(t *testing.T) {
	allocation := [][]int{[]int{1, 0, 2, 1, 1},
		[]int{2, 0, 1, 1, 0},
		[]int{1, 1, 0, 1, 0},
		[]int{1, 1, 1, 1, 0}}

	need := [][]int{[]int{1, 1, 2, 3, 1},
		[]int{2, 2, 2, 1, 0},
		[]int{2, 1, 3, 1, 0},
		[]int{1, 1, 2, 2, 1}}

	for x := 0; x < 6; x++ {
		if ok, schedule := Safe([]int{0, 0, x, 1, 1}, allocation, need); ok {
			t.Error("Failed ", x, schedule)
		}
	}
}
