package service

import (
	"nipun.io/brew_machine/domain"
	"nipun.io/brew_machine/repository"
)

type LocalDispenserService struct {
	BeverageRepository   repository.BeverageRepository
	IngredientRepository repository.IngredientRepository
	OutletCnt            int
}

func (lds *LocalDispenserService) MakeBeverage(name string, outlet int) (domain.Beverage, error) {
	// check if given outlet is available for providing the beverage.
	beverage, err := lds.BeverageRepository.Get(name)
	if err != nil {
		return domain.Beverage{}, err
	}
	// TODO : add lock on ingredientRepo for update operation
	acquiredIngredientsList := []string{}
	for ingredient, quantity := range beverage.IngredientsQuantityMap {
		err = lds.IngredientRepository.UpdateQuantity(ingredient, -quantity)
		acquiredIngredientsList = append(acquiredIngredientsList, ingredient)
		if err != nil {
			// rolling back  the change.
			for _, acquiredIngredient := range acquiredIngredientsList {
				lds.IngredientRepository.UpdateQuantity(acquiredIngredient, quantity)
			}
			return domain.Beverage{}, err
		}
	}
	// END TODO : add lock on ingredientRepo for update operation

	return *beverage, nil
}
