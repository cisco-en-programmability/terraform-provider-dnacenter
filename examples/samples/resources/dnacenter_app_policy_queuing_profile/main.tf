
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.17-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_app_policy_queuing_profile" "example" {
  provider = dnacenter
  parameters {
    payload {
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
      clause {

        instance_id = 337340
        interface_speed_bandwidth_clauses {

          instance_id     = 1
          interface_speed = "ALL"
          tc_bandwidth_settings {

            bandwidth_percentage = 50
            instance_id          = 1
            traffic_class        = "NETWORK_CONTROL"
          }
          tc_bandwidth_settings {

            bandwidth_percentage = 10
            instance_id          = 1
            traffic_class        = "BROADCAST_VIDEO"
          }
          tc_bandwidth_settings {

            bandwidth_percentage = 10
            instance_id          = 1
            traffic_class        = "SIGNALING"
          }
          tc_bandwidth_settings {

            bandwidth_percentage = 5
            instance_id          = 1
            traffic_class        = "VOIP_TELEPHONY"
          }
          tc_bandwidth_settings {

            bandwidth_percentage = 5
            instance_id          = 1
            traffic_class        = "BULK_DATA"
          }

          tc_bandwidth_settings {

            bandwidth_percentage = 5
            instance_id          = 1
            traffic_class        = "CONSUMER"
          }

          tc_bandwidth_settings {

            bandwidth_percentage = 2
            instance_id          = 1
            traffic_class        = "REAL_TIME_INTERACTIVE"
          }

          tc_bandwidth_settings {

            bandwidth_percentage = 2
            instance_id          = 1
            traffic_class        = "MULTIMEDIA_STREAMING"
          }

          tc_bandwidth_settings {

            bandwidth_percentage = 2
            instance_id          = 1
            traffic_class        = "MULTIMEDIA_CONFERENCING"
          }

          tc_bandwidth_settings {

            bandwidth_percentage = 2
            instance_id          = 1
            traffic_class        = "BEST_EFFORT"
          }

          tc_bandwidth_settings {

            bandwidth_percentage = 1
            instance_id          = 1
            traffic_class        = "SCAVENGER"
          }

          tc_bandwidth_settings {

            bandwidth_percentage = 6
            instance_id          = 1
            traffic_class        = "TRANSACTIONAL_DATA"
          }
        }
        is_common_between_all_interface_speeds = "true"
        # tc_dscp_settings {

        #   dscp          = "string"
        #   instance_id   = 1
        #   traffic_class = "string"
        # }
        type = "DSCP_CUSTOMIZATION"
      }
      description = "Cisco Validated Design Queuing Profile 2"
      # id          = "42e3255f-304c-4f3f-8f01-fc7e813721c9"
      name = "CVD_QUEUING_PROFILE2"
    }
  }
}

output "dnacenter_app_policy_queuing_profile_example" {
  value = dnacenter_app_policy_queuing_profile.example
}
