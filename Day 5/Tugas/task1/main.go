package main

import (
	"fmt"
	"task1/product"
)

func main() {

	listProduct := product.GetListProducts()
	money := 100000
	remainedMoney := money

	// Tugas Week 1 - Soal Nomor 1 - Poin A
	paidProducts := listProduct.BuyProducts(&remainedMoney)

	fmt.Println("========================================")

	fmt.Println("Uang Yang Dibayarkan :", money)
	fmt.Println("Total Belanja :", money-remainedMoney)
	fmt.Println("Sisa Uang :", remainedMoney)

	fmt.Println("========================================")

	fmt.Println("Daftar Produk Yang Dibeli :")
	fmt.Println("----------------------------------------")

	for _, product := range paidProducts {
		fmt.Printf("%s | %d\n", product.Name, product.Price)
	}

	fmt.Println("========================================")

	price := 10000

	// Tugas Week 1 - Soal Nomor 1 - Poin C
	productInPrice := listProduct.GetProductsByPrice(price)

	fmt.Printf("Daftar Produk Dengan Harga %d :\n", price)
	fmt.Println("----------------------------------------")

	for _, product := range productInPrice {
		fmt.Printf("%s | %d\n", product.Name, product.Price)
	}

	fmt.Println("========================================")

	// Tugas Week 1 - Soal Nomor 1 - Poin B
	lowestPriceProduct := listProduct.GetLowestPriceProduct()
	highestPriceProduct := listProduct.GetHighestPriceProduct()

	fmt.Printf("Produk Termurah : %s | %d\n", lowestPriceProduct.Name, lowestPriceProduct.Price)
	fmt.Printf("Produk Termahal : %s | %d\n", highestPriceProduct.Name, highestPriceProduct.Price)

	fmt.Println("========================================")

}
