package res

// TODO: Refactor for string literal for tokenization and jit-construct support. <rbt>

type IdsIndex struct {
	AppSecGroup                                         string
	ClientConfig                                        string
	CosmosDBAccount                                     string
	CosmosDBPostgresDatabase                            string
	KubernetesCluster                                   string
	NetworkInterface                                    string
	NetworkInterfaceApplicationSecurityGroupAssociation string
	NetworkInterfaceNetworkSecurityGroupAssociation     string
	NetworkSecurityGroup                                string
	PostgresFlexibleServer                              string
	PostgresAdmin                                       string
	PrivateDNSZone                                      string
	PrivateDNSZoneGroup                                 string
	PrivateDNSZoneVNetLink                              string
	PrivateEndpoint                                     string
	PublicIP                                            string
	ResourceGroup                                       string
	Subnet                                              string
	SubnetNSGAssoc                                      string
	VirtualMachine                                      string
	VirtualNetwork                                      string
}

var Ids = IdsIndex{
	AppSecGroup:              "application_security_group",
	ClientConfig:             "client_config",
	CosmosDBAccount:          "cosmosdb_account",
	CosmosDBPostgresDatabase: "cosmosdb_postgres_database",
	KubernetesCluster:        "kubernetes_cluster",
	NetworkInterface:         "network_interface",
	NetworkInterfaceApplicationSecurityGroupAssociation: "network_interface_asg_association",
	NetworkInterfaceNetworkSecurityGroupAssociation:     "network_interface_nsg_association",
	NetworkSecurityGroup:                                "network_security_group",
	PostgresFlexibleServer:                              "postgresql_flexible_server",
	PostgresAdmin:                                       "postgresql_flexible_server_active_directory_admin",
	PrivateDNSZone:                                      "private_dns_zone",
	PrivateDNSZoneGroup:                                 "private_dns_zone_group",
	PrivateDNSZoneVNetLink:                              "private_dns_zone_virtual_network_link",
	PrivateEndpoint:                                     "private_endpoint",
	PublicIP:                                            "public_ip",
	ResourceGroup:                                       "resource_group",
	Subnet:                                              "subnet",
	SubnetNSGAssoc:                                      "subnet_nsg_assoc",
	VirtualMachine:                                      "virtual_machine",
	VirtualNetwork:                                      "virtual_network",
}
