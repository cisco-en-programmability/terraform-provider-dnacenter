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
data "dnacenter_configuration_template" "example" {
  provider                     = dnacenter
  filter_conflicting_templates = "false"
  product_family               = "string"
  product_series               = "string"
  product_type                 = "string"
  project_id                   = "string"
  project_names                = ["string"]
  software_type                = "string"
  software_version             = "string"
  sort_order                   = "asc"
  tags                         = ["string"]
  un_committed                 = "false"
}

output "dnacenter_configuration_template_example" {
  value = data.dnacenter_configuration_template.example.items
}
*/

data "dnacenter_configuration_template" "example" {
  provider       = dnacenter
  latest_version = "true"
  template_id    = "2121f1d7-d9ea-4eea-8f83-45af2941f5f9"
}

output "dnacenter_configuration_template_example" {
  value = data.dnacenter_configuration_template.example.item
}
