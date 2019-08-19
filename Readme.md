# golang + Try/Catch/Finally = GoTry

goTry (Golang for Try/Catch/Finally) is a Golang implementation for Exception handling.
You can see an extended doc in [godocs](https://godoc.org/github.com/fmorenovr/goTry).

## Try/Catch/Finally

This method is to Exception handling, in this case, when occurs a panic(), we can run a function instead that panic, maybe to solve that.  
But is not recommendable to use this library, because golang have an explicit error handling.  
See more info [here](https://golang.org/doc/effective_go.html).

## goTry

* First, You should download my library:

      go get github.com/fmorenovr/gotry/

* Then, you should use for differents implements in Go.
        
      gotry.Try(func() {
        gotry.Throw("New exception")
      }).Catch(func(e gotry.Exception) {
        // Print crash
        fmt.Println("Catch ---> exception catched #1:", e)
      }).Finally(func() {
        fmt.Println("Finally ---> This always print after all try block #1")
      })

* Could be a nested Try/Catch block

      gotry.Try(func() {
        gotry.Throw("New exception")
      }).Catch(func(e gotry.Exception) {
        // Print crash
        fmt.Println("Catch ---> exception catched #1:", e)
        gotry.Try(func(){
          gotry.Throw("New exception here!!")
        }).Catch(func(e gotry.Exception) {
          // Print crash
          fmt.Println("Catch ---> exception catched #2:", e)
        }).Finally(func() {
          fmt.Println("Finally ---> This always print after all try block #2")
        })
      }).Finally(func() {
        fmt.Println("Finally ---> This always print after all try block #1")
      })
