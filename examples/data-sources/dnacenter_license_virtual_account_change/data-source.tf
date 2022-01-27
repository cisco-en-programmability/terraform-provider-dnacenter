
data "dnacenter_license_virtual_account_change" "example" {
  provider             = dnacenter
  smart_account_id     = "string"
  virtual_account_name = "string"
  device_uuids         = ["string"]
}