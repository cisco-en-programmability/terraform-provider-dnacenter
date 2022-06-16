
resource "dnacenter_network_device_list" "example" {
  provider = dnacenter
  parameters {

    cli_transport           = "string"
    compute_device          = "false"
    enable_password         = "string"
    extended_discovery_info = "string"
    http_password           = "string"
    http_port               = "string"
    http_secure             = "false"
    http_user_name          = "string"
    ip_address              = ["string"]
    meraki_org_id           = ["string"]
    netconf_port            = "string"
    password                = "******"
    serial_number           = "string"
    snmp_auth_passphrase    = "string"
    snmp_auth_protocol      = "string"
    snmp_mode               = "string"
    snmp_priv_passphrase    = "string"
    snmp_priv_protocol      = "string"
    snmp_ro_community       = "string"
    snmp_rw_community       = "string"
    snmp_retry              = 1
    snmp_timeout            = 1
    snmp_user_name          = "string"
    snmp_version            = "string"
    type                    = "string"
    update_mgmt_ipaddress_list {

      exist_mgmt_ip_address = "string"
      new_mgmt_ip_address   = "string"
    }
    user_name = "string"
  }
}

output "dnacenter_network_device_list_example" {
  value = dnacenter_network_device_list.example
}