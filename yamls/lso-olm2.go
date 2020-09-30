package yamls

import (
	"fmt"

	"github.com/ghodss/yaml"
)

// LsoOperatorGroup is used to define OperatorGroup for LSO
func LsoOperatorGroup() {
	oG := []byte(`{
		"apiVersion":"operators.coreos.com/v1alpha2",
	     "kind":"OperatorGroup",
	    "metadata":
	            {
					 "name": "local-operator-group",
					"namespace":"openshift-local-storage"
				},
		"spec":
		      {
				  "tagetNamespaces":
				  ["openshift-local-storage"]
				},
	 }`)

	oG, err := yaml.JSONToYAML(oG)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(oG))
	}
}
