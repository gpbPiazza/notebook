# Notes

- Kubectl possuí configs lá iremos setar a localização dos nossos cluters via contexts, logo:

- `kubectl config get-contexts` -> lista todos os contexts configurados
- `kubectl config get-clusters` -> lista todos os clusters configurados
- `-n <namespace>` -> flag para colocar namespace que você está lidando no cluster.
- `kubectl logs <pod-name> -n payment-api -f` -> comando para visualizar os logs de um pod, a flag `-f` mantem o terminal em real time com os logs, sem a flag ele apenas printa os logs do momento do comando
- `kubectl logs <pod-name> --all-containers=true` -> irá listar todos os logs de todos so containers de um pod.
- 
### Setup 

```zsh
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