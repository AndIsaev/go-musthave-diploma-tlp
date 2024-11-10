package exception

import "errors"

var OrderAlreadyExistsAnotherUser = errors.New("the order already set for another user")

var OrderAlreadyExists = errors.New("the order already exists")

var InvalidOrderNumber = errors.New("invalid order number format")
