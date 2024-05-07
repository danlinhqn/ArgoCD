## ArgoCD

Tool same is Tekton

### Setup in K8S

```
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

### Get Password

```
kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | base64 -d
```

### Public ArgoCD with NodePord

```
kubectl edit svc argocd-server -n argocd
```
