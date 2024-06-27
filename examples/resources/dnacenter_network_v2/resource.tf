terraform {
  required_providers {
    dnacenter = {
      version = "1.1.33-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_network_v2" "example" {
  provider = dnacenter
  parameters {

    settings {
      dns_server {
        domain_name          = "devhf.local"
        primary_ip_address   = "abc.def.ghi.jkl"
        secondary_ip_address = "abc.def.ghi.jkl"
      }
      syslog_server {
        ip_addresses = [
          "abc.def.ghi.jkl"
        ]
        configure_dnac_ip = true
      }
      snmp_server {
        ip_addresses = [
          "abc.def.ghi.jkl"
        ]
        configure_dnac_ip = true
      }
      ntp_server = [
        "abc.def.ghi.jkl",
        "abc.def.ghi.jkl",
        "abc.def.ghi.jkl",
        "abc.def.ghi.jkl"
      ]
      timezone = "Europe/Oslo"
      message_of_theday {
        banner_message         = "\n\nAccess to this system is restricted to authorized personell only.\nAll access attempts are logged.\n\n"
        retain_existing_banner = "false"
      }
      network_aaa {
        servers       = "ISE"
        ip_address    = "abc.def.ghi.jkl"
        network       = "abc.def.ghi.jkl"
        protocol      = "TACACS"
        shared_secret = "REDACTED"
      }
      client_and_endpoint_aaa {
        servers       = "ISE"
        ip_address    = "abc.def.ghi.jkl"
        network       = "abc.def.ghi.jkl"
        protocol      = "RADIUS"
        shared_secret = "REDACTED"
      }
    }
    site_id = "string"
  }
}

output "dnacenter_network_v2_example" {
  value = dnacenter_network_v2.example
}