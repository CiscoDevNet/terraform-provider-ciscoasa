resource "aws_security_group" "asav" {
  name        = "asav"
  description = "only SSH ingress"
  vpc_id      = "${aws_vpc.main.id}"

  ingress {
    protocol  = "tcp"
    from_port = 22
    to_port   = 22

    cidr_blocks = ["${local.workstation_external_cidr}"]
  }

  ingress {
    protocol  = "tcp"
    from_port = 443
    to_port   = 443

    cidr_blocks = ["${local.workstation_external_cidr}"]
  }
}

#
# Usage Instructions: Once the instance has launched, connect to 
# your instance using an SSH client with the username 'admin' and 
# the ssh-key you selected during launch (no password required 
# with ssh-key). Please see Cisco ASAv Online Documentation for 
# further configuration information.
#
# https://www.cisco.com/c/en/us/td/docs/security/asa/asa99/asav/quick-start/asav-quick/asav-aws.html
#

resource "aws_network_interface" "mgmt" {
  subnet_id       = "${aws_subnet.main.id}"
  security_groups = ["${aws_security_group.asav.id}"]
}

resource "aws_eip" "mgmt" {
  vpc               = true
  network_interface = "${aws_network_interface.mgmt.id}"
}

resource "aws_network_interface" "if" {
  subnet_id         = "${aws_subnet.main.id}"
  source_dest_check = false
}

#
# Note, you have to manually subscribe to this first in 
# the console.
#
# https://aws.amazon.com/marketplace/pp/B00WH2LGM0
#
data "aws_ami" "asav" {
  most_recent = true
  owners      = ["aws-marketplace"]

  filter {
    name   = "name"
    values = ["asav*"]
  }

  filter {
    name   = "product-code"
    values = ["80uds1joqwlz35hw1lx5h1bcc"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}

resource "aws_instance" "asav" {
  ami           = "${data.aws_ami.asav.id}"
  instance_type = "c4.large"
  key_name      = "${aws_key_pair.ssh.key_name}"

  root_block_device {
    delete_on_termination = true
    volume_size           = "10"
    volume_type           = "standard"
  }

  network_interface {
    network_interface_id = "${aws_network_interface.mgmt.id}"
    device_index         = 0
  }

  network_interface {
    network_interface_id = "${aws_network_interface.if.id}"
    device_index         = 1
  }

  user_data = <<ZERODAY
! ASA 9.9(2)1
interface management0/0
management-only
nameif management
security-level 100
ip address dhcp setroute
no shut
interface gigabitethernet0/0
nameif inside
security-level 100
ip address dhcp setroute
no shut
!
same-security-traffic permit inter-interface
same-security-traffic permit intra-interface
!
crypto key generate rsa modulus 2048
ssh 0 0 management
ssh timeout 30
username admin password ${var.admin_password} privilege 15
username admin attributes
service-type admin
! required config end
! example dns configuration
dns domain-lookup management
DNS server-group DefaultDNS
! where this address is the.2 on your public subnet
name-server 172.19.0.2
! example ntp configuration
name 129.6.15.28 time-a.nist.gov
name 129.6.15.29 time-b.nist.gov
name 129.6.15.30 time-c.nist.gov
ntp server time-c.nist.gov
ntp server time-b.nist.gov
ntp server time-a.nist.gov
! REST API stuff
rest-api image boot:/asa-restapi-132300-lfbff-k8.SPA
aaa authentication http console LOCAL
http server enable
http 0.0.0.0 0.0.0.0 management
rest-api agent
ZERODAY

  tags = {
    Name = "terraform-provider-ciscoasa acc test"
  }
}
