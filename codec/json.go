package codec

import (
	"encoding/json"

	"github.com/go-board/std/result"
)

func MarshalJSON[T any](t T) result.Result[[]byte] {
	return result.Of(json.Marshal(t))
}

func UnmarshalJSON[T any](data []byte) result.Result[T] {
	var t T
	err := json.Unmarshal(data, &t)
	return result.Of(t, err)
}

func UnmarshalJSONPtr[T any](data []byte) result.Result[*T] {
	var t *T
	err := json.Unmarshal(data, &t)
	return result.Of(t, err)
}

func Jsonify[T any](data T) string {
	b, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(b)
}
