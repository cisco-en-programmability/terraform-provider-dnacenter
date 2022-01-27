
data "dnacenter_event_count" "example" {
  provider = dnacenter
  event_id = "string"
  tags     = "string"
}

output "dnacenter_event_count_example" {
  value = data.dnacenter_event_count.example.item
}
