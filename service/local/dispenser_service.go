package service

import (
	"fmt"
	"nipun.io/brew_machine/domain"
	"nipun.io/brew_machine/repository"
	"nipun.io/brew_machine/service"
)

type LocalDispenserService struct {
	BeverageRepository     repository.BeverageRepository
	IngredientRepository   repository.IngredientRepository
	OutletCnt              int
	TransactionLockManager service.TransactionLockManager
}

func (lds *LocalDispenserService) MakeBeverage(name string, outlet int) (domain.Beverage, error) {
	// check if given outlet is available for providing the beverage.
	outletNumberStr := fmt.Sprintf("%d", outlet)
	lockKeys := []string{"LocalDispenserService", "Outlet", outletNumberStr}
	lds.TransactionLockManager.AcquireLock(lockKeys)

	beverage, err := lds.BeverageRepository.Get(name)
	if err != nil {
		lds.TransactionLockManager.ReleaseLock(lockKeys)
		return domain.Beverage{}, err
	}
	acquiredIngredientsList := []string{}

	for ingredient, quantity := range beverage.IngredientsQuantityMap {
		// remove the X units of an ingredient
		err = lds.IngredientRepository.UpdateQuantity(ingredient, -quantity)
		if err != nil {
			// rolling back  the change.
			for _, acquiredIngredient := range acquiredIngredientsList {
				// ignoring the error deliberately, will need to add more complexity to handle this.
				// add the X units of an ingredient that was taken
				lds.IngredientRepository.UpdateQuantity(acquiredIngredient, quantity)

			}
			lds.TransactionLockManager.ReleaseLock(lockKeys)
			return domain.Beverage{}, err
		}
		acquiredIngredientsList = append(acquiredIngredientsList, ingredient)
	}
	lds.TransactionLockManager.ReleaseLock(lockKeys)
	return *beverage, nil
}
