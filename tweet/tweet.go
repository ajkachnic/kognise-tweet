package tweet

type Tweet struct {
	Text string `json:"text"`
}

type TwitterResponse struct {
	Tweets []Tweet `json:"statuses"`
}