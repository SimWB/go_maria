package go_maria
import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql" )

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}
