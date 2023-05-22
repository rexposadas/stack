
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




go test -v -r NameOfTest // running a specific  test

go func() { // anonymous func as a go routine
}()         // the last () is so that the func is called. 