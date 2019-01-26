package rest

import (
  "testing"
  "log"
  )

func TestMain(t *testing.T) {
  code_fail, _ := Do_post("foo")
  if(code_fail != 400){
    log.Fatalln("code_fail most by 400 and not: ", code_fail)
  }

  code_success, resp := Do_post("{\"first_name\":\"Katie\",\"last_name\":\"Bugs\",\"email\":\"kbugsk3@springer.com\",\"phone\":\"5341197577\"}")
  if(code_success == 400 && resp != "xxxxxxxx"){
      log.Println(code_success)
      log.Println(resp)
    log.Fatalln("code_success most by 200 and not: ", code_success)
  }
}
