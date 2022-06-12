package main

import (
	"context"
	"sync"
	"time"
)

// Bucket 统计桶
type Bucket struct {
	Success    int
	Failure    int
	Timeout    int
	Rejection  int
	Total      int
	LatestTime time.Time
}

// Summary 统计信息
type Summary struct {
	Success   int
	Failure   int
	Timeout   int
	Rejection int

	Total int
}

// Window 滑动窗口
type Window struct {
	ctx context.Context
	mu  sync.Mutex

	buckets    []*Bucket
	cap        int
	duration   time.Duration
	subscriber chan *Summary

	LatestTime time.Time
}

// NewSlidingWindow 返回滑动窗口，该窗口包含 numBuckets 个 Bucket，每个 Bucket 统计 duration 时长内的 Event
func NewSlidingWindow(ctx context.Context, duration int, numBuckets int) *Window {
	buckets := make([]*Bucket, numBuckets)
	for i := 0; i < numBuckets; i++ {
		buckets[i] = &Bucket{}
	}
	return &Window{
		buckets:    buckets,
		mu:         sync.Mutex{},
		cap:        numBuckets,
		duration:   time.Second * time.Duration(duration),
		subscriber: make(chan *Summary, numBuckets),
		ctx:        ctx,
	}
}

// Start 启动滑动窗口订阅事件流
func (w *Window) Start(events <-chan *Event) {
	ticker := time.NewTicker(w.duration)
	go func() {
		for {
			select {
			case <-w.ctx.Done():
				goto Label
			case <-ticker.C:
				go w.statistic()
			case event := <-events:
				now := time.Now()
				switch event.Status {
				case Success:
					go w.handleSuccess(now)
				case Failure:
					go w.handleFailure(now)
				case Timeout:
					go w.handleTimeout(now)
				case Rejection:
					go w.handleRejection(now)
				}
			}
		}
	Label:
		close(w.subscriber)
	}()
}

func (w *Window) Subscribe() <-chan *Summary {
	return w.subscriber
}

func (w *Window) handleSuccess(now time.Time) {
	w.mu.Lock()
	w.LatestTime = now
	w.getCurrentBucket(now).Success++
	w.getCurrentBucket(now).Total++
	w.mu.Unlock()
}

func (w *Window) handleFailure(now time.Time) {
	w.mu.Lock()
	w.LatestTime = now
	w.getCurrentBucket(now).Failure++
	w.getCurrentBucket(now).Total++
	w.mu.Unlock()
}

func (w *Window) handleTimeout(now time.Time) {
	w.mu.Lock()
	w.LatestTime = now
	w.getCurrentBucket(now).Timeout++
	w.getCurrentBucket(now).Total++
	w.mu.Unlock()
}

func (w *Window) handleRejection(now time.Time) {
	w.mu.Lock()
	w.LatestTime = now
	w.getCurrentBucket(now).Rejection++
	w.getCurrentBucket(now).Total++
	w.mu.Unlock()
}

func (w *Window) statistic() {
	summary := &Summary{
		Success:   0,
		Failure:   0,
		Timeout:   0,
		Rejection: 0,
		Total:     0,
	}
	w.mu.Lock()
	for _, bucket := range w.buckets {
		summary.Success += bucket.Success
		summary.Failure += bucket.Failure
		summary.Timeout += bucket.Timeout
		summary.Rejection += bucket.Rejection
		summary.Total += bucket.Total
	}
	w.mu.Unlock()
	w.subscriber <- summary
}

func (w *Window) getCurrentBucket(now time.Time) *Bucket {
	idx := now.Second() % w.cap
	currentBucket := w.buckets[idx]

	if currentBucket == nil {
		currentBucket = &Bucket{}
		w.buckets[idx] = currentBucket
	} else {
		if now.Unix() != currentBucket.LatestTime.Unix() {
			currentBucket.Reset()
		}
	}

	currentBucket.LatestTime = now
	return currentBucket
}

func (b *Bucket) Reset() {
	b.Success = 0
	b.Failure = 0
	b.Timeout = 0
	b.Rejection = 0
	b.Total = 0
}
