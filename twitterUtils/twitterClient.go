// package twitterUtils

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/dghubble/go-twitter/twitter"
// 	"github.com/dghubble/oauth1"
// )

// var client twitter.Client

// func getTwitterClient() *twitter.Client {
// 	var config = oauth1.NewConfig("consumerKey", "consumerSecret")
// 	var token = oauth1.NewToken("accessToken", "accessSecret")
// 	var httpClient = config.Client(oauth1.NoContext, token)

// 	// Twitter client
// 	client := twitter.NewClient(httpClient)
// 	return client
// }

// func GetNewTweet() (*twitter.Search, error) {
// 	// Search Tweets
// 	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
// 		Query: "gopher",
// 	})

// 	if err != nil && resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("Error thrown while searching for tweets %s", err)
// 	}

// 	return search, nil
// }
package twitterutils

func notafunc() {

}
