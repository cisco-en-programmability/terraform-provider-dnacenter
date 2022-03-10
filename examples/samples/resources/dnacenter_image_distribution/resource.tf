
terraform {
  required_providers {
    dnacenter = {
      version = "0.2.0-beta"
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
  lifecycle {
    create_before_destroy = true
  }
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