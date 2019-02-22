output "asav_public_ip" {
  value = "${aws_eip.mgmt.public_ip}"
}

output "asav_username" {
  value = "${var.admin_username}"
}

output "asav_password" {
  value = "${var.admin_password}"
}

#
# until $(terraform output ssh); do; sleep 5; done
#
output "ssh" {
  value = "ssh -i ssh.pem -oConnectTimeout=5 -oKexAlgorithms=+diffie-hellman-group1-sha1 admin@${aws_instance.asav.public_ip}"
}
