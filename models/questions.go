package dmas

import (
	"dmas/config"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Question structure
type Question struct {
	_ID    int64  `db:"_id" json:"id"`
	Title  string `db:"title" json:"title"`
	Answer string `db:"answer" json:"answer"`
	A1     string `db:"a1" json:"a1"`
	A2     string `db:"a2" json:"a2"`
	A3     string `db:"a3" json:"a3"`
	A4     string `db:"a4" json:"a4"`
}

//GetAllQuestions will return all objects from questions collection
func GetAllQuestions() []Question {
	var result []Question //so results is accutally array of questions

	session := dmas.MgoSession.Clone()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	session.DB(dmas.DbName).C("questions").Find(nil).All(&result)

	return result
}

//GetQuestion function should return actual question based on id or something - note that u can use mockup down bellow
func GetQuestion() Question {

	session := dmas.MgoSession.Clone()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	result := Question{}
	session.DB(dmas.DbName).C("questions").Find(bson.M{"title": "Majkl Dzordan je bio koji pik na draftu"}).One(&result)

	return result
}

//NOTE THAT FUNCTIONS FOR EXPORT IN PACKAGE MUST START WITH UPPERCASE
