package repository

import (
	"github.com/stretchr/testify/assert"
	"nipun.io/brew_machine/domain"
	"testing"
)

func TestInMemoryBeverageRepository_AddNewAndGet(t *testing.T) {
	inMemoryBeverageRepository := InMemoryBeverageRepository{
		data: map[string]*domain.Beverage{},
		// TODO : add transactionalLockManager Mock objects here.
	}
	chai := "chai"
	inMemoryBeverageRepository.AddNew(domain.Beverage{
		Name: chai,
		IngredientsQuantityMap: map[string]int{
			"water": 50,
			"tea":   10,
			"sugar": 10,
			"milk":  50,
		},
	})

	got, err := inMemoryBeverageRepository.Get(chai)
	assert.True(t, err == nil)
	assert.True(t, got != nil)
	assert.True(t, got.Name == chai)
	assert.True(t, len(got.IngredientsQuantityMap) == 4)
}
