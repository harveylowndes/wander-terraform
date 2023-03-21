package test

import (
	///"os"
	//"path/filepath"
	//"log"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"

	//"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/tools/clientcmd"
)

func TestEndToEnd(t *testing.T) {
	defer terraform.Destroy(t, Setup(t))

	// IAM
	terraformOptionsIAM := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/project_infrastructure/iam",
		VarFiles: []string{"../../../config/default.tfvars", "../../../config/test.tfvars"},
	})
	
	defer terraform.Destroy(t, terraformOptionsIAM)

	terraform.WorkspaceSelectOrNew(t, terraformOptionsIAM, "test")
	terraform.InitAndApply(t, terraformOptionsIAM)

	// VPC
	terraformOptionsVPC := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/project_infrastructure/vpc",
		VarFiles: []string{"../../../config/default.tfvars", "../../../config/test.tfvars"},
	})
	
	defer terraform.Destroy(t, terraformOptionsVPC)

	terraform.WorkspaceSelectOrNew(t, terraformOptionsVPC, "test")
	terraform.InitAndApply(t, terraformOptionsVPC)

	// Network
	terraformOptionsNetwork := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/region_infrastructure/network",
		VarFiles: []string{"../../../config/default.tfvars", "../../../config/test.tfvars"},
	})
	
	defer terraform.Destroy(t, terraformOptionsNetwork)

	terraform.WorkspaceSelectOrNew(t, terraformOptionsNetwork, "us-west1")
	terraform.InitAndApply(t, terraformOptionsNetwork)

	// Cluster
	terraformOptionsCluster := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/region_infrastructure/cluster",
		VarFiles: []string{"../../../config/default.tfvars", "../../../config/test.tfvars"},
	})
	
	defer terraform.Destroy(t, terraformOptionsCluster)

	terraform.WorkspaceSelectOrNew(t, terraformOptionsCluster, "us-west1")
	terraform.InitAndApply(t, terraformOptionsCluster)

	// Deploy app and curl endpoint
	terraformOptionsApp := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/application_deployment",
		VarFiles: []string{"../../../config/default.tfvars", "../../../config/test.tfvars"},
	})
	
	defer terraform.Destroy(t, terraformOptionsApp)

	terraform.WorkspaceSelectOrNew(t, terraformOptionsApp, "us-west1")
	terraform.InitAndApply(t, terraformOptionsApp)

	// TODO (hlowndes) finish tests. Just need to query kube api 
	// for the external IP of the loadbalancer created by the ingress
	// then curl expecting the "Hello world!" string in response

	//kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config",)
   	//config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
   	//if err != nil {
	//   log.Fatal(err)
   	//}
	//clientset, err := kubernetes.NewForConfig(config)

	//_ := clientset.CoreV1()
}