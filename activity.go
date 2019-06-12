package incidentparseractivity

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-incidentparseractivity")

// MyActivity is a stub for your Activity implementation
type XMLParserActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &XMLParserActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *XMLParserActivity) Metadata() *activity.Metadata {
	return a.metadata
}

//XSD
type Incidentsinfo struct {
	XMLName      xml.Name   `xml:"IncidentsInfo" json:"-"`
	IncidentList []Incident `xml:"Incident" json:"Incident"`
}

type Incident struct {
	XMLName xml.Name `xml:"Incident" json:"-"`
	XCoor   string   `xml:"XCoor" json:"XCoor"`
	YCoor   string   `xml:"YCoor" json:"YCoor"`
	Type    string   `xml:"Type" json:"Type"`
	Message string   `xml:"Message" json:"Message"`
}

// end of XSD

// Eval implements activity.Activity.Eval
func (a *XMLParserActivity) Eval(ctx activity.Context) (done bool, err error) {

	XMLString := ctx.GetInput("xmlString").(string)

	activityLog.Debugf("XML String is : [%s]", XMLString)
	//fmt.Println("XML String is : ", XMLString)

	if len(XMLString) == 0 {
		activityLog.Debugf("value in the field is empty ")
		//fmt.Println("value in  the field is empty ")

	}
	//	XMLString = (string(XMLString))

	xml_data := Incidentsinfo{}
	err = xml.Unmarshal([]byte(XMLString), &xml_data)

	jsondata, _ := json.Marshal(xml_data)
	if err != nil {
		activityLog.Debugf("Error ", err)
		fmt.Println("error: ", err)
		return
	}

	//fmt.Println(" JSON String ")
	//fmt.Println(string(jsondata))

	// Set the output as part of the context
	activityLog.Debugf("Activity has parsed XML Successfully")
	fmt.Println("Activity has parsed XML Successfully")

	ctx.SetOutput("output", string(jsondata))

	return true, nil
}
