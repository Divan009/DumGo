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

func Create(newUser User) (string, error) {
	stmt, err := db.Prepare("INSERT INTO users(username,password) VALUES(?,?)")
	print(stmt)
	if err != nil {
		log.Print(err)
		return "", err
	}
	hashedPswd, err := HashPassword(newUser.Password)
	_, err = stmt.Exec(newUser.Username, hashedPswd)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return "success", nil
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

//GetUserIdByUsername check if a user exists in database by given username

func GetUserIdByUsername(username string) (int, error) {
	sqlQuery := "SELECT id FROM users WHERE username = ?"
	statement, err := db.Prepare(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)
	//row := statement.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
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
