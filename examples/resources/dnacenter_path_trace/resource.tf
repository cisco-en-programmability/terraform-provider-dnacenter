
resource "dnacenter_path_trace" "example" {
  provider = dnacenter
  item {

    detailed_status {



    }

    network_elements {

      accuracy_list {



      }
      detailed_status {



      }
      device_statistics {

        cpu_statistics {





        }
        memory_statistics {




        }
      }


      egress_physical_interface {

        acl_analysis {


          matching_aces {


            matching_ports {

              ports {



              }

            }

          }

        }

        interface_statistics {















        }



        path_overlay_info {








          vxlan_info {



          }
        }
        qos_statistics {











        }




      }
      egress_virtual_interface {

        acl_analysis {


          matching_aces {


            matching_ports {

              ports {



              }

            }

          }

        }

        interface_statistics {















        }



        path_overlay_info {








          vxlan_info {



          }
        }
        qos_statistics {











        }




      }
      flex_connect {



        egress_acl_analysis {


          matching_aces {


            matching_ports {

              ports {



              }

            }

          }

        }
        ingress_acl_analysis {


          matching_aces {


            matching_ports {

              ports {



              }

            }

          }

        }


      }

      ingress_physical_interface {

        acl_analysis {


          matching_aces {


            matching_ports {

              ports {



              }

            }

          }

        }

        interface_statistics {















        }



        path_overlay_info {








          vxlan_info {



          }
        }
        qos_statistics {











        }




      }
      ingress_virtual_interface {

        acl_analysis {


          matching_aces {


            matching_ports {

              ports {



              }

            }

          }

        }

        interface_statistics {















        }



        path_overlay_info {








          vxlan_info {



          }
        }
        qos_statistics {











        }




      }





      perf_mon_statistics {



















      }





    }
    network_elements_info {

      accuracy_list {



      }
      detailed_status {



      }
      device_statistics {

        cpu_statistics {





        }
        memory_statistics {




        }
      }


      egress_interface {

        physical_interface {

          acl_analysis {


            matching_aces {


              matching_ports {

                ports {



                }

              }

            }

          }

          interface_statistics {















          }



          path_overlay_info {








            vxlan_info {



            }
          }
          qos_statistics {











          }




        }
        virtual_interface {

          acl_analysis {


            matching_aces {


              matching_ports {

                ports {



                }

              }

            }

          }

          interface_statistics {















          }



          path_overlay_info {








            vxlan_info {



            }
          }
          qos_statistics {











          }




        }
      }
      flex_connect {



        egress_acl_analysis {


          matching_aces {


            matching_ports {

              ports {



              }

            }

          }

        }
        ingress_acl_analysis {


          matching_aces {


            matching_ports {

              ports {



              }

            }

          }

        }


      }

      ingress_interface {

        physical_interface {

          acl_analysis {


            matching_aces {


              matching_ports {

                ports {



                }

              }

            }

          }

          interface_statistics {















          }



          path_overlay_info {








            vxlan_info {



            }
          }
          qos_statistics {











          }




        }
        virtual_interface {

          acl_analysis {


            matching_aces {


              matching_ports {

                ports {



                }

              }

            }

          }

          interface_statistics {















          }



          path_overlay_info {








            vxlan_info {



            }
          }
          qos_statistics {











          }




        }
      }





      perf_monitor_statistics {



















      }





    }

    request {














    }
  }
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