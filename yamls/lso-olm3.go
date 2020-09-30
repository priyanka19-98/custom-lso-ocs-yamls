package yamls

import (
	"fmt"

	"github.com/ghodss/yaml"
)

// LsoOperatorSource is used to create OperatorSource for LSO
func LsoOperatorSource() {
	oS := []byte(`{
        "apiVersion": "operators.coreos.com/v1",
		"kind":"OperatorSource",
		"metadata":{
			"name":"redhat-local-storage-src",
			"namespace":"openshift-marketplace"
		},
		"spec":{
			"type":"appregistry",
			"endpoint":"https://quay.io/cnr",
			"registryNamespace":"hekumar",
			"displayName":"Red Hat Storage operators",
			"publisher":"Red Hat Storage",
		},
	}`)

	oS, err := yaml.JSONToYAML(oS)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(oS))
	}
}
