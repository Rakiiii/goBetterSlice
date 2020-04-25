package gobetterslice

func AppendWithOutRepeatInt(s1 []int,s2 ...int)[]int{
	res := make([]int,len(s1))
	copy(res,s1)
	for _,j := range s2{
		flag := true
		for _,i := range s1{
			if j == i{
				flag = false
			}
		}

		if flag {
			res = append(res,j)
		}
	}

	return res
}

func AppendWithOutRepeatInt64(s1 []int64,s2 ...int64)[]int64{
	res := make([]int64,len(s1))
	copy(res,s1)
	for _,j := range s2{
		flag := true
		for _,i := range s1{
			if j == i{
				flag = false
			}
		}

		if flag {
			res = append(res,j)
		}
	}

	return res
}

func AppendWithOutRepeatFloat(s1 []float64,s2 ...float64)[]float64{
	res := make([]float64,len(s1))
	copy(res,s1)
	for _,j := range s2{
		flag := true
		for _,i := range s1{
			if j == i{
				flag = false
			}
		}

		if flag {
			res = append(res,j)
		}
	}

	return res
}

func AppendWithOutRepeat(s1 []interface{},s2 ...interface{})[]interface{}{
	res := make([]int,len(s1))
	copy(res,s1)
	for _,j := range s2{
		flag := true
		for _,i := range s1{
			if j == i{
				flag = false
			}
		}

		if flag {
			res = append(res,j)
		}
	}

	return res
}

