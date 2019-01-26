package types
import (
  "time"
)
type Person_DAS struct {
  Id            int64      `db:"id"`
  Status        int64      `db:"status"`
  Crm           string     `db:"crm_result"`
  Person_id     int64      `db:"person_id"`
  First_name		string	   `db:"first_name"`
  Last_name			string	   `db:"last_name"`
  E_mail				string	   `db:"email"`
  Phone					string	   `db:"phone"`
  Report				string	   `db:"report"`
  Created_at    time.Time  `db:"created_at"`
  Updated_at    time.Time  `db:"updated_at"`
}

type public_Data interface {
  GetPublicPerson()
}

type Public_Person struct {
  First_name		string	   `json:"first_name"`
  Last_name			string	   `json:"last_name"`
  E_mail				string	   `json:"email"`
  Phone					string	   `json:"phone"`
}

func (person Person_DAS) GetPublicPerson() Public_Person  {
  return Public_Person{
    First_name: person.First_name,
    Last_name: person.Last_name,
    E_mail: person.E_mail,
    Phone: person.Phone,
  }
}
