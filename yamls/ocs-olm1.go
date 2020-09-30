package yamls

import (
	"fmt"

	"github.com/ghodss/yaml"
)

// OcsNamespace is used for creating namespace for OCS operator
func OcsNamespace() {

	ns := []byte(`{
		  "apiVersion":"v1",
		  "kind":"Namespace",
		  "metadata":{
			  "name":"openshift-storage",
		  }
	  }`)

	ns, err := yaml.JSONToYAML(ns)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(ns))
	}

}
