
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.28-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_image_distribution" "example" {
  provider = dnacenter

  parameters {
    payload {
      device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a5"
      image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
    }
    payload {
      device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a4"
      image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
    }
    payload {
      device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a6"
      image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
    }
    payload {
      device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a7"
      image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
    }
    payload {
      device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a8"
      image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
    }
  }
}

output "dnacenter_image_distribution_example" {
  value = dnacenter_image_distribution.example
}

data "dnacenter_task" "example" {
  depends_on = [dnacenter_image_distribution.example]
  provider   = dnacenter
  task_id    = dnacenter_image_distribution.example.item.0.task_id
}

output "dnacenter_task_example" {
  value = data.dnacenter_task.example.item
}
