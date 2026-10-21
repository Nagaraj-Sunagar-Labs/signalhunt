package scan

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Company struct {
	ID   int    `yaml:"value"`
	Name string `yaml:"name"`
}

var companies = []Company{}

func loadCompanyDetails() {
	/*
		01. In Bluetooth.com they maintain lookup of the company id and name lookup
		02. I have downloaded the file as of date 10-Oct-2026
		03. Kept in the data repo in the scan package
		04. It has lookup for the number of company and name of the company.
	*/
	//Users/nagarajsunagar/Developer/Engineering/coding-challenges/140/signalhunt/scan/data/company-lookup.yml
	data, err := os.ReadFile("/Users/nagarajsunagar/Developer/Engineering/coding-challenges/140/signalhunt/scan/data/company-lookup.yml")
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(data, &companies)
	println("Total Company Details Loaded: ", len(companies))
}

var companylookup = make(map[int]string)

func Createcompanieslookup() {
	loadCompanyDetails()
	for _, company := range companies {
		companylookup[company.ID] = company.Name
	}
}
