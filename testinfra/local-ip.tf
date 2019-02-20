#
# Workstation External IP
#
# This configuration is not required and is
# only provided as an example to easily fetch
# the external IP of your local workstation to
# configure inbound EC2 Security Group access.
#

data "http" "workstation_external_ip" {
  url = "http://ipv4.icanhazip.com"
}

# Override with variable or hardcoded value if necessary
locals {
  workstation_external_cidr = "${chomp(data.http.workstation_external_ip.body)}/32"
}
