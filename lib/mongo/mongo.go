package mongo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var /* const */ MONGOURI = os.Getenv("MONGO_URI")

var Ctx = context.Background()

type BlogoMongo struct {
	MongoClient *mongo.Client
	collection  *mongo.Collection
}

func NewBlogoMongo() *BlogoMongo {
	client, err := mongo.Connect(Ctx, options.Client().ApplyURI(MONGOURI))
	if err != nil {
		panic(err)
	}

	return &BlogoMongo{
		MongoClient: client,
		collection:  client.Database("blogo").Collection("blogs"),
	}
}

func (bm *BlogoMongo) NewPost(id, title, hook, content string, tags []string) error {
	_, err := bm.collection.InsertOne(Ctx, bson.D{
		{Key: "_id", Value: id},
		{Key: "title", Value: title},
		{Key: "hook", Value: hook},
		{Key: "tags", Value: tags},
		{Key: "content", Value: content},
		{Key: "date", Value: time.Now().Unix()},

		{Key: "views", Value: 0},
		{Key: "likes", Value: 0},
		{Key: "cookies", Value: 0},

		{Key: "private", Value: true},
	})
	return err
}

func (bm *BlogoMongo) GetPost(id string) (Post, error) {
	var post Post
	err := bm.collection.FindOne(Ctx, bson.D{{Key: "_id", Value: id}}).Decode(&post)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (bm *BlogoMongo) GetAllPosts() ([]Post, error) {
	cur, err := bm.collection.Find(Ctx, bson.D{{}})
	if err != nil {
		return []Post{}, err
	}

	defer cur.Close(Ctx)

	var posts []Post
	err = cur.All(Ctx, &posts)

	return posts, err
}
