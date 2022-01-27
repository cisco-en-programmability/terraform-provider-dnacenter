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
/*
data "dnacenter_tag" "example" {
  provider                   = dnacenter
  additional_info_attributes = "string"
  additional_info_name_space = "string"
  field                      = "1"
  level                      = "0"
  limit                      = "1"
  name                       = "WAN"
  offset                     = "10"
  order                      = "string"
  size                       = "1"
  sort_by                    = "string6"
  system_tag                 = "true"
}

output "dnacenter_tag_example" {
  value = data.dnacenter_tag.example.items
}
*/
data "dnacenter_tag" "example" {
  provider = dnacenter
  id       = "eacdbb1f-da5f-4acc-978f-bc723a5f29bc"
}

output "dnacenter_tag_example" {
  value = data.dnacenter_tag.example.item
}
