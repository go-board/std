package adapters

import (
	"bufio"
	"io"

	"github.com/go-board/std/iterator"
	"github.com/go-board/std/optional"
	"github.com/go-board/std/result"
)

type LineReader struct {
	r *bufio.Scanner
}

func (i *LineReader) Next() optional.Optional[result.Result[[]byte]] {
	if i.r.Scan() {
		if i.r.Err() != nil {
			return optional.Some(result.Err[[]byte](i.r.Err()))
		}
		return optional.Some(result.Ok(i.r.Bytes()))
	}
	return optional.None[result.Result[[]byte]]()
}

var _ iterator.Iterator[result.Result[[]byte]] = (*LineReader)(nil)

func NewLineReader(r io.Reader) *LineReader {
	return &LineReader{r: bufio.NewScanner(r)}
}
