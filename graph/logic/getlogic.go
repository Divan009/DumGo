package logic

import (
	"fmt"
	"strconv"

	"github.com/Divan009/DumGo/graph/model"
	"github.com/Divan009/DumGo/graph/postgres"
)

// getting the link

func GetLinks() ([]*model.Link, error) {
	link, err := postgres.GetLink(postgres.GetDb())
	if err != nil {

		fmt.Println(err)
	}
	return link, nil

}

func AddLink(title string, address string) (*model.Link, error) {
	link := model.Link{Title: title, Address: address}
	rowsAffected, err := postgres.InsertLink(link)
	// if err == nil && rowsAffected > 0
	link.ID = strconv.FormatInt(rowsAffected, 10)
	return &link, err
}
