
data "dnacenter_license_device_registration" "example" {
  provider             = dnacenter
  virtual_account_name = "string"
  device_uuids         = ["string"]
}