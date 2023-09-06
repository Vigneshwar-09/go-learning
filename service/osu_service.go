package service

import (
	"context"
	"encoding/json"
	"errors"
	"example/hello/configs"
	"example/hello/model"
	"fmt"
	"io"

	// "io"
	"os"
	"time"

	"net/http"
	"net/url"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/exp/slices"
)

var tokenCollection *mongo.Collection = configs.GetCollection("osuAuth")
var userCollection *mongo.Collection = configs.GetCollection("user")

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
	cur, err := tokenCollection.DeleteMany(context.TODO(), bson.M{"expiresIn": bson.M{
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

func FetchUserData(userId string) (resp model.User, err error) {
	userResponse := new(model.User)

	err = userCollection.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&userResponse)
	if err != nil && err == mongo.ErrNoDocuments {
		token, _ := FetchAuthToken()
		client := &http.Client{}
		request, _ := http.NewRequest("GET", "https://osu.ppy.sh/api/v2/users/"+userId+"/osu", nil)
		request.Header = http.Header{"Authorization": {token.TokenType + " " + token.AccessToken}, "Content-Type": {"application/json"}, "Accept": {"application/json"}}
		response, _ := client.Do(request)

		json.NewDecoder(response.Body).Decode(&userResponse)
		userResponse.ID = cast.ToString(userResponse.ID)
		userCollection.InsertOne(context.TODO(), userResponse)

	}
	// body, _ := io.ReadAll(response.Body)
	// json.Unmarshal(body, &resp)
	return *userResponse, nil

}

func LineChartForUser(userString string) error {
	user, _ := FetchUserData(userString)
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		// charts.WithXAxisOpts(opts.XAxis{
		// 	// Show: false,
		// 	Min: user.RankHistory.Data[0],
		// }),
		charts.WithYAxisOpts(opts.YAxis{
			// Show: false,
			Min: user.RankHistory.Data[0],
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    user.Username,
			Subtitle: user.RankHistory.Mode,
			Link:     user.AvatarURL,
		}))

	fmt.Println(generateLineData(user.RankHistory.Data))

	dateRange := make([]int64, 89)
	for i := range dateRange {
		dateRange[i] = 89 - int64(i)
	}
	slices.Reverse(user.RankHistory.Data)
	fmt.Println("Date Range : ", dateRange)
	line.SetXAxis(generateLineData(dateRange)).
		AddSeries("Rank History", generateLineData(user.RankHistory.Data)).
		// AddSeries("Category B", generateLineData(user.)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: false, ShowSymbol: true}))

	page := components.NewPage()
	page.AddCharts(line)
	f, err := os.Create("line.html")
	if err != nil {
		fmt.Println("Fuck me" + err.Error())
	}
	page.Render(io.MultiWriter(f))
	return nil
}

func generateLineData(data []int64) []opts.LineData {
	items := make([]opts.LineData, 0)
	// data[0]=0
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{
			Name:       cast.ToString(data[i]),
			Value:      data[i],
			Symbol:     "circle",
			SymbolSize: 10,
			// XAxisIndex: i,
			// YAxisIndex: i,
		})
	}
	return items
}

// func generateLineItems() []opts.LineData {
// 	items := make([]opts.LineData, 0)
// 	for i := 0; i < itemCntLine; i++ {
// 		items = append(items, opts.LineData{Value: rand.Intn(300)})
// 	}
// 	return items
// }
