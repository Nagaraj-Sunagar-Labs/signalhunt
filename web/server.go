package web

import (
	"fmt"
	"net/http"
)

func getBluetoothDetails(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Get Request for Bluetooth Device Details Nearby")
}

func main() {
	http.HandleFunc("/getdetails", getBluetoothDetails)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
