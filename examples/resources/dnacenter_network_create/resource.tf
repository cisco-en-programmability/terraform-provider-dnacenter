
resource "dnacenter_network_create" "example" {
  provider = dnacenter
  parameters {

    settings {

      client_and_endpoint_aaa {

        ip_address    = "string"
        network       = "string"
        protocol      = "string"
        servers       = "string"
        shared_secret = "string"
      }
      dhcp_server = ["string"]
      dns_server {

        domain_name          = "string"
        primary_ip_address   = "string"
        secondary_ip_address = "string"
      }
      message_of_theday {

        banner_message         = "string"
        retain_existing_banner = "string"
      }
      netflowcollector {

        ip_address = "string"
        port       = 9090
      }
      network_aaa {

        ip_address    = "string"
        network       = "string"
        protocol      = "string"
        servers       = "string"
        shared_secret = "string"
      }
      ntp_server = ["string"]
      snmp_server {

        configure_dnac_ip = "false"
        ip_addresses      = ["string"]
      }
      syslog_server {

        configure_dnac_ip = "false"
        ip_addresses      = ["string"]
      }
      timezone = "string"
    }
    site_id = "string"
  }
}

output "dnacenter_network_create_example" {
  value = dnacenter_network_create.example
}