package main

import (
	yamls "main/yamls"
)

func main() {
	yamls.LsoNamespace()
	yamls.LsoOperatorGroup()
	yamls.LsoOperatorSource()
	yamls.LsoSubscription()
	yamls.LsoLocalVolumeSet()
	yamls.OcsNamespace()
	yamls.OcsOperatorGroup()
	yamls.OcsCatalogSource()
	yamls.OcsSubscription()
	yamls.OcsStorageCluster()
}
