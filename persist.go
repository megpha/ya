package mytube

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	MONGO_URL          = "mongodb://youtube:welcome123@ds147079.mlab.com:47079/youtube"
	DB_NAME            = "youtube"
	VIDEO_COLLECTION   = "videos"
	CHANNEL_COLLECTION = "channels"
)

func VideosByChannel(channel string) (videos []Video) {
	session, err := mgo.Dial(MONGO_URL)

	if err != nil {
		fmt.Println(err)
	}

	defer session.Close()

	collectionHandler := session.DB(DB_NAME).C(VIDEO_COLLECTION)
	collectionHandler.Find(bson.M{"channel": channel}).All(&videos)
	return
}

func Channels() (results []Channel) {
	session, err := mgo.Dial(MONGO_URL)

	if err != nil {
		fmt.Println(err)
	}

	defer session.Close()

	collectionHandler := session.DB(DB_NAME).C(CHANNEL_COLLECTION)
	collectionHandler.Find(bson.M{}).All(&results)
	return
}

func Persist(channel string, results []Video) {
	session, err := mgo.Dial(MONGO_URL)

	if err != nil {
		fmt.Println(err)
	}

	defer session.Close()

	collectionHandler := session.DB(DB_NAME).C(VIDEO_COLLECTION)
	collectionHandler.RemoveAll(bson.M{"channel": channel})

	video_collection := collectionHandler.Bulk()

	for _, item := range results {
		video_collection.Insert(item)
	}

	video_collection.Run()
}
