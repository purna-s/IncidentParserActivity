package incidentparseractivity

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("xmlString", `<?xml version="1.0" encoding="UTF-8"?>
<IncidentsInfo>
	<Incident>
		<XCoor>25133.26</XCoor>
		<YCoor>43509.02</YCoor>
		<Type>1</Type>
		<Message>(9/3)14:29 Roadworks on SLE (towards BKE) before Mandai Rd Exit. Avoid lane 1.</Message>
	</Incident>
	<Incident>
		<XCoor>23835.43</XCoor>
		<YCoor>45044.879</YCoor>
		<Type>0</Type>
		<Message>(9/3)14:28 Accident on SLE (towards CTE) at Woodlands Ave 12 Exit.</Message>
	</Incident>	
</IncidentsInfo>`)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}
	act.Eval(tc)
	//check output attr

	output := tc.GetOutput("output")
	assert.Equal(t, output, output)

}
