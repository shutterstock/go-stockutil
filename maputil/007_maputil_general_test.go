package maputil

import (
    "testing"
    "strings"
)

func TestMapJoin(t *testing.T) {
    input := map[string]interface{}{
        `key1`: `value1`,
        `key2`: true,
        `key3`: 3,
    }

    output := Join(input, `=`, `&`)

    fmt.Println(output)

    if output == `` {
        t.Error("Output should not be empty")
    }

    if !strings.Contains(output, `key1=value1`) {
        t.Errorf("Output should contain '%s'", `key1=value1`)
    }

    if !strings.Contains(output, `key2=true`) {
        t.Errorf("Output should contain '%s'", `key2=true`)
    }

    if !strings.Contains(output, `key3=3`) {
        t.Errorf("Output should contain '%s'", `key3=3`)
    }
}

