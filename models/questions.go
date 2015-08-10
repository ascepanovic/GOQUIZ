package models

import (
	"dmas/config"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Question structure
type Question struct {
	ID    bson.ObjectId  `db:"_id" json:"id"`
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

//GetQuestionByID function should return actual question based on id
func GetQuestionByID(id int64) Question {

	session := dmas.MgoSession.Clone()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	result := Question{}
	session.DB(dmas.DbName).C("questions").FindId(id).One(&result)

	return result
}

// CreateQuestion create new question
func CreateQuestion(title, answer, a1, a2, a3, a4 string) (bool, error) {

	session := dmas.MgoSession.Clone()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	var newQuestion Question
	newQuestion.Title = title
	newQuestion.Answer = answer
	newQuestion.A1 = a1
	newQuestion.A2 = a2
	newQuestion.A3 = a3
	newQuestion.A4 = a4

	err := session.DB(dmas.DbName).C("questions").Insert(newQuestion)

	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateQuestion updating question in db
func UpdateQuestion(id int64, params bson.M) (bool, error) {

	session := dmas.MgoSession.Clone()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	update := bson.M{"$set": params}
	err := session.DB(dmas.DbName).C("questions").Update(bson.M{"_id": id}, update)

	if err != nil {
		return false, err
	}

	return true, nil
}

// DeleteQuestion removes question from collection
func DeleteQuestion(id int64) bool {
	session := dmas.MgoSession.Clone()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	err := session.DB(dmas.DbName).C("questions").RemoveId(id)

	if err != nil {
		return false
	}

	return true
}

//NOTE THAT FUNCTIONS FOR EXPORT IN PACKAGE MUST START WITH UPPERCASE
