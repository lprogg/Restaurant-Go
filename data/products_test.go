package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Ion",
		Price: 1,
		SKU:   "a-ab-abc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
