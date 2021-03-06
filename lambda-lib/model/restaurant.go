package model

type Restaurant struct {
	Name   string              `dynamodbav:"name" json:"name,omitempty"`
	Image  string              `dynamodbav:"image" json:"image,omitempty"`
	Themes []map[string]string `dynamodbav:"themes" json:"themes,omitempty"`
}
