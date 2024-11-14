
resource "dnacenter_pnp_server_profile_update" "example" {
  provider = dnacenter
  parameters {

    cco_user = "string"
    profile {

      address_fqdn  = "string"
      address_ip_v4 = "string"
      address_ip_v6 = "string"
      cert          = "string"
      make_default  = "false"
      name          = "string"
      port          = 9090
      profile_id    = "string"
      proxy         = "false"
    }
    smart_account_id   = "string"
    virtual_account_id = "string"
  }
}

output "dnacenter_pnp_server_profile_update_example" {
  value = dnacenter_pnp_server_profile_update.example
}