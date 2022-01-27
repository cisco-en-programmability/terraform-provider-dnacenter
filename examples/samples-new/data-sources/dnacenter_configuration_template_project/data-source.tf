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


data "dnacenter_configuration_template_project" "example" {
  provider   = dnacenter
  name       = "Onboarding Configuration"
  sort_order = "asc"
}

output "dnacenter_configuration_template_project_example" {
  value = data.dnacenter_configuration_template_project.example.items
}
/*
data "dnacenter_configuration_template_project" "example" {
  provider   = dnacenter
  project_id = "2128b364-8751-45f5-95cb-f5ecaa2e9085"
}

output "dnacenter_configuration_template_project_example" {
  value = data.dnacenter_configuration_template_project.example.item
}
*/