package types

type Details struct {
	Message    string            `json:"message,omitempty"`
	Events     interface{}       `json:"events,omitempty"`
	Logs       interface{}       `json:"logs,omitempty"`
	Object     interface{}       `json:"object,omitempty"`
	Describes  []*Describe       `json:"describes,omitempty"`
	Resources  []*ResourceStatus `json:"resources,omitempty,omitempty"`
	Conditions []*Condition      `json:"conditions,omitempty"`
}

type Describe struct {
	Name   string      `json:"name,omitempty"`
	Kind   string      `json:"kind,omitempty"`
	Events interface{} `json:"events,omitempty"`
	Logs   interface{} `json:"logs,omitempty"`
	Object interface{} `json:"object,omitempty"`
}

type Condition struct {
	// Type of condition.
	Type ConditionType `json:"type,omitempty"`
	// Status of the condition, one of True, False, Unknown.
	Status ConditionStatus `json:"status,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"last_transition_time,omitempty"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"last_update_time,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`
}

type ConditionType string

const (
	ConditionGetClient             ConditionType = "GetClient"
	ConditionGetTemplate           ConditionType = "GetTemplate"
	ConditionApplied               ConditionType = "Applied"
	ConditionDestroy               ConditionType = "Destroy"
	ConditionOriginDestroyAndCheck ConditionType = "OriginDestroyAndCheck"
	ConditionPreResourceClear      ConditionType = "PreResourceClear"
	ConditionMiddleResourceClear   ConditionType = "MiddleResourceClear"
	ConditionPostResourceClear     ConditionType = "PostResourceClear"
	ConditionHealthyChecked        ConditionType = "HealthyChecked"
	ConditionMigrateDataSave       ConditionType = "MigrateDataSave"
	ConditionDataSave              ConditionType = "DataSave"
	ConditionFinish                ConditionType = "Finish"
)

type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
const (
	ConditionTrue        ConditionStatus = "true"
	ConditionFalse       ConditionStatus = "false"
	ConditionUnknown     ConditionStatus = "unknown"
	ConditionProgressing ConditionStatus = "progressing"
)

type ResourceStatus struct {
	Group           string         `json:"group,omitempty" `
	Version         string         `json:"version,omitempty" `
	Kind            ResourceKind   `json:"kind,omitempty"`
	Namespace       string         `json:"namespace,omitempty" `
	Name            string         `json:"name,omitempty"`
	Status          SyncStatusCode `json:"status,omitempty"`
	Health          *HealthStatus  `json:"health,omitempty"`
	Hook            bool           `json:"hook,omitempty"`
	RequiresPruning bool           `json:"requiresPruning,omitempty"`
}

const (
	ClusterKind            ResourceKind = "Cluster"
	CloneSetKind           ResourceKind = "CloneSet"
	ClusterRoleBindingKind ResourceKind = "ClusterRoleBinding"
	ClusterRoleKind        ResourceKind = "ClusterRole"
	ConfigMapKind          ResourceKind = "ConfigMap"
	CronJobKind            ResourceKind = "CronJob"
	DeploymentKind         ResourceKind = "Deployment"
	NetworkPolicyKind      ResourceKind = "NetworkPolicy"
	IngressKind            ResourceKind = "Ingress"
	JobKind                ResourceKind = "Job"
	ApplicationKind        ResourceKind = "Application"
	NamespaceKind          ResourceKind = "Namespace"
	PVCKind                ResourceKind = "PersistentVolumeClaim"
	ResourceQuotaKind      ResourceKind = "ResourceQuota"
	RoleKind               ResourceKind = "Role"
	TCPIngressKind         ResourceKind = "TCPIngress"
	SidecarSetKind         ResourceKind = "SidecarSet"
	ServiceAccountKind     ResourceKind = "ServiceAccount"
	ServiceKind            ResourceKind = "Service"
	SecretKind             ResourceKind = "Secret"
	RoleBindingKind        ResourceKind = "RoleBinding"
	OamApplicationKind     ResourceKind = "Application"
)

type ResourceKind string

type SyncStatusCode string

// Possible comparison results
const (
	SyncStatusCodeUnknown   SyncStatusCode = "Unknown"
	SyncStatusCodeSynced    SyncStatusCode = "Synced"
	SyncStatusCodeOutOfSync SyncStatusCode = "OutOfSync"
)

type HealthStatus struct {
	Status  HealthStatusCode `json:"status,omitempty" protobuf:"bytes,1,opt,name=status"`
	Message string           `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
}

// HealthStatusCode Represents resource health status
type HealthStatusCode string

const (
	// HealthStatusUnknown Indicates that health assessment failed and actual health status is unknown
	HealthStatusUnknown HealthStatusCode = "Unknown"
	// HealthStatusProgressing Progressing health status means that resource is not healthy but still have a chance to reach healthy state
	HealthStatusProgressing HealthStatusCode = "Progressing"
	// HealthStatusHealthy Resource is 100% healthy
	HealthStatusHealthy HealthStatusCode = "Healthy"
	// HealthStatusSuspended Assigned to resources that are suspended or paused. The typical example is a
	// [suspended](https://kubernetes.io/docs/tasks/job/automated-tasks-with-cron-jobs/#suspend) CronJob.
	HealthStatusSuspended HealthStatusCode = "Suspended"
	// HealthStatusDegraded Degrade status is used if resource status indicates failure or resource could not reach healthy state
	// within some timeout.
	HealthStatusDegraded HealthStatusCode = "Degraded"
	// HealthStatusMissing Indicates that resource is missing in the cluster.
	HealthStatusMissing HealthStatusCode = "Missing"
)

type CilDeployChain struct {
	TraceId      string `json:"trace_id,omitempty"`
	Action       string `json:"action,omitempty"`
	ParentUniqid string `json:"-"`
	Uniqid       string `json:"-"`
}
