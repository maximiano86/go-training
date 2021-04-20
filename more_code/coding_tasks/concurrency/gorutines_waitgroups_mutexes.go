Coding Exercise #1

There is an error in the following Go Program. Even though the goroutine is correctly launched, it doesnt print any message.

     package main
     
    import (
        "fmt"
    )
     
    func sayHello(n string) {
        fmt.Printf("Hello, %s!\n", n)
    }
     
    func main() {
        go sayHello("Mr. Wick")
    }

Your task is to synchronize main and the goroutine using WaitGroups. The program should print the string received as argument by sayHello().

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/UfQrVmlwrvS.


Coding Exercise #2

1. Create a function called sum() that calculates and then prints out  the sum of 2 float numbers it receives as arguments.

Format the result with 2 decimal points.

2. From main launch 3 goroutines that execute the function you have just created (sum)

3. Synchronize the goroutines and the main function using WaitGroups

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/2ArrVTSj55g.


Coding Exercise #3

1. Create an anonymous function that calculates and prints out the square root of a float value it receives as argument.

2. Launch the function as a goroutine and synchronize it with main using WaitGroups

Note: You calculate the square root of a float named f using the Sqrt() function from math package like this:

    x := math.Sqrt(f)
    fmt.Println(x)

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/4Z_oVsvLyk1.


Coding Exercise #4

Change the code from Exercise #3 and launch 50 goroutines that calculate concurrently the square root of all the numbers between 100 and 149 (both included).

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/W0bUu1wD9Df.


Coding Exercise #5

You work at a Banking Application and have created 2 functions: one that deposits a value into an account and another that withdraws a value from the account.

You want to simulate many deposits and withdraws that take place simultaneously  and start some goroutines.

During testing you notice that a date race occurred.

Your task is to change the code in order to protect the accounts balance from simultaneously writing using a mutex.

This is the initial program that has errors:

    package main
     
    import (
        "fmt"
        "sync"
    )
     
    func deposit(b *int, n int, wg *sync.WaitGroup) {
        *b += n
        wg.Done()
    }
     
    func withdraw(b *int, n int, wg *sync.WaitGroup) {
        *b -= n
        wg.Done()
    }
     
    func main() {
        var wg sync.WaitGroup
     
        wg.Add(200)
     
        balance := 100
     
        for i := 0; i < 100; i++ {
            go deposit(&balance, i, &wg)
            go withdraw(&balance, i, &wg)
        }
        wg.Wait()
     
        fmt.Println("Final balance value:", balance)
    }

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/R6Cn4td3Z_R.