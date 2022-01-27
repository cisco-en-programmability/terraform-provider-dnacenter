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

resource "dnacenter_event_subscription" "example" {
    provider = dnacenter
    parameters {
      
      description = "Test REST subscription 4"
      filter {
        
        # event_ids = ["LICMGMT-DEV-VA-CHANGE-FAILURE"]
        event_ids = ["LICMGMT-DEV-VA-CHANGE-FAILURE", "LICMGMT-DEV-VA-CHANGE-SUCCESS", "LICMGMT-DLC-FAILURE"]
      }
      name = "Test REST subscription 4"
      subscription_endpoints {
        
        instance_id = "b4b841cf-cffe-4837-b88f-cea33a3a19ba"
        subscription_details {
          
          connector_type = "REST"
          method = "POST"
          name = "TestWebhook"
          url = "https://enprxnc1g2b3p.x.pipedream.net"
        }
      }
      # subscription_id = "string"
      # version = "string"
    }
}

output "dnacenter_event_subscription_example" {
    value = dnacenter_event_subscription.example
}