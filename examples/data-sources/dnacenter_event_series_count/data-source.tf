
data "dnacenter_event_series_count" "example" {
  provider   = dnacenter
  category   = "string"
  domain     = "string"
  end_time   = "hh:mm"
  event_ids  = "string"
  severity   = "string"
  source     = "string"
  start_time = "hh:mm"
  sub_domain = "string"
  type       = "string"
}

output "dnacenter_event_series_count_example" {
  value = data.dnacenter_event_series_count.example.item
}
