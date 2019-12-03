package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

// TwitterAccount はTwitterの認証用の情報
type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

func main() {
	// Json読み込み
	raw, error := ioutil.ReadFile("./twitterAccount.json")
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	var twitterAccount TwitterAccount
	// 構造体にセット
	json.Unmarshal(raw, &twitterAccount)

	// 認証
	api := anaconda.NewTwitterApiWithCredentials(twitterAccount.AccessToken, twitterAccount.AccessTokenSecret, twitterAccount.ConsumerKey, twitterAccount.ConsumerSecret)
	v := url.Values{}
	v.Set("count","10000")
	// v.Set("exclude","retweets")
	// 検索
	searchResult, _ := api.GetSearch(`"アオキ大好き！最高！"`, v)
	fmt.Println(strconv.Itoa(len(searchResult.Statuses))+"件ヒットしました！！")
	for _, tweet := range searchResult.Statuses {
		fmt.Printf("%d\n", tweet.User.Id)
		api.FollowUserId(tweet.User.Id,v)
		fmt.Println("--------------------------------------------------------------")
	}
}
