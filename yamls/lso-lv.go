package yamls

import (
	"fmt"

	"github.com/ghodss/yaml"
)

// LsoLocalVolumeSet is used to create LoalVolumeSet for LSO (Local Storage Operator)
func LsoLocalVolumeSet() {
	lvs := []byte(`{
		"apiVersion":"local.storage.openshift.io/v1alpha1",
        "kind": "LocalVolumeSet",
		"metadata":{
			"name": "local-disks",
			"namespace": "openshift-local-storage",
		},
		"spec":{
			"storageClassName":"localblock-sc",
			"volumeMode":"Block",
			"deviceInclusionSpec":
			{"deviceTypes": ["disk"]},
		},
	}`)

	lvs, err := yaml.JSONToYAML(lvs)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(lvs))
	}
}
