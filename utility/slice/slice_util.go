package utility

type sliceUtil struct {
}

var sliceUtilCon *sliceUtil

func NewSliceUtil() *sliceUtil {
	sliceUtilCon = new(sliceUtil)
	return sliceUtilCon
}

// Union 求并集
func (s sliceUtil) Union(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// Intersect 求交集
func (s sliceUtil) Intersect(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int)
	nn := make([]int64, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// Difference 求差集 slice1-并集
func (s sliceUtil) Difference(slice1, slice2 []int64) []int64 {
	m := make(map[int64]int)
	nn := make([]int64, 0)
	inter := s.Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}
