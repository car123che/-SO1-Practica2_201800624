package controllers

import (
	// "context"
	// "mongo-backend/configs"
	// "mongo-backend/models"

//	"bytes"
//	"encoding/json"
	"client/responses"
	//"io/ioutil"
//	"log"
	"net/http"
//	"time"
	"context"
	"fmt"
	pb "github.com/tomiok/grpc-test-wishlist/proto"
	"google.golang.org/grpc"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

// Get result from mongo db

func Inicio(c *fiber.Ctx) error {




	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewWishListServiceClient(conn)

	res, err := serviceClient.Create(context.Background(), &pb.CreateWishListReq{
		WishList: &pb.WishList{
			Id:   generateID(),
			Name: "CHE LE GUSTA LA PALOMA ",
		},
	})

	if err != nil {
		panic("wishlist is not created  " + err.Error())
	}

	fmt.Println(res.WishListId)

	return c.Status(http.StatusOK).JSON(responses.Inicio{
		//Status: http.StatusOK,
		Bienvenido: "chinga tu madre",
		
	})
	


}


func generateID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}     

