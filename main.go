package main

import (
	"go_test/controllers"
	"go_test/rmq"
)

func main() {
	//router := mux.NewRouter()
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "3000"
	//}
	//router.HandleFunc("/api/log/http/", controllers.CreateHttpLog).Methods("POST")
	var rc rmq.RabbitClient
	rc.Consume("test_queue", controllers.CreateHttpLog)
	//err := http.ListenAndServe(":"+port, router)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
