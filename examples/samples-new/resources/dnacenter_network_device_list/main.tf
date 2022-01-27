
terraform {
  required_providers {
  dnacenter = {
    version = "0.0.3"
    source  = "hashicorp.com/edu/dnacenter"
    # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
  }
  }
}

resource "dnacenter_network_device_list" "example" {
  provider = dnacenter
  parameters {
    /*
    cli_transport           = "string"
    compute_device          = "false"
    enable_password         = "string"
    extended_discovery_info = "string"
    ip_address              = ["string"]
    meraki_org_id           = ["string"]
    netconf_port            = "string"
    password                = "******"
    */
    http_user_name          = "User"
    http_password           = "123"
    http_port               = "8081"
    http_secure             = "true"
    ip_address              = ["10.121.1.1"]
    # serial_number           = "FLM2213W05R"
    #type                    = "Cisco 4331 Integrated Services Router"
    type                    = "COMPUTE_DEVICE"
    /*
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
    type                    = "COMPUTE_DEVICE"
    update_mgmt_ipaddress_list {

      exist_mgmt_ip_address = "string"
      new_mgmt_ip_address   = "string"
    }
    user_name = "string"
    */
  }
}

output "dnacenter_network_device_list_example" {
  value = dnacenter_network_device_list.example
  sensitive = true
}