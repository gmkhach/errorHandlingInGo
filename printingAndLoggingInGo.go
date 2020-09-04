package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    // Recover is another built-in function of the log package that helps you to regain control over the panicinc goroutine.
    // If a goroutin in panicin you can call recover to capture the value given to panic and continue normal execution.
    // Recover can only be used inside defered functions.
    defer func(){
        if r := recover(); r != nil {
            fmt.Println("Recovered")
        }
    }()

    // With log you can are not limited to printing out just to standard out. 
    // For example, you can log your errors into a log file.
    _, err := os.Open("thisFileDoesNotExist.txt")
    if err != nil {
        fmt.Println("Error:", err)
        
        // With log we get the same result as with fmt, but we also get the datetime stamp of the error.
        log.Println("Error:", err)
    }

    // Now let's create a log file and log our errors in that file using the log package.
    logFile, err := os.Create("log.txt")
    if err != nil {
        // We shouldn't use the log file to print this error yet because we are not sure if it's been successfully created yet.
        // So, instead we use a regular fmt.Println()
        fmt.Println(err)
    }
    // Always close the file imediately with a deferred call.
    defer logFile.Close()

    // Setting the default location of where log package should print.
    log.SetOutput(logFile)

    // Now let's try the same call again and this time try to print the error into the logFile.
    _, err = os.Open("thisFileDoesNotExist.txt")
    if err != nil {
        log.Println("Error:", err)

        // Panic stops normal execution of the current goroutine and the execution of the cuntion is terminated.
        // After it is returned to the caller of the function all of the defered functions are ran and then the execution of the caller is terminated.
        // This process condinues cascading up until all function in the current go routine have stopped, and the program is terminated.
        log.Panicln(err)
        // You can also invoke panic by simply calling it directly and passing the error to it. 
        // panic(err)

        // The Fatal functionj will write the error in the logFile and call an os.Exit(1), which terminates the programs without making any of the deferred calls.
        log.Fatalln("Fatal Error:", err, "Exiting the program.")
    }
}