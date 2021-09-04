package main

import (
	"fmt"
	"Intl.NumberFormat"
)
func main() {
	var a = 2000;
	var options = { style: 'currency', currency: 'USD'};
	
	console.log(new Intl.NumberFormat('en-US', options).format(a));

}
