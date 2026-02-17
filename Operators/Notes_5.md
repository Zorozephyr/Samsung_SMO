Controller-Runtime:
Manager -> Coordinates multiple controllers, Manages client connections and caching, handles leader election and provides unified entry point
Reconciler
Client is the specific Go interface that acts as the "bridge" between your code and the Kubernetes API server
Reads go to local informer based cache and writes go directly to API server. Handles Optimistic Concurrency when updating Resources

Manager Responsibilities
Manages Controllers: Registers and runs controllers
Manages Cache: Maintains local cache of resources
Manages Client: Provides client for API access
Manages Scheme: Handles API type registration
Leader Election: Ensures only one instance runs

func (r *MyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error)
ctx => context for cancellation and timeouts
req => Request containing namespace and name
Returns -> ctrl.Result(What to do next), error

Empty Result (ctrl.Result{}):

Reconciliation succeeded
No requeue needed
Controller will wait for next event
Requeue (ctrl.Result{Requeue: true}):

Reconciliation needs to run again
Requeues immediately
Use when you need to retry
RequeueAfter (ctrl.Result{RequeueAfter: time.Duration}):

Requeue after a delay
Useful for rate limiting
Example: ctrl.Result{RequeueAfter: 30 * time.Second}

func main() {
    // Create manager
    mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
        Scheme:                 scheme,
        MetricsBindAddress:     metricsAddr,
        Port:                   9443,
        HealthProbeBindAddress: probeAddr,
        LeaderElection:          enableLeaderElection,
    })
    
    // Setup reconciler
    if err := (&controllers.MyReconciler{
        Client: mgr.GetClient(),
        Scheme: mgr.GetScheme(),
    }).SetupWithManager(mgr); err != nil {
        // Handle error
    }
    
    // Start manager
    mgr.Start(ctrl.SetupSignalHandler())
}

Internal Dependencies (Kubernetes Resources)
This is when your Custom Resource (CR) depends on another object inside the same cluster.

Example: Your EC2Instance operator needs to pull a password from a Secret. If the user hasn't created that Secret yet, the Secret is a missing dependency.

Action: The controller sees the Secret is gone, logs a message like "Waiting for Secret," and returns RequeueAfter: 5 * time.Second.

External Dependencies (Cloud/Third-Party)
Since you are working on an EC2 Operator, this is your most common scenario.

Example: To create an EC2 instance, you need a VPC ID and a Subnet ID. If the AWS VPC is still in the "Pending" or "Creating" state, your instance cannot be launched yet.

Action: Your operator checks the AWS API, sees the VPC isn't "Available," and decides to wait.

Version Strategy
v1: Stable, production-ready
v1beta1: Beta, may change
v1alpha1: Alpha, experimental
Versioning Rules
Start with v1alpha1 for new APIs
Promote to v1beta1 when stable
Promote to v1 when production-ready
Support multiple versions during transition

Validation with markers
graph TB
    MARKER[Marker] --> VALIDATION[Validation Rule]
    MARKER --> DOC[Documentation]
    MARKER --> DISPLAY[Display Column]
    
    VALIDATION --> REQUIRED[Required]
    VALIDATION --> MINMAX[Min/Max]
    VALIDATION --> PATTERN[Pattern]
    VALIDATION --> ENUM[Enum]
    
    style VALIDATION fill:#FFB6C1


Example Validations:
// +kubebuilder:validation:Required
Message string `json:"message"`

// +kubebuilder:validation:Minimum=1
Replicas int32 `json:"replicas"`

// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`
Name string `json:"name"`