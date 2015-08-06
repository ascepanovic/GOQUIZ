package dmas

import (
  "dmas/config"
  "time"

  "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User structure
type User struct {
	_ID    string  `db:"_id" json:"id"`
	Firstname  string `db:"firstname" json:"firstname"`
	Lastname string `db:"firstname" json:"lastname"`
	Username     string `db:"username" json:"username"`
	Email     string `db:"email" json:"email"`
	Password     string `db:"password" json:"password"`
	Salt     string `db:"salt" json:"salt"`
  FacebookId     string `db:"facebook_id" json:"facebook_id"`
  Active     bool `db:"active" json:"active"`
  CreationDate     time.Time `db:"creation_date" json:"creation_date"`
}

// GetAllUsers returns all users from database
func GetAllUsers() []User {
  var result []User //so results is accutally array of questions

	session := dmas.MgoSession.Clone()
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	session.DB(dmas.DbName).C("users").Find(nil).All(&result)

	return result
}

// GetUserById returns single user from database
func GetUserById(id string) User {
  session := dmas.MgoSession.Clone()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	result := User{}
	session.DB(dmas.DbName).C("users").FindId(id).One(&result)

	return result
}

// CreateUser in database
func CreateUser(data map[string]string) (bool, error) {
  session := dmas.MgoSession.Clone()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

  var isActive = false
  if data["is_active"] == "true" {
    isActive = true
  }

  var newUser User
  newUser.Firstname = data["firstname"]
  newUser.Lastname = data["lastname"]
  newUser.Username = data["username"]
  newUser.Email = data["email"]
  newUser.FacebookId = data["facebook_id"]
  newUser.Active = isActive
  newUser.CreationDate = time.Now()

  err := session.DB(dmas.DbName).C("users").Insert(newUser)

	if err != nil {
		return false, err
	}

	return true, nil
}

// UpdateUser updating user in db
func UpdateUser(id string, params bson.M) (bool, error) {

	session := dmas.MgoSession.Clone()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	update := bson.M{"$set": params}
	err := session.DB(dmas.DbName).C("users").Update(bson.M{"_id": id}, update)

	if err != nil {
		return false, err
	}

	return true, nil
}

// DeleteUser removes users from collection
func DeleteUser(id string) bool {
	session := dmas.MgoSession.Clone()
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	err := session.DB(dmas.DbName).C("users").RemoveId(id)

	if err != nil {
		return false
	}

	return true
}
