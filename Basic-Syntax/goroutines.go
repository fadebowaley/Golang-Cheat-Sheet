// Use a goroutine to run a function asynchronously
go myFunction()

// Use a channel to send and receive data between goroutines
ch := make(chan int)
go func() {
    ch <- 5
}()
x := <-ch
fmt.Println(x)