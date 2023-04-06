package kubernetes

// podSessionOptions provides a struct to assign parameters to an exec session
type PodSessionOptions struct {
	Command    []string
	Namespace  string
	PodName    string
	Stdin      bool
	Stdout     bool
	Stderr     bool
	TtyEnabled bool
}

type SyncEcrCmdOptions struct {
	Namespace           string
	Region              string
	RegistryURL         string
	KubeInClusterConfig bool
}
