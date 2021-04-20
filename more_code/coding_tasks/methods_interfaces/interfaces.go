Coding Exercise #1

Consider the following type and interface declaration.

    type vehicle interface {
        License() string
        Name() string
    }
     
    type car struct {
        licenceNo string
        brand     string
    }

1. Create a Go program where car type implements the vehicle interface.

2. Create a variable of type vehicle that holds a car struct value.

3. Call the methods (Licence and Name) on the interface value declared at step 2

4. Run the program without errors.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/WrIXdWhqaS4.


Coding Exercise #2

1. Declare a variable called empty of type empty interface. Print out its type.

2. Assign an int value to the variable called empty. Print out its type.

3. Assign a float64 value to empty. Print out its type.

4. Assign an int slice value to empty. Print out its type.

5. Add a new int value to the slice (empty variable).

6. Print out the slice (empty variable).

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/0yLRzyhZpxM.


Coding Exercise #3

There is an error in the following Go program. Try to identify the error, change the code and run the program without errors.

    package main
     
    import "fmt"
     
    type cube struct {
        edge float64
    }
     
    func volume(c cube) float64 {
        return c.edge * c.edge * c.edge
    }
     
    func main() {
        var x interface{}
        x = cube{edge: 5}
        v := volume(x)
        fmt.Printf("Cube Volume: %v\n", v)
    }

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/-c4ws0uRGre.