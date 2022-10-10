package response

import (
	"encoding/json"
	"io"
)

type DocumentResultStatus int

const (
	DocumentResultStatusOK                    DocumentResultStatus = 1
	DocumentResultStatusAuthenticationFailure DocumentResultStatus = 2
	DocumentResultStatusCreditExceeded        DocumentResultStatus = 3
	DocumentResultStatusServiceBusy           DocumentResultStatus = 4
	DocumentResultStatusValidationError       DocumentResultStatus = 5
	DocumentResultStatusGenericError          DocumentResultStatus = 6
)

type ErrorMessage string

type Response struct {
	Status                      *DocumentResultStatus `json:"Status,omitempty"`
	ErrorMessage                *ErrorMessage         `json:"ErrorMessage,omitempty"`
	DocSentimentResultString    *string               `json:"DocSentimentResultString,omitempty"`
	DocSentimentValue           *float64              `json:"DocSentimentValue,omitempty"`
	DocSentimentPolarity        *string               `json:"DocSentimentPolarity,omitempty"`
	Magnitude                   *float64              `json:"Magnitude,omitempty"`
	Entities                    []SentencePart        `json:"Entities,omitempty"`
	Themes                      []SentencePart        `json:"Themes,omitempty"`
	Keywords                    []SentencePart        `json:"Keywords,omitempty"`
	Citations                   []SentencePart        `json:"Citations,omitempty"`
	SlangWords                  []SlangWord           `json:"SlangWords,omitempty"`
	SwearWords                  []SlangWord           `json:"SwearWords,omitempty"`
	CoreSentences               []Sentence            `json:"CoreSentences,omitempty"`
	PartsOfSpeech               []PartOfSpeech        `json:"PartsOfSpeech,omitempty"`
	AutoCategories              []Category            `json:"AutoCategories,omitempty"`
	UserCategories              []Category            `json:"UserCategories,omitempty"`
	Subjectivity                *string               `json:"Subjectivity,omitempty"`
	DetectedLanguage            *string               `json:"DetectedLanguage,omitempty"`
	CloudTagHTML                *string               `json:"CloudTagHTML,omitempty"`
	ResultTextHTML              *string               `json:"ResultTextHtml,omitempty"`
	TransactionTotalCreditsLeft *int                  `json:"TransactionTotalCreditsLeft,omitempty"`
	TransactionUseByDate        *string               `json:"TransactionUseByDate,omitempty"`
	TransactionCurrentDay       *int                  `json:"TransactionCurrentDay,omitempty"`
	TransactionDailyLimit       *int                  `json:"TransactionDailyLimit,omitempty"`
	StorageInfo                 *StorageInfo          `json:"StorageInfo,omitempty"`
	Timestamp                   *int64                `json:"Timestamp,omitempty"`
}

type SentencePart struct {
	Text              *string  `json:"Text,omitempty"`
	SentenceText      *string  `json:"SentenceText,omitempty"`
	Mentions          *int     `json:"Mentions,omitempty"`
	SentencePartType  *string  `json:"SentencePartType,omitempty"`
	KeywordType       *string  `json:"KeywordType,omitempty"`
	SentimentResult   *string  `json:"SentimentResult,omitempty"`
	Magnitude         *float64 `json:"Magnitude,omitempty"`
	SentimentValue    *float64 `json:"SentimentValue,omitempty"`
	SentimentPolarity *string  `json:"SentimentPolarity,omitempty"`
}

type SlangWord struct {
	SlangWordText        *string `json:"SlangWordText,omitempty"`
	SlangWordTranslation *string `json:"SlangWordTranslation,omitempty"`
}

type Sentence struct {
	SentenceNumber        *int     `json:"SentenceNumber,omitempty"`
	Text                  *string  `json:"Text,omitempty"`
	SentimentResultString *string  `json:"SentimentResultString,omitempty"`
	SentimentValue        *float64 `json:"SentimentValue,omitempty"`
	SentimentPolarity     *string  `json:"SentimentPolarity,omitempty"`
	Magnitude             *float64 `json:"Magnitude,omitempty"`
}

type PartOfSpeech struct {
	Text                        *string  `json:"Text,omitempty"`
	Subject                     *string  `json:"Subject,omitempty"`
	Action                      *string  `json:"Action,omitempty"`
	Object                      *string  `json:"Object,omitempty"`
	ObjectSentimentResultString *string  `json:"ObjectSentimentResultString,omitempty"`
	ObjectSentimentValue        *float64 `json:"ObjectSentimentValue,omitempty"`
	ObjectSentimentPolarity     *string  `json:"ObjectSentimentPolarity,omitempty"`
}

type Category struct {
	CategoryName string  `json:"CategoryName,omitempty"`
	Score        float64 `json:"Score,omitempty"`
}

type StorageInfo struct {
	CreateDate            string  `json:"CreateDate,omitempty"`
	IP                    string  `json:"IP,omitempty"`
	IsTwitterMode         bool    `json:"IsTwitterMode,omitempty"`
	IsExcel               bool    `json:"IsExcel,omitempty"`
	IsGSExcel             bool    `json:"IsGSExcel,omitempty"`
	UserCategoryModelName *string `json:"UserCategoryModelName,omitempty"`
	RequestIdentifier     *string `json:"RequestIdentifier,omitempty"`
	PrivateKey            string  `json:"PrivateKey,omitempty"`
	DocumentText          string  `json:"DocumentText,omitempty"`
}

func (c *Response) Decode(body io.ReadCloser) error {
	defer body.Close()
	return json.NewDecoder(body).Decode(c)
}

func (s *DocumentResultStatus) IsSuccess() bool {
	return *s == DocumentResultStatusOK
}

func (e *ErrorMessage) Error() string {
	return string(*e)
}
