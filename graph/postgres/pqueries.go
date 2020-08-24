package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Divan009/DumGo/graph/model"
)

// func InsertToOrderQueue(db *sql.DB, orderQueueId string, timeToLive string, queueInput *models.NewOrderQueue) (string, error) {
// 	query := `insert into public.orderQueue ("addedToQueueAt", "orderCreateTime", "orderQueueId", requestid, soldto, shipto, payload, timetolive, remarks, reason, type) values
// 				('` + queueInput.AddedToQueueAt + `','` + queueInput.OrderCreatedTime + `','` + orderQueueId + `','` + queueInput.Requestid + `','` + queueInput.SoldTo + `','` + queueInput.ShipTo + `','` + queueInput.Payload + `','` + timeToLive + `','` + *queueInput.Remarks + `','` + queueInput.Reason + `','` + queueInput.Type + `')`
// 	log.Println(query)
// 	loguid := GetLoggerID()
// 	_, err := db.Exec(query)
// 	if err != nil {
// 		Logs(WARN, loguid, "InsertToOrderQueue: "+err.Error()+"\nQuery used: "+query)
// 		return "", err
// 	}
// 	return orderQueueId, nil

// func InsertLink(db *sql.DB) int64 {
// 	stmt, err := db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	res, err := stmt.Exec(link.Title, link.Address)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		log.Fatal("Error:", err.Error())
// 	}
// 	log.Print("Row inserted!")
// 	return id
// }

func GetLink(db *sql.DB) ([]*model.Link, error) {

	sqlStatement := `SELECT id, "Title", "Address" FROM links;`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	linkss := make([]*model.Link, 0)
	for rows.Next() {
		role := new(model.Link)
		err := rows.Scan(&role.ID, &role.Title, &role.Address)
		if err != nil {
			fmt.Println("Unable to extract the data")
			return linkss, err
		} else {
			inputUser := model.Link{
				role.ID,
				role.Title,
				role.Address,
				role.User}
			log.Println()

			linkss = append(linkss, &inputUser)
		}
	}
	return linkss, nil
}
