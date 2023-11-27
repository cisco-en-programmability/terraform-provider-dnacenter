terraform {
  required_providers {
    dnacenter = {
      version = "1.1.29-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_network_create" "global" {
  provider = dnacenter

  parameters {

    settings {
      dhcp_server = ["1.1.1.1"]

      dns_server{
        domain_name = "hola"
        primary_ip_address = "1.1.1.1"
      }
      message_of_theday{
        banner_message="Have a good day"
      }
    }

    site_id = "771662ca-cb7e-47ff-9eaf-9b8c85e8e389"
  }
}
output "dnacenter_network_create_example" {
  value = dnacenter_network_create.global
}
