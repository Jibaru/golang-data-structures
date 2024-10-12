package structures

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrVertexNotFound = errors.New("vertex not found")
	ErrEdgeNotFound   = errors.New("edge not found")
)

// NumericMaxValue returns the maximum possible value for a given numeric type W.
func NumericMaxValue[W Numeric]() W {
	var zero W

	switch v := any(zero).(type) {
	case int:
		var a int = math.MaxInt
		return W(a)
	case int8:
		var a int8 = math.MaxInt8
		return W(a)
	case int16:
		var a int16 = math.MaxInt16
		return W(a)
	case int32:
		var a int32 = math.MaxInt32
		return W(a)
	case int64:
		var a int64 = math.MaxInt64
		return W(a)
	case uint:
		var a uint = ^uint(0)
		return W(a) // Maximum uint value
	case uint8:
		var a uint8 = math.MaxUint8
		return W(a)
	case uint16:
		var a uint16 = math.MaxUint16
		return W(a)
	case uint32:
		var a uint32 = math.MaxUint32
		return W(a)
	case uint64:
		var a uint64 = ^uint64(0)
		return W(a) // Maximum uint64 value
	case float32:
		var a float32 = math.MaxFloat32
		return W(a)
	case float64:
		var a float64 = math.MaxFloat64
		return W(a)
	default:
		panic(fmt.Sprintf("unsupported type for maxValue: %T", v))
	}
}

// NumericZeroValue returns the zero value for the numeric type W.
func NumericZeroValue[W Numeric]() W {
	var zero W
	return zero
}

// reconstructPathOfGraph reconstructs the path from the 'previous' map.
func reconstructPathOfGraph[T comparable](previous map[T]*T, to T) []T {
	var path []T
	for at := &to; at != nil; at = previous[*at] {
		path = append([]T{*at}, path...)
	}
	return path
}
