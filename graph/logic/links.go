package logic

import (
	"log"
)

// type Link struct {
// 	ID      string
// 	Title   string
// 	Address string
// 	User    *users.User
// }

//#2
func (link Link) Save() int64 {
	//#3
	stmt, err := database.db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	//#4
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}
	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}
