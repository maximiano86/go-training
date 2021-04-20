Coding Exercise #1

1. Using the var keyword declare a string called name and initialize it with your name.

2. Using short declaration syntax declare a string called country and assign the country you are living in to the string variable.

3. Print the following string on multiple lines like this:

Your name: `here the value of name variable`

Country: `here the value of country variable`

4. Print out the following strings:

a) He says: "Hello"

b) C:\Users\a.txt

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/StcNEb5UaLr.


Coding Exercise #2

1. Declare a rune called r that stores the non-ascii letter ă

2. Print out the type of r

3. Declare 2 strings that contains the values ma and m

4. Concatenate the strings and the rune in a new string (the new string will hold the value mamă ).

BTW: mamă means mother in Romanian.

Note: You should convert rune to string to contatenate it to another string.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/4tExzsxi1h4.


Coding Exercise #3

Consider the following string declaration: s1 := "țară means country in Romanian"

1. Iterate over the string and print out byte by byte

2. Iterate  over the string and print out rune by rune and the byte index where the rune starts in the string

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/nI9xPzYc2HQ.


Coding Exercise #4

There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

    package main
     
    import "fmt"
     
    func main() {
    	s1 := "Go is cool!"
    	x := s1[0]
    	fmt.Println(x)
     
    	s1[0] = 'x'
     
    	// printing the number of runes in the string
    	fmt.Println(len(s1))
    }

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/QLmPAfn8ylx.


Coding Exercise #5

Consider the following string declaration:s := "你好 Go!"

1. Convert the string to a byte slice.

2. Print out the byte slice

3. Iterate over the byte slice and print out each index and byte in the byte slice

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/XxTCxSxDdmT.


Coding Exercise #6

Consider the following string declaration:s := "你好 Go!"

1. Convert the string to a rune slice.

2. Print out the rune slice

3. Iterate over the rune slice and print out each index and rune in the rune slice

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/Y9F0VXJhg9j.