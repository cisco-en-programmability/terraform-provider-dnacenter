

terraform {
  required_providers {
    dnacenter = {
      version = "1.3.0-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_application_policy_application_set" "example" {
  provider = dnacenter
  parameters {
    payload {
      default_business_relevance = "default"
      # id                             = "string"
      name = "Terraform Test"
      # namespace                      = "string"
      # qualifier                      = "string"
      # scalable_group_external_handle = "string"
      # scalable_group_type            = "string"
      # type                           ="string"
    }
  }
}

output "dnacenter_application_policy_application_set_example" {
  value = dnacenter_application_policy_application_set.example
}
