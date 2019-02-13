package crm_database

import(
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"

)

type Credentials struct {
	Host string
	Port int
	User string
	Dbname string
}

func Init_pq_database_connection(creds Credentials)(db *sql.DB){
     //dev purposes only

	psql_info := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		creds.Host, creds.Port, creds.User, creds.Dbname)

	db, err := sql.Open("postgres", psql_info)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("it worked")
	return db
}

