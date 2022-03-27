# Preface
Initial task description could be found in *pleo-antaeus* folder.
Source files for the new payment service (**Payment Provider service**) are located at *pleo-payment-provider* folder.

## How to run

### Locally (Docker Compose)
If you want to run both services locally, just run:
```
docker-compose up
```
### Kubernetes (Minikube or any other K8s)
If you want to run both services at K8s, deploy the resources first:
```
kubectl apply -f ./pleo-antaeus/deploy/ -f ./pleo-payment-provider/deploy/
```
Make shure it works by performing smock tests:
```
curl 192.168.39.72/pp/rest/health
> ok
curl 192.168.39.72/rest/health
> "ok"
```
## How to test
Get Ingress address and replace all commands below with your IP (replace 192.168.39.72)
```
kubectl get ing -o jsonpath="{.items[0].status.loadBalancer.ingress[0].ip}"
```

Enshure there are unpaid invoices (status PENDING) by calling Antheus:
```
curl 192.168.39.72/rest/v1/invoices

> [{"id":1,"customerId":1,"amount":{"value":184.04,"currency":"USD"},"status":"PAID"},{"id":2,"customerId":1,"amount":{"value":238.25,"currency":"USD"},"status":"PENDING"} ...
```
Call Antheus service to pay the invoices:
```
curl 192.168.39.72/rest/v1/invoices/pay -X POST
> true
```
Antheus will call Payment Provider service (/pp/rest/v1/charge) which randomly succeed with particular invoice payment, so if you call Antheus pay endpoint (/rest/v1/invoices/pay) multiple times eventially you will find all infoices at status PAID (/rest/v1/invoices)

## How to improve
1. Q: How would a new deployment look like for these services? What kind of tools would you use?
   A: We may use Helm charts to pack related set of K8s resouces for each microservice. Any CI/CD tools to integrate changes and deploy package to environments (GitHub Actions, Jenkins, GitLab, ...)
2. Q: If a developers needs to push updates to just one of the services, how can we grant that permission without allowing the same developer to deploy any other services running in K8s?
   A: Developers should not have an access to deploy services to environments manualy, this should be done by CI/CD tools according with GitOps flow. Each servise's source files should be located at separate repository. Than we can restrict the access for particular team to particular repository. Each commit/PR to repository triggers pipeline and the service builded from this repository deployes to environments.
3. Q: How do we prevent other services running in the cluster to talk to your service. Only Antaeus should be able to do it.
   A: We can use Network Policies (K8s kind: NetworkPolicy) to strict the egress communications between pods inside the cluster. Or we can use some advanced tools like Calico or Istio for service mesh.
