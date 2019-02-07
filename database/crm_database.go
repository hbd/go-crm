package crm_database

import(
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"

)

func init_database_connection(){
     //dev purposes only
     const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		dbname   = "crm"
   )

   psql_info := fmt.Sprintf("host=%s port=%d user=%s "+
   "dbname=%s sslmode=disable",
   host, port, user, dbname)

   db, err := sql.Open("postgres", psql_info)
   if err != nil {
   panic(err)
   }
   defer db.Close()

   err = db.Ping()
   if err != nil {
   panic(err)
   }
}

