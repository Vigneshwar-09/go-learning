package service

import (
	"context"
	"encoding/json"
	"errors"
	"example/hello/configs"
	"example/hello/model"
	"fmt"
	"os"
	"time"

	"net/http"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var tokenCollection *mongo.Collection = configs.GetCollection(configs.DB, "osuAuth")
var clientId string = os.Getenv("OSUCLIENT")
var clientSecert string = os.Getenv("OSUSECERET")

func GenerateToken() (model.OAuth, error) {

	fourmData := url.Values{"client_id": {clientId}, "client_secret": {clientSecert}, "grant_type": {"client_credentials"}, "scope": {"public"}}

	var err error
	var response *http.Response

	response, err = http.PostForm("https://osu.ppy.sh/oauth/token", fourmData)

	if err != nil {
		fmt.Println(err)
		return model.OAuth{}, err
	}
	responseData := new(model.OAuthResponse)

	json.NewDecoder(response.Body).Decode(&responseData)

	finalData := model.NewOAuth(*responseData)

	tokenCollection.InsertOne(context.TODO(), finalData)

	return *finalData, nil

}

func RemoveExpiredToken() error {
	cur, err := tokenCollection.DeleteMany(context.TODO(), bson.M{"createdAt": bson.M{
		"$lt": primitive.NewDateTimeFromTime(time.Now()),
	}})
	if err != nil {
		return err
	}
	fmt.Println("Deleted expired token count : ", cur.DeletedCount)
	return nil
}

func FetchAuthToken() (oAuth model.OAuth, err error) {

	if RemoveExpiredToken() != nil {
		return oAuth, errors.New("error while removing expired token")
	}

	var tokenList []model.OAuth
	finalCursor, err := tokenCollection.Find(context.TODO(), bson.M{"createdAt": bson.M{
		"$gte": primitive.NewDateTimeFromTime(time.Now())}})

	if err != nil {
		return oAuth, err
	}

	finalCursor.All(context.TODO(), tokenList)

	if len(tokenList) > 0 {
		return tokenList[0], nil
	}
	
	return GenerateToken()

}
