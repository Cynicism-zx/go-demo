package main

import (
	"testing"

	"github.com/duke-git/lancet/v2/slice"
	"github.com/emirpasic/gods/lists/arraylist"
	doublyLinkedList "github.com/emirpasic/gods/lists/doublylinkedlist"
	singleLinkedList "github.com/emirpasic/gods/lists/singlylinkedlist"
	"github.com/emirpasic/gods/utils"
)

// GoDS (Go Data Structures) - Sets, Lists, Stacks, Maps, Trees, Queues, and much more
// go数据结构包,包含了集合,列表,栈,映射,树,队列等等
// https://github.com/emirpasic/gods

func TestArrayList(t *testing.T) {
	list := arraylist.New()
	list.Add("a")      // ["a"]
	list.Add("c", "b") // ["a","c","b"]
	t.Log(list)

	list.Sort(utils.StringComparator) // ["a","b","c"]
	t.Log(list)

	_, _ = list.Get(0)                    // "a",true
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Swap(0, 1)                       // ["b","a",c"]
	list.Remove(2)                        // ["b","a"]
	list.Remove(1)                        // ["b"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	t.Log(list)

	_ = list.Empty()    // true
	_ = list.Size()     // 0
	list.Add("a")       // ["a"]
	list.Clear()        // []
	list.Insert(0, "b") // ["b"]
	list.Insert(0, "a") // ["a","b"]
	t.Log(list)
}

func TestSinglyLinkedList(t *testing.T) {
	list := singleLinkedList.New()
	list.Add("a")                     // ["a"]
	list.Add("c", "b")                // ["a","c","b"]
	list.Sort(utils.StringComparator) // ["a","b","c"]
	t.Log(list)

	_, _ = list.Get(0)                    // "a",true
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	t.Log(list)

	list.Swap(0, 1) // ["b","a",c"]
	list.Remove(2)  // ["b","a"]
	list.Remove(1)  // ["b"]
	list.Remove(0)  // []
	list.Remove(0)  // [] (ignored)
	t.Log(list)

	_ = list.Empty()    // true
	_ = list.Size()     // 0
	list.Add("a")       // ["a"]
	list.Clear()        // []
	list.Insert(0, "b") // ["b"]
	list.Insert(0, "a") // ["a","b"]
	t.Log(list)
}

func TestDoublyLinkedList(t *testing.T) {
	list := doublyLinkedList.New()
	list.Add("a")                     // ["a"]
	list.Add("c", "b")                // ["a","c","b"]
	list.Sort(utils.StringComparator) // ["a","b","c"]
	t.Log(list)

	_, _ = list.Get(0)                    // "a",true
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Swap(0, 1)                       // ["b","a",c"]
	t.Log(list)

	list.Remove(2)      // ["b","a"]
	list.Remove(1)      // ["b"]
	list.Remove(0)      // []
	list.Remove(0)      // [] (ignored)
	_ = list.Empty()    // true
	_ = list.Size()     // 0
	list.Add("a")       // ["a"]
	list.Clear()        // []
	list.Insert(0, "b") // ["b"]
	list.Insert(0, "a") // ["a","b"]
	t.Log(list)
}

func BenchmarkGodsListContains(b *testing.B) {
	list := arraylist.New()
	list.Add("a", "b", "c")
	for i := 0; i < b.N; i++ {
		list.Contains("a")
	}
	// BenchmarkGodsListContains-16      	96690255	        12.40 ns/op
	// Lancet is 1.6x faster than Gods
}

func BenchmarkLancetListContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice.Contain([]string{"a", "b", "c"}, "a")
	}
	// BenchmarkLancetListContains-16    	159258054	         7.533 ns/op
	// Lancet is 1.6x faster than Gods
}
