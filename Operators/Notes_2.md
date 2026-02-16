nerdctl ps
nerdctl build .
nerdctl pull nginx
nerdclt run nginx

server is master
agents is workers

Virtualization (The "Matrix")

Your Laptop (The Host):
You have, say, 1 physical CPU with 8 cores and 16GB of RAM.

The Slice (The VM/WSL):
Rancher Desktop creates a Virtual Machine (VM) (running Linux) on Windows. It carves out a slice of your laptop—say, 2 CPUs and 4GB RAM—and dedicates it to this Linux world.

The Illusion (Containers as Nodes):
Inside that Linux slice, the tool you used (k3d) pulls a magic trick. It starts 3 Docker Containers.

Container 1: Pretends to be the "Master Node".

Container 2: Pretends to be "Worker Node A".

Container 3: Pretends to be "Worker Node B".

To your laptop: It just sees 3 heavy programs running. It divides its processing power between them.
To Kubernetes: It sees 3 completely separate "computers" networked together.

. The "All-in-One" (Local / Dev)
What you have now (Rancher Desktop / Minikube).

Setup: The Master and the Worker are the same computer (or virtual machine).

Use Case: Testing, learning, coding on your laptop.

Risk: If your laptop dies, the whole cluster dies.

B. The "Standard" (Production)
Setup: distinct physical (or virtual) machines.

3 computers are only Masters (for safety).

5+ computers are only Workers.

Use Case: Running real companies (Netflix, Uber, etc.).

Why: If a Worker computer catches fire, the Masters notice and move the work to the other 4 Workers. The app never goes down.

C. The "Managed" (Cloud)
Examples: EKS (AWS), AKS (Azure), GKE (Google).

Setup: You don't even see the Master nodes. The cloud provider manages the "Brain" for you completely. You only pay for the "Worker" computers that run your specific apps.

k3d cluster list
k3d cluster start clusterName
k3d cluster delete clusterName
k3d cluster create --config cluster.yaml

kubectk cluster-info
docker ps
