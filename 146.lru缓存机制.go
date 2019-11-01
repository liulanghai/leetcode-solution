/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU缓存机制
 */

// @lc code=start
package cn

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	data     map[int]*list.Element
	l        *list.List
	capacity int
	Mu       sync.Mutex
}
type node struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	l := list.New()
	var cache LRUCache
	cache.data = make(map[int]*list.Element)
	cache.l = l
	cache.capacity = capacity
	return cache
}

func (this *LRUCache) Get(key int) int {
	this.Mu.Lock()
	defer this.Mu.Unlock()
	val, ok := this.data[key]
	if ok {
		this.l.MoveToFront(val)
		n := val.Value.(node)
		return n.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	this.Mu.Lock()
	defer this.Mu.Unlock()

	val, ok := this.data[key]
	if ok {
		var n node
		n.key = key
		n.val = value
		val.Value = n
		this.l.MoveToFront(val)
		return
	}
	var n node
	n.key = key
	n.val = value
	e := this.l.PushFront(n)
	this.data[key] = e
	if this.l.Len() > this.capacity { //删除一个
		last := this.l.Back()
		v := last.Value.(node)
		delete(this.data, v.key)
		this.l.Remove(last)
	}
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
