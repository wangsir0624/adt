package adt

type Element struct {
    prev, next *Element
    
    Value  interface {}  
}

func (e *Element) Prev() *Element {
    return e.prev
}

func (e *Element) Next() *Element {
    return e.next
}

type List struct {
    root *Element
    
    length int
}

func NewList() *List {
    list := new(List)
    list.Reset()
    
    return list
}

func ToList(values []interface{}) *List {
    list := NewList()
    
    for _, v := range values {
        list.InsertAtBack(v)
    }
       
    return list
}

func (l *List) Reset() {
    l.root = new(Element)
    l.length = 0
}

func (l *List) First() *Element {
    return l.root.next
}

func (l *List) Last() *Element {
    return l.root.prev
}

func (l *List) Len() int {
    return l.length
}

func (l *List) IsEmpty() bool {
    return l.length == 0
}

func (l *List) InsertAfter(v interface{}, at *Element) *Element {
    if at == nil {
        return l.InsertAtBack(v)
    }
    
    e := new(Element)    
    e.Value = v
    
    return l.insertAfter(e, at)
}

func (l *List) InsertBefore(v interface{}, at *Element) *Element {
    if at == nil {
        return l.InsertAtFront(v)
    }
    
    e := new(Element)    
    e.Value = v
    
    return l.insertBefore(e, at)   
}

func (l *List) InsertAtFront(v interface{}) *Element {
    e := new(Element)    
    e.Value = v
    
    if l.length == 0 {
        l.root.prev = e
        l.root.next = e
    } else {
        tmp := l.root.next
        l.root.next = e
        e.next = tmp
        tmp.prev = e
    }
    l.length++
    
    return e
}

func (l *List) InsertAtBack(v interface{}) *Element {
    e := new(Element)
    e.Value = v
    
    if l.length == 0 {
        l.root.prev = e
        l.root.next = e
    } else {
        tmp := l.root.prev
        l.root.prev = e
        e.prev = tmp
        tmp.next = e
    }
    l.length++
    
    return e
}

func (l *List) insertAfter(e, at *Element) *Element {
    n := at.next
    at.next = e
    e.prev = at
    e.next = n
    if n != nil {
        n.prev = e
    }
    l.length++
    
    return e   
}

func (l *List) insertBefore(e, at *Element) *Element {
    n := at.prev
    at.prev = e
    e.next = at
    e.prev = n
    if n != nil {
        n.next = e
    }
    l.length++
    
    return e  
}

func (l *List) MoveAfter(e, at *Element) *Element {
    v := l.Remove(e)
    return l.InsertAfter(v, at)
}

func (l *List) MoveBefore(e, at *Element) *Element {
    v := l.Remove(e)
    return l.InsertBefore(v, at)
}

func (l *List) MoveToFront(e *Element) *Element {
    v := l.Remove(e)
    return l.InsertAtFront(v)
}

func (l *List) MoveToBack(e *Element) *Element {
    v := l.Remove(e)
    return l.InsertAtBack(v)
}

func (l *List) Remove(e *Element) interface{} {
    if e == l.root.next {
        l.root.next = e.next  
    } 
    if e == l.root.prev {
        l.root.prev = e.prev
    }
    if e.prev != nil {
        e.prev.next = e.next
    }
    if e.next != nil {
        e.next.prev = e.prev
    }
    l.length--
    
    return e.Value
}