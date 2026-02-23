Config Sync:

```markdown
- **RootSync**: A cluster-scoped CRD used by cluster administrators to sync cluster-wide configurations (e.g., RBAC, Namespaces, Quotas) from a central "root" repository. It runs with high privileges (`cluster-admin`) to manage resources across the entire cluster.
- **RepoSync**: A namespace-scoped CRD that allows namespace owners to manage resources within their specific namespace. It syncs from a "namespace" repository and operates with restricted permissions, typically limited to the namespace where it is defined.
- **Mechanism**: Both resources use a reconciler pod that monitors a Git, OCI, or Helm source. When changes are detected, the reconciler applies the manifests to the cluster, ensuring the live state matches the desired state defined in the source of truth.
```

