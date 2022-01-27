
data "dnacenter_threat_summary" "example" {
  provider     = dnacenter
  end_time     = 1
  site_id      = ["string"]
  start_time   = 1
  threat_level = ["string"]
  threat_type  = ["string"]
}