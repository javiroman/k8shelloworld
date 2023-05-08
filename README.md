# k8shelloworld

## Podman

```
podman build -f Dockerfile.alpine . -t javiroman/helloworld:v3
podman run -d -p 8080:8080 --name hello --env FILE="/etc/motd" javiroman/helloworld:v3
curl localhost:8080
podman stop hello
podman rm hello
podman run -ti --entrypoint="/bin/bash" -p 8080:8080 --name hello javiroman/helloworld:v3
podman login -u javiroman
podman push javiroman/zookeeper:v3
```

## Kubernetes

```
apiVersion: v1
kind: Pod
metadata:
  name: helloworld
  labels:
    app: helloworld-app
spec:
  containers:
  - name: helloworld
    image: javiroman/helloworld:v3
    env:
    - name: FILE
      value: "/etc/hosts"
---
apiVersion: v1
kind: Service
metadata:
  name: helloworld-service
spec:
  type: NodePort
  ports:
    - targetPort: 8080
      port: 8080
      nodePort: 30000
  selector:
    app: helloworld-app
```

```
k apply -f helloworld.yaml
k get pod -o wide
curl k8s-worker03.kubernetes.lan:30000
```

