package maputil

import (
    "strings"
    "strconv"
    "sort"
    "github.com/shutterstock/go-stockutil/stringutil"
)


func StringKeys(input map[string]interface{}) []string {
    keys := make([]string, 0)

    for k, _ := range input {
        keys = append(keys, k)
    }

    return keys
}

// Take a flat (non-nested) map keyed with fields joined on fieldJoiner and return a
// deeply-nested map
//
func DiffuseMap(data map[string]interface{}, fieldJoiner string) (map[string]interface{}, error) {
    output     := make(map[string]interface{})

//  get the list of keys and sort them because order in a map is undefined
    dataKeys := StringKeys(data)
    sort.Strings(dataKeys)

//  for each data item
    for _, key := range dataKeys {
        value, _ := data[key]
        keyParts := strings.Split(key, fieldJoiner)

        output = DeepSet(output, keyParts, value).(map[string]interface{})
    }

    return output, nil
}


// func GetPathWithIndex(data *simplejson.Json, path []string) *simplejson.Json {
//     out := data

//     for _, part := range path {
//         if i, err := strconv.Atoi(part); err == nil {
//             out = out.GetIndex(i)
//         }else{
//             out = out.Get(part)
//         }
//     }

//     return out
// }



func DeepSet(data interface{}, path []string, value interface{}) interface{} {
    if len(path) == 0 {
        return data
    }

    var first = path[0]
    var rest    = make([]string, 0)

    if len(path) > 1 {
        rest = path[1:]
    }

//  Leaf Nodes
//    this is where the value we're setting actually gets set/appended
    if len(rest) == 0 {
        switch data.(type) {
        // ARRAY
        case []interface{}:
            return append(data.([]interface{}), value)

        // MAP
        case map[string]interface{}:
            dataMap := data.(map[string]interface{})
            dataMap[first] = value

            return dataMap
        }

    }else{
    //  Array Embedding
    //    this is where keys that are actually array indices get processed
    //  ================================
    //  is `first' numeric (an array index)
        if stringutil.IsInteger(rest[0]) {
            dataMap := data.(map[string]interface{})

        //  is the value at `first' in the map isn't present or isn't an array, create it
        //  -------->
            curVal, _ := dataMap[first]

            switch curVal.(type) {
            case []interface{}:
            default:
                dataMap[first] = make([]interface{}, 0)
                curVal, _ = dataMap[first]
            }
        //  <--------|


        //  recurse into our cool array and do awesome stuff with it
            dataMap[first] = DeepSet(curVal.([]interface{}), rest, value).([]interface{})
            return dataMap


    //  Intermediate Map Processing
    //    this is where branch nodes get created and populated via recursion
    //    depending on the data type of the input `data', non-existent maps
    //    will be created and either set to `data[first]' (the map)
    //    or appended to `data[first]' (the array)
    //  ================================
        }else{
            switch data.(type) {
        //  handle arrays of maps
            case []interface{}:
                dataArray := data.([]interface{})

                if curIndex, err := strconv.Atoi(first); err == nil {
                    if curIndex >= len(dataArray) {
                        for add := len(dataArray); add <= curIndex; add++ {
                            dataArray = append(dataArray, make(map[string]interface{}))
                        }
                    }

                    if curIndex < len(dataArray) {
                        dataArray[curIndex] = DeepSet(dataArray[curIndex], rest, value)
                        return dataArray
                    }
                }

        //  handle good old fashioned maps-of-maps
            case map[string]interface{}:
                dataMap := data.(map[string]interface{})

            //  is the value at `first' in the map isn't present or isn't a map, create it
            //  -------->
                curVal, _ := dataMap[first]

                switch curVal.(type) {
                case map[string]interface{}:
                default:
                    dataMap[first] = make(map[string]interface{})
                    curVal, _ = dataMap[first]
                }
            //  <--------|

                dataMap[first] = DeepSet(dataMap[first], rest, value)
                return dataMap
            }
        }
    }

    return data
}
