package sync

import (
	"sync/atomic"
	"time"
)

type ChannelSender[T any] struct {
	inner  chan<- T
	closed *atomic.Bool
}

func (cs *ChannelSender[T]) Send(v T) bool {
	if cs.closed.Load() {
		return false
	}
	cs.inner <- v
	return true
}

func (cs *ChannelSender[T]) SendTimeout(v T, duration time.Duration) bool {
	if cs.closed.Load() {
		return false
	}
	select {
	case cs.inner <- v:
		return true
	case <-time.NewTimer(duration).C:
		return false
	}
}

func (cs *ChannelSender[T]) Close() {
	cs.closed.Store(true)
}

type ChannelReceiver[T any] struct {
	inner  <-chan T
	closed *atomic.Bool
}

func (r *ChannelReceiver[T]) Receive() (v T, ok bool) {
	if r.closed.Load() {
		return
	}
	v, ok = <-r.inner
	return
}

func (r *ChannelReceiver[T]) ReceiveTimeout(duration time.Duration) (v T, ok bool) {
	if r.closed.Load() {
		return
	}
	select {
	case v = <-r.inner:
		return v, true
	case <-time.NewTimer(duration).C:
		return
	}
}

func (r *ChannelReceiver[T]) Range(f func(T)) {
	for t := range r.inner {
		f(t)
	}
}

func BufferedChannel[T any](n int) (sender *ChannelSender[T], receiver *ChannelReceiver[T]) {
	channel := make(chan T, n)
	closed := &atomic.Bool{}
	return &ChannelSender[T]{inner: channel, closed: closed}, &ChannelReceiver[T]{inner: channel, closed: closed}
}

func UnbufferedChannel[T any]() (sender *ChannelSender[T], receiver *ChannelReceiver[T]) {
	return BufferedChannel[T](0)
}
