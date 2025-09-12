package c_base

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Value interface {
	~bool | ~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

type AnyPointer interface {
	~*any | ~*bool | ~*string |
		~*int | ~*int8 | ~*int16 | ~*int32 | ~*int64 |
		~*uint | ~*uint8 | ~*uint16 | ~*uint32 | ~*uint64 |
		~*float32 | ~*float64 | ~*complex64 | ~*complex128
}
