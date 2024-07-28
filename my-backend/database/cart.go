package database

import (
	"github.com/gin-gonic/gin"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"my-backend/models"
)

var (
	ErrCantFindProduct    = errors.New("can't find product")
	ErrCantDecodeProducts = errors.New("can't find product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("can't add this product to the cart")
	ErrCantRemoveItemCart = errors.New("unable to remove this item from the cart")
	ErrCantGetItem        = errors.New("unable to get item from the cart")
	ErrCantBuyCartItem    = errors.New("can't update the purchase")
)

func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, primitive.ObjectID, userID string) error {
	searchfromdb, err := prodCollection.Find(ctx, bson.M{"_id": productID})
	if err != nil {
		log.Println(err)
		return ErrCantFindProduct
	}
	var productCart []models.ProductUser
	err = searchfromdb.All(ctx, &productcart)
	if err != nil {
		log.Println	(err)
		return ErrCantDecodeProducts
	}

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}

	filter := bson.D{primitive.E{Key:"_id", Value: id}} // to update somthing, we need the id
	update := bson.D{{Key:"$push", Value: bson.D{primitive.E{Key:"usercart", Value: bson.D{{Key:"$each", Value: productCart}}}}}}

	_, err =userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ErrCantUpdateUser
	}
	return nil
}
func RemoveCartItem(ctx, context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	id , err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return ErrUserIdIsNotValid
	}
	filter := bson.D(primitive.E{Key:"_id", Value: id})
	update := bson.M("$pull", bson.M{"usercart": bson.M{"product_id": productID}})
	_, err = UpdateMany(ctx, filter, update)
	if err 1+ nil {
		return ErrCantRemoveItemCart
	}
	return nil
}

func BuyItemFromCart() {

}

func InstantBuyer() {

}
