package slice

import (
	"github.com/TangTangHC/basic-go-study/errors"
)

type T any

func NewSlice(len, cap int) []T {
	return make([]T, len, cap)
}

func Len(sli []T) int {
	return len(sli)
}

func Cap(sli []T) int {
	return cap(sli)
}

func Append(sli []T, v T) error {
	sli = append(sli, v)
	return nil
}

func Set(sli []T, i int, v T) error {
	if i < 0 || i >= Len(sli) {
		return errors.IndexOutOfBoundsError
	}
	sli[i] = v
	return nil
}

func Add(sli []T, i int, v T) error {
	if i < 0 || i > Len(sli) {
		return errors.IndexOutOfBoundsError
	}
	sli = append(sli, v)
	copy(sli[:i+1], sli[:i])
	sli[i] = v
	return nil
}

func Find(sli []T, tar T, wn int) (int, bool, error) {
	if wn <= 0 || wn >= Len(sli) {
		return -1, false, errors.IndexOutOfBoundsError
	}
	i := 1
	for k, v := range sli {
		if v == tar {
			if i == wn {
				return k, true, nil
			}
			i++
		}
	}
	return -1, false, nil
}

func FindFirst(sli []T, tar T) (int, bool, error) {
	return Find(sli, tar, 1)
}

func FindLast(sli []T, tar T) (int, bool, error) {
	for i := Len(sli) - 1; i >= 0; i-- {
		if sli[i] == tar {
			return i, true, nil
		}
	}
	return -1, false, nil
}

func FindFunc(sli []T, tar T, wn int, f func(T, T) bool) (int, bool, error) {
	if wn <= 0 || wn >= Len(sli) {
		return -1, false, errors.IndexOutOfBoundsError
	}
	i := 1
	for k, v := range sli {
		if f(v, tar) {
			if i == wn {
				return k, true, nil
			}
			i++
		}
	}
	return -1, false, nil
}

func FindFuncFirst(sli []T, tar T, f func(T, T) bool) (int, bool, error) {
	return FindFunc(sli, tar, 0, f)
}

func FindFuncLast(sli []T, tar T, f func(T, T) bool) (int, bool, error) {
	for i := Len(sli); i >= 0; i-- {
		if f(sli[i], tar) {
			return i, true, nil
		}
	}
	return -1, false, nil
}

// Delete

func Delete(sli []T, index int) ([]T, T, error) {
	if index < 0 || index >= Len(sli) {
		var n T
		return nil, n, errors.IndexOutOfBoundsError
	}
	sli = shrink(sli)
	res := sli[index]
	copy(sli[index:], sli[index+1:])
	sli = sli[:Len(sli)-1]
	return sli, res, nil
}

// - 如果容量 > 2048，并且长度小于容量一半，那么就会缩容为原本的 5/8
// - 如果容量 (64, 2048]，如果长度是容量的 1/4，那么就会缩容为原本的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU去执行缩容
func shrink(sli []T) []T {
	c := Cap(sli)
	l := Len(sli)
	if c > 64 && c <= 2048 && l <= c/4 {
		factor := 0.625
		nCap := int(float64(c) * factor)
		nSli := make([]T, l, nCap)
		copy(nSli, sli)
		return nSli
	}
	if c > 2048 && l <= c/2 {
		nSli := make([]T, l, c/2)
		copy(nSli, sli)
		return nSli
	}
	return sli
}
