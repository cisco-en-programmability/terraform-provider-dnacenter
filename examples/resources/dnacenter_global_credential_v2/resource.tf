
resource "dnacenter_global_credential_v2" "example" {
  provider = dnacenter

  parameters {

    cli_credential {

      description     = "string"
      enable_password = "string"
      id              = "string"
      password        = "******"
      username        = "string"
    }
    https_read {

      description = "string"
      id          = "string"
      password    = "******"
      port        = 1
      username    = "string"
    }
    https_write {

      description = "string"
      id          = "string"
      password    = "******"
      port        = 1
      username    = "string"
    }
    id = "string"
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

output "dnacenter_global_credential_v2_example" {
  value = dnacenter_global_credential_v2.example
}