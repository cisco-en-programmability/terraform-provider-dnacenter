

terraform {
  required_providers {
    dnacenter = {
      version = "1.1.33-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

data "dnacenter_device_details" "example" {
  provider   = dnacenter
  identifier = "uuid"
  search_by  = "57d9e4e4-e655-4512-a137-8f8c90e59ab1"
#   timestamp  = "string"
}

output "dnacenter_device_details_example" {
  value = data.dnacenter_device_details.example.item
}