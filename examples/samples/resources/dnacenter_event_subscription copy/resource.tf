terraform {
  required_providers {
    dnacenter = {
      version = "1.0.14-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_event_subscription" "example" {
  provider = dnacenter
  parameters {
    payload {
      description = "Test REST subscription 4"
      filter {

        //categories = ["string"]
        /*domains_subdomains {

        domain      = "string"
        sub_domains = ["string"]
      }*/
        event_ids = ["LICMGMT-DEV-VA-CHANGE-FAILURE", "LICMGMT-DEV-VA-CHANGE-SUCCESS", "LICMGMT-DLC-FAILURE"]
        /*severities = ["string"]
      site_ids   = ["string"]
      sources    = ["string"]
      types      = ["string"]*/
      }
      name = "Test REST subscription 4"
      subscription_endpoints {

        instance_id = "b4b841cf-cffe-4837-b88f-cea33a3a19ba"
        subscription_details {

          connector_type = "REST"
        }
      }
      //subscription_id = "string"
      //version         = "string"
    }
  }
}

output "dnacenter_event_subscription_example" {
  value = dnacenter_event_subscription.example
}