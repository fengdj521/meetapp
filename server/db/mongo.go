package db

import (
	"fmt"

	"os"

	"golang.org/x/net/context"
	mgo "gopkg.in/mgo.v2"
)

const mongoDBName = "meetapp"

type mongodb string

func MongoDB(ctx context.Context) *mgo.Session {
	key := mongodb(mongoDBName)
	db, _ := ctx.Value(key).(*mgo.Session)
	return db
}

func Name() string {
	return mongoDBName
}

func OpenMongoDB(ctx context.Context) context.Context {
	url := os.Getenv("MONGOLAB_URI")
	if url == "" {
		url = fmt.Sprintf("%s:%d", "localhost", 27017)
	}
	sesh, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	ctx = context.WithValue(ctx, mongodb(mongoDBName), sesh)
	return ctx
}

func CloseMongoDB(ctx context.Context) context.Context {
	sesh := MongoDB(ctx)
	if sesh == nil {
		fmt.Println("not found mongoDB")
	}
	sesh.Close()
	ctx = context.WithValue(ctx, mongodb(mongoDBName), nil)
	return ctx
}
