# Wander Terraform Project

This is just a project to showcase a hello world demo on GCP from building the infrastructure to deploying the application and testing. This is very bare bones and not production ready. There are many TODO's here including:
* Better test coverage
* Helm deployment through terraform
* CI/CD Configurations
* Management of external IPs and load balancers
* Dynamic cert rotation for secure TCP connections
* Locked down IAM policy permissions
* Improved tooling
* Further locking down of firewall rules
* Developer account break-glass escalation
* Helm pre/post hooks


## Prerequisites

Local provisioning and test execution:
* Terraform 13+
* Go

Development:
* Terraform 13+
* Go
* Kubectl
* Helm
* Docker

## Project Structure

* app - the demo app code, dockerfile and helm chart
* config - terraform variable files for each environment type
* terraform - All terraform "modules"
* test - Terraform go tests (using terra test)

### Terraform directory

In order there are 3 sub directories:
* project_infrastructure - Project level infrastructure
* region_infrastructure - Region level infrastructure
* application_deployment - Deploying of the application

Within each directory there is a folder named based on the group of resources it manages. There are many ways this could have been done but i have chosen to have a state file per resource. From experience, having large state files can be very difficult to maintain. They can grow to a size where refreshes start to be slow and it also stops asynchronous workflows where they might be possible. They project has not yet used modules, although this could have also been used and if the project was bigger i would assume they would be used.

## Provisioning Infrastructure

Since each resource is essentially its own project, its directory can be entered and treated as such. To apply in these directories takes some understanding into how this works. Firstly there is slight differences between the project and region terraform. Note: Backend in project infrastructure sets up the backend for tfstate files. This manages its own tfstate locally.

### Project
This project utilizes terraform workspaces as well as the terraform variables. In project infrastructure you will need to set the workspace to the environment you are going to be working in:
```bash
terraform workspace select --or-create <development/staging/production>
```
Then it is a matter of performing a standard terraform apply
```bash
terraform apply --var-file $project_root/config/default.tfvars --var-file $project_root/config/<development/staging/production>.tfvars
```

### Region
For region infrastructure it is the same as project, except the workspace selected must be the region to deploy
```bash
terraform workspace select --or-create <e.g. us-west1>
```

### Helper tool
From the root project there is a provision script which will enter the correct module, pick up the correct tfvars file from the config directory and perform a terraform apply
```bash
./provision -t <type> -m <module> -r <region> -e <environment>
```

## Application

### Building
Just a very basic go project
```bash
cd app
go build -o bin/server go/*.go
docker build -t gcr.io/<project id>/hello-app:latest
docker push gcr.io/<project id>/hello-app:latest
helm package chart/
helm push hello-app-<version>.tgz
```

### Deploying
Should end up being the same as the region infrastructure terraform but in the application_deployment directory

Ensure your gcloud config is pointing to the right project
```bash
gcloud config set project <project id>
cd terraform/application_deployment
terraform apply $project_root/config/default.tfvars --var-file $project_root/config/<development/staging/production>.tfvars
```

This will deploy to the cluster.
### Tests

Run
```bash
go test <test file>
```
to run the test of choice

These are very unfinished due to time constraints. The e2e test is very much almost there, however its missing the actual deploy part. This would be easier if the external ip and loadbalancer was managed through the terraform.