
terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

resource "dnacenter_app_policy_queuing_profile" "example" {
  provider = dnacenter
  parameters {

    # clause {

    #   instance_id = 337341
    #   # interface_speed_bandwidth_clauses {

    #   #   instance_id     = 1
    #   #   interface_speed = "string"
    #   #   tc_bandwidth_settings {

    #   #     bandwidth_percentage = 1
    #   #     instance_id          = 1
    #   #     traffic_class        = "string"
    #   #   }
    #   # }
    #   is_common_between_all_interface_speeds = "true"
    #   # tc_dscp_settings {

    #   #   dscp          = "string"
    #   #   instance_id   = 1
    #   #   traffic_class = "string"
    #   # }
    #   type = "BANDWIDTH"
    # }
    # clause {

    #   instance_id = 337340
    #   # interface_speed_bandwidth_clauses {

    #   #   instance_id     = 1
    #   #   interface_speed = "string"
    #   #   tc_bandwidth_settings {

    #   #     bandwidth_percentage = 1
    #   #     instance_id          = 1
    #   #     traffic_class        = "string"
    #   #   }
    #   # }
    #   is_common_between_all_interface_speeds = "true"
    #   # tc_dscp_settings {

    #   #   dscp          = "string"
    #   #   instance_id   = 1
    #   #   traffic_class = "string"
    #   # }
    #   type = "DSCP_CUSTOMIZATION"
    # }
    description = "Cisco Validated Design Queuing Profile 3"
    # id          = "42e3255f-304c-4f3f-8f01-fc7e813721c9"
    name = "CVD_QUEUING_PROFILE"
  }
}

output "dnacenter_app_policy_queuing_profile_example" {
  value = dnacenter_app_policy_queuing_profile.example
}
