package LRUCache

import (
	"fmt"
	"testing"

	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLRUCache(t *testing.T) {
	Convey("test normal lru cache", t, func() {
		lruCache := Constructor(2)
		lruCache.Put(1, 1)
		lruCache.Put(2, 2)
		data := lruCache.Get(1)
		So(data, ShouldEqual, 1)
		lruCache.Put(3, 3)
		data = lruCache.Get(2)
		So(data, ShouldEqual, -1)
		lruCache.Put(4, 4)
		data = lruCache.Get(1)
		So(data, ShouldEqual, -1)
		data = lruCache.Get(3)
		So(data, ShouldEqual, 3)
		data = lruCache.Get(4)
		So(data, ShouldEqual, 4)
	})
}

func TestZeroLRUCache(t *testing.T) {
	Convey("Test zero capacity lru cache", t, func() {
		lruCache := Constructor(0)

		lruCache.Put(1, 1)
		actual := lruCache.Get(1)
		So(actual, ShouldEqual, -1)
	})
}

func TestNone(t *testing.T) {
	i := 0
	fmt.Println(time.Now().Unix())
	for i < 100000 {
		i++
		time.Parse("2006-01-02 15:04:05", time.Now().String())
	}
	fmt.Println(time.Now().Unix())
}
