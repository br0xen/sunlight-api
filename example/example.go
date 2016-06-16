package main

import (
	"fmt"
	"os"

	"github.com/br0xen/sunlight-api/openstates"
)

func main() {
	var apiKey string
	if len(os.Args) > 1 {
		apiKey = os.Args[1]
	} else {
		fmt.Println("Usage: example <apikey>")
		return
	}
	o := openstates.Create(apiKey)
	//d, err := o.StateMetadata(states.Kansas)
	//d, err := o.GetBillDetailFromID("KSB00002165")
	//d, err := o.GetBillDetail(states.Kansas, "2013-2014", "HR 6020")
	//d, err := o.GetLegislatorsForState(states.Kansas)
	//d, err := o.GetLegislatorDetail("KSL000018")
	d, err := o.GetLegislatorsForGeo(35.79, -78.78)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v\n", d)
}
