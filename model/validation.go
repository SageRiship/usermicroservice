package model

type Validation struct {
	Uname    string `json:"uname,omitempty" bson:"uname,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}
