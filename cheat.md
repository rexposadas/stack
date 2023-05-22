
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


go test -v -r NameOfTest // running a specific  test