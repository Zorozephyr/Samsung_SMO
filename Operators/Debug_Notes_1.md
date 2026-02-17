docker logs k3d-ec2Operator-server-0 | tail -n 20

How to check if docker bridge network is overlapping with local WIFI ip range?
1.ip addr show | grep "inet "
Look for your primary interface (usually wlan0 for Wi-Fi or eth0). You'll see something like 192.168.1.15/24
2.docker network inspect k3d-ec2Operator | grep Subnet
If the first three numbers of your Local IP and the Docker Subnet match (e.g., both are 192.168.1.x), you have an overlap.


k3d cluster create ec2Operator --k3s-arg "--tls-san=127.0.0.1@server:0"
When k3d creates a cluster, it runs inside a Docker container.

The Internal IP: Inside Docker, the API server might think its name is k3d-ec2Operator-server-0.

The External IP: Your Ubuntu host sees the cluster at 127.0.0.1 (localhost).

By adding --tls-san=127.0.0.1, you are telling the K3s certificate generator: "Even though you are inside a container, please trust requests coming from 127.0.0.1."