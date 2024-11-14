
resource "dnacenter_pnp_global_settings" "example" {
  provider = dnacenter

  parameters {

    accept_eula = "string"
    default_profile {

      cert           = "string"
      fqdn_addresses = ["string"]
      ip_addresses   = ["string"]
      port           = "string"
      proxy          = "string"
    }
    id = "string"
    sava_mapping_list {

      cco_user = "string"
      expiry   = "string"
      profile {

        address_fqdn  = "string"
        address_ip_v4 = "string"
        cert          = "string"
        make_default  = "string"
        name          = "string"
        port          = "string"
        profile_id    = "string"
        proxy         = "string"
      }
      smart_account_id   = "string"
      virtual_account_id = "string"
    }
  }
}

output "dnacenter_pnp_global_settings_example" {
  value = dnacenter_pnp_global_settings.example
}