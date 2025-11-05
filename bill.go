package main

import "fmt"

type Bill struct {
	name string
	items map[string]float64
	tip float64
}

func newBill(name string) Bill {
	b := Bill {
		name: name,
		items: map[string]float64{},
		tip: 0,
	}

	return b
}

func (b Bill) format() string {
	fs := "Bill Breakdown \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%v ... %v \n", k + ":", v )
		total += v
	}

	fs += fmt.Sprintf("%v ... $%0.2f \n", "Tip:", b.tip)
	fs += fmt.Sprintf("%v ... $%0.2f \n", "Total:", total+b.tip)

	return fs
}

func (b *Bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *Bill) addItem(name string, price float64) {
	b.items[name] = price
}