```
cfssl genkey -initca csr.json | cfssljson -bare ca
cfssl gencert -ca ca.pem -ca-key ca-key.pem csr.json | cfssljson -bare cert

minikube start --extra-config apiserver.Admission.PluginNames=NamespaceLifecycle,LimitRanger,ServiceAccount,PersistentVolumeLabel,DefaultStorageClass,DefaultTolerationSeconds,MutatingAdmissionWebhook,ValidatingAdmissionWebhook,ResourceQuota

kubectl create secret generic webhook-cert --from-file=./cert.pem --from-file=./cert-key.pem
kubectl describe secret webhook-cert

kubectl apply -f manifest.yml
kubectl apply -f webhook_configration.yml
kubectl apply -f pod.yml

kubectl logs po/webhook-...
```
