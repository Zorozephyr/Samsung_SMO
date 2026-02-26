Open Network Automation Platform (ONAP)
Anuket 
Sylva
Camara
Paraglider
CNTi
SONiC & DENT
OpenDaylight, OVS/OVN, FD.io, DPDK

Infrastructure CRs
These define the physical or virtual clusters where the network functions will run.

WorkloadCluster
Functionality: Represents the target Kubernetes cluster (e.g., an edge site).

How it works:

It contains connectivity details (SecretRef for kubeconfig) and metadata (labels like region: us-west).

Nephio uses this to know where to push configurations and how to specialize them (e.g., "This cluster uses SR-IOV, so configure the UPF accordingly").

3. Resource & Specialization CRs (The "Hydration" Layer)
This is Nephio's "secret sauce." Instead of hardcoding IPs, Nephio uses "Claims" inside the package.

IPClaim / VLANClaim
Functionality: A request for a network resource (IP address or VLAN ID) without specifying the exact value.

How it works:

Draft: The Blueprint contains an IPClaim (e.g., "I need an IP for the N3 interface").

Hydration: When the PackageVariant creates a draft, a specialized function (IPAM injector) sees this claim.

Allocation: The function talks to a backend (IPAM system), reserves an IP, and writes the specific IP into the package status or replaces the claim with the actual value.

Result: The final YAML that hits the cluster has a static IP, but the user never had to type it.

Capacity
Functionality: Defines the throughput or resource requirements for a Network Function (e.g., "I need 10 Gbps throughput").

How it works: Operators read this intent and translate it into low-level Kubernetes resources, such as CPU/Memory requests or SR-IOV Virtual Function allocations.

4. Workload CRs (The Network Functions)
These CRs end up on the Workload Cluster. They act as the "Instruction Manual" for the specific Network Function Operator.

NFDeployment (e.g., UPFDeployment, SMFDeployment)
Functionality: The declarative intent for a specific Network Function. It replaces standard K8s Deployments for Telco workloads.

How it works:

The UPFDeployment CR specifies the provider (e.g., free5gc), the version, and the interfaces (N3, N4, N6) linked to the IPClaims.

The Operator Loop: The specific NF Operator (running on the workload cluster) watches this CR. When it sees it, it generates the actual K8s Deployment, Service, ConfigMap, and NetworkAttachmentDefinition (Multus) required to bring up the 5G Core function.

The Question:
The SMO connects to the O-Cloud (Hardware/K8s) via O2.
But the Network Function (the App itself, like a DU) connects to the SMO via O1.

Why doesn't the App just talk to the O-Cloud directly? Or why doesn't the O2 interface handle the App configuration?

O1 is mainly for business logic, while O2 is for technical logic. O2 knows and sets up the technical details of the App, while O1 handles the business logic and policies.
