
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.24-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_deploy_template" "example" {
  provider = dnacenter

  parameters {
    force_push_template = "false"
    //is_composite                    = "true"
    //member_template_deployment_info = []
    target_info {
      //host_name             = "C9K-Branch-SFO"
      id = "3923aed0-16e5-4ed0-b430-ff6dcfd9c517"
      # params = {
      #   int  = "g1/01/12"
      #   desc = "auto"
      # }
      //resource_params       = []
      type = "MANAGED_DEVICE_IP"
      //versioned_template_id = "c49c6179-0631-4b2c-8f94-3216929af040"
    }
    template_id = "fcfd4d19-99e2-494e-9c6f-0d85cf3094e5"
  }
}

output "dnacenter_deploy_template_example" {
  value = dnacenter_deploy_template.example
}

# data "dnacenter_task" "example" {
#   depends_on = [dnacenter_deploy_template.example]
#   provider   = dnacenter
#   task_id    = dnacenter_deploy_template.example.item.0.task_id
# }

# output "dnacenter_task_example" {
#   value = data.dnacenter_task.example.item
# }
