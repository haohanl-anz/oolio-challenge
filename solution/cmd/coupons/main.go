package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	// Files containing coupon codes
	CouponBase1 = "../.coupons/couponbase1"
	CouponBase2 = "../.coupons/couponbase2"
	CouponBase3 = "../.coupons/couponbase3"
)

// The purpose of this script is to resolve and find the list of valid coupon codes.
// High level approach:
// 1. Create hashmap of coupon code as key, and count as int of number of files it appears in
// 2. For each file, read line by line, and output a hash set
// 3. At the end, filter out the coupon codes that have count larger or equal to 2
// 4. Export valid coupon codes to a file
func main() {
	fmt.Println("Running coupon resolver...")

	// Print out first 10 lines of each file
	var m = make(map[string]int)
	ingestFile(CouponBase1, m)
	ingestFile(CouponBase2, m)
	ingestFile(CouponBase3, m)

	// Create and open a new file to write valid coupons
	outputFile, err := os.Create("valid_coupons.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Check for count > 1 and write to file
	for k, v := range m {
		if v > 1 {
			_, err := outputFile.WriteString(fmt.Sprintf("%s\n", k))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("Coupon resolver completed. Valid coupons written to valid_coupons.txt")

}

func ingestFile(filename string, m map[string]int) map[string]int {
	// Set up line counter to track progress
	lc := 0
	// Create a hash set to track unique coupon codes
	uniqueCoupons := make(map[string]struct{})

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create scanner
	scanner := bufio.NewScanner(file)

	// Scan over file and count coupon
	for scanner.Scan() {
		line := scanner.Text()
		lc++
		uniqueCoupons[line] = struct{}{}

		// Print progress
		if lc%1000000 == 0 {
			fmt.Printf("Scanning %s at %d lines...\n", filename, lc)
		}
	}

	// Merge unique coupons into main map
	lc = 0
	for coupon := range uniqueCoupons {
		m[coupon]++
		lc++
		// Print progress
		if lc%1000000 == 0 {
			fmt.Printf("Saving coupons from %s at %d coupons...\n", filename, lc)
		}
	}

	return m
}
