package main

import (
	"encoding/json"
  "../types"
  "../crud"
  "../rest"
  "database/sql"
)

func main() {
  db := crud.Crud_initer_person()
  response := crud.Person_reader(db)
  for len(response) > 0{
    cicle_manager(response, db)
    response = crud.Person_reader(db)
  }
}

func cicle_manager(persons []types.Person_DAS, db *sql.DB) {
  for _, person := range persons{
    go manage_crm_communication(db, person)
  }
}

func manage_crm_communication(db *sql.DB, person types.Person_DAS) (error, int64){
	// get a json from the person struct
  person_structure := person_json_parser(person)
	// try to send a post with the json
  result, report := rest.Do_post(person_structure.Crm)
  if(result == 400){
		// if the message fail, report it to the ddbb and retry it till succes the message
    person_structure.Status = 3
    person_structure.Report = report
    crud.Report_person_state(person_structure, db)
    return manage_crm_communication(db, person_structure)
  }
	// if the message succes, report it and finnish the routine
  person_structure.Status = 2
  person_structure.Report = "Send success"
  return crud.Report_person_state(person_structure, db)
}

func person_json_parser(person types.Person_DAS) types.Person_DAS{
  public := person.GetPublicPerson()
  x, err := json.Marshal(public)
  if(err == nil){
    person.Crm = string(x)
  }
  return person
}
