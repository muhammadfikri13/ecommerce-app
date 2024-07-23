package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can't find product")
	ErrCantDecodeProducts = errors.New("can't find product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("can't add this product to the cart")
	ErrCantRemoveItemCart = errors.New("unable to remove this item from the cart")
	ErrCantGetItem        = errors.New("unable to get item from the cart")
	ErrCantBuyCartItem    = errors.New("can't update the purchase")
)

func AddProductToCart() {

}
func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
