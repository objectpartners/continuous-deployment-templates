output "rancher_server_url" {
  value = "${module.server.instance_public_dns}:${var.port}"
}
