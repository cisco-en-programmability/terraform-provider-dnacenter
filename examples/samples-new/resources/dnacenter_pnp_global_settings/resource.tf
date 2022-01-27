

terraform {
  required_providers {
  dnacenter = {
    version = "0.0.3"
    source  = "hashicorp.com/edu/dnacenter"
    # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
  }
  }
}
resource "dnacenter_pnp_global_settings" "example" {
  provider = dnacenter
  parameters {
    aaa_credentials {

      password = "a"
      username = "string"
    }
    /*
    id = "string"
    
    accept_eula = "false"
    default_profile {

      cert           = "string"
      fqdn_addresses = ["string"]
      ip_addresses   = ["string"]
      port           = 1
      proxy          = "false"
    }
    sava_mapping_list {

      auto_sync_period = 1
      cco_user         = "string"
      expiry           = 1
      last_sync        = 1
      profile {

        address_fqdn  = "string"
        address_ip_v4 = "string"
        cert          = "string"
        make_default  = "false"
        name          = "string"
        port          = 1
        profile_id    = "string"
        proxy         = "false"
      }
      smart_account_id = "string"
      sync_result {

        sync_list {

          device_sn_list = ["string"]
          sync_type      = "string"
        }
        sync_msg = "string"
      }
      sync_result_str    = "string"
      sync_start_time    = 1
      sync_status        = "string"
      tenant_id          = "string"
      token              = "string"
      virtual_account_id = "string"
    }
    task_time_outs {

      config_time_out         = 1
      general_time_out        = 1
      image_download_time_out = 1
    }*/
    tenant_id = "1"
    version   = 1
  }
}

output "dnacenter_pnp_global_settings_example" {
  value = dnacenter_pnp_global_settings.example
  sensitive = true
}