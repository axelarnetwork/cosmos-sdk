package rosetta

import (
	"encoding/json"
	"time"

	crgerrs "github.com/cosmos/cosmos-sdk/server/rosetta/lib/errors"
)

// timeToMilliseconds converts time to milliseconds timestamp
func timeToMilliseconds(t time.Time) int64 {
	return t.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// unmarshalMetadata unmarshals the given meta to the target
func unmarshalMetadata(meta map[string]interface{}, target interface{}) error {
	b, err := json.Marshal(meta)
	if err != nil {
		return crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	err = json.Unmarshal(b, target)
	if err != nil {
		return crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	return nil
}

// marshalMetadata marshals the given interface to map[string]interface{}
func marshalMetadata(o interface{}) (meta map[string]interface{}, err error) {
	b, err := json.Marshal(o)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}
	meta = make(map[string]interface{})
	err = json.Unmarshal(b, &meta)
	if err != nil {
		return nil, err
	}

	return
}

func filter[T any](source []T, predicate func(T) bool) []T {
	var out []T

	for i := range source {
		if predicate(source[i]) {
			out = append(out, source[i])
		}
	}

	return out
}

func filterIndex[T any](source []T, predicate func(T) bool) []int {
	var out []int

	for i := range source {
		if predicate(source[i]) {
			out = append(out, i)
		}
	}

	return out
}

func and[T any](predicate ...func(T) bool) func(T) bool {
	return func(t T) bool {
		for i := range predicate {
			if !predicate[i](t) {
				return false
			}
		}

		return true
	}
}
