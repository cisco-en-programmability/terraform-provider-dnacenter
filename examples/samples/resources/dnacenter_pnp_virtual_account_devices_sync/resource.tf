

terraform {
  required_providers {
    dnacenter = {
      version = "1.1.22-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
resource "dnacenter_pnp_virtual_account_devices_sync" "example" {
  provider = dnacenter

  parameters {
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
}
