
data "dnacenter_device_credential_update" "example" {
  provider = dnacenter
  settings {

    cli_credential {

      description     = "string"
      enable_password = "string"
      id              = "string"
      password        = "******"
      username        = "string"
    }
    https_read {

      id       = "string"
      name     = "string"
      password = "******"
      port     = "string"
      username = "string"
    }
    https_write {

      id       = "string"
      name     = "string"
      password = "******"
      port     = "string"
      username = "string"
    }
    snmp_v2c_read {

      description    = "string"
      id             = "string"
      read_community = "string"
    }
    snmp_v2c_write {

      description     = "string"
      id              = "string"
      write_community = "string"
    }
    snmp_v3 {

      auth_password    = "string"
      auth_type        = "string"
      description      = "string"
      id               = "string"
      privacy_password = "string"
      privacy_type     = "string"
      snmp_mode        = "string"
      username         = "string"
    }
  }
}