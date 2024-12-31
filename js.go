package webgl

import (
	"fmt"
	"runtime"
	"syscall/js"
)

func TypedArrayOf(s interface{}) (result js.Value) {
	switch s := s.(type) {
	case []int8:
		result = js.Global().Get("Int8Array").New(len(s))
		for i, v := range s {
			result.SetIndex(i, v)
		}
	case []int16:
		result = js.Global().Get("Int16Array").New(len(s))
		for i, v := range s {
			result.SetIndex(i, v)
		}
	case []int32:
		result = js.Global().Get("Int32Array").New(len(s))
		for i, v := range s {
			result.SetIndex(i, v)
		}
	case []uint8:
		result = js.Global().Get("Uint8Array").New(len(s))
		js.CopyBytesToJS(result, s)
		runtime.KeepAlive(s)
	case []uint16:
		result = js.Global().Get("Uint16Array").New(len(s))
		for i, v := range s {
			result.SetIndex(i, v)
		}
	case []uint32:
		result = js.Global().Get("Uint32Array").New(len(s))
		for i, v := range s {
			result.SetIndex(i, v)
		}
	case []float32:
		result = js.Global().Get("Float32Array").New(len(s))
		for i, v := range s {
			result.SetIndex(i, v)
		}
	case []float64:
		result = js.Global().Get("Float64Array").New(len(s))
		for i, v := range s {
			result.SetIndex(i, v)
		}
	default:
		panic(fmt.Sprintf("jsutil: unexpected value at SliceToTypedArray: %T", s))
	}

	return
}
