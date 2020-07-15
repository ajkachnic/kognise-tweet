# kognise-tweet

Generates a tweet from kognise...

Uses a markov chain and generates tweets based off his tweeting history

## Usage

```sh
git clone https://github.com/ajkachnic/kognise-tweet.git

cd kognise-tweet

go run main.go -train

go run main.go
```

## FAQ

> What? Why would I want this?

tbh i don't really know, but there are a lot of weirdo's out there

> How does it work?

why would i know that?

> Can it work for another person's tweets?

Maybe... you'd need to download all of their tweets (via [GetOldTweets](https://github.com/Mottl/GetOldTweets3))

And then convert that from a CSV file to a JSON file. Replace `data.json` with that file. Note that there might be some cleaing up to do, like removing tweets with only numbers or booleans.

## Notes

You can run with the `-count` flag to set a number of tweets.

Also note that I have no idea if this will work well