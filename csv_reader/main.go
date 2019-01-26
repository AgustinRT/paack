package main

import (
	"bufio"
	"encoding/csv"
  "database/sql"
	"strings"
  "../types"
  "../crud"
  "strconv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"bytes"
	// "time"
)

func main() {
	// theorical path read the csv
  path := "../resources/MOCK_DATA.csv"
  choose_behaviour(crud.Crud_initer_person(), path)
}

func choose_behaviour(db *sql.DB, path string)  {
	// check if the DDBB is avaliable. if isn't, launch a script for restore it
  if crud.Check_connection(db) == nil{
    executer_csv(path, db)
  }else{
		cmd := exec.Command("bash", "../postgres_workspace.sh")
		cmd.Stdin = strings.NewReader("");
		var out bytes.Buffer;
		cmd.Stdout = &out;
		cmd.Run();
		choose_behaviour(crud.Crud_initer_person(), path)
  }
}

func executer_csv(path string, db *sql.DB) {
  csvFile, _ := os.Open(path)
  reader := csv.NewReader(bufio.NewReader(csvFile))
  work_pile := []types.Person{}
  for {
    record, err := reader.Read()
    if err == io.EOF {
      fmt.Println("is the end: ", err)
      break
    }
    if err != nil {
      log.Fatal(err)
    }
    new_person, err_p := parser_to_Person(record)

    if err_p != nil {
      fmt.Println("error in person", err_p)
    }else{
      crud.Person_inserter(new_person, db)
      work_pile = append(work_pile, new_person)
    }
  }

  fmt.Println("persons inserted: ", len(work_pile))
}

func parser_to_Person(line_input []string) (types.Person, error){
		new_phone := strings.Replace(line_input[4], " ", "", -1)
    new_id, err := strconv.ParseInt(line_input[0], 10, 64)
    _, err2 := strconv.ParseInt(line_input[0], 10, 64)
		if(err != nil){
			fmt.Println(err)
			return types.Person{}, err
		}
		if(err2 != nil){
			fmt.Println(err2)
			return types.Person{}, err2
		}

		return types.Person{
			Id: new_id,
			First_name: line_input[1],
			Last_name: line_input[2],
			E_mail: line_input[3],
			Phone: new_phone,
		}, err
}
