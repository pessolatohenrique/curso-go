package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("field ID is required")
	}

	if o.Price < 0 {
		return errors.New("field Price must be positive")
	}

	if o.Tax < 0 {
		return errors.New("field Tax must be positive")
	}

	return nil
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	errors := order.Validate()

	if errors != nil {
		return nil, errors
	}

	return &order, nil
}

func (o *Order) CalculateFinalPrice() {
	o.FinalPrice = o.Price + o.Tax
}
