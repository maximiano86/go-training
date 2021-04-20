Coding Exercise #1

1. Using the var keyword, declare a bidirectional unbuffered channel called c1 that works with values of type float64

2. Using the make() built-in function declare and initialize a receive-only channel called c2 and a send-only channel called c3. Both work with data of type rune.

3. Declare a bidirectional buffered channel  called c4 with a capacity of 10 ints.

4. Print out the type of all the channels declared.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/AAy5k4dp_3t.


Coding Exercise #2

Create a function literal (a.k.a. anonymous function) that sends the string value if receives as argument to main func using a channel.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/BJ5I1RlZemV.


Coding Exercise #3

There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

    package main
     
    import (
        "fmt"
    )
     
    func main() {
        c := make(<-chan int)
     
        go func(n int) {
            c <- n
        }(100)
     
        fmt.Println(<-c)
    }

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/n803fVTWSKi.


Coding Exercise #4

Create a goroutine named power() that has one parameter of type int, calculates the square value of its parameter and then sends  the result into a channel.

In the main function launch 50 goroutines that calculate the square values of all numbers between 1 and 50 included.

Print out the square values. 

A square(or raising to power 2) is the result of multiplying a number by itself. e.g., 25 is the square of 5.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/hcR7GUZlpn1.


Coding Exercise #5

Change the program from Exercise #4 and calculate the square of all values between 1 and 50 included using an anonymous function.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/NufTb2AaDhy.