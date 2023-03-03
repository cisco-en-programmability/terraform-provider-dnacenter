
terraform {
  required_providers {
    dnacenter = {
      version = "1.0.19-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_image_device_activation" "example" {
  provider = dnacenter

  parameters {
    //client_type=                    "string"
    //client_url=                     "string"
    //schedule_validate=              "false"
    payload {
      activate_lower_image_version = "true"
      device_upgrade_mode          = "string"
      device_uuid                  = "3eb928b8-2414-4121-ac35-1247e5d666a4"
      distribute_if_needed         = "true"
      image_uuid_list = [
        "6af2b040-a312-4f57-8c8e-21f5e3e07598"
      ]
      smu_image_uuid_list = [
        "6af2b040-a312-4f57-8c8e-21f5e3e07597"
      ]
    }
  }
}

output "dnacenter_image_device_activation_example" {
  value = dnacenter_image_device_activation.example
}

data "dnacenter_task" "example" {
  depends_on = [dnacenter_image_device_activation.example]
  provider   = dnacenter
  task_id    = dnacenter_image_device_activation.example.item.0.task_id
}

output "dnacenter_task_example" {
  value = data.dnacenter_task.example.item
}
