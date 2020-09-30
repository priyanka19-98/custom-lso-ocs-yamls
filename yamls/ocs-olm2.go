package yamls

import (
	"fmt"

	"github.com/ghodss/yaml"
)

// OcsOperatorGroup is used to create OperatorGroup for OCS Operator
func OcsOperatorGroup() {
	og := []byte(`{
		"apiVersion":"operators.coreos.com/v1",
		"kind":"OperatorGroup",
		"metadata":{
			"name":"openshift-storage-operatorgroup",
			"namespace":"openshift-storage",
		},
		"spec":{
			"targetNamespaces":["openshift-storage"]
		}
	}`)

	og, err := yaml.JSONToYAML(og)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(og))
	}
}
