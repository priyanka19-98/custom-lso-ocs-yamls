package yamls

import (
	"fmt"

	"github.com/ghodss/yaml"
)

//OcsStorageCluster is used to create StorageCluster for OCS Operator
func OcsStorageCluster() {
	sc := []byte(`{
	    "apiVersion":"ocs.openshift.io/v1",
		"kind":"StorageCluster",
        "metadata":{
          "name": "ocs-storagecluster",
		  "namespace":"openshift-storage",
		},
		"spec":{
			"labelSelector": {
				"matchExpressions":[],
			},
          "monDataDirHostPath":"/var/lib/rook",
		  "storageDeviceSets":[
			  {"count": 1,
			    "dataPVCTemplate":
				{
					"spec":{
						"accessModes":["ReadWriteOnce"],
						"resources": {
							"requests":{
								"storage": 30Gi,
							},
						},
						"storageClassName":"localblock-sc",
						"volumeMode":"Block",

					},


				},
				"name":"ocs-deviceset",
				"portable": true,
				"replica":3,
			  },
		  ]
		 	
		},
   }`)

	sc, err := yaml.JSONToYAML(sc)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(sc))
	}
}
