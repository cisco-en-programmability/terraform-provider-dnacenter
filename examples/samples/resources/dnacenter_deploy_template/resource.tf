
terraform {
  required_providers {
    dnacenter = {
      version = "0.2.0-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_deploy_template" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    force_push_template             = "true"
    is_composite                    = "true"
    member_template_deployment_info = []
    target_info {
      host_name             = "C9K-Branch-SFO"
      id                    = "3923aed0-16e5-4ed0-b430-ff6dcfd9c517"
      params                = []
      resource_params       = []
      type                  = "DEFAULT"
      versioned_template_id = "c49c6179-0631-4b2c-8f94-3216929af040"
    }
    template_id = "fe2bd8b9-2cf0-4b73-b7dc-755ff0f26363"
  }
}

output "dnacenter_deploy_template_example" {
  value = dnacenter_deploy_template.example
}
