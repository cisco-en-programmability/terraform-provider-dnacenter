# Configure provider with your Cisco Catalyst Center SDK credentials
provider "dnacenter" {
  # Cisco Catalyst Center user name
  username = "admin"
  # it can be set using the environment variable DNAC_BASE_URL

  # Cisco Catalyst Center password
  password = "admin123"
  # it can be set using the environment variable DNAC_USERNAME

  # Cisco Catalyst Center base URL, FQDN or IP
  base_url = "https://172.168.196.2"
  # it can be set using the environment variable DNAC_PASSWORD

  # Boolean to enable debugging
  debug = "false"
  # it can be set using the environment variable DNAC_DEBUG

  # Boolean to enable or disable SSL certificate verification
  ssl_verify = "false"
  # it can be set using the environment variable DNAC_SSL_VERIFY
}