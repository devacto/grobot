package data

import (
	"fmt"
	"os"

	"github.com/devacto/grobot/Godeps/_workspace/src/gopkg.in/mgo.v2"
)

// Nutrition is the kind of nutrition and the amount.
type Nutrition struct {
	Name     string
	Quantity string
}

// Food is s type.
type Food struct {
	Name       string
	Company    string
	Nutritions []Nutrition
}

// Col is the Collection returned from the DB.
var Col *mgo.Collection

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
	if err := Col.Find(nil).All(&result); err != nil {
		panic(err)
	}
	return result
}
