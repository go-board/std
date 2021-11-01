package stream

type Stream[T any] struct {
	ch chan T
}

func New[T any]() *Stream[T] {
	return &Stream{ch: make(chan T)}
}

type streamState struct{}

func (s *Stream[T]) Sender() *Sender[T] {
	return &Sender[T]{ch: s.ch}
}

func (s *Stream[T]) Receiver() *Receiver[T] {
	return &Receiver[T]{ch: s.ch}
}

type Sender[T any] struct {
	ch    chan<- T
	state *streamState
}

func (s *Sender[T]) Send(t T) {
	s.ch <- t
}

func (s *Sender[T]) Close() {
	close(s.ch)
}

type Receiver[T any] struct {
	ch    <-chan T
	state *streamState
}

func (r *Receiver[T]) Closed() bool {
	return false
}

func (r *Receiver[T]) Receive() (T, bool) {
	t, ok := <-r.ch
	return t, ok
}
