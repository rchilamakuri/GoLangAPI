//This edit is only made for checking the branches for this reppository is 
//working or not
package main

import(
   "html/template"
   "log"
   "net/http"
   "time"
)

type PageVariables struct {

       Date    string
       Time    string


}

func main(){

         http.HandleFunc("/", HomePage)
         log.Fatal(http.ListenAndServe(":8080", nil))   


}

func HomePage(w http.ResponseWriter, r *http.Request){

       now := time.Now()
       HomePageVars := PageVariables{
         Date: now.Format("02-01-2006"), 
         Time: now.Format("15:04:05"),   

}

       t, err := template.ParseFiles("homepage.html")
       if err != nil{
              log.Print("tmeplate parsing error: ", err)
              
}   

      err = t.Execute(w, HomePageVars)
      if err !=nil {
               log.print("template executing error: ", err)

 } 



}
