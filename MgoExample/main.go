package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Person struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Phone string
	Timestamp time.Time
}

var (
	IsDrop = true
)

func main() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	if IsDrop {
		err = session.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}
	c := session.DB("test").C("people")

	index := mgo.Index {
		Key: []string{"name", "phone"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
	err = c.Insert(&Person{Name:"Ale", Phone:"19008198", Timestamp:time.Now()},
	&Person{Name:"Cla", Phone:"123456789", Timestamp:time.Now()})
	if err != nil {
		panic(err)
	}

	// Query one
	result := Person{}
	err = c.Find(bson.M{"name" : "Ale"}).Select(bson.M{"phone": 0}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone", result)

	// Query all
	var results []Person
	err = c.Find(bson.M{"name" : "Ale"}).Sort("-timestamp").All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

	// Update
	colQuerier := bson.M{"name" : "Ale"}
	change := bson.M{"$set" : bson.M{"phone" : "+86 99 8888 7777", "timestamp" : time.Now()}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}

	// Query All
	// Query All
	err = c.Find(bson.M{}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)

}
