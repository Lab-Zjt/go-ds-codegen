# go-ds-codegen

## Introdution

Support: Vector, Stack, Queue, Deque, List, Priority Queue, Map

NOT recommend to use iterator because iterator may keep the object alive. If use, please call Illegal() method to avoid this problem.

## Usage

```
-t type name, like int, float, string and so on
-c container name, support: Vector, Stack, Queue, Deque, List, Priority Queue, Map
-g generated name, like IntVector, FloatList or what you like
-i iterator name, like ItInt, ItString or what you like
-k key type name, using in Map
-v value type name, using in Map
-ft filter function typename
-mp map function typename
-rd reduce function typename
-fe foreach function typename
```

`./codegen -t string -c PriorityQueue -g StringPQ >> ./StringPQ.go`
`./codegen -k int -v string -i Iterator -c Map -g ISmap >> ./ISmap.go`

## API

```
Vector: Size,Empty,Begin,End,PushBack,PopBack,PushFront,PopFront,Front,Back,At,Filter,Map,ForEach,Reduce
List: Size,Empty,PushBack,PopBack,PushFront,PopFront,Begin,End,Front,Back,Clear,Erase,Find,Filter,Map,ForEach,Reduce
Deque: Size,Empty,PushBack,PushFront,PopBack,PopFront,Clear
Queue: Size,Empty,Push,Pop,Front,Back,Clear
PriorityQueue: Size,Empty,Push,Pop,Top,Size,Clear
Stack: Size,Empty,Push,Pop,Top,Clear
Map: Size,Empty,Insert,Erase,Begin,End,At,Filter,Map,ForEach,Reduce,Clear
```

## Attention

Map, PriorityQueue need compare function `less($typename,$typename)bool`, you need to pass a function when you create a map or priority queue, or you can modify the constructor `New$container` of the generated file.

NOT recommend to use iterator because iterator may keep the object alive. If use, please call Illegal() method to avoid this problem or promise iterator will be destroyed immediately.

Maybe you shoule modify the package name of the generated file.
