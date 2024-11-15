package exception

import "errors"

var ErrOrderAlreadyExistsAnotherUser = errors.New("the order already set for another user")

var ErrOrderAlreadyExists = errors.New("the order already exists")

var ErrInvalidOrderNumber = errors.New("invalid order number format")

var ErrNotContentOfUser = errors.New("the user has no orders")
