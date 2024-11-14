
data "dnacenter_license_device_license_summary" "example" {
  provider             = dnacenter
  device_type          = "string"
  device_uuid          = "string"
  dna_level            = "string"
  limit                = 1
  order                = "string"
  page_number          = 1.0
  registration_status  = "string"
  smart_account_id     = "string"
  sort_by              = "string"
  virtual_account_name = "string"
}

output "dnacenter_license_device_license_summary_example" {
  value = data.dnacenter_license_device_license_summary.example.items
}
