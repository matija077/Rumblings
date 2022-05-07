package mongoDriver

import (
	"context"
	"fmt"
	"log"
	"mongodb/source/faker"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type MongoDatabaseDriver struct {
	Collection *mongo.Collection
}

func (db *MongoDatabaseDriver) Find(ctx *context.Context, key string, value interface{}) (result interface{}) {
	findOneOptions := options.FindOneOptions{}
	filter := bson.M{
		key: value,
	}

	result = db.Collection.FindOne(*ctx, filter, &findOneOptions).Decode(&struct{ key string }{})
	return result
}

func (db *MongoDatabaseDriver) FindOne(ctx *context.Context, key string, value interface{}) (result interface{}) {
	findOneOptions := options.FindOneOptions{}
	filter := bson.M{
		key: value,
	}
	fmt.Println(filter)

	result = db.Collection.FindOne(*ctx, filter, &findOneOptions).Decode(faker.Person{})
	return result
}

func (dD *MongoDatabaseDriver) ensureIndex(ctx *context.Context, collection *mongo.Collection) bool {
	listIndexOptions := options.ListIndexesOptions{}
	indexesCursor, err := collection.Indexes().List(*ctx, &listIndexOptions)

	var indexModel = mongo.IndexModel{
		Keys: bson.M{
			"GUID": 1,
		}, Options: nil,
	}

	s, err := collection.Indexes().CreateOne(*ctx, indexModel)
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Println(s)

	if err != nil {
		log.Fatal("indexes can't be found")
		return false
	}

	var result []bson.M
	if err = indexesCursor.All(*ctx, &result); err != nil {
		log.Fatal(err)
	}

	for _, v := range result {
		for k1, v1 := range v {
			fmt.Printf("%v: %v\n", k1, v1)
		}
		fmt.Println()
	}

	return true
}
