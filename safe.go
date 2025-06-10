package php2go

import "sync"

type GoTool struct {
	ch chan int
	wg sync.WaitGroup
}

func NewGoTool(num int) *GoTool {
	return &GoTool{
		ch: make(chan int, num),
		wg: sync.WaitGroup{},
	}
}

func (gt *GoTool) Add() {
	gt.ch <- 1
	gt.wg.Add(1)
}

func (gt *GoTool) Done() {
	<-gt.ch
	gt.wg.Done()
}

func (gt *GoTool) Wait() {
	gt.wg.Wait()
	close(gt.ch)
}

// SafeSlice 是一个线程安全的切片封装
type SafeSlice[T any] struct {
	mu    sync.Mutex
	slice []T
}

// NewSafeSlice 创建并返回一个新的 SafeSlice 实例
func NewSafeSlice[T any]() *SafeSlice[T] {
	return &SafeSlice[T]{
		slice: make([]T, 0),
	}
}

// Append 添加元素到切片中，确保线程安全
func (ss *SafeSlice[T]) Append(value ...T) {
	ss.mu.Lock()         // 加锁
	defer ss.mu.Unlock() // 函数结束时解锁
	ss.slice = append(ss.slice, value...)
}

// GetSlice 返回当前切片的副本，确保线程安全
func (ss *SafeSlice[T]) GetSlice() []T {
	ss.mu.Lock()         // 加锁
	defer ss.mu.Unlock() // 函数结束时解锁
	// 返回切片的副本以避免外部修改
	newSlice := make([]T, len(ss.slice))
	copy(newSlice, ss.slice)
	return newSlice
}

// SafeMap 是一个线程安全的Map封装，支持泛型
type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// NewSafeMap 创建并返回一个新的SafeMap实例
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// Set 设置键值对，确保线程安全
func (sm *SafeMap[K, V]) Set(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Get 获取值，确保线程安全
func (sm *SafeMap[K, V]) Get(key K) (V, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.data[key]
	return val, ok
}

// Delete 删除键值对，确保线程安全
func (sm *SafeMap[K, V]) Delete(key K) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.data, key)
}

// Len 返回Map长度，确保线程安全
func (sm *SafeMap[K, V]) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}

// GetMap 返回Map的副本，确保线程安全
func (sm *SafeMap[K, V]) GetMap() map[K]V {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	newMap := make(map[K]V, len(sm.data))
	for k, v := range sm.data {
		newMap[k] = v
	}
	return newMap
}
