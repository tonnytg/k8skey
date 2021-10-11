# K8SKey
<hr />

Example using with interactive menu: 

    ➜  k8skey git:(master) ✗ go run main.go --select true
    Project[0]: localhost1
    Project[1]: ultra-sound-324019
    Choose Project: 1
    Project: 1
    ---
    Cluster[0]: autopilot-cluster-1
    Choose Cluster: 0
    Cluster: 0
    ---
    Project[1]: ultra-sound-324019
    Cluster[0]: autopilot-cluster-1
    Fetching cluster endpoint and auth data.
    kubeconfig entry generated for autopilot-cluster-1.

Update Database

    --database update

Select Menu

    --select true

Set Project, Cluster and Region

    --project PROJECT_ID --cluster CLUSTER_ID --region REGION_ID