package logic

import (
	"fmt"

	"github.com/Divan009/DumGo/graph/model"
	"github.com/Divan009/DumGo/graph/postgres"
)

func GetLinks() ([]*model.Link, error) {
	link, err := postgres.GetLink(postgres.GetDb())
	if err != nil {

		fmt.Println(err)
	}
	return link, nil

}

// func InsertToOrderQueue(db *sql.DB) (string, error) {
// 	db := postgres.GetDb()
// 	timeToLive := os.Getenv("TIME_TO_LIVE")
// 	addOrderQueueResponse, err := postgres.InsertLink(title, address)
// 	if err != nil {
// 		// Logs(ERROR,err.Error())
// 		return string, err
// 	}
// 	response := string
// 	response.OrderQueueID = addOrderQueueResponse
// 	return response, nil
// }
