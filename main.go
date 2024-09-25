package main

import (
    "log"
    "os"
    "math/rand"
)

func genKey string (){
  return ""
}
 
func main() {
    psw_correct := "SueRSaV3pa$$W0rD"

    file, err := os.Open("password.txt")
    
    if err != nil{
        log.Fatalln(err)  
    }
    
    defer file.Close() 
     
    data := make([]byte, 64)
    n, err := file.Read(data)

    buf := string(data[:n])

    if(buf == psw_correct){
      log.Println("correct")
      data := []byte("KEY&&")
      file.Write()
    }else{
      log.Println("incorrect")
    }
}