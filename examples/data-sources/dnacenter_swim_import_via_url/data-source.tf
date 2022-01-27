
data "dnacenter_swim_import_via_url" "example" {
  provider         = dnacenter
  schedule_at      = "string"
  schedule_desc    = "string"
  schedule_origin  = "string"
  application_type = "string"
  image_family     = "string"
  source_url       = "string"
  third_party      = "false"
  vendor           = "string"
}