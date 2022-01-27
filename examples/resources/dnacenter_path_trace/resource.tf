
resource "dnacenter_path_trace" "example" {
  provider = dnacenter
  parameters {

    control_path     = "false"
    dest_ip          = "string"
    dest_port        = "string"
    flow_analysis_id = "string"
    inclusions       = ["string"]
    periodic_refresh = "false"
    protocol         = "string"
    source_ip        = "string"
    source_port      = "string"
  }
}

output "dnacenter_path_trace_example" {
  value = dnacenter_path_trace.example
}