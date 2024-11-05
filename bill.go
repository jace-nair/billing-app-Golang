package main

// Custom Struct Type acts as a blue print for any bill object
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Invoke this function to create and return a new bill object
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}
