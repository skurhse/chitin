package resources

import "github.com/aws/jsii-runtime-go"

type IdsIndex struct {
	AppSecGroup                                         *string
	CosmosDBAccount                                     *string
	CosmosDBMongoDatabase                               *string
	NetworkInterface                                    *string
	NetworkInterfaceApplicationSecurityGroupAssociation *string
	NetworkInterfaceNetworkSecurityGroupAssociation     *string
	NetworkSecurityGroup                                *string
	PrivateDNSZone                                      *string
	PrivateDNSZoneGroup                                 *string
	PrivateDNSZoneVNetLink                              *string
	PrivateEndpoint                                     *string
	PublicIP                                            *string
	ResourceGroup                                       *string
	Subnet                                              *string
	VirtualMachine                                      *string
	VirtualNetwork                                      *string
}

var Ids = IdsIndex{
	AppSecGroup:      jsii.String("application_security_group"),
	CosmosDBAccount:  jsii.String("cosmosdb_account"),
	NetworkInterface: jsii.String("network_interface"),
	NetworkInterfaceApplicationSecurityGroupAssociation: jsii.String("network_interface_asg_association"),
	NetworkInterfaceNetworkSecurityGroupAssociation:     jsii.String("network_interface_nsg_association"),
	NetworkSecurityGroup:                                jsii.String("network_security_group"),
	PrivateDNSZone:                                      jsii.String("private_dns_zone"),
	PrivateDNSZoneGroup:                                 jsii.String("private_dns_zone_group"),
	PrivateDNSZoneVNetLink:                              jsii.String("private_dns_zone_virtual_network_link"),
	PrivateEndpoint:                                     jsii.String("private_endpoint"),
	PublicIP:                                            jsii.String("public_ip"),
	ResourceGroup:                                       jsii.String("resource_group"),
	Subnet:                                              jsii.String("subnet"),
	VirtualMachine:                                      jsii.String("virtual_machine"),
	VirtualNetwork:                                      jsii.String("virtual_network"),
}
