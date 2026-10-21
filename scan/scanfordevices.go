package scan

import (
	"fmt"
	"time"

	"tinygo.org/x/bluetooth"
)

var (
	adapter          = bluetooth.DefaultAdapter
	DEFAULT_SCAN_SEC = int64(5)
)

func ScanForDevices() {
	/*
		Loading a Compay details from the yml file downloaded for bluetooth dataset
		01. loadCompayDetails is method to load the data to the struct called Company
		02. And that is loaded as list of the companies, with values ID and Name of that company
	*/
	println("Loading Company Details!")
	Createcompanieslookup()
	println("Company details loaded correctly!")
	/* TEST: to check company details loaded correctly
	count := 2
	for _, company := range companies {
		println("CompanyID: ", company.ID)
		println("CompanyName: ", company.Name)
		count++
		if count == 1 {
			break
		}
	}*/
	println("Scanning for devices...")
	/*Enabled the bluetooth stack to scan for the all the visible bluetooth devices near this device*/
	must("Enable BLE Stack", adapter.Enable())
	/*
		Requirement is need to store a unique list of bluetoth devices hashmap (datastructure)
		01. initialize the map
		02. map with addressstring as key, localname ans the value
		03. if the address already visible in the scan don't add it
		04. if the witoought name divices is getting detected everytime add it make name is NONAME
		05. once most of them are added stop the scan from the bluetooth adapter scan.
	*/
	uniquedevices := make(map[string]string)
	starttime := time.Now().Unix()
	endtime := starttime + DEFAULT_SCAN_SEC
	fmt.Println("Start time of scan: ", starttime)
	fmt.Println("End time of scan: ", endtime)
	err := adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		address := result.Address.String()
		devicename := result.LocalName()
		signalstrength := result.RSSI
		if signalstrength >= -60 {
			println("---------------------")
			println("Device Address: ", result.Address.String())
			println("Signal Strength: ", result.RSSI) // RSSI [Received Signal Strength Indicator] - it shows how strong bluetooth signals
			/*
				RSSI - is measured in decibels relative to 1milliwat(dBm)
				-30 dBm [Stronger] -80 dBm[Weaker] , it means value close to O are the one which are stronger bluetooth signal
				why the values are in negative?
				- signals received from bluetooth devices are lessar then 1mW
				- to represent them in some proper representation it is measured interms of the 1mW with formula
				  dBm = 10 log10(p/1mw)
			*/
			if result.LocalName() != "" {
				println("Local Name of Device: ", result.LocalName())
			}
			for _, manufactor := range result.ManufacturerData() {
				/*fmt.Println("Company ID: ", manufactor.CompanyID)*/
				fmt.Println("Company Name: ", companylookup[int(manufactor.CompanyID)])
			}
			println("----------------------")
			uniquedevices[address] = devicename
		}
		if time.Now().Unix() > endtime {
			adapter.StopScan()
		}

	})
	if err != nil {
		panic("Error while getting details!")
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
