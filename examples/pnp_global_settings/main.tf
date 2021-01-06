
terraform {
  required_providers {
    dnacenter = {
      versions = ["0.2"]
      source   = "hashicorp.com/edu/dnacenter"
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

