
terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

# data "dnacenter_configuration_template" "example_source" {
#   provider       = dnacenter
#   latest_version = "true"
#   template_id    = "2121f1d7-d9ea-4eea-8f83-45af2941f5f9"
# }

# output "dnacenter_configuration_template_example_source" {
#   value = data.dnacenter_configuration_template.example_source.item
# }

resource "dnacenter_configuration_template" "example" {
  lifecycle {
    ignore_changes = [parameters.0.author]
  }
  provider = dnacenter
  parameters {
    project_id       = "55c807d9-138c-4ddf-a9af-5cb3fd572104"
    template_content = "if a > b \n hola22"
    language         = "JINJA"
    name             = "Saludo Terraform"
    software_type    = "IOS-XE"
    author           = "altus"
    description      = "Created by Terraform automation"
    device_types {
      product_family = "Switches and Hubs"
      product_series = "Cisco Catalyst 9300 Series Switches"
    }
  }

}

output "dnacenter_configuration_template_example" {
  value = dnacenter_configuration_template.example
}
