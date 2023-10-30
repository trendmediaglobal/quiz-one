package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductPurchase(t *testing.T) {
	expect := Purchase{
		Products: []Product{
			{Name: "Orbit Star Z2", Price: 449000},
			{Name: "Orbit Star G1", Price: 399000},
		},
		PaymentMethod: "CC",
		GrandTotal:    763200,
	}

	_instance := Purchase{
		Products:      []Product{},
		PaymentMethod: "",
		GrandTotal:    0,
	}

	t.Run("Set Product", func(t *testing.T) {
		_instance.SetProduct(Product{Name: "Orbit Star Z2", Price: 449000})
		_instance.SetProduct(Product{Name: "Orbit Star G1", Price: 399000})
		assert.Equal(t, expect.Products, _instance.Products)
	})

	t.Run("Set Payment Method", func(t *testing.T) {
		_instance.SetPaymentMethod("CC")
		assert.Equal(t, expect.PaymentMethod, _instance.PaymentMethod)
	})

	t.Run("Calculate Grand Total", func(t *testing.T) {
		_instance.CalculateGrandTotal()
		assert.Equal(t, expect.GrandTotal, _instance.GrandTotal)
	})
}
