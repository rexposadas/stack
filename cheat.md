
x := 9

m := []string{"my","slice"}
for _, v := range m {
    fmt.Printf(v) // strings in slice
}

m := make(chan int)  // default, needs both producer and receiver
m := make(chan int, 3) // buffered
m<- 1 // add to channel
<-m   // fetch from channel
x := <-m // assign what you fetched
select {
    case x := <-m:      // get chan value then assign 
        fmt.Println(x)  // Print what was got from channel
}

m := make(chan int)
quit := make(chan int)
select{
    case <-m:
    case <-quit: 
        // getting something here means we need to quit something. 
        // This is a way to quit threads or apps.
}

reflect.DeepEqual(gotStruct, wantStruct)
ff, ok := w.(*os.File) // 2 result type avoids panic
i = int(mi) // convert to int
assert.True(t, ok)

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ { // always need this
        Add(i, i+1)
    }
}

go test -v -run NameOfTest // running a specific  test

go func() { // anonymous func as a go routine
}()         // the last () is so that the func is called. 


What is an interface: A type defined using a set of methods.
What is a struct: A user-defined type to store a collection of fields.
What is a receiver: a method bound to a type
What is reflection: the ability of a program to examine its own structure. This is done by the compiler leaving behind type information in the executable.
What is benchmarking: Go runs a op as many times as possible for a single second. 
