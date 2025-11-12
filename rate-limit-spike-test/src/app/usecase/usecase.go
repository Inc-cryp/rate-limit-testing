package usecase

import (
	bookUC "rate-limit-spike-test/src/app/usecase/books"
	// pickUpUC "rate-limit-spike-test/src/app/usecase/pickup"
)

type AllUseCases struct {
	BookUC bookUC.BooksUCInterface
	// PickUpUC pickUpUC.PickUpUCInterface
}
