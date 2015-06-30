package restdemo

import (
 "net/http"
 "fmt"
)


func main(){
  http.ListenAndServe("localhost:6666")

}
