package maputil

import (
    "testing"
)

func TestDeepGetScalar(t *testing.T) {
    input := make(map[string]interface{})
    
    input = DeepSet(input, []string{"deeply", "nested", "value"}, 1.4).(map[string]interface{})

    if v := DeepGet(input, []string{"deeply", "nested", "value"}, nil); v == nil {  
        t.Errorf("%s\n", v)
    }
}

func TestDeepGetArrayElement(t *testing.T) {
    input := make(map[string]interface{})
    
    input = DeepSet(input, []string{"tags", "0"}, "base").(map[string]interface{})
    input = DeepSet(input, []string{"tags", "1"}, "other").(map[string]interface{})

    if v := DeepGet(input, []string{"tags", "0"}, nil); v != "base" {  
        t.Errorf("%s\n", v)
    }

    if v := DeepGet(input, []string{"tags", "1"}, nil); v != "other" {
        t.Errorf("%s\n", v)
    }
}


func TestDeepGetMapKeyInArray(t *testing.T) {
    input := make(map[string]interface{})
    
    input = DeepSet(input, []string{"devices", "0", "name"}, "lo").(map[string]interface{})
    input = DeepSet(input, []string{"devices", "1", "name"}, "eth0").(map[string]interface{})

    if v := DeepGet(input, []string{"devices", "0", "name"}, nil); v != "lo" {  
        t.Errorf("%s\n", v)
    }

    if v := DeepGet(input, []string{"devices", "1", "name"}, nil); v != "eth0" {
        t.Errorf("%s\n", v)
    }
}



func TestDeepGetMapKeyInDeepArray(t *testing.T) {
    input := make(map[string]interface{})
    
    input = DeepSet(input, []string{"devices", "0", "switch", "peers", "0"}, "0.0.0.0").(map[string]interface{})
    input = DeepSet(input, []string{"devices", "0", "switch", "peers", "1"}, "0.0.1.1").(map[string]interface{})
    input = DeepSet(input, []string{"devices", "1", "switch", "peers", "0"}, "1.1.0.0").(map[string]interface{})
    input = DeepSet(input, []string{"devices", "1", "switch", "peers", "1"}, "1.1.1.1").(map[string]interface{})

    if v := DeepGet(input, []string{"devices", "0", "switch", "peers", "0"}, nil); v != "0.0.0.0" {
        t.Errorf("%s\n", v)
    }

    if v := DeepGet(input, []string{"devices", "0", "switch", "peers", "1"}, nil); v != "0.0.1.1" {
        t.Errorf("%s\n", v)
    }

    if v := DeepGet(input, []string{"devices", "1", "switch", "peers", "0"}, nil); v != "1.1.0.0" {
        t.Errorf("%s\n", v)
    }

    if v := DeepGet(input, []string{"devices", "1", "switch", "peers", "1"}, nil); v != "1.1.1.1" {
        t.Errorf("%s\n", v)
    }
}