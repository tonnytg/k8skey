# K8SKey
<hr />


Export Token and run command.

    ➜  k8skey git:(master) ✗ export GCP_TOKEN=`gcloud auth application-default print-access-token`
    ➜  k8skey git:(master) ✗ go run main.go
    List Projects
    statusCode: 200
    Project: ultra-sound-324019
    Cluster[0]: autopilot-cluster-1 - us-central1
