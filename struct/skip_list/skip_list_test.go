package skip_list_test

import (
	"testing"

	"github.com/0RAJA/Rutils/struct/skip_list"
	"github.com/stretchr/testify/require"
)

type M int

func (m M) CompareTo(comparable skip_list.Comparable) int {
	c, _ := comparable.(M)
	return int(m - c)
}

func TestNewSkipListFetMaxMin(t *testing.T) {
	sList := skip_list.NewSkipList()
	for i := 1; i <= 10; i++ {
		sList.Put(skip_list.NewKV(M(i), i))
	}
	min := sList.GetMin()
	require.NotEmpty(t, min)
	require.Equal(t, *min, skip_list.KV{K: M(1), V: 1})
	max := sList.GetMax()
	require.NotEmpty(t, max)
	require.Equal(t, *max, skip_list.KV{K: M(10), V: 10})
	result, ok := sList.Get(M(0))
	require.False(t, ok)
	require.NotEmpty(t, result)
	require.Equal(t, *result, skip_list.KV{K: M(1), V: 1})
	result, ok = sList.Get(M(11))
	require.False(t, ok)
	require.Empty(t, result)
}
