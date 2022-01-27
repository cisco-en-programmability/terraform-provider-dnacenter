
terraform {
  required_providers {
  dnacenter = {
    version = "0.0.3"
    source  = "hashicorp.com/edu/dnacenter"
    # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
  }
  }
}
resource "dnacenter_path_trace" "example" {
  provider = dnacenter
  parameters {

    control_path     = "true"
    dest_ip          = "10.1.1.2"
    dest_port        = "8080"
    #flow_analysis_id = "false"
    #inclusions       = ["string"]
    periodic_refresh = "false"
    protocol         = "udp"
    source_ip        = "10.1.1.3"
    source_port      = "8082"
  }
}

output "dnacenter_path_trace_example" {
  value = dnacenter_path_trace.example
}