package maputil

import (
    "testing"
    "fmt"
    _ "encoding/json"
)

func TestDeepSetString(t *testing.T) {
    input := make(map[string]interface{})
    testValue := "test-string"

    input = DeepSet(input, []string{"str"}, testValue).(map[string]interface{})

    if value, ok := input["str"]; !ok {
        t.Errorf("want key 'str' to exist, it does not")
    }else if value != testValue {
        t.Errorf("want 'str' == %q, got: %q", testValue, value)
    }
}


func TestDeepSetBool(t *testing.T) {
    input := make(map[string]interface{})
    testValue := true

    input = DeepSet(input, []string{"bool"}, testValue).(map[string]interface{})

    if value, ok := input["bool"]; !ok {
        t.Errorf("want key 'bool' to exist, it does not")
    }else if value != testValue {
        t.Errorf("want 'bool' == %s, got: %s", testValue, value)
    }
}


func TestDeepSetArray(t *testing.T) {
    input := make(map[string]interface{})
    testValues := []string{"first", "second"}

    for i, tv := range testValues {
        input = DeepSet(input, []string{"top-array", fmt.Sprint(i) }, tv).(map[string]interface{})
    }

    // input = DeepSet(input, []string{"top-array"}, 3.4).(map[string]interface{})

    if topArray, ok := input["top-array"]; !ok {
        t.Errorf("want key 'topArray' to exist, it does not")
    }else{
        switch topArray.(type) {
        case []interface{}:
            for i, val := range topArray.([]interface{}) {
                if val != testValues[i] {
                    t.Errorf("want v[%d] == %q, got: %q", i, testValues[i], val)
                }
            }
        default:
            t.Errorf("want topArray to be []string, got: %T", topArray)
        }
    }
}


func TestDeepSetNestedMapCreation(t *testing.T) {
    input := make(map[string]interface{})

    input = DeepSet(input, []string{"deeply", "nested", "map"}, true).(map[string]interface{})
    input = DeepSet(input, []string{"deeply", "nested", "count"}, 2).(map[string]interface{})

    if deeply, ok := input["deeply"]; !ok {
        t.Errorf("want key 'deeply' to exist, it does not")
    }else{
        deeplyMap := deeply.(map[string]interface{})

        if nested, ok := deeplyMap["nested"]; !ok {
            t.Errorf("want key 'deeply.nested' to exist, it does not")
        }else{
            nestedMap := nested.(map[string]interface{})

            if v, ok := nestedMap["map"]; !ok {
                t.Errorf("want key 'deeply.nested.map' to exist, it does not")
            }else if v != true {
                t.Errorf("want key 'deeply.nested.map' == true, got: %q", v)
            }

            if v, ok := nestedMap["count"]; !ok {
                t.Errorf("want key 'deeply.nested.count' to exist, it does not")
            }else if v != 2 {
                t.Errorf("want key 'deeply.nested.count' == 2, got: %q", v)
            }
        }
    }
}


func TestDiffuseMap(t *testing.T) {
    input := make(map[string]interface{})

    input["name"]                    = "test.thing.name"
    input["enabled"]                 = true
    input["cool.beans"]              = "yep"
    input["tags.0"]                  = "base"
    input["tags.1"]                  = "other"
    input["devices.0.name"]          = "lo"
    input["devices.1.name"]          = "eth0"
    input["devices.1.peers.0"]       = "0.0.0.0"
    input["devices.1.peers.1"]       = "1.1.1.1"
    input["devices.1.peers.2"]       = "2.2.2.2"
    input["devices.1.switch.0.name"] = "aa:bb:cc:dd:ee:ff"
    input["devices.1.switch.0.ip"]   = "111.222.0.1"
    input["devices.1.switch.1.name"] = "cc:dd:ee:ff:bb:dd"
    input["devices.1.switch.1.ip"]   = "111.222.0.2"


    if output, err := DiffuseMap(input, "."); err != nil {
        t.Errorf("Error diffusing map: %s", err)
    }else{
    //  name
        if v, _ := output["name"]; v != "test.thing.name" {
            t.Errorf("want 'name' == %q, got: %q", "test.thing.name", v)
        }

    //  enabled
        if v, _ := output["enabled"]; v != true {
            t.Errorf("want 'enabled' == %s, got: %q", true, v)
        }

    //  tags[]
        if v, ok := output["tags"]; !ok {
            t.Errorf("want 'tags' to exist, it does not")
        }else if l := len(v.([]interface{})); l != 2{
            t.Errorf("want 'tags' to have 2 elements, got: %d", l)
        }else{
            vArray := v.([]interface{})

            if vArray[0] != "base" {
               t.Errorf("want 'tags[0]' == %q, got: %q", "base", vArray[0])
            }else if vArray[1] != "other" {
                t.Errorf("want 'tags[1]' == %q, got: %q", "other", vArray[1])
            }
        }
    }
}
