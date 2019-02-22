resource "tls_private_key" "ssh" {
  algorithm = "RSA"
}

resource "aws_key_pair" "ssh" {
  key_name_prefix = "terraform-provider-ciscoasa"
  public_key      = "${tls_private_key.ssh.public_key_openssh}"
}

resource "local_file" "public_key_openssh" {
  content  = "${tls_private_key.ssh.public_key_openssh}"
  filename = "ssh.pub"
}

resource "local_file" "private_key_pem" {
  content  = "${tls_private_key.ssh.private_key_pem}"
  filename = "ssh.pem"

  provisioner "local-exec" {
    command = "${format("chmod 0600 %v", self.filename)}"
  }
}
