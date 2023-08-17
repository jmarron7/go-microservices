package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "testProd",
		Price: 1.25,
		SKU:   "abd-abdf-adsf",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
