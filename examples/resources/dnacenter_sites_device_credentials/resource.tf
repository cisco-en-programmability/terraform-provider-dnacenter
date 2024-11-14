
resource "dnacenter_sites_device_credentials" "example" {
  provider = dnacenter
  parameters {

    cli_credentials_id {

      credentials_id = "string"
    }
    http_read_credentials_id {

      credentials_id = "string"
    }
    http_write_credentials_id {

      credentials_id = "string"
    }
    id = "string"
    snmpv2c_read_credentials_id {

      credentials_id = "string"
    }
    snmpv2c_write_credentials_id {

      credentials_id = "string"
    }
    snmpv3_credentials_id {

      credentials_id = "string"
    }
  }
}

output "dnacenter_sites_device_credentials_example" {
  value = dnacenter_sites_device_credentials.example
}