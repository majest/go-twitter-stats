package tweet

import "labix.org/v2/mgo/bson"

type Tweet struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Text       string        `bson:"text" json:"text"`
	UserId       int        `bson:"userId" json:"userId"`
	UserName       string        `bson:"userName" json:"userName"`
}