# Terraform for Azure

Udemy course: [Link]()

## Notes

Useful terraform CLI commands:

* **terraform fmt <*/file> **: Formats all .tf files in the directory or a sigle file when given as parameter.

* **terraform taing <instance> **: Destroyes and recreates the instance in the next terrafrom apply.

* **terraform import <resource_type.resource_name unique-identifier> **: Imports already existing resources.

* **terraform workspace <new/list/show/select/delete> <name> **: Manage workspaces

* **terraform state <move(vm)/remove(rm)/list/pull/push/replace(provider)/show> **: Allows the manipulation of the terraform state. 

* **terrafrom refresh**: Updates the state file to reflect the "actual" state of the infrastructure.

Enable more logging by setting TF_LOG env variable (ex. TF_LOG=DEBUG teraform apply).
The valid log leves are the following:

* TRACE
* DEBUG
* INFO
* WARN
* ERROR 

Secrets can be stored in the state file (ex. initial database password)

To get a list of the virtual machines use `az vm image list -p <publisher>`.
For a linux vm you can use `az vm image list -p "Canonical"`