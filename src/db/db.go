package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var GDb *Db

// NewDatabase create connection with mongodb, it works in Singleton manner
func NewDatabase() *Db {
	if GDb == nil {
		GDb = new(Db)
		clientOptions := options.Client().ApplyURI(connectionUrl)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Println("[DB] Fatal Error. Killing process.")
			log.Fatal(err)
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Println("[DB] Fatal Error. Killing process.")
			log.Fatal(err)
		}
		log.Println("[DB] Connected to MongoDB")

		collection := client.Database(dbName).Collection(collecName)
		log.Println(collection)
		log.Println("[DB] Database collection successfully initialized.")

		GDb = &Db{
			DbName:     dbName,
			CollecName: collecName,
			Collection: collection,
		}
	}

	return GDb
}

// Read reads documents from database
func (db *Db) Read() []Doc {
	log.Println("[DB] Start read operation...")

	var result []Doc
	var res []Doc

	opts := options.Find().SetSort(bson.M{"created": -1})
	cursor, err := db.Collection.Find(context.TODO(), bson.M{"userId": "root"}, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &res); err != nil {
		log.Fatal(err)
	}

	log.Println("[DB] Found:", res[0])

	result = append(result, res[0])

	log.Println("[DB] Read operation success")

	return result
}

// Write writes document on database
func (db *Db) Write(userid string, contents Doc) error {
	log.Println("[DB] Start write operation...")
	log.Println("[DB] Writing contents:", contents)

	res, err := db.Collection.InsertOne(context.TODO(), contents)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("[DB] Writing operation success")
	log.Println("[DB] Result:", res)

	return nil
}

// Count counts how many documents are left in database
func (db *Db) Count(userid string) int64 {
	log.Println("[DB] Start counting documents...")

	cnt, err := db.Collection.CountDocuments(context.TODO(), bson.M{"userId": "root"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[DB] Counting operation done.")
	log.Println("[DB] Document Number:", cnt)

	return cnt
}

// Delete deletes specified or every user document
func (db *Db) Delete(userid string) error {
	log.Println("[DB] Start delete operation...")

	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: true,
	})

	cnt := db.Count("root")
	for i := 0; i < int(cnt); i++ {
		res, err := db.Collection.DeleteOne(context.TODO(), bson.M{"userId": "root"}, opts)
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Printf("[DB] Deleted %v documents\n", res.DeletedCount)
	}

	return nil
}
