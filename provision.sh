#! /bin/bash
set +x

while getopts t:m:r:e: flag
do
    case "${flag}" in
        t) type=${OPTARG};;
        m) module=${OPTARG};;
        r) region=${OPTARG};;
        e) environment=${OPTARG};;
    esac
done

workspace=$region
if [[ "$type" == "project" ]]; then
    workspace=$environment
fi
pushd "./terraform/${type}_infrastructure/${module}"
    vfp="../../../config/$environment.tfvars"
    terraform init -reconfigure
    terraform workspace select --or-create $workspace
    terraform apply --var-file ../../../config/default.tfvars --var-file ${vfp}
popd

