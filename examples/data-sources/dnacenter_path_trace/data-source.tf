
data "dnacenter_path_trace" "example" {
  provider         = dnacenter
  dest_ip          = "string"
  dest_port        = "string"
  gt_create_time   = "string"
  last_update_time = "string"
  limit            = "string"
  lt_create_time   = "string"
  offset           = "string"
  order            = "string"
  periodic_refresh = "false"
  protocol         = "string"
  sort_by          = "string"
  source_ip        = "string"
  source_port      = "string"
  status           = "string"
  task_id          = "string"
}

output "dnacenter_path_trace_example" {
  value = data.dnacenter_path_trace.example.items
}

data "dnacenter_path_trace" "example" {
  provider         = dnacenter
  flow_analysis_id = "string"
}

output "dnacenter_path_trace_example" {
  value = data.dnacenter_path_trace.example.item
}
