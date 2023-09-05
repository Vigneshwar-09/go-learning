package service

import (
	"context"
	"encoding/json"
	"example/hello/configs"
	"example/hello/model"
	"fmt"
	"os"

	"net/http"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
)

var osuCollection *mongo.Collection = configs.GetCollection(configs.DB, "osuAuth")
var clientId string = os.Getenv("OSUCLIENT")
var clientSecert string = os.Getenv("OSUSECERET")

func FetchAuthToken() model.OAuth {

	fourmData := url.Values{"client_id": {clientId}, "client_secret": {clientSecert}, "grant_type": {"client_credentials"}, "scope": {"public"}}

	response, err := http.PostForm("https://osu.ppy.sh/oauth/token", fourmData)

	if err != nil {
		fmt.Println(err)
		return *new(model.OAuth)
	}
	responseData := new(model.OAuthResponse)

	json.NewDecoder(response.Body).Decode(&responseData)

	finalData := model.NewOAuth(*responseData)

	osuCollection.InsertOne(context.TODO(), finalData)

	return *finalData

}

func 