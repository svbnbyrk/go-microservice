package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "nices",
		Price: 1.00,
		SKU:"as-dsa-dsa",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
