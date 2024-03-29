
data "dnacenter_event_artifact" "example" {
  provider  = dnacenter
  event_ids = "string"
  limit     = 1
  offset    = 1
  order     = "string"
  search    = "string"
  sort_by   = "string"
  tags      = "string"
}

output "dnacenter_event_artifact_example" {
  value = data.dnacenter_event_artifact.example.items
}
