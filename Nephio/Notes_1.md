Nephio Concepts - Porch: Understand that Porch is a Kubernetes API extension. It makes "Git Repos" look like Kubernetes resources.

The kpt Book - Functions: This is the "logic" of Nephio. You need to understand how a containerized function can read a YAML, modify it, and spit it back out.

You can't easily programmatically "edit" a Helm chart after it's rendered without making the template incredibly complex.

Nephio Solution(Configuration As Data):
Nephio tread YAML like a databse. Instead of templates we use Functions. A function reads a YAML file and modifid, modifies a field and spits YAML back out...This is called KRM(Kubernetes Resource Model Pipeline)

KPT(Kubernetes Package Transformation):
makes configuration data the source of truth, stored separately from the live state
uses a uniform, serializable data model to represent configuration
separates code that acts on the configuration from the data and from packages / bundles of the data
abstracts configuration file structure and storage from operations that act upon the configuration data; clients manipulating configuration data don’t need to directly interact with storage (git, container images)

Components of KPT Toolchain:
1.KPT CLI
2.Function SDK -> Can be used to create functions to transform or validate YAML KRM oinput/output format
3.Function Catalog -> A catalog of off the shelf test functions, kpt makes configuration easy to create and transform via resuable functions. The functions need to be idempotent

KPT manages KRM resouces in bundles called package
A packages is explictly decalred using a file name KptFile containing a KRM resource of kind Kptfile

A workflow in kpt can be best modelled as performing some verbs on the noun package. For example, when consuming an upstream package, the initial workflow can look like this:

Get: Using kpt pkg get
Explore: Using an editor or running commands such as kpt pkg tree
Edit: Customize the package either manually or automatically using kpt fn eval. This may involve editing the functions pipeline in the Kptfile which is executed in the next stage.
Render: Using kpt fn render
Update: Using kpt pkg update
Create: Initialize a directory using kpt pkg init.
Initialize: One-time process using kpt live init
Preview: Using kpt live apply --dry-run
Apply: Using kpt live apply
Observe: Using kpt live status

KRM function is a containerized program that can perform CRUD operations on KRM resources stored in the loacl file system.kpt functions are the extensible mechanism to automate mutation and validation of KRM resources.

kpt fn eval: Executes a given function on the package. The image to run and the functionConfig is specified as CLI argument. This is an imperative way to run functions.

pipeline:
  mutators:
    - image: ghcr.io/kptdev/krm-functions-catalog/set-labels:latest
      configMap:
        app: wordpress
  validators:
    - image: ghcr.io/kptdev/krm-functions-catalog/kubeconform:latest

There are two differences between mutators and validators:
    Validators are not allowed to modify resources.
    Validators are always executed after mutators.

When you invoke render function, kpt performs the following steps:
1.Sequentially executes the list of muutators declared in mysql package.
2.Similarly execute validators

Specifying functionConfig:
functionConfig is an optional meta resource containing the arguments to a particular invocation of the function. There are two different ways to declare the functionConfig.

configPath:
apiVersion: kpt.dev/v1
kind: Kptfile
metadata:
    name: mysql
pipeline:
    mutators:
        - image: set-labels:latest
        configPath: labels.yaml
configgMap:

# wordpress/mysql/Kptfile
apiVersion: kpt.dev/v1
kind: Kptfile
metadata:
  name: mysql
pipeline:
  mutators:
    - image: set-labels:latest
      configMap:
        tier: mysql

Specifying function name
Specifying selectors
Specifying exclude

By default, functions cannot access the network. You can enable network access using the --network flag.
By default, functions cannot access the host file system. You can use the --mount flag to mount host volumes. kpt accepts the same options to --mount specified on the Docker Volumes page.


“Porch” is short for “Package Orchestration”.
Porch provides you with:

Kubernetes-native package management
Manage packages through Kubernetes resources (PackageRevision, Repository) instead of direct Git operations. Use kubectl, client-go, or any Kubernetes tooling, or Porch’s porchctl CLI utility.

Approval workflows
Packages move through lifecycle stages (Draft → Proposed → Published) with explicit approval gates. Prevent accidental publication of unreviewed changes.

Automatic package discovery
Register a Git or OCI repository once, and Porch automatically discovers all packages within it. No manual inventory management.

Function execution
Apply KRM functions to transform and validate packages. Functions run in isolated containers with results tracked in package history.

Package cloning and upgrades
Clone packages from upstream sources and automatically upgrade them when new upstream versions are published. Three-way merge handles local customizations.

GitOps integration
All changes are committed to Git with full history. Works seamlessly with Flux, Config Sync and other GitOps deployment tools.

Multi-repository orchestration
Manage packages across multiple Git and OCI repositories from a single control plane. Controllers can automate cross-repository operations.

Collaboration and governance
Multiple users and automation can work on packages concurrently. Draft revisions provide workspace isolation before publication.

Repository synchronization
Porch detects changes made directly in Git (outside Porch) and synchronizes its cache. Supports both Porch-managed and externally-managed packages.

Standard kpt packages
Packages remain standard kpt packages. No vendor lock-in or Porch specific DSL “code” in kpt packages.

Porch Components:
Porch Server:
The kubernetes aggregated api server that provides the Package Revision and Repository APIs. It contains the engine(orchestration logic), cache(repository content) and repository adapters(GIT/OCI abstraction)

Function Runner: Seperate GRPC service that executes the KRM functions in containers. IT can run both functions supplied by Porch and externally developed function images

Controllers: Kubernetes controllers automate package operations. The PackageVariant controller clones and updates packages. The PackageVariantSet controller manages sets of package variants.

Cache: A storage backend for caching of repository content for performance reasons. Porch supports a CR-based cache (using a Kubernetes custom resources) or a PostgreSQL-based cache for larger deployments.


Porch Core Concepts:
A porch package encapsulates te orchestration of a single kpt package, stored in a git repository.
Package Revision is a single version of a package.

Lifecycle:
Stage1: Draft : the package revision is being authored (created or edited). The package revision contents can be modified, but the package revision is not ready to be used/deployed.
Stage2: Proposed : the package revision’s author has completed the preparation of the package revision and its files and has proposed that the package revision be published.
Stage3: Published :  the package revision has been approved and is ready to be used. Published package revisions may be deployed. A published package revision may be copied to create a new package revision of the same package, in which development of the package may continue. A published package revision may also be cloned to create the first package revision of an entirely new package.
Stage4: Deletion Proposed: a user has proposed that this package revision be deleted from the repository. A package revision must be proposed for deletion before it can be deleted from Porch.

A workspace is the unique identifier of a package revision within a package.
A Revision number on a package revision identifies the order in which package revisions of a package were published. 

Placeholder package revision: A dummy package revision reference that points at a package’s latest package revision. The placeholder package revision is created by Porch simultaneously with the first package revision for a particular package. Each time a new package revision is published on the package, the placeholder package revision is updated (actually deleted and recreated).

The following rules apply:

there is always at most one placeholder package revision for a package
it always has a revision number of -1
its workspace name is always the branch in the Git repository on which the package revision exists - usually (though not always) main
its naming comvention is {repository-name}.{package-name}.{branch-name}, where {branch-name} is the branch in Git on which the package revision exists

Upstream are downstream are like dependencies

Think of kpt and Porch like the relationship between a Git CLI and GitHub.

kpt (The "Engine"): It is a client-side CLI tool. It treats Kubernetes YAML as "data." Its magic trick is the Pipeline: it can take a package, run a containerized function (like set-annotations) to modify the YAML, and output the result. It is manual and local.

Porch (The "Orchestrator"): It is "kpt-as-a-service." Porch is a Kubernetes API extension that lives inside your cluster. It manages your Git repositories and package lifecycles (Draft → Proposed → Published) through Kubernetes Custom Resources (PackageRevision).

The Connection: When you tell Porch to "modify a package," Porch actually triggers kpt functions in the background. Nephio uses Porch to automate thousands of these "kpt" edits across many clusters simultaneously.

Hands On:
sudo apt update && sudo apt install -y docker.io
sudo usermod -aG docker $USER && newgrp docker

# Install Kind (Kubernetes in Docker)
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind

# Create a cluster
kind create cluster --name nephio-lab

Install Porchctl
curl -L https://github.com/nephio-project/porch/releases/download/v1.5.6/porchctl_linux_amd64.tar.gz | tar -xz
sudo mv porchctl /usr/local/bin/

In O-RAN, the O2 interface needs to provision resources.

The Request: Someone asks for a "Small Cell" deployment.

The Porch Action: Porch pulls a "Small Cell Package" (standardized YAML).

The kpt Function: A function automatically injects the specific VLAN IDs or IP addresses required for that specific Samsung hardware.

The Result: The O2 interface sees this finalized "Package" and spins up the hardware.

mkdir ~/my-packages
cd ~/my-packages
git init -b main
git config --global user.email "you@example.com"
git config --global user.name "Your Name"
touch README.md
git add . && git commit -m "initial commit"

# Tell Porch to use this directory (Mapping it as a 'Repository' resource)
porchctl repo register --name training-repo --directory ~/my-packages

kpt live init && kpt live apply => Similar to helm upgrade/install