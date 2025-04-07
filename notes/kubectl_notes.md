# Notes

- Kubectl possuí configs lá iremos setar a localização dos nossos cluters via contexts, logo:

- Use `kubectl describe <source-name> -n <namespace>` -> to see the full infra yaml of the resource.
- `kubectl config get-contexts` -> lista todos os contexts configurados
- Use `kubectl config current-context` -> to know each context you are using now.
- Use `kubectl config get-clusters` -> lista todos os clusters configurados
- Use `-n <namespace>` -> flag para colocar namespace que você está lidando no cluster.
- Use `kubectl logs <pod-name> -n payment-api -f` -> comando para visualizar os logs de um pod, a flag `-f` mantem o terminal em real time com os logs, sem a flag ele apenas printa os logs do momento do comando
- Use `kubectl logs <pod-name> --all-containers=true` -> irá listar todos os logs de todos so containers de um pod.
- Use `kubectl rollout restart deployment <nome-do-deployment> -n <nome-do-namespace>` to restart all pods from a deployment
- Use `kubectl scale <statefulset || deployment || job> <resource_name> --replicas=1 -n <name_space>` -> to scale up or down resources.
- Use `kubectl exec -it <pod_name> -- /bin/sh` -> to connect in pod shell if the pod use `sh` if he uses bash change to `/bin/bash`.
- Use `kubectl get configmap || cm` -> to list all configmaps of the cluster, you also can use `kubectl get configmap <config-map-name> -n <namespace> -o yaml` -> to see all his values in yaml
### Basic operations

- Use `kubectl describe` to see full yaml of the resource
- Use `kubectl get` to list the name of th resource
-

### Setup 
-  Use `gcloud container clusters get-credentials production-gke-cluster --region=<region> --project <gcp_project> && \
  kubectl config rename-context "<cluster_name>" "old-gke-production"` ->   This command config a new ./kube/confg

```sh
############################################################
# Adiciona contexto do Kubernetes para o cluster solicitado.
# Arguments:
#   ID do projeto Gcloud
#   Região do cluster GKE
#   Nome do cluster GKE
############################################################
function setup_access_to_cluster() {
  local project_id="$1"
  local region="$2"
  local cluster_name="$3"
  local k8s_cluster_context="gke_${project_id}_${region}_${cluster_name}"

    i "Configurando acesso ao cluster $cluster_name na region $region e projeto $project_id..."
    gcloud container clusters get-credentials "$cluster_name" --region "$region" --project "$project_id" || {
      e "Falha ao obter credenciais do cluster $cluster_name na region $region e projeto $project_id."
    }

    kubectl config rename-context "$k8s_cluster_context" "$cluster_name" || e "[ERRO] Falha ao renomear contexto Kubernetes."
  else
    i "Cluster $cluster_name na region $region e projeto $project_id já está configurado."
  fi
}
```