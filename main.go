package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
    fmt.Println(prompt)

    input, err := r.ReadString('\n')

    return strings.TrimSpace(input), err
}

func createNew() Bill {
    reader := bufio.NewReader(os.Stdin)
    name, _ := getInput("Create a new bill name", reader)
    b := newBill(name)

    fmt.Println("Created the bill -", b.name)

    return b
}

func printOptions(b Bill) {
    reader := bufio.NewReader(os.Stdin)
    opt, _ := getInput("Choose Option (a - Add item, s - Save bill, t - Add tip):", reader)

    switch opt {
    case "a":
        fmt.Println("You chose a")

        name, _ := getInput("Item name", reader)
        price, _ := getInput("Item price", reader)

        p, err := strconv.ParseFloat(price, 64)

        if err != nil {
            fmt.Println("Price must be a number")
            printOptions(b)
        }

        b.addItem(name, p)

        fmt.Println("Item added -", name, price)
        printOptions(b)
        
    case "t":
        fmt.Println("You chose t")
        
        tip, _ := getInput("Add Tip", reader)

        p, err := strconv.ParseFloat(tip, 64)

        if err != nil {
            fmt.Println("Tip must be a number")
            printOptions(b)
        }

        b.updateTip(p)

        fmt.Println("Tip add -", tip)
        printOptions(b)

    case "s":
        fmt.Println("You chose to save the bill")
        b.saveBill()

    default:
        fmt.Println("Not Valid option")
        printOptions(b)
    }
}

func main() {
	myBill := createNew()
    printOptions(myBill)
}