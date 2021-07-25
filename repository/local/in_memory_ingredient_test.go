package repository

import (
	"github.com/stretchr/testify/assert"
	"nipun.io/brew_machine/domain"
	"testing"
)

func TestInMemoryIngredientRepository_AddNew_Update_Get(t *testing.T) {
	inMemoryIngredientRepository := InMemoryIngredientRepository{
		data: map[string]*domain.Ingredient{},
	}
	inMemoryIngredientRepository.AddNew(domain.Ingredient{
		Name:              "water",
		AvailableQuantity: 1000,
	})
	inMemoryIngredientRepository.AddNew(domain.Ingredient{
		Name:              "milk",
		AvailableQuantity: 2000,
	})
	inMemoryIngredientRepository.AddNew(domain.Ingredient{
		Name:              "tea",
		AvailableQuantity: 200,
	})
	inMemoryIngredientRepository.AddNew(domain.Ingredient{
		Name:              "sugar",
		AvailableQuantity: 500,
	})

	got, err := inMemoryIngredientRepository.Get("water")

	assert.True(t, err == nil)
	assert.True(t, got != nil)
	assert.True(t, got.Name == "water")
	assert.True(t, got.AvailableQuantity == 1000)

	currentQuantity := got.AvailableQuantity
	delta := 1000
	expectedWaterQuantity := currentQuantity + delta
	inMemoryIngredientRepository.UpdateQuantity("water", 1000)
	got, err = inMemoryIngredientRepository.Get("water")

	assert.True(t, err == nil)
	assert.True(t, got != nil)
	assert.True(t, got.Name == "water")
	assert.True(t, got.AvailableQuantity == expectedWaterQuantity)
}
