
# 🦊 xenia

a Terraform CDK project that creates and provisions Azure Kubernetes Service clusters, managed via GitHub Actions.

## providers
  - [azuread](https://registry.terraform.io/providers/hashicorp/azuread/latest)
  - [azurerm](https://registry.terraform.io/providers/hashicorp/azurerm/latest)

## modules
- [naming](https://registry.terraform.io/modules/Azure/naming/azurerm/latest)

## resources
- [core](src/pkg/stk/core.go)
  - [azurerm_virtual_network](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_network)
  - [azurerm_application_security_group](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/application_security_group)
  - [azurerm_network_security_group](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_security_group)
  - [azuread_group](https://registry.terraform.io/providers/hashicorp/azuread/latest/docs/resources/group)
- [jumpbox](src/pkg/stk/jump.go)
  - [azurerm_public_ip](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/public_ip)
  - [azurerm_network_interface](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_interface)
  - [azurerm_linux_virtual_machine](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/linux_virtual_machine)
- [postgres](src/pkg/stk/postgres.go)
  - [azurerm_postgresql_flexible_server](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_flexible_server)
  - [azurerm_private_dns_zone](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/private_dns_zone)
  - [azurerm_private_dns_zone_virtual_network_link](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/private_dns_zone_virtual_network_link)
  - [azurerm_postgresql_flexible_server_database](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_flexible_server_database)
  - [azurerm_postgresql_flexible_server_firewall_rule](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_flexible_server_firewall_rule)
  - [azurerm_private_endpoint](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/private_endpoint)
  - [azurerm_postgresql_flexible_server_active_directory_administrator](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/postgresql_flexible_server_active_directory_administrator)
  - []()
- [cluster](src/pkg/stk/cluster.go)
  - [azurerm_kubernetes_cluster](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/kubernetes_cluster)
  - 

## actions
- [terraform-cdk-action](https://github.com/marketplace/actions/terraform-cdk-action)
- [deploy-to-kubernetes-cluster](https://github.com/marketplace/actions/deploy-to-kubernetes-cluster)
