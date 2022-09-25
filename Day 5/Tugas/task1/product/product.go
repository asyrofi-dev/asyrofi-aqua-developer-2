package product

import (
	"sort"
)

type Product struct {
	Id    int
	Name  string
	Price int
}

type Products []Product

func GetListProducts() Products {
	return listProduct
}

func (products Products) SortProductByPrice() Products {

	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})

	return products
}

func (products Products) BuyProducts(remainedMoney *int) Products {
	var result Products

	products.SortProductByPrice()

	for _, product := range listProduct {
		if *remainedMoney >= product.Price {
			*remainedMoney -= product.Price
			result = append(result, product)
		} else {
			break
		}
	}

	return result
}

func (products Products) GetHighestPriceProduct() Product {
	result := products[0]

	for i := 1; i < len(products); i++ {
		if products[i].Price > result.Price {
			result = products[i]
		}
	}

	return result
}

func (products Products) GetLowestPriceProduct() Product {
	result := products[0]

	for i := 1; i < len(products); i++ {
		if products[i].Price < result.Price {
			result = products[i]
		}
	}

	return result
}

func (products Products) GetProductsByPrice(price int) Products {
	var result Products

	for _, product := range products {
		if product.Price == price {
			result = append(result, product)
		}
	}

	return result
}

var listProduct = Products{
	{
		Id:    1,
		Name:  "Benih Lele",
		Price: 50000,
	},
	{
		Id:    2,
		Name:  "Pakan Lele Cap Menara",
		Price: 25000,
	},
	{
		Id:    3,
		Name:  "Probiotik A",
		Price: 75000,
	},
	{
		Id:    4,
		Name:  "Probiotik Nila B",
		Price: 10000,
	},
	{
		Id:    5,
		Name:  "Pakan Nila",
		Price: 20000,
	},
	{
		Id:    6,
		Name:  "Benih Nila",
		Price: 20000,
	},
	{
		Id:    7,
		Name:  "Cupang",
		Price: 5000,
	},
	{
		Id:    8,
		Name:  "Benih Nila",
		Price: 30000,
	},
	{
		Id:    9,
		Name:  "Benih Cupang",
		Price: 10000,
	},
	{
		Id:    10,
		Name:  "Probiotik B",
		Price: 10000,
	},
}
