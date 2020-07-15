package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ajkachnic/tweet-from/tweet"
	"github.com/mb-14/gomarkov"
)

func getTweetsFromFile() tweet.TwitterResponse {
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}
	tweets := []tweet.Tweet {}
	err = json.Unmarshal(file, &tweets)
	if err != nil {
		fmt.Println(err)
	}
	return tweet.TwitterResponse{
		Tweets: tweets,
	}
}

func buildModel() (*gomarkov.Chain, error) {
	chain := gomarkov.NewChain(1)
	tweets := getTweetsFromFile()
	for _, tweet := range tweets.Tweets {
		if tweet.Text != "" {
			chain.Add(strings.Split(tweet.Text, " "))
		}
	}
	return chain, nil
}
func saveModel(chain *gomarkov.Chain) {
	jsonObj, _ := json.Marshal(chain)
	err := ioutil.WriteFile("model.json", jsonObj, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
func loadModel() (*gomarkov.Chain, error) {
	var chain gomarkov.Chain
	data, err := ioutil.ReadFile("model.json")
	if err != nil {
		return &chain, err
	}
	err = json.Unmarshal(data, &chain)
	if err != nil {
		return &chain, err
	}
	return &chain, nil
}

func generateTweet(chain *gomarkov.Chain) {
	tokens := []string{gomarkov.StartToken}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}
	fmt.Println(strings.Join(tokens[1:len(tokens)-1], " "))
}

func main() {
	train := flag.Bool("train", false, "Train the markovet chain")
	count := flag.Int("count", 1, "Set the amount of tweets to generate")
	flag.Parse()
	if *train {
		chain, err := buildModel()
		
		if err != nil {
			fmt.Println(err)
			return
		}
		saveModel(chain)
	} else {
		chain, err := loadModel()
		if err != nil {
			fmt.Println(err)
			return
		}
		var ct int = *count;
		
		for i := 0; i < ct; i ++ {
			generateTweet(chain)
		}
	}
}