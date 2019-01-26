package crud

import (
    "testing"
    "log"
    "../types"
    // "database/sql"
    // _ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "person_master"
// 	password = "QWERTYU"
// 	dbname   = "person_register"
// )

func TestMain(t *testing.T) {
  db := Crud_initer_person()
  err := Check_connection(db)
  if err != nil{
    log.Fatalln("fail at connecion db by: ", err)
  }

  fail_person := types.Person{
    Id: int64(-1),
    First_name: "aaa",
    Last_name: "bbb",
    E_mail: "sdf@sdg.com",
    Phone: "6589246348t2g3408tg32408tg2304tg208g2t304t8g3408tgq3408tgq235",}
  err_forced := Person_inserter(fail_person, db)
  if err_forced == nil {
    log.Fatalln("fail at forced error db by: ", err_forced)
  }

  dummy_person := types.Person{
    Id: int64(-1),
    First_name: "aaa",
    Last_name: "bbb",
    E_mail: "sdf@sdg.com",
    Phone: "65892465",
  }
  err_insert := Person_inserter(dummy_person, db)
  if err_insert != nil {
    log.Fatalln("fail at insert by: ", err_insert)
  }

  empty_person := types.Person{}
  err_empty := Person_inserter(empty_person, db)
  if err_empty != nil {
    log.Fatalln("fail at err_empty by: ", err_empty)
  }

  sample_check := Person_reader(db)
  if len(sample_check) == 0{
    log.Fatalln("error for not select persons: ")
  }

  updateble_error := sample_check[0]
  updateble_error.Status = 99
  updateble_error.Crm = "fooooo"
  updateble_error.Report = "nothing"

  err_update, id_zero := Report_person_state(updateble_error, db)
  if err_update == nil && id_zero == 0 {
    log.Fatalln("fail at forced by: ", err_update)
  }

  updateble_person := sample_check[0]
  updateble_person.Status = 99
  updateble_person.Crm = "{\"first_name\":\"Katie\",\"last_name\":\"Bugs\",\"email\":\"kbugsk3@springer.com\",\"phone\":\"5341197577\"}"
  updateble_person.Report = "nothing"

  err_update, id_back := Report_person_state(updateble_person, db)
  if err_update != nil && id_zero != 0 {
    log.Fatalln("fail at update by: ", err_update)
  }
}
