package hipchat

import "encoding/json"

var BaseURL = "https://api.hipchat.com/v2/"

type apiResponse struct {
	*Error `json:"error"`

	Items json.RawMessage `json:"items"`

	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	MaxResults float64 `json:"maxResults,omitempty"`
	StartIndex float64 `json:"startIndex,omitempty"`
}

type roomResponse struct {
	*Room
	*Error `json:"error"`
}

type userResponse struct {
	*User
	*Error `json:"error"`
}

func (a *apiResponse) IsError() bool {
	return a.Error != nil
}
