package crud

import (
    "../types"
  	"fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "person_master"
	password = "QWERTYU"
	dbname   = "person_register"
)

func Crud_initer_person() *sql.DB{
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  return db
}

func Check_connection(db *sql.DB) error{
  return db.Ping()
}

func Person_inserter(person types.Person, db *sql.DB) (error){
	sqlStInsert := `
									INSERT INTO person_pile (person_id, first_name, last_name, email, phone)
									VALUES ($1, $2, $3, $4, $5)
									`
  stmt, err := db.Prepare(sqlStInsert)
  if(err == nil){
    _, errf := stmt.Exec(person.Id, person.First_name, person.Last_name, person.E_mail, person.Phone)
  	return errf
  }else{
    return err
  }
}

func Report_person_state(person types.Person_DAS, db *sql.DB) (error, int64){
  report_statement := `
  UPDATE person_pile
     SET
         status = $2,
     crm_result = $3,
         report = $4
   WHERE
             id = $1
   RETURNING id;
  `

  var id int64
  err := db.QueryRow(report_statement, person.Id, person.Status, person.Crm, person.Report).Scan(&id)

	fmt.Println("resp update: ", id)
  return err, id
}

func Person_reader(db *sql.DB) []types.Person_DAS{
  selection_statement := `UPDATE person_pile SET status = 1 FROM (select id FROM person_pile where status = 0 limit 5) AS work_bloq where work_bloq.id = person_pile.id RETURNING person_pile.id, person_pile.status, person_pile.person_id, person_pile.first_name, person_pile.last_name, person_pile.email, person_pile.phone, person_pile.created_at, person_pile.updated_at;`
  rows, err := db.Query(selection_statement)
  var response []types.Person_DAS
  for rows.Next() {
      var work types.Person_DAS
      err = rows.Scan(&work.Id, &work.Status, &work.Person_id, &work.First_name, &work.Last_name, &work.E_mail, &work.Phone, &work.Created_at, &work.Updated_at)
      if(err == nil){
       response = append(response, work)
      }else{
       fmt.Println("err: ", err)
      }
   }
   return response
}
