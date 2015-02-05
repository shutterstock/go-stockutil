package stringutil

import (
    "strconv"
)


func IsInteger(in string) bool {
    if _, err := strconv.Atoi(in); err == nil {
        return true
    }

    return false
}