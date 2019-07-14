package utils

func AssignNotEmpty(src *string, dst *string) {
	if *src != "" {
		*dst = *src
	}
}

func AssignNotZero(src interface{}, dst interface{}) {
	switch v := dst.(type) {
	case *int:
		if *src.(*int) != 0 {
			*v = *src.(*int)
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
