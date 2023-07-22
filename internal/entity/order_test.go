package entity_test

import (
	"curso-go/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New_Order_Without_Error(t *testing.T) {
	order, err := entity.NewOrder("123", 50, 10)
	assert.Equal(t, order.ID, "123")
	assert.Equal(t, err, nil)
}

func Test_Should_Validate_Required_ID(t *testing.T) {
	order := entity.Order{ID: "", Price: 50, Tax: 10}
	assert.Error(t, order.Validate(), "field ID is required")
}

func Test_Should_Validate_Price(t *testing.T) {
	order := entity.Order{ID: "123", Price: -50, Tax: 10}
	assert.Error(t, order.Validate(), "field Price must be positive")
}

func Test_Should_Validate_Tax(t *testing.T) {
	order := entity.Order{ID: "123", Price: 50, Tax: -10}
	assert.Error(t, order.Validate(), "field Tax must be positive")
}
