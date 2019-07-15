package utils

func AssignNotEmpty(src *string, dst *string) {
	if *src != "" {
		*dst = *src
	}
}

func AssignNotZero(src interface{}, dst interface{}) {
	switch v := dst.(type) {
	case *int32:
		if *src.(*int32) != 0 {
			*v = *src.(*int32)
		}
	case *float32:
		if *src.(*float32) != 0 {
			*v = *src.(*float32)
		}
	case *float64:
		if *src.(*float64) != 0 {
			*v = *src.(*float64)
		}
	default:
		panic("wrong type")
	}
}
