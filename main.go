package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

type Purchase struct {
	Products      []Product
	PaymentMethod string
	GrandTotal    float64
}

func (p *Purchase) SetProduct(product Product) {
	p.Products = append(p.Products, product)
}

func (p *Purchase) SetPaymentMethod(paymentMethod string) {
	p.PaymentMethod = paymentMethod
}

func (p *Purchase) CalculateGrandTotal() {
	total := 0.0

	for _, product := range p.Products {
		total += product.Price
	}

	var discount float64

	if p.PaymentMethod == "CC" {
		discount = 0.1
	} else if p.PaymentMethod == "VA" {
		discount = 0.05
	}

	p.GrandTotal = total - (total * discount)
}

func main() {
	purchase := Purchase{
		Products:      []Product{},
		PaymentMethod: "",
		GrandTotal:    0,
	}

	purchase.SetProduct(Product{Name: "Orbit Star Z2", Price: 449000})
	purchase.SetProduct(Product{Name: "Orbit Star G1", Price: 399000})
	purchase.SetPaymentMethod("CC")
	purchase.CalculateGrandTotal()

	fmt.Printf("Grand Total: %.f\n\n", purchase.GrandTotal)
	fmt.Println("Product Purchases :")

	for _, product := range purchase.Products {
		fmt.Printf("Name: %s, Price: %.f\n", product.Name, product.Price)
	}
}
