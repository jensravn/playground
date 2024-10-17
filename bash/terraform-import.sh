# Variables
export PROJECT="my-project"
export NEW_RESOURCE="resource_type.resource_name"
export OLD_RESOURCE="resource_type.resource_name"
export BUCKET="my-terraform-state-bucket"
export PREFIX="terraform-state"
export RESOURCE_PATH="projects/$PROJECT/..."

# Init with bucket and prefix (folder in bucket)
# This will create the files and folders .terraform + .terraform.lock.hcl
terraform init -backend-config="bucket=$BUCKET" -backend-config="prefix=$PREFIX"

# Import into resource defined in terraform
terraform import -var-file=../env/production.tfvars "$NEW_RESOURCE" "$RESOURCE_PATH"

# Check resource is now in the terraform state
terraform state list $NEW_RESOURCE

# Check old state before removing it
terraform state list $OLD_RESOURCE

# Remove old object from state
terraform state rm $OLD_RESOURCE

# Check state was removed
terraform state list $OLD_RESOURCE
