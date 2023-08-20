package slice_sum

import "unsafe"

func ClassicSum(xs []int64) int64 {
	total := int64(0)
	for _, x := range xs {
		total += x
	}

	return total
}

func Vectorized1Sum(xs []int64) int64 {
	var acc0 int64

	var i int
	for i = 0; i < len(xs); i++ {
		x := (*[1]int64)(unsafe.Pointer(&xs[i]))

		acc0 += x[0]
	}

	return acc0
}

func Vectorized2Sum(xs []int64) int64 {
	var acc0, acc1 int64

	var i int
	for i = 0; i < len(xs); i += 2 {
		x := (*[2]int64)(unsafe.Pointer(&xs[i]))

		acc0 += x[0]
		acc1 += x[1]
	}

	for j := i; j < len(xs); j++ {
		acc0 += xs[j]
	}

	return acc0 + acc1
}

func Vectorized4Sum(xs []int64) int64 {
	var acc0, acc1, acc2, acc3 int64

	var i int
	for i = 0; i < len(xs); i += 4 {
		x := (*[4]int64)(unsafe.Pointer(&xs[i]))

		acc0 += x[0]
		acc1 += x[1]
		acc2 += x[2]
		acc3 += x[3]
	}

	for j := i; j < len(xs); j++ {
		acc0 += xs[j]
	}

	return acc0 + acc1 + acc2 + acc3
}

func Vectorized4Sum2Acc(xs []int64) int64 {
	var acc0, acc1 int64

	var i int
	for i = 0; i < len(xs); i += 4 {
		x := (*[4]int64)(unsafe.Pointer(&xs[i]))

		acc0 += x[0] + x[1]
		acc1 += x[2] + x[3]
	}

	for j := i; j < len(xs); j++ {
		acc0 += xs[j]
	}

	return acc0 + acc1
}

//go:noinline
func Vectorized4Sum2AccNoInline(xs []int64) int64 {
	var acc0, acc1 int64

	var i int
	for i = 0; i < len(xs); i += 4 {
		x := (*[4]int64)(unsafe.Pointer(&xs[i]))

		acc0 += x[0] + x[1]
		acc1 += x[2] + x[3]
	}

	for j := i; j < len(xs); j++ {
		acc0 += xs[j]
	}

	return acc0 + acc1
}

func Vectorized8Sum(xs []int64) int64 {
	var acc0, acc1, acc2, acc3, acc4, acc5, acc6, acc7 int64

	var i int
	for i = 0; i < len(xs); i += 8 {
		x := (*[8]int64)(unsafe.Pointer(&xs[i]))

		acc0 += x[0]
		acc1 += x[1]
		acc2 += x[2]
		acc3 += x[3]
		acc4 += x[4]
		acc5 += x[5]
		acc6 += x[6]
		acc7 += x[7]
	}

	for j := i; j < len(xs); j++ {
		acc0 += xs[j]
	}

	return acc0 + acc1 + acc2 + acc3 + acc4 + acc5 + acc6 + acc7
}

func Vectorized8Sum2Acc(xs []int64) int64 {
	var acc0, acc1 int64

	var i int
	for i = 0; i < len(xs); i += 8 {
		x := (*[8]int64)(unsafe.Pointer(&xs[i]))

		acc0 += x[0] + x[1] + x[2] + x[3]
		acc1 += x[4] + x[5] + x[6] + x[7]
	}

	for j := i; j < len(xs); j++ {
		acc0 += xs[j]
	}

	return acc0 + acc1
}

func Vectorized16Sum(xs []int64) int64 {
	var (
		acc0, acc1, acc2, acc3, acc4, acc5, acc6, acc7       int64
		acc8, acc9, acc10, acc11, acc12, acc13, acc14, acc15 int64
	)

	var i int
	for i = 0; i < len(xs); i += 16 {
		x := (*[16]int64)(unsafe.Pointer(&xs[i]))

		acc0 += x[0]
		acc1 += x[1]
		acc2 += x[2]
		acc3 += x[3]
		acc4 += x[4]
		acc5 += x[5]
		acc6 += x[6]
		acc7 += x[7]
		acc8 += x[8]
		acc9 += x[9]
		acc10 += x[10]
		acc11 += x[11]
		acc12 += x[12]
		acc13 += x[13]
		acc14 += x[14]
		acc15 += x[15]
	}

	for j := i; j < len(xs); j++ {
		acc0 += xs[j]
	}

	return acc0 + acc1 + acc2 + acc3 + acc4 + acc5 + acc6 + acc7 +
		acc8 + acc9 + acc10 + acc11 + acc12 + acc13 + acc14 + acc15
}

func Vectorized16Sum2Acc(xs []int64) int64 {
	var acc0, acc1 int64

	var i int
	for i = 0; i < len(xs); i += 16 {
		x := (*[16]int64)(unsafe.Pointer(&xs[i]))

		acc0 += x[0] + x[1] + x[2] + x[3] + x[4] + x[5] + x[6] + x[7]
		acc1 += x[8] + x[9] + x[10] + x[11] + x[12] + x[13] + x[14] + x[15]
	}

	for j := i; j < len(xs); j++ {
		acc0 += xs[j]
	}

	return acc0 + acc1
}
