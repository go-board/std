package iterator

import (
	"github.com/go-board/std/optional"
)

type Iterator[T any] interface {
	Next() *optional.Optional[T]
}
