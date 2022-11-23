package cast

// func name ref https://github.com/spf13/cast/blob/master/cast.go
// but we do NOT copy data

// ToStringMap casts an interface to a map[string]interface{} type.
func ToStringMap(i interface{}) map[string]interface{} {
	if m, ok := i.(map[string]interface{}); ok {
		return m
	}
	logger.Printf("unable to cast %#v of type %T to map[string]interface{}", i, i)
	return map[string]interface{}{}
}

func ToSlice(i interface{}) []interface{} {
	if m, ok := i.([]interface{}); ok {
		return m
	}
	logger.Printf("unable to cast %#v of type %T to []interface{}", i, i)
	return []interface{}{}
}

func ToArray(i interface{}, size int) []interface{} {
	if m, ok := i.([]interface{}); ok {
		if len(m) == size {
			return m
		}
	}
	logger.Printf("unable to cast %#v of type %T to []interface{}", i, i)
	return make([]interface{}, size)
}

func ToFloat64(i interface{}) float64 {
	if m, ok := i.(float64); ok {
		return m
	}
	logger.Printf("unable to cast %#v of type %T to float64", i, i)
	return 0.0
}

func ToString(i interface{}) string {
	if m, ok := i.(string); ok {
		return m
	}
	logger.Printf("unable to cast %#v of type %T to string", i, i)
	return ""
}

// ToInt64 in golang, mongodb driver only use int32 and int64 for int
func ToInt64(i interface{}) int64 {
	if v, ok := i.(int64); ok {
		return v
	}
	if v, ok := i.(int32); ok {
		return int64(v)
	}
	logger.Printf("unable to cast %#v of type %T to int64/int32", i, i)
	return 0
}
