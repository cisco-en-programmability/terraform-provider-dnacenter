
data "dnacenter_license_device_count" "example" {
  provider             = dnacenter
  device_type          = "string"
  dna_level            = "string"
  registration_status  = "string"
  smart_account_id     = "string"
  virtual_account_name = "string"
}

output "dnacenter_license_device_count_example" {
  value = data.dnacenter_license_device_count.example.item
}
