
terraform {
  required_providers {
    dnacenter = {
      version = "1.0.6-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_service_provider" "example" {
  provider = dnacenter
  parameters {

    settings {
      qos {
        profile_name = "Test_tf_new"
        model        = "8-class-model"
        wan_provider = "test1-provider"
      }
    }
  }
}

output "dnacenter_service_provider_example" {
  value = dnacenter_service_provider.example
}
