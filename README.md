
# 🦊 xenia

a Terraform CDK project that creates and provisions Azure Kubernetes Service clusters, managed via GitHub Actions.

## docs

- stacks
  - [core](src/pkg/stk/core.go)
    - [azurerm_virtual_network](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/virtual_network)
    - [azurerm_application_security_group](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/application_security_group)
    - [azurerm_network_security_group](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/network_security_group)
  - [jumpbox](src/pkg/stk/jump.go)
  - [postgres](src/pkg/stk/postgres.go)
  - [mongo](src/pkg/stk/mongo.go)
  - [cluster](src/pkg/stk/cluster.go)
- providers
  - [azurerm](https://github.com/hashicorp/terraform-provider-azurerm)
  - [kubernetes](https://github.com/hashicorp/terraform-provider-kubernetes)
  - [github](https://github.com/integrations/terraform-provider-github)
- modules
  - [azurerm-naming](https://github.com/Azure/terraform-azurerm-naming)
- actions
  - [terraform-cdk-action](https://github.com/marketplace/actions/terraform-cdk-action)
  - [deploy-to-kubernetes-cluster](https://github.com/marketplace/actions/deploy-to-kubernetes-cluster)
