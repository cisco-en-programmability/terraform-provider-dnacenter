
terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source   = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

resource "dna_pnp_global_settings" "response" {
  provider = dnacenter
  item {
    version = 1
    aaa_credentials {
      username = ""
      password = ""
    }
    task_time_outs {
      config_time_out         = 10
      image_download_time_out = 120
      general_time_out        = 20
    }
    sava_mapping_list {
    }
    accept_eula = true
    default_profile {
      ip_addresses = [
        "",
        "192.168.196.2",
        "10.121.1.5"
      ]
      fqdn_addresses = []
      port           = 443
      cert           = "-----BEGIN CERTIFICATE-----\n...\n-----END CERTIFICATE-----"
      proxy          = false
    }
  }
}

output "dna_pnp_global_settings_response" {
  value = dna_pnp_global_settings.response
}

