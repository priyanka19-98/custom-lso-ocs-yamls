package yamls

import (
	"fmt"

	"github.com/ghodss/yaml"
)

// LsoSubscription is used to create Subsciption for LSO
func LsoSubscription() {
	s := []byte(`{
		"apiVersion":"operators.coreos.com/v1alpha1",
		"kind":"Subscription",
		"metadata":{
			"name": "local-storage-operator",
			"namespace":"openshift-local-storage",
		},
		"spec":{
			"channel": "4.6",
  "installPlanApproval": "Automatic",
  "name": "local-storage-operator",
  "source": "redhat-local-storage-src",
  "sourceNamespace": "openshift-marketplace",
		},

	}`)

	s, err := yaml.JSONToYAML(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(s))
	}
}
