terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacdnacenter_swim_trigger_distribution" "example" {
  provider = dnacenter
  payload {

    device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a4"
    image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
  }
}