
resource "dnacenter_discovery" "example" {
  provider = dnacenter
  parameters {

    attribute_info            = ["string"]
    cdp_level                 = 1
    device_ids                = "string"
    discovery_condition       = "string"
    discovery_status          = "string"
    discovery_type            = "string"
    enable_password_list      = ["string"]
    global_credential_id_list = ["string"]
    http_read_credential {

      comments           = "string"
      credential_type    = "string"
      description        = "string"
      id                 = "string"
      instance_tenant_id = "string"
      instance_uuid      = "string"
      password           = "******"
      port               = 1
      secure             = "false"
      username           = "string"
    }
    http_write_credential {

      comments           = "string"
      credential_type    = "string"
      description        = "string"
      id                 = "string"
      instance_tenant_id = "string"
      instance_uuid      = "string"
      password           = "******"
      port               = 1
      secure             = "false"
      username           = "string"
    }
    id                      = "string"
    ip_address_list         = "string"
    ip_filter_list          = ["string"]
    is_auto_cdp             = "false"
    lldp_level              = 1
    name                    = "string"
    netconf_port            = "string"
    num_devices             = 1
    parent_discovery_id     = "string"
    password_list           = ["******"]
    preferred_mgmt_ipmethod = "string"
    protocol_order          = "string"
    retry                   = 1
    retry_count             = 1
    snmp_auth_passphrase    = "string"
    snmp_auth_protocol      = "string"
    snmp_mode               = "string"
    snmp_priv_passphrase    = "string"
    snmp_priv_protocol      = "string"
    snmp_ro_community       = "string"
    snmp_ro_community_desc  = "string"
    snmp_rw_community       = "string"
    snmp_rw_community_desc  = "string"
    snmp_user_name          = "string"
    snmp_version            = "string"
    time_out                = 1
    timeout                 = 1
    update_mgmt_ip          = "false"
    user_name_list          = ["string"]
  }
}

output "dnacenter_discovery_example" {
  value = dnacenter_discovery.example
}