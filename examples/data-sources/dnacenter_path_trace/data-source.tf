
data "dnacenter_path_trace" "example" {
  provider         = dnacenter
  dest_ip          = "string"
  dest_port        = 1.0
  gt_create_time   = 1.0
  last_update_time = 1.0
  limit            = 1
  lt_create_time   = 1.0
  offset           = 1
  order            = "string"
  periodic_refresh = "false"
  protocol         = "string"
  sort_by          = "string"
  source_ip        = "string"
  source_port      = 1.0
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
