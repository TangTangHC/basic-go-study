package slice

import (
	"github.com/TangTangHC/basic-go-study/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	tcs := []struct {
		name    string
		sli     []T
		index   int
		want    []T
		wantOne T
		wantErr error
	}{
		{
			name:    "空切片",
			sli:     nil,
			index:   0,
			want:    nil,
			wantErr: errors.IndexOutOfBoundsError,
		},
		{
			name:    "移出第一个",
			sli:     []T{1, 2, 3, 4, 5, 6},
			index:   0,
			want:    wantSli([]T{2, 3, 4, 5, 6}, 6),
			wantOne: 1,
			wantErr: nil,
		},
		{
			name:    "移出最后一个",
			sli:     []T{1, 2, 3, 4, 5, 6},
			index:   5,
			want:    wantSli([]T{1, 2, 3, 4, 5}, 6),
			wantOne: 6,
			wantErr: nil,
		},
		{
			name:    "移出中间一个",
			sli:     []T{1, 2, 3, 4, 5, 6},
			index:   3,
			want:    wantSli([]T{1, 2, 3, 5, 6}, 6),
			wantOne: 4,
			wantErr: nil,
		},
		{
			name:    "容量在64-2048直接触发缩容",
			sli:     newSli(10, 65),
			index:   5,
			want:    wantSli([]T{0, 1, 2, 3, 4, 6, 7, 8, 9}, 40),
			wantOne: 5,
			wantErr: nil,
		},
		{
			name:    "容量在64-2048直接不触发缩容",
			sli:     newSli(17, 65),
			index:   5,
			want:    wantSli([]T{0, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, 65),
			wantOne: 5,
			wantErr: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, one, err := Delete(tc.sli, tc.index)
			assert.Equal(t, err, tc.wantErr)
			if err != nil {
				return
			}
			assert.Equal(t, res, tc.want)
			assert.Equal(t, one, tc.wantOne)

			assert.Equal(t, Cap(res), Cap(tc.want))
		})
	}
}

func newSli(l, c int) []T {
	sli := make([]T, 0, c)
	for i := 0; i < l; i++ {
		sli = append(sli, i)
	}
	return sli
}

func wantSli(sli []T, c int) []T {
	target := make([]T, Len(sli), c)
	copy(target, sli)
	return target
}

func TestShrink(t *testing.T) {
	// 64以内
	sli64 := make([]T, 0, 64)
	nsli64 := shrink(sli64)
	assert.Equal(t, sli64, nsli64)
	assert.Equal(t, Cap(sli64), Cap(nsli64))
	assert.Equal(t, Len(sli64), Len(nsli64))

	// 64到1024
	sliMid := make([]T, 0, 65)
	for i := 0; i < 10; i++ {
		sliMid = append(sliMid, i)
	}
	nsliMid := shrink(sliMid)
	assert.Equal(t, sli64, nsli64)
	assert.Equal(t, int(float64(Cap(sliMid))*0.625), Cap(nsliMid))
	assert.Equal(t, Len(sliMid), Len(nsliMid))

	sliMidNo := make([]T, 0, 65)
	for i := 0; i < 33; i++ {
		sliMidNo = append(sliMidNo, i)
	}
	nsliMidNo := shrink(sliMidNo)
	assert.Equal(t, sliMidNo, nsliMidNo)
	assert.Equal(t, Cap(sliMidNo), Cap(nsliMidNo))
	assert.Equal(t, Len(sliMidNo), Len(nsliMidNo))

	// 1024以上
	sliMore := make([]T, 0, 2049)
	for i := 0; i < 10; i++ {
		sliMore = append(sliMore, i)
	}
	nsliMore := shrink(sliMore)
	assert.Equal(t, sliMore, nsliMore)
	assert.Equal(t, Cap(sliMore)/2, Cap(nsliMore))
	assert.Equal(t, Len(sliMore), Len(nsliMore))

	sliMoreMo := make([]T, 0, 2050)
	for i := 0; i < 1026; i++ {
		sliMoreMo = append(sliMoreMo, i)
	}
	nsliMoreMo := shrink(sliMoreMo)
	assert.Equal(t, sliMoreMo, nsliMoreMo)
	assert.Equal(t, Cap(sliMoreMo), Cap(nsliMoreMo))
	assert.Equal(t, Len(sliMoreMo), Len(nsliMoreMo))
}
