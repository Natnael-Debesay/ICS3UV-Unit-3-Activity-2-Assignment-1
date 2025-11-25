// Author: Natnael Debesay
// Version: 1.0.0
// Date: 2025-11-24
// Fileoverview: This program rounds off numbers to a specific number of decimal places.

package main

import "fmt"

func main() {
// variables
const NUMBER1 float64 = 8.5467
const NUMBER2 float64 = 9.6382
const NUMBER3 float64 = 18.5146
const NUMBER4 float64 = 125.496

// output
fmt.Printf("Number 1 = %.3f\n", NUMBER1)
fmt.Printf("Number 2 = %.5f\n", NUMBER2)
fmt.Printf("Number 3 = %.1f\n", NUMBER3)
fmt.Printf("Number 4 = %.1f\n", NUMBER4)

fmt.Println("\nDone.")
}