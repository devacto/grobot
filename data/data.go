package data

import (
	"fmt"
	"log"
	"os"

	"github.com/devacto/grobot/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"github.com/devacto/grobot/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
	"github.com/devacto/grobot/Godeps/_workspace/src/gopkg.in/olivere/elastic.v2"
)

// Nutrition is the kind of nutrition and the amount.
type Nutrition struct {
	Name     string
	Quantity string
}

// Food is s type.
type Food struct {
	Id         string
	Name       string
	Company    string
	Nutritions []Nutrition
}

// Col is the Collection returned from the DB.
var Col *mgo.Collection

// NewNutrition returns a Nutrition.
func NewNutrition(name string, quantity string) Nutrition {
	n := Nutrition{Name: name, Quantity: quantity}
	return n
}

// NewFood returns a Food.
func NewFood(id, name, company string, nut []Nutrition) Food {
	f := Food{Id: id, Name: name, Company: company, Nutritions: nut}
	return f
}

// FoodWithIdExists will return true or false if food with that id exists.
func FoodWithIdExists(id string) bool {
	session, err := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	if err != nil {
		log.Fatalf("Can't connect to mongo, go error %v\n", err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	Col = session.DB("").C("foods")

	result := new(Food)
	if err = Col.Find(bson.M{"id": id}).One(&result); err != nil {
		return false
	}

	return true
}

// GetAllFoods fetches all foods from the database.
func GetAllFoods() []Food {
	session, err := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	Col = session.DB("").C("foods")

	var result []Food
	if err = Col.Find(nil).All(&result); err != nil {
		panic(err)
	}
	return result
}



// InsertFood inserts one food into the database.
func InsertFood(f Food) {
	session, err := mgo.Dial(os.Getenv("MONGOLAB_URI"))
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	Col = session.DB("").C("foods")

	if err = Col.Insert(f); err != nil {
		panic(err)
	}

	es, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(os.Getenv("BONSAI_URL")),
		elastic.SetMaxRetries(1))
	if err != nil {
		log.Fatalf("Cannot connect to elastic at %s error: %v\n", os.Getenv("BONSAI_URL"), err)
	}

	// If index does not exist then create food index.
	exists, err := es.IndexExists("food").Do()
	if err != nil {
		log.Panicf("Cannot check if index exists error %v\n", err)
	}
	if !exists {
		_, err = es.CreateIndex("food").Do()
		if err != nil {
			log.Panicf("Cannot create index error %v", err)
		}
	}

	log.Printf("Indexing %s", f.Name)
	_, err = es.Index().
		Index("food").
		Type("food").
		Id(f.Id).
		BodyJson(f).
		Do()
	if err != nil {
		log.Panicf("Cannot index a food error %v\n", err)
	}
}
