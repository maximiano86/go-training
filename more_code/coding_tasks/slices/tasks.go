Coding Exercise #1

Using a composite literal declare and initialize a slice of type string with 3 elements.

Iterate over the slice and print each element in the slice and its index.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/CRPEvm-A31h.


Coding Exercise #2

There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

    package main
     
    import "fmt"
     
    func main() {
        mySlice := []float64{1.2, 5.6}
     
        mySlice[2] = 6
     
        a := 10
        mySlice[0] = a
     
        mySlice[3] = 10.10
     
        mySlice = append(mySlice, a)
     
        fmt.Println(mySlice)
     
    }

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/gVw5_PnrM-h.


Coding Exercise #3

1. Declare a slice called nums with three float64 numbers.

2. Append the value 10.1 to the slice

3. In one statement append to the slice the values: 4.1,  5.5 and 6.6

4. Print out the slice

5. Declare a slice called n with two float64 values

6. Append n to nums

7. Print out the nums slice

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/RNWABjmI2Le.


Coding Exercise #4

Create a Go program that reads some numbers from the command line and then calculates the sum and the product of all the numbers given at command line.

The user should give between 2 and 10 numbers.

Notes:

- the program should be run in a terminal (go run main.go) not in Go Playground

- example: 

go run main.go 3 2 5

Expected output: Sum: 10, Product: 30

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/HEaHKVj30qa.


Coding Exercise #5

Consider the following slice declaration: nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.

Print those elements and their sum.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/GbHEibG1t3p.


Coding Exercise #6

Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}

Using copy() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/wx9qjwBN-qC.


Coding Exercise #7

Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}

Using append() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/v_uzg5V_OvH.


Coding Exercise #8

Consider the following slice declaration:

 years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}

Using a slice expression and append() function create a new slice called newYears that contains the first 3 and the last 3 elements of the slice.  newYears should be []int{2000, 2001, 2002, 2008, 2009, 2010}

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/QyGtkDU_w_j.


