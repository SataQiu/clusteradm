// Copyright Contributors to the Open Cluster Management project
package clustermanager

import (
	"k8s.io/cli-runtime/pkg/genericclioptions"
	genericclioptionsclusteradm "open-cluster-management.io/clusteradm/pkg/genericclioptions"
)

//Options: The structure holding all the command-line options
type Options struct {
	//ClusteradmFlags: The generic optiosn from the clusteradm cli-runtime.
	ClusteradmFlags *genericclioptionsclusteradm.ClusteradmFlags

	values          Values
	//The file to output the resources will be sent to the file.
	registry string
	//version of predefined compatible image versions
	bundleVersion string
	//If set, the command will hold until the OCM control plane initialized
	wait bool
	
	Streams genericclioptions.IOStreams
}

type BundleVersion struct {
	// registation image version
	RegistrationImageVersion string
	// placement image version
	PlacementImageVersion string
	// work image version
	WorkImageVersion string
	// operator image version
	OperatorImageVersion string
}

//Values: The values used in the template
type Values struct {
	//The values related to the hub
	Registry string `json:"registry"`
	//bundle version
	BundleVersion BundleVersion
	//Hub: Hub information
	Hub Hub
}

type Hub struct {
	//APIServer: The API Server external URL
	APIServer string
	//KubeConfig: The kubeconfig of the boostrap secret to connect to the hub
	KubeConfig string
	//image registry
	Registry string
}


func newOptions(clusteradmFlags *genericclioptionsclusteradm.ClusteradmFlags, streams genericclioptions.IOStreams) *Options {
	return &Options{
		ClusteradmFlags: clusteradmFlags,
		Streams:         streams,
	}
}
