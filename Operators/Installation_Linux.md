# Installation Steps (Linux/Ubuntu)

## 1. Docker
Pre-requisites and installation for Docker Engine.

```bash
# Update package index and install certificates
sudo apt-get update
sudo apt-get install -y ca-certificates curl gnupg

# Add Docker's official GPG key
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

# Set up the repository
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Update and Install Docker
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Post-installation (Run without sudo)
sudo usermod -aG docker $USER
newgrp docker
```

## 2. k3d (Lightweight Kubernetes in Docker)
Used for creating local Kubernetes clusters.

```bash
# Install k3d using the official install script
curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash

# Verify installation
k3d --version
```

## 3. k9s (Kubernetes CLI UI)
Terminal-based UI to interact with your Kubernetes clusters.

### Option A: Via Snap (Ubuntu/Debian)
```bash
sudo snap install k9s
```

### Option B: Binary Download (Universal)
```bash
# Download the latest release (check github for newer versions)
curl -L -s https://github.com/derailed/k9s/releases/download/v0.32.4/k9s_Linux_amd64.tar.gz -o k9s.tar.gz

# Extract and move to /usr/local/bin
tar -xvzf k9s.tar.gz
sudo mv k9s /usr/local/bin/

# Cleanup
rm k9s.tar.gz
```

## 4. Go (Golang)
Required for developing Operators.

```bash
# Remove any existing Go installation
sudo rm -rf /usr/local/go

# Download Go (adjust version as needed, e.g., 1.22.0)
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz

# Extract to /usr/local
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz

# Add Go to PATH (Persistent)
# Add these lines to ~/.bashrc or ~/.zshrc
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Verify installation
go version
```