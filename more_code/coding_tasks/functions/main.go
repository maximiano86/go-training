Coding Exercise #1

Create a function called cube() that takes a parameter of type float64 and returns the cube of that parameter (the parameter to the power of 3).

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/4UMwXWkYDyy.


Coding Exercise #2

Create a Go program with a function called f1() that takes a parameter of type uint and returns 2 values:

a) the factorial of n

b) the sum of all integer numbers greater than zero (>0) and less than or equal to n (<=n)

Test the program by calling the function.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/9UhNsDOxoEN.


Coding Exercise #3

Write a function called myFunc() that takes exactly one argument which is an int number written between double quotes (this is in fact a string). If the argument is integer 'n', the function should return the result of n + nn + nn

Example: myFunc('5') returns 5 + 55 + 555 which is 615 and myFunc('9') returns 9 + 99 + 999 which is 1107

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/EA7rzjXv_vl.


Coding Exercise #4

Create a function with the identifier sum that takes in a variadic parameter of type int and returns the sum of all values of type int passed in.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/3HsGVDvfNBB.


Coding Exercise #5

Change the function from the previous exercise and use a `naked return`.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/oF6_6ViX2pF.


Coding Exercise #6

Create a function called searchItem() that takes 2 parameters: a) a string slice and b) a string.

The function should search for the string (the second parameter) in the slice (the first parameter) and returns true if it finds the string in the slice and false otherwise. Do function does an case-sensitive search.

Call the function and see how it works.

Example: 

    animals := []string{"lion", "tiger", "bear"}
    result := searchItem(animals, "bear")
    fmt.Println(result) // => true
    result = searchItem(animals, "pig")
    fmt.Println(result)     // => false

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/3u9bzvzb_Cc.


Coding Exercise #7

Change the function from the previous exercise to do a case-insensitive search.

Example: 

    animals := []string{"Lion", "tiger", "bear"}
    result := searchItem(animals, "beaR")
    fmt.Println(result) // => true
    result = searchItem(animals, "lion")
    fmt.Println(result)     // => true

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/icJ_ovYXxCc.


Coding Exercise #8

Consider the following Go program that prints out:

The Go gopher is the iconic mascot of the Go project.

Hello, Go playground!

 

    package main
     
    import "fmt"
     
    func print(msg string) {
    	fmt.Println(msg)
    }
     
    func main() {
    	print("The Go gopher is the iconic mascot of the Go project.")
    	fmt.Println("Hello, Go playground!")
    }

Modify only the line in the main() body function where the print() function is invoked so that the program will print out Hello, Go playground! and then The Go gopher is the iconic mascot of the Go project.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/kGmrBDov_3B.


Coding Exercise #9

Create a function that takes in an int value and prints out that value.

Assign the function to a variable, print out the type of the variable and then call that function through the variable name.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/2bZEmgHQu3u.