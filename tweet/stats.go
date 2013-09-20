package tweet

import (
	m "github.com/majest/go-user-service/db"
	"labix.org/v2/mgo/bson"
)

const collection = "tweets"

type Model interface{
	Save() Tweet
}

func (d *Tweet) Save() {
	m.Insert(collection, &d)
}

func Get(id string) (d *Tweet) {
	d = &Tweet{}
	m.Get(collection, &d, id)
	return
}

func All() (d *[]Tweet) {
	d = &[]Tweet{}
	m.All(collection, d)
	return
}

func New() *Tweet {
	return &Tweet{Id:bson.NewObjectId()}
}
