
```go
m := []string{"my","slice"}
for _, v := range m {
    fmt.Printf(v) // strings in slice
}
```

```go
m := make(chan int)  // default, needs both producer and receiver
m := make(chan int, 3) // buffered
m<- 1 // add to channel
<-m   // fetch from channel
x := <-m // assign what you fetched
select {
    case x := <-m:      // get chan value then assign 
        fmt.Println(x)  // Print what was got from channel
}
```

```go
m := make(chan int)
quit := make(chan int)
select{
    case <-m:
    case <-quit: 
        // getting something here means we need to quit something. 
        // This is a way to quit threads or apps.
}
```

reflect.DeepEqual(gotStruct, wantStruct)
ff, ok := w.(*os.File) // 2 result type avoids panic
i = int(mi) // convert to int
assert.True(t, ok)

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ { // always need this
        Add(i, i+1)
    }
}
```go
go test -v -run NameOfTest // running a specific  test
```

```go
    ticker := time.NewTicker(1 * time.Second)
	<- ticker.C  // tick every second
```

```go
i := 42

j := float64(i)
```

```go
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)
```

```go
switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	}
```


What is an interface: A type defined using a set of methods.
What is a struct: A user-defined type to store a collection of fields.
What is a receiver: a method bound to a type
What is reflection: the ability of a program to examine its own structure. This is done by the compiler leaving behind type information in the executable.
What is benchmarking: Go runs a operation as many times as possible for a single second. 
