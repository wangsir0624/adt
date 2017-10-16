package adt

import (
    "testing"
    "fmt"
)

func TestToList(t *testing.T) {
    data := []interface{}{2, 3, 6, 1, 8}
    length := len(data)
    list := ToList(data)
    if list.Len() != length {
        t.Error("ToList() failed\r\n")
        t.SkipNow()
    }
    i := 0
    for tmp := list.First(); tmp != nil; tmp = tmp.Next() {
        if tmp.Value != data[i] {
            t.Error("ToList() failed\r\n")
            t.SkipNow()
        }
        
        i++
    }
}

func ExampleList() {
    list := NewList()
    fmt.Println(list.First())
    list.InsertAfter(4, nil)
    //output: <nil>
}