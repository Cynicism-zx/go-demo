package lo

import (
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

// Path: https://github.com/samber/lo

//唯一过滤
func TestUniq(t *testing.T) {
	names := lo.Uniq[string]([]string{"Samuel", "John", "Samuel"})
	t.Log(names)
}

//乱排
func TestShuffle(t *testing.T) {
	re := lo.Shuffle([]int{10, 2, 3, 15, 23})
	t.Log(re)
}

//条件过滤
func TestFilter(t *testing.T) {
	res := lo.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(x int, index int) bool {
		return x > 5
	})
	t.Log(res)
}

//类型转换
func TestMap(t *testing.T) {
	re := lo.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, index int) string {
		return strconv.FormatInt(x, 10)
	})
	t.Log(re)
}

// 条件分组
func TestGroupBy(t *testing.T) {
	re := lo.GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	t.Log(re)
}

func TestPartitionBy(t *testing.T) {
	oddEven := func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	}

	re := lo.PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, oddEven)
	// 异步回调
	re1 := lop.PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, oddEven)
	t.Log(re)
	t.Log(re1)
}

func TestForEach(t *testing.T) {
	var counter uint64
	collection := []int{1, 2, 3, 4}
	lo.ForEach(collection, func(x int, index int) {
		atomic.AddUint64(&counter, 1)
	})
	t.Log(counter)
}

// 集合简化为单个值
func TestReduce(t *testing.T) {
	re := lo.Reduce([]int{1, 2, 3, 4, 7, 8}, func(x int, y int, index int) int {
		return x + y
	}, 0)
	t.Log(re)
}

func TestReduceRight(t *testing.T) {
	re := lo.ReduceRight([][]int{{0, 1}, {2, 3}, {4, 5}}, func(agg []int, item []int, _ int) []int {
		return append(agg, item...)
	}, []int{})
	t.Log(re)
}

// 条件唯一
func TestUniqBy(t *testing.T) {
	re := lo.UniqBy[int, int]([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	t.Log(re)
}

// 返回一个元素数组，元素被分成长度为size的组
func TestChunk(t *testing.T) {
	re := lo.Chunk[int]([]int{0, 1, 2, 3, 4, 5, 6}, 2)
	t.Log(re)
}

func TestInterleave(t *testing.T) {
	re := lo.Interleave[int]([]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9})
	t.Log(re)
}

// 反转数组
func TestReverse(t *testing.T) {
	re := lo.Reverse[int]([]int{0, 1, 2, 3, 4, 5})
	t.Log(re)
}

type foo struct {
	bar string
}

func (f foo) Clone() foo {
	return foo{f.bar}
}

// 填充默认值
func TestFill(t *testing.T) {
	re := lo.Fill[foo]([]foo{foo{"a"}, foo{"a"}}, foo{"b"})
	t.Logf("%+v", re)
}
