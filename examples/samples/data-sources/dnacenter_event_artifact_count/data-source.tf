
data "dnacenter_event_artifact_count" "example" {
  provider = dnacenter
}

output "dnacenter_event_artifact_count_example" {
  value = data.dnacenter_event_artifact_count.example.item
}
