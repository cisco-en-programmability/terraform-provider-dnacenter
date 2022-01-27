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

data "dna_template_project" "response" {
  provider = dnacenter
  name     = "Cloud DayN Templates"
}
output "dna_template_project_response" {
  value = data.dna_template_project.response.items.0.id
}


data "dna_template" "response" {
  provider   = dnacenter
  project_id = data.dna_template_project.response.items.0.id
  # product_series = "Cisco Cloud Services Router 1000V Series"
  # product_family = "Routers"
  # software_type  = "IOS"
}
output "dna_template_response" {
  value = data.dna_template.response
}


data "dna_template_details" "response" {
  provider    = dnacenter
  template_id = "82def961-26b7-4828-a7a2-579a4e10e446"
  # latest_version = true
}
output "dna_template_details_response" {
  value = data.dna_template_details.response
}


data "dna_template_version" "response" {
  provider    = dnacenter
  template_id = "82def961-26b7-4828-a7a2-579a4e10e446"
}
output "dna_template_version_response" {
  value = data.dna_template_version.response
}


data "dna_template_preview" "response" {
  provider    = dnacenter
  template_id = "82def961-26b7-4828-a7a2-579a4e10e446"
  params = {
    NetworkId = 1
  }
}

output "dna_template_preview_response" {
  value = data.dna_template_preview.response
}

data "dna_template_deploy" "response" {
  provider = dnacenter
  template_deployment_info {
    template_id = "5f5e3eca-2b43-4228-b7b6-3b957b56c110"
    target_info {
      hostname = "10.121.1.1"
      id       = "1a23a341-7ea2-41e9-8814-989a7d10c4be"
      params = {
        NetworkId = 1
      }
      type = "MANAGED_DEVICE_IP"
    }
  }
  member_templates_deployment_info {
    template_id = "5f5e3eca-2b43-4228-b7b6-3b957b56c110"
    target_info {
      hostname = "10.121.1.1"
      id       = "1a23a341-7ea2-41e9-8814-989a7d10c4be"
      params = {
        NetworkId = 1
      }
      type = "MANAGED_DEVICE_IP"
    }
  }
}
output "dna_template_deploy_response" {
  value = data.dna_template_deploy.response
}

data "dna_template_deploy_status" "response" {
  provider      = dnacenter
  deployment_id = data.dna_template_deploy.response.item.0.deployment_id
}
output "dna_template_deploy_status_response" {
  value = data.dna_template_deploy_status.response
}

