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

func InsertLink(link model.Link) (int64, int64, error) {
	sqlQuery := "INSERT link SET Title = ?, Address = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer closeStmt(stmt)

	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		return 0, 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return rowsAffected, lastInsertedId, err
}

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

// func closeRows(rows *sql.Rows) {
// 	if rows != nil {
// 		rows.Close()
// 	}
// }

func closeStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}
