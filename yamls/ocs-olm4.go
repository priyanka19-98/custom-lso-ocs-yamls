package yamls

import (
	"fmt"

	"github.com/ghodss/yaml"
)

// OcsSubscription is used to create Subscription for OCS Operator
func OcsSubscription() {
	s := []byte(`{
		"apiVersion": "operators.coreos.com/v1alpha1",
		"kind": "Subscription",
		"metadata":{
			"name": "ocs-subscription",
			"namespace":"openshift-storage",
		},
		"spec":{
			"channel":"alpha",
			"config":{
				"resources":{},
			},
			"name":"ocs-operator",
			"source": "ocs-catalogsource",
			"sourceNamespace":"openshift-marketplace",
		},
	}`)

	s, err := yaml.JSONToYAML(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(s))
	}
}
