package yamls

import (
	"fmt"

	"github.com/ghodss/yaml"
)

//LsoNamespace is used to create namespace for LSO
func LsoNamespace() {
	n := []byte(`{"apiVersion":"v1",
	"kind":"Namespace",
	"metadata": {
		"name": "openshift-local-storage",
	}
	}`)

	n, err := yaml.JSONToYAML(n)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(n))
	}
}
