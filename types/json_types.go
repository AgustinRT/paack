package types

type Person struct {
	Id						int64		`json:"id"`
	First_name		string	`json:"first_name"`
	Last_name			string	`json:"last_name"`
	E_mail				string	`json:"email"`
	Phone					string	`json:"phone"`
}
