
resource "dnacenter_license_device" "example" {
  provider = dnacenter
  parameters {

    device_uuids         = ["string"]
    smart_account_id     = "string"
    virtual_account_name = "string"
  }
}

output "dnacenter_license_device_example" {
  value = dnacenter_license_device.example
}