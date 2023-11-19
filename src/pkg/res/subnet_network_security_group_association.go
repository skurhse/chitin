package res

import (
	"fmt"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	nsg "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/networksecuritygroup"
	sn "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/subnet"
	as "github.com/transprogrammer/xenia/generated/hashicorp/azurerm/subnetnetworksecuritygroupassociation"
)

func NewSubnetNSGAssoc(stk cdktf.TerraformStack, subnet sn.Subnet, nsg nsg.NetworkSecurityGroup, token string) as.SubnetNetworkSecurityGroupAssociation {

	id := fmt.Sprintf("%s_%s", Ids.SubnetNSGAssoc, token)

	input := as.SubnetNetworkSecurityGroupAssociationConfig{
		SubnetId:               subnet.Id(),
		NetworkSecurityGroupId: nsg.Id(),
	}

	return as.NewSubnetNetworkSecurityGroupAssociation(stk, &id, &input)
}
