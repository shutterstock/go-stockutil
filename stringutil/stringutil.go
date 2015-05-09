package stringutil

import (
    "fmt"
    "reflect"
    "strconv"
)


func IsInteger(in string) bool {
    if _, err := strconv.Atoi(in); err == nil {
        return true
    }

    return false
}


func IsFloat(in string) bool {
    if _, err := strconv.ParseFloat(in, 64); err == nil {
        return true
    }

    return false
}


func ToString(in interface{}) (string, error) {
    switch reflect.TypeOf(in).Kind() {
    case reflect.Int:
        return strconv.FormatInt(int64(in.(int)), 10), nil
    case reflect.Int8:
        return strconv.FormatInt(int64(in.(int8)), 10), nil
    case reflect.Int16:
        return strconv.FormatInt(int64(in.(int16)), 10), nil
    case reflect.Int32:
        return strconv.FormatInt(int64(in.(int32)), 10), nil
    case reflect.Int64:
        return strconv.FormatInt(in.(int64), 10), nil
    case reflect.Uint:
        return strconv.FormatUint(uint64(in.(uint)), 10), nil
    case reflect.Uint8:
        return strconv.FormatUint(uint64(in.(uint8)), 10), nil
    case reflect.Uint16:
        return strconv.FormatUint(uint64(in.(uint16)), 10), nil
    case reflect.Uint32:
        return strconv.FormatUint(uint64(in.(uint32)), 10), nil
    case reflect.Uint64:
        return strconv.FormatUint(in.(uint64), 10), nil
    case reflect.Float32:
        return strconv.FormatFloat(float64(in.(float32)), 'f', -1, 32), nil
    case reflect.Float64:
        return strconv.FormatFloat(in.(float64), 'f', -1, 64), nil
    case reflect.Bool:
        return strconv.FormatBool(in.(bool)), nil
    case reflect.String:
        return in.(string), nil
    default:
        return "", fmt.Errorf("Unable to convert type '%T' to string", in)
    }
}