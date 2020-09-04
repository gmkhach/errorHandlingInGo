package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
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

        // The Fatal function will write the error in the logFile and call an os.Exit(1), which terminates the programs without making any of the deferred calls.
        log.Fatalln("Fatal Error:", err, "Exiting the program.")
    }
}