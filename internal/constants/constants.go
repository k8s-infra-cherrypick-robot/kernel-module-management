package constants

const (
	ModuleNameLabel      = "kmm.node.kubernetes.io/module.name"
	NodeLabelerFinalizer = "kmm.node.kubernetes.io/node-labeler"
	TargetKernelTarget   = "kmm.node.kubernetes.io/target-kernel"
	PodType              = "kmm.node.kubernetes.io/pod-type"
	PodHashAnnotation    = "kmm.node.kubernetes.io/last-hash"
	KernelLabel          = "kmm.node.kubernetes.io/kernel-version.full"
	DaemonSetRole        = "kmm.node.kubernetes.io/role"

	WorkerPodVersionLabelPrefix    = "beta.kmm.node.kubernetes.io/version-worker-pod"
	DevicePluginVersionLabelPrefix = "beta.kmm.node.kubernetes.io/version-device-plugin"
	ModuleVersionLabelPrefix       = "kmm.node.kubernetes.io/version-module"

	ModuleFinalizer   = "kmm.node.kubernetes.io/module-finalizer"
	JobEventFinalizer = "kmm.node.kubernetes.io/job-event-finalizer"

	ManagedClusterModuleNameLabel  = "kmm.node.kubernetes.io/managedclustermodule.name"
	KernelVersionsClusterClaimName = "kernel-versions.kmm.node.kubernetes.io"
	DockerfileCMKey                = "dockerfile"
	PublicSignDataKey              = "cert"
	PrivateSignDataKey             = "key"

	ModuleLoaderRoleLabelValue = "module-loader"

	OperatorNamespaceEnvVar = "OPERATOR_NAMESPACE"
)
