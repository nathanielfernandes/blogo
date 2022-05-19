package mongo

type Post struct {
	ID string `bson:"_id" json:"-"`

	Title   string   `bson:"title" json:"title"`
	Hook    string   `bson:"hook" json:"hook"`
	Tags    []string `bson:"tags" json:"tags"`
	Content string   `bson:"content" json:"-"`
	Date    int64    `bson:"date" json:"date"`

	Views   int64 `bson:"views" json:"views"`
	Likes   int64 `bson:"likes" json:"likes"`
	Cookies int64 `bson:"cookies" json:"cookies"`

	Private bool `bson:"private" json:"-"`

	// HTML []byte `bson:"-" json:"-"`
}
