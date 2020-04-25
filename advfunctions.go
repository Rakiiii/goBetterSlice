package gobetterslice

//AppendWithOutRepeatInt return slice of int constaractde of elements from @s1 and @s2 without repeat,if @s1 or @s2 contains
//repeats they wouldn't be removed
func AppendWithOutRepeatInt(s1 []int, s2 ...int) []int {
	res := make([]int, len(s1))
	copy(res, s1)
	for _, j := range s2 {
		flag := true
		for _, i := range s1 {
			if j == i {
				flag = false
			}
		}

		if flag {
			res = append(res, j)
		}
	}

	return res
}

//AppendWithOutRepeatInt return slice of int64 constaractde of elements from @s1 and @s2 without repeat,if @s1 or @s2 contains
//repeats they wouldn't be removed
func AppendWithOutRepeatInt64(s1 []int64, s2 ...int64) []int64 {
	res := make([]int64, len(s1))
	copy(res, s1)
	for _, j := range s2 {
		flag := true
		for _, i := range s1 {
			if j == i {
				flag = false
			}
		}

		if flag {
			res = append(res, j)
		}
	}

	return res
}

//AppendWithOutRepeatInt return slice of float64 constaractde of elements from @s1 and @s2 without repeat,if @s1 or @s2 contains
//repeats they wouldn't be removed
func AppendWithOutRepeatFloat(s1 []float64, s2 ...float64) []float64 {
	res := make([]float64, len(s1))
	copy(res, s1)
	for _, j := range s2 {
		flag := true
		for _, i := range s1 {
			if j == i {
				flag = false
			}
		}

		if flag {
			res = append(res, j)
		}
	}

	return res
}

//AppendWithOutRepeatInt return slice of float64 constaractde of elements from @s1 and @s2 without repeat,if @s1 or @s2 contains
//repeats they wouldn't be removed
func AppendWithOutRepeatString(s1 []string, s2 ...string) []string {
	res := make([]string, len(s1))
	copy(res, s1)
	for _, j := range s2 {
		flag := true
		for _, i := range s1 {
			if j == i {
				flag = false
			}
		}

		if flag {
			res = append(res, j)
		}
	}

	return res
}

//AppendWithOutRepeatInt return slice of interface{} constaractde of elements from @s1 and @s2 without repeat,if @s1 or @s2 contains
//repeats they wouldn't be removed
func AppendWithOutRepeat(s1 []interface{}, s2 ...interface{}) []interface{} {
	res := make([]interface{}, len(s1))
	copy(res, s1)
	for _, j := range s2 {
		flag := true
		for _, i := range s1 {
			if j == i {
				flag = false
			}
		}

		if flag {
			res = append(res, j)
		}
	}

	return res
}
