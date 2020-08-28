package postgres

import (
	"database/sql"
	//"golang.org/x/crypto/bcrypt"
	"github.com/Divan009/DumGo/graph/model"
	"golang.org/x/crypto/bcrypt"

	"fmt"
	"log"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (newUser *User) Create() {
	stmt, err := db.Prepare("INSERT INTO users(username,password) VALUES(?,?)")
	print(stmt)
	if err != nil {
		log.Fatal(err)
	}
	hashedPswd, err := HashPassword(newUser.Password)
	_, err = stmt.Exec(newUser.Username, hashedPswd)
	if err != nil {
		log.Fatal(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// INSERT INTO user_new (name) VALUES ($1);

// sqlQuery := "INSERT INTO todo (text, done) VALUES ($1, $2);"
// 	stmt, err := db.Query(sqlQuery, input.Text, input.Done)

//issue with the return

func InsertLink(link model.Link) (int64, error) {
	sqlQuery := "INSERT INTO links (title, address) VALUES ($1, $2);"
	stmt, err := db.Prepare(sqlQuery)
	defer closeStmt(stmt)

	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	// lastInsertedId, err := res.LastInsertId()
	// fmt.Println(lastInsertedId)
	return rowsAffected, err
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
