Regional Represents the cloud cluster and control plane
Edge01 represents User plane and closer to end user cluster

UE(End User) -> End User.In a lab/sandbox, this is usually a software simulator (like UERANSIM) pretending to be a 5G smartphone.

gNodeB(5G base station) -> 5G Base Station. In a lab/sandbox, this is usually a software simulator (like Free5GC) pretending to be a 5G base station.

UPF(user plane function) -> high speed router, SIM -> GNB -> UPF -> Internet. UPF tries to make this flow as fast as possible

AMF(Access and Mobility Management Function) -> Authentication, Authorization, Registration, Mobility...So when you turn ur phone on, it will decide if ur phone is allowed to connect to the network

SMF(Session Management Function) -> It sets up tunnel for the user. This user is watching a video; please route their traffic to the internet

AUSF (Authentication Server Function): Security. It runs the complex cryptography to prove the SIM card is valid.

NRF (Network Repository Function): The Phonebook. It keeps a list of all running network functions so they can find each other (e.g., the AMF asks NRF, "Where is the SMF?").

NSSF (Network Slice Selection Function): The Slicer. If the network is cut into "slices" (one for gaming, one for IoT), this component decides which slice the user gets put into.

UDM (Unified Data Management): The Frontend Database. It processes user data (subscription info, allowed services).

UDR (Unified Data Repository): The Backend Storage. The actual database where the UDM and PCF store their data.

PCF (Policy Control Function): The Lawyer. It decides the rules: "User X has a Gold plan, give them high speed," or "User Y is out of data, throttle them."

WebUI: This is likely a specific dashboard for the free5gc lab to visualize subscribers; it is not a standard 3GPP network function.


porchctl rpkg clone -n default catalog-infra-capi-b0ae9512aab3de73bbae623a3b554ade57e15596 --repository mgmt regional
The Action: It commands Porch to clone a generic Kubernetes cluster blueprint (infra-capi...) from the public catalog into your local mgmt repository.
The Naming: It renames this specific copy "regional" so it serves as the dedicated configuration for your Regional site (the green box in your diagram).
The Result: It creates a "Draft" status package in Git, which is now ready for you to customize (e.g., set IP addresses) and approve to trigger the actual cluster creation.

1. Context (Context: kind-kind)
"The Connection Setting"

What it is: This is a configuration in your kubeconfig file. It tells your terminal (and k9s) where to send commands.

In your screen: It is the top-left line. It means "I am currently talking to the Management Cluster."

Analogy: This is the Phone Number you dialed to reach the operator.

2. Cluster (Cluster: kind-kind)
"The Management Cluster" (The Mother Ship)

What it is: This is the actual, running Kubernetes infrastructure you are logged into right now. This is the "Brain" of Nephio.

In your screen: It represents the VM/Container that is running the Nephio software.

Analogy: This is the Factory where the robots live.

3. Cluster Resource (NAME: regional)
"The Workload Cluster" (The Product)

What it is: This is a Custom Resource (CR). In standard Kubernetes, you have resources like Pod, Service, or Deployment. In Nephio (via Cluster API), you have a resource type called Cluster.

The Big Difference: This is not a cluster you are inside. It is a text file (YAML) sitting inside the Management Cluster that says: "Please build me a new cluster named 'regional'."

In your screen: The line regional | Provisioned tells you that the Management Cluster has successfully built the "regional" cluster.

Will it be in my Contexts?
No, not automatically.

If you run kubectl config get-contexts, you will likely still only see kind-kind (the Management Cluster).

Why? The Management Cluster created the new cluster, but it kept the "keys" (the kubeconfig) to itself inside a Secret. It doesn't touch your personal ~/.kube/config file.

kubectl get secret regional-kubeconfig -o jsonpath='{.data.value}' | base64 -d > $HOME/.kube/regional-kubeconfig
export KUBECONFIG=$HOME/.kube/config:$HOME/.kube/regional-kubeconfig

Featuremgmt Repo (The Order Form)mgmt-staging Repo (The Factory Floor)AnalogyThe Menu. You order "1 Regional Cluster" here.The Kitchen. The chefs chop vegetables, cook the steak, and plate the food here.What's inside?High-level "Intent" (e.g., "Build a cluster named regional").Low-level "Details" (e.g., The specific CNI config, the exact storage class, the IP pool logic).Who writes to it?You (via Porch/CLI). You approve the high-level design.The Automation (Porch Controllers). You rarely touch this manually.Who reads it?Porch (to know what to build).Config Sync (to apply the final result to the cluster).

