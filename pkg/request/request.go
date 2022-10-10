package request

import "encoding/json"

type Request struct {
	PrivateKey            string `json:"PrivateKey,omitempty"`
	Secret                string `json:"Secret,omitempty"`
	DocumentText          string `json:"DocumentText,omitempty"`
	IsTwitterContent      bool   `json:"IsTwitterContent,omitempty"`
	UserCategoryModelName string `json:"UserCategoryModelName,omitempty"`
	DocumentLanguage      string `json:"DocumentLanguage,omitempty"`
	RequestIdentifier     string `json:"RequestIdentifier,omitempty"`
}

func (a *Request) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}
