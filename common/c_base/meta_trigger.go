package c_base

var IsNotZero = func(value any) bool {
	return value.(int) != 0
}

var IsZero = func(value any) bool {
	return value.(int) == 0
}
