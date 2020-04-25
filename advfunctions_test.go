package gobetterslice

import (
	"fmt"
	"testing"
)

func TestAppenWithOutRepeatInt(t *testing.T) {
	fmt.Println("Start TestAppendWithOutRepeatInt")

	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 1, 6, 5, 4}
	right := []int{1, 2, 3, 4, 6, 5}

	res := AppendWithOutRepeatInt(slice1, slice2...)

	checkFlag := true

	for i, j := range res {
		if j != right[i] {
			t.Error("Wrong value:", j, " at position:", i, " expected:", right[i])
			checkFlag = false
		}
	}

	if checkFlag {
		fmt.Println("TestAppendWithOutRepeatInt=[ok]")
	}
}

func TestAppenWithOutRepeatInt64(t *testing.T) {
	fmt.Println("Start TestAppendWithOutRepeatInt64")

	slice1 := []int64{1, 2, 3, 4}
	slice2 := []int64{3, 4, 1, 6, 5, 4}
	right := []int64{1, 2, 3, 4, 6, 5}

	res := AppendWithOutRepeatInt64(slice1, slice2...)

	checkFlag := true

	for i, j := range res {
		if j != right[i] {
			t.Error("Wrong value:", j, " at position:", i, " expected:", right[i])
			checkFlag = false
		}
	}

	if checkFlag {
		fmt.Println("TestAppendWithOutRepeatInt64=[ok]")
	}
}
