
terraform {
  required_providers {
    dnacenter = {
      versions = ["0.2"]
      source   = "hashicorp.com/edu/dnacenter"
    }
  }
}

provider "dnacenter" {
}

resource "dna_pnp_global_settings" "response" {
  provider = dnacenter
  item {
    version = 1
    aaa_credentials {
      username = ""
      password = ""
    }
    task_time_outs {
      config_time_out         = 10
      image_download_time_out = 120
      general_time_out        = 20
    }
    sava_mapping_list {
    }
    accept_eula = true
    default_profile {
      ip_addresses = [
        "",
        "192.168.196.2",
        "10.121.1.5"
      ]
      fqdn_addresses = []
      port           = 443
      cert           = "-----BEGIN CERTIFICATE-----\nMIIDlzCCAn+gAwIBAgIJAIS4OLqH1S8xMA0GCSqGSIb3DQEBCwUAMGIxLTArBgNV\nBAMMJGViNTk2NjkxLWIzMDItNTIyNS1kNTY1LTAyMzdhYmRhYjdiNDEWMBQGA1UE\nCgwNQ2lzY28gU3lzdGVtczEZMBcGA1UECwwQQ2lzY28gRE5BIENlbnRlcjAeFw0x\nOTA1MTUxODMyMjlaFw0yMjAyMDgxODMyMjlaMGIxLTArBgNVBAMMJGViNTk2Njkx\nLWIzMDItNTIyNS1kNTY1LTAyMzdhYmRhYjdiNDEWMBQGA1UECgwNQ2lzY28gU3lz\ndGVtczEZMBcGA1UECwwQQ2lzY28gRE5BIENlbnRlcjCCASIwDQYJKoZIhvcNAQEB\nBQADggEPADCCAQoCggEBALtEGfJuPCle+vP+Sox46CVVg+k19m6ALNv+us1WIAH6\nC4d0tq/5XCQXmHnZtoqU7XFi52etQ9rQPUzERkx9vHQfh+DPemx4aI0YyHhXaE3N\naKQa4NebKQx4lop0Lsrcow/199wE6/SrtAScIEFTm1zHz165LSUNeDbJk8CD0A+G\nbudx4exRErGTMah4OSSGkbGd6nU0GsQsZ9YeW4GBMh0QdJYsObXEbTmLKhO08tHL\nl8kvJ63mMXAeJpsC0j96keOKBthuLvG2kpMn4KLph/z8+fH0oQbFBUUTyHlK/fgm\nBYOjC3l8nIgJYT42M3XhMA7ZcjGTukQvVjuurfZcgGECAwEAAaNQME4wHQYDVR0O\nBBYEFL3xxCmimsHnz7Yn99McQf8udIsqMB8GA1UdIwQYMBaAFL3xxCmimsHnz7Yn\n99McQf8udIsqMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBABpFDiwD\nR+QuOCowHkm4ygyy4pZchiJ+7LPsCFKEaGEfeZnMI1IDLz/IcDnG++yh0QVTCK/Z\niMaLX3lzEli4JRjy1Somvgfb1pqeb63PCAZJmJliqo14nC4daoaBPdJbtEIKNuDZ\nK/Ly8NuRhTwN6QjQLi7zHwDwAkhFSfiVs1Iocy5/dmPpNVuQ1HXU/Wog7v0FLKtp\nM9XcJgwrQlpPCwU1i3awoD67pcbyW7/fIuC/aK4EAi+7yIvb6acyfnlDWoArd2qr\n0IXVMPhZYVA1REedFkVNEBP3VQGf20WFlYhIuyOHH/PEOq4RXC6jYcB3gr+NJ2KG\nUsHkw3PBc8KbCtI=\n-----END CERTIFICATE-----"
      proxy          = false
    }
  }
}

output "dna_pnp_global_settings_response" {
  value = dna_pnp_global_settings.response
}

