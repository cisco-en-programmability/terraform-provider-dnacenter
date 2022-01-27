
data "dnacenter_event_series" "example" {
  provider   = dnacenter
  category   = "string"
  domain     = "string"
  end_time   = "hh:mm"
  event_ids  = "string"
  limit      = "#"
  offset     = "#"
  order      = "string"
  severity   = "string"
  sort_by    = "string"
  source     = "string"
  start_time = "hh:mm"
  sub_domain = "string"
  type       = "string"
}

output "dnacenter_event_series_example" {
  value = data.dnacenter_event_series.example.items
}
