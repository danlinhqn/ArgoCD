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
yy8vR4WP6JohhC9s

### Public ArgoCD with NodePord

```
kubectl edit svc argocd-server -n argocd
```

----

## Cho chạy ra ngoài

k port-forward service/argocd-server -n argocd-server -n argocd 8080:443