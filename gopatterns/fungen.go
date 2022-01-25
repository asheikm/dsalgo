package main
import "fmt"

func main() {
    numofReceiveGoRoutines := 2 // configurable
    funCh := make(chan func())
    
    // receive funs and exxecute
    for i := 0; i < numofReceiveGoRoutines; i++ {
        go receiveFuncs(funCh)
    }
    
    funcGenerator(funCh, sendFunction)
}

func sendFunction() {
    fmt.Println("welcome")
}


// Function generator or provider which sends funs 
func funcGenerator(funCh chan func(), f func()) {
    for {
        funCh  <-  f
    }   
}

// func exxecutor
func receiveFuncs(funCh chan func()) {
        for {
            e := <-funCh
            e()
        }
}

// I have a input provider that sends me funcs that needs to be executed.
// My system is limited and low on hardware so number of cores is 2
// this number can be provided from outside.
