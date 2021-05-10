package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type (
	Product struct {
		Id          int     `json:"id"`
		Name        string  `json:"name" validate:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" validate:"gt=0"`
		SKU         string  `json:"sku" validate:"required,sku"`
	}

	Products []*Product
)

var (
	productList = Products{
		&Product{
			Id:          1,
			Name:        "Pulpe de pui cu cartofi prajiti",
			Description: "Meniu principal",
			Price:       2.5,
			SKU:         "abc123",
		},
		&Product{
			Id:          2,
			Name:        "Paste Carbonara",
			Description: "Paste",
			Price:       2.00,
			SKU:         "abcd1234",
		},
	}

	ErrorProductNotFound = fmt.Errorf("Product not found")
)

func GetProducts() Products { return productList }

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := regex.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	} else {
		return true
	}
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}

func GetNextId() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.Id + 1
}

func (p *Product) InsertProduct() {
	p.Id = GetNextId()
	productList = append(productList, p)
}

func FindProduct(id int) (*Product, int, error) {
	for i, product := range productList {
		if product.Id == id {
			return product, i, nil
		}
	}

	return nil, -1, ErrorProductNotFound
}

func (p *Product) PutProduct(id int) error {
	_, pos, err := FindProduct(id)

	if err != nil {
		return err
	}

	p.Id = id
	productList[pos] = p

	return nil
}
