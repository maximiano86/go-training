Coding Exercise #1

1. Create a new type called money. Its underlying type is float64.

2. Create a method called print that prints out the money value with exact 2 decimal points.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/Ft-MyMuwp8H.


Coding Exercise #2

Consider the money type declared at exercise #1. Create a new method for the money type called printStr that returns the money value as a string with 2 decimal points.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/ww34lKghjLR.


Coding Exercise #3

1. Create a new struct type called book with 2 fields: price and title of type float64 and string.

2. Create a method for the newly defined type called vat that returns the vat value if vat is 9%.

3. Create a method called discount that takes a book value as receiver and decreases the price of the book by 10%.

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/LriHS1yGcW_U.


Coding Exercise #4

A junior Go Programmer has just developed the following Go Program.   You want the change the books price by calling changePrice method. However, you notice that after calling the method the price is not changed.

You task is to change the code in order for the changePrice method to work as expected.

    package main
     
    import "fmt"
     
    type book struct {
        title string
        price float64
    }
     
    // method for book type
    func (b book) changePrice(p float64) {
        b.price = p
    }
     
    func main() {
        // book value
        bestBook := book{title: "The Trial by Kafka", price: 9.9}
     
        // changing the price
        bestBook.changePrice(11.99)
     
        fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change
    }

Are you stuck? Do you want to see the solution for this exercise? Click https://play.golang.org/p/w-3uhoIAU0l.