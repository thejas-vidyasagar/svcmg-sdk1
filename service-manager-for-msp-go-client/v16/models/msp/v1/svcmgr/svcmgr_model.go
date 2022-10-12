/*
 * Generated file models/msp/v1/svcmgr/svcmgr_model.go.
 *
 * Product version: 16.7.0.970-SNAPSHOT
 *
 * Part of the Service Manager for MSP
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
Module msp.v1.svcmgr of Service Manager for MSP
*/
package svcmgr

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/thejas-vidyasagar/svcmg-sdk1/service-manager-for-msp-go-client/v16/models/common/v1/config"
	import3 "github.com/thejas-vidyasagar/svcmg-sdk1/service-manager-for-msp-go-client/v16/models/common/v1/response"
	import2 "github.com/thejas-vidyasagar/svcmg-sdk1/service-manager-for-msp-go-client/v16/models/msp/v1/error"
	import4 "github.com/thejas-vidyasagar/svcmg-sdk1/service-manager-for-msp-go-client/v16/models/prism/v4/config"
	"time"
)

/*
*
Ahv client configuration.
*/
type AhvClientConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  client config CA (Certificate Authority) chain.
	*/
	CaChain *string `json:"caChain,omitempty"`
	/**
	  X.509 certificate.
	*/
	Certificate *string `json:"certificate,omitempty"`
	/**
	  client key.
	*/
	Key *string `json:"key,omitempty"`
	/**
	  client password for the username.
	*/
	Password *string `json:"password,omitempty"`

	PrismIp *import1.IPv4Address `json:"prismIp,omitempty"`
	/**
	  Prism Element Port.
	*/
	PrismPort *int `json:"prismPort,omitempty"`
	/**
	  client username.
	*/
	Username *string `json:"username,omitempty"`
}

func NewAhvClientConfig() *AhvClientConfig {
	p := new(AhvClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.AhvClientConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.AhvClientConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Network resources for Ahv.
*/
type AhvNetworkResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Cluster Id of the Prism Element which hosts the Kubernetes cluster.
	*/
	ClusterId *string `json:"clusterId,omitempty"`
	/**
	  Network id to be used for this vm config.
	*/
	Network *string `json:"network"`

	NicType *NicType `json:"nicType,omitempty"`
}

func (p *AhvNetworkResourceConfig) MarshalJSON() ([]byte, error) {
	type AhvNetworkResourceConfigProxy AhvNetworkResourceConfig
	return json.Marshal(struct {
		*AhvNetworkResourceConfigProxy
		Network *string `json:"network,omitempty"`
	}{
		AhvNetworkResourceConfigProxy: (*AhvNetworkResourceConfigProxy)(p),
		Network:                       p.Network,
	})
}

func NewAhvNetworkResourceConfig() *AhvNetworkResourceConfig {
	p := new(AhvNetworkResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.AhvNetworkResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.AhvNetworkResourceConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
AHV platform specific config.
*/
type AhvResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NetworkResource *AhvNetworkResourceConfig `json:"networkResource"`

	VmResource *VmResourceConfig `json:"vmResource"`
}

func (p *AhvResourceConfig) MarshalJSON() ([]byte, error) {
	type AhvResourceConfigProxy AhvResourceConfig
	return json.Marshal(struct {
		*AhvResourceConfigProxy
		NetworkResource *AhvNetworkResourceConfig `json:"networkResource,omitempty"`
		VmResource      *VmResourceConfig         `json:"vmResource,omitempty"`
	}{
		AhvResourceConfigProxy: (*AhvResourceConfigProxy)(p),
		NetworkResource:        p.NetworkResource,
		VmResource:             p.VmResource,
	})
}

func NewAhvResourceConfig() *AhvResourceConfig {
	p := new(AhvResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.AhvResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.AhvResourceConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type AppOrServiceHistory struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *string `json:"action,omitempty"`

	Apptype *string `json:"apptype,omitempty"`

	CreatedTimestamp *time.Time `json:"createdTimestamp,omitempty"`

	Message *string `json:"message,omitempty"`

	Name *string `json:"name,omitempty"`

	ParentAppUuid *string `json:"parentAppUuid,omitempty"`

	Status *string `json:"status,omitempty"`

	TaskUuid *string `json:"taskUuid,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	Version *string `json:"version,omitempty"`
}

func NewAppOrServiceHistory() *AppOrServiceHistory {
	p := new(AppOrServiceHistory)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.AppOrServiceHistory"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.AppOrServiceHistory"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Application struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *string `json:"action,omitempty"`

	Apptype *string `json:"apptype,omitempty"`

	CreatedTimestamp *time.Time `json:"createdTimestamp,omitempty"`

	LastUpdatedTimestamp *time.Time `json:"lastUpdatedTimestamp,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	SubServices []Service `json:"subServices,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	Version *string `json:"version,omitempty"`
}

func NewApplication() *Application {
	p := new(Application)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.Application"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.Application"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Applications struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ApplicationList []Application `json:"applicationList,omitempty"`
}

func NewApplications() *Applications {
	p := new(Applications)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.Applications"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.Applications"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
base client configuration.
*/
type BaseClientConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  X.509 certificate.
	*/
	Certificate *string `json:"certificate,omitempty"`
	/**
	  client key.
	*/
	Key *string `json:"key,omitempty"`
	/**
	  client password for the username.
	*/
	Password *string `json:"password,omitempty"`
	/**
	  client username.
	*/
	Username *string `json:"username,omitempty"`
}

func NewBaseClientConfig() *BaseClientConfig {
	p := new(BaseClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.BaseClientConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.BaseClientConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
client configuration for accessing ahv, esx platform.
*/
type ClientConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AhvClientConfig *AhvClientConfig `json:"ahvClientConfig,omitempty"`

	EsxClientConfig *EsxClientConfig `json:"esxClientConfig,omitempty"`
	/**
	  client configuration name.
	*/
	Name *string `json:"name"`
}

func (p *ClientConfig) MarshalJSON() ([]byte, error) {
	type ClientConfigProxy ClientConfig
	return json.Marshal(struct {
		*ClientConfigProxy
		Name *string `json:"name,omitempty"`
	}{
		ClientConfigProxy: (*ClientConfigProxy)(p),
		Name:              p.Name,
	})
}

func NewClientConfig() *ClientConfig {
	p := new(ClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.ClientConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.ClientConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Cluster struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Annotations are key value pair of strings which are meant to given to customize cluster deployment.
	Supported annotations are -
	  - Select storage provisioner service. Accepted values are csi/volume_provisioner.
	    Key: "msp/storage/service"
	    Value: "volume_provisioner/csi"

	  - Disable logging addons on the cluster.
	    Key: "msp/logging/disable"
	    Value: "true"

	  - Disable monitoring addons on the cluster.
	    Key: "msp/monitoring/disable"
	    Value: "true"
	*/
	Annotations []import1.KVStringPair `json:"annotations,omitempty"`

	AvailabilityType *ClusterAvailabilityType `json:"availabilityType"`
	/**
	  Client spec needed for VM/Network CRUD operations.
	*/
	ClientConfigs []ClientConfig `json:"clientConfigs,omitempty"`

	ClusterMode *ClusterMode `json:"clusterMode"`

	DeploymentType *ClusterDeploymentType `json:"deploymentType"`
	/**
	  Brief description about this cluster.
	*/
	Description *string `json:"description,omitempty"`

	Domain *import1.FQDN `json:"domain"`
	/**
	  Configuration for Load balancers.
	*/
	LoadBalancerConfigs []LoadBalancerConfig `json:"loadBalancerConfigs,omitempty"`

	ManagementIps *ClusterManagementIps `json:"managementIps,omitempty"`
	/**
	  Name of the cluster. Name should be a valid hostname consisting of alpha numeric characters and/or dash.
	*/
	Name *string `json:"name"`
	/**
	  Cluster network configuration along with other static IPs needed.
	*/
	NetworkConfigs []ClusterNetworkConfig `json:"networkConfigs,omitempty"`
	/**
	  The storage class specification used for creation of ABS volumes attached to worker VMs.
	*/
	PersistentVolume []PersistentVolumeConfig `json:"persistentVolume,omitempty"`
	/**
	  Resource configuration list for the VM resources.
	*/
	ResourceConfigs []ClusterResourceConfig `json:"resourceConfigs,omitempty"`
	/**
	  Default ssh user to login to msp vms.
	*/
	SshUser *string `json:"sshUser,omitempty"`
	/**
	  The storage class specification used for creation of volumes needed for stateful services.
	*/
	StorageClasses []ClusterStorageClass `json:"storageClasses"`

	WorkerConfigs []WorkerConfig `json:"workerConfigs,omitempty"`
	/**
	  Specifies whether to deploy xtrim service or not. Nutanix Xtrim is a tool which runs fstrim on a
	mount point continuously to discard blocks which are not in use by the filesystem.
	*/
	XtrimDeploy *bool `json:"xtrimDeploy,omitempty"`
}

func (p *Cluster) MarshalJSON() ([]byte, error) {
	type ClusterProxy Cluster
	return json.Marshal(struct {
		*ClusterProxy
		AvailabilityType *ClusterAvailabilityType `json:"availabilityType,omitempty"`
		ClusterMode      *ClusterMode             `json:"clusterMode,omitempty"`
		DeploymentType   *ClusterDeploymentType   `json:"deploymentType,omitempty"`
		Domain           *import1.FQDN            `json:"domain,omitempty"`
		Name             *string                  `json:"name,omitempty"`
		StorageClasses   []ClusterStorageClass    `json:"storageClasses,omitempty"`
	}{
		ClusterProxy:     (*ClusterProxy)(p),
		AvailabilityType: p.AvailabilityType,
		ClusterMode:      p.ClusterMode,
		DeploymentType:   p.DeploymentType,
		Domain:           p.Domain,
		Name:             p.Name,
		StorageClasses:   p.StorageClasses,
	})
}

func NewCluster() *Cluster {
	p := new(Cluster)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.Cluster"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.Cluster"}
	p.UnknownFields_ = map[string]interface{}{}

	p.XtrimDeploy = new(bool)
	*p.XtrimDeploy = true

	return p
}

/*
*
Availability type of Cluster in terms of cluster components like Master, Etcd, Workers.
HA : HA means components like Master, Etcd, Workers are HA, i.e. more than 1 instances run for each.
NonHA : NonHA may mean there maybe perhaps single instance of some components like Master, Etcd.
*/
type ClusterAvailabilityType int

const (
	CLUSTERAVAILABILITYTYPE_UNKNOWN  ClusterAvailabilityType = 0
	CLUSTERAVAILABILITYTYPE_REDACTED ClusterAvailabilityType = 1
	CLUSTERAVAILABILITYTYPE_HA       ClusterAvailabilityType = 2
	CLUSTERAVAILABILITYTYPE_NON_HA   ClusterAvailabilityType = 3
)

// returns the name of the enum given an ordinal number
func (e *ClusterAvailabilityType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HA",
		"NON_HA",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *ClusterAvailabilityType) index(name string) ClusterAvailabilityType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"HA",
		"NON_HA",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterAvailabilityType(idx)
		}
	}
	return CLUSTERAVAILABILITYTYPE_UNKNOWN
}

func (e *ClusterAvailabilityType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterAvailabilityType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterAvailabilityType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterAvailabilityType) Ref() *ClusterAvailabilityType {
	return &e
}

/*
*
Type of deployment whether it is combined or scale out.
COMBINED: Master/etcd components maybe clubbed together with worker nodes. Use this if there are lack of resources.
SCALE_OUT:  Master/etcd components will run in their own instances.
*/
type ClusterDeploymentType int

const (
	CLUSTERDEPLOYMENTTYPE_UNKNOWN   ClusterDeploymentType = 0
	CLUSTERDEPLOYMENTTYPE_REDACTED  ClusterDeploymentType = 1
	CLUSTERDEPLOYMENTTYPE_COMBINED  ClusterDeploymentType = 2
	CLUSTERDEPLOYMENTTYPE_SCALE_OUT ClusterDeploymentType = 3
)

// returns the name of the enum given an ordinal number
func (e *ClusterDeploymentType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COMBINED",
		"SCALE_OUT",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *ClusterDeploymentType) index(name string) ClusterDeploymentType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"COMBINED",
		"SCALE_OUT",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterDeploymentType(idx)
		}
	}
	return CLUSTERDEPLOYMENTTYPE_UNKNOWN
}

func (e *ClusterDeploymentType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterDeploymentType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterDeploymentType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterDeploymentType) Ref() *ClusterDeploymentType {
	return &e
}

/*
*
cluster management related ips.
*/
type ClusterManagementIps struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Name of the cluster network config.
	*/
	ClusterNetworkConfigName *string `json:"clusterNetworkConfigName,omitempty"`

	MasterVip *import1.IPv4Address `json:"masterVip,omitempty"`

	MspDnsIp *import1.IPv4Address `json:"mspDnsIp,omitempty"`
}

func NewClusterManagementIps() *ClusterManagementIps {
	p := new(ClusterManagementIps)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.ClusterManagementIps"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.ClusterManagementIps"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Cluster Mode specifies how the parameters will be derived
DEFAULT : Based on the availabilityType, the default configuration will be used.

	HA: Default configuration of 1 Master, 3 Etcd and 3 worker VMs.
	Non-HA: Everything is hosted on 1 VM.

NUM_WORKERS : Client specifies number of worker VMs + environment config.
*/
type ClusterMode int

const (
	CLUSTERMODE_UNKNOWN     ClusterMode = 0
	CLUSTERMODE_REDACTED    ClusterMode = 1
	CLUSTERMODE_DEFAULT     ClusterMode = 2
	CLUSTERMODE_NUM_WORKERS ClusterMode = 3
)

// returns the name of the enum given an ordinal number
func (e *ClusterMode) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DEFAULT",
		"NUM_WORKERS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *ClusterMode) index(name string) ClusterMode {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"DEFAULT",
		"NUM_WORKERS",
	}
	for idx := range names {
		if names[idx] == name {
			return ClusterMode(idx)
		}
	}
	return CLUSTERMODE_UNKNOWN
}

func (e *ClusterMode) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ClusterMode:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ClusterMode) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ClusterMode) Ref() *ClusterMode {
	return &e
}

/*
*
Cluster Network configuration.
*/
type ClusterNetworkConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  List of DNS servers for this network.
	*/
	DnsServers []import1.IPv4Address `json:"dnsServers,omitempty"`

	EsxLocation *VcenterLocation `json:"esxLocation,omitempty"`

	GatewayIp *import1.IPv4Address `json:"gatewayIp,omitempty"`

	IpRanges []IpRange `json:"ipRanges,omitempty"`
	/**
	  Name of the cluster network config.
	*/
	Name *string `json:"name"`
	/**
	  IP address representing the subnet mask.
	*/
	Netmask *string `json:"netmask,omitempty"`
	/**
	  Network managed by the IPAM (IP address management) entity.
	*/
	Network *string `json:"network"`
}

func (p *ClusterNetworkConfig) MarshalJSON() ([]byte, error) {
	type ClusterNetworkConfigProxy ClusterNetworkConfig
	return json.Marshal(struct {
		*ClusterNetworkConfigProxy
		Name    *string `json:"name,omitempty"`
		Network *string `json:"network,omitempty"`
	}{
		ClusterNetworkConfigProxy: (*ClusterNetworkConfigProxy)(p),
		Name:                      p.Name,
		Network:                   p.Network,
	})
}

func NewClusterNetworkConfig() *ClusterNetworkConfig {
	p := new(ClusterNetworkConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.ClusterNetworkConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.ClusterNetworkConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ClusterResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AhvResourceConfig *AhvResourceConfig `json:"ahvResourceConfig,omitempty"`
	/**
	  Name of client config to use.
	*/
	Client *string `json:"client,omitempty"`

	EsxResourceConfig *EsxResourceConfig `json:"esxResourceConfig,omitempty"`
	/**
	  Labels for this resource config. Its a key value pair of strings
	*/
	Labels []import1.KVStringPair `json:"labels,omitempty"`
	/**
	  Name of resource config.
	*/
	Name *string `json:"name"`
}

func (p *ClusterResourceConfig) MarshalJSON() ([]byte, error) {
	type ClusterResourceConfigProxy ClusterResourceConfig
	return json.Marshal(struct {
		*ClusterResourceConfigProxy
		Name *string `json:"name,omitempty"`
	}{
		ClusterResourceConfigProxy: (*ClusterResourceConfigProxy)(p),
		Name:                       p.Name,
	})
}

func NewClusterResourceConfig() *ClusterResourceConfig {
	p := new(ClusterResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.ClusterResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.ClusterResourceConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Cluster storage class.
*/
type ClusterStorageClass struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Nutanix Cluster Reference.
	*/
	ClusterId *string `json:"clusterId"`

	DataServiceIp *import1.IPv4Address `json:"dataServiceIp,omitempty"`
	/**
	  Data Service Port.
	*/
	DataServicePort *int `json:"dataServicePort,omitempty"`
	/**
	  File system of the storage volume.
	*/
	FileSystem *string `json:"fileSystem,omitempty"`
	/**
	  Whether the storage class is default.
	*/
	IsDefaultStorageClass *bool `json:"isDefaultStorageClass,omitempty"`

	IsFlashMode *bool `json:"isFlashMode,omitempty"`
	/**
	  Name of the storage class. Used as a key to refer to it.
	*/
	Name *string `json:"name"`
	/**
	  Password for Administrator username of prism element.
	*/
	Password *string `json:"password,omitempty"`

	PrismIp *import1.IPv4Address `json:"prismIp,omitempty"`
	/**
	  Prism Element Port.
	*/
	PrismPort *int `json:"prismPort,omitempty"`
	/**
	  Name of the storage container to be used.
	*/
	StorageContainer *string `json:"storageContainer,omitempty"`
	/**
	  Administrator username of prism element.
	*/
	User *string `json:"user,omitempty"`
}

func (p *ClusterStorageClass) MarshalJSON() ([]byte, error) {
	type ClusterStorageClassProxy ClusterStorageClass
	return json.Marshal(struct {
		*ClusterStorageClassProxy
		ClusterId *string `json:"clusterId,omitempty"`
		Name      *string `json:"name,omitempty"`
	}{
		ClusterStorageClassProxy: (*ClusterStorageClassProxy)(p),
		ClusterId:                p.ClusterId,
		Name:                     p.Name,
	})
}

func NewClusterStorageClass() *ClusterStorageClass {
	p := new(ClusterStorageClass)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.ClusterStorageClass"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.ClusterStorageClass"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ComponentDetails struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Error details if any.
	*/
	ErrorMessage *string `json:"errorMessage,omitempty"`
	/**
	  Time when the first health failure is seen.
	*/
	FirstFailureTime *time.Time `json:"firstFailureTime,omitempty"`
	/**
	  Messages published by the component.
	*/
	Messages []string `json:"messages,omitempty"`
	/**
	  Number of failures since the last success.
	*/
	NumFailures *int `json:"numFailures,omitempty"`
	/**
	  Time when the health is checked.
	*/
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

func NewComponentDetails() *ComponentDetails {
	p := new(ComponentDetails)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.ComponentDetails"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.ComponentDetails"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type CustomValue struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Service *string `json:"service,omitempty"`

	Values []CustomValueItem `json:"values,omitempty"`
}

func NewCustomValue() *CustomValue {
	p := new(CustomValue)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.CustomValue"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.CustomValue"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type CustomValueItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Key *string `json:"key,omitempty"`

	ValueItemDiscriminator_ *string `json:"$valueItemDiscriminator,omitempty"`

	Value *OneOfCustomValueItemValue `json:"value,omitempty"`
}

func NewCustomValueItem() *CustomValueItem {
	p := new(CustomValueItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.CustomValueItem"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.CustomValueItem"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CustomValueItem) GetValue() interface{} {
	if nil == p.Value {
		return nil
	}
	return p.Value.GetValue()
}

func (p *CustomValueItem) SetValue(v interface{}) error {
	if nil == p.Value {
		p.Value = NewOneOfCustomValueItemValue()
	}
	e := p.Value.SetValue(v)
	if nil == e {
		if nil == p.ValueItemDiscriminator_ {
			p.ValueItemDiscriminator_ = new(string)
		}
		*p.ValueItemDiscriminator_ = *p.Value.Discriminator
	}
	return e
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/applications/{appUUID} Delete operation
*/
type DeleteAppApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteAppApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteAppApiResponse() *DeleteAppApiResponse {
	p := new(DeleteAppApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.DeleteAppApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.DeleteAppApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteAppApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteAppApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteAppApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/services/{serviceUUID} Delete operation
*/
type DeleteServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteServiceApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewDeleteServiceApiResponse() *DeleteServiceApiResponse {
	p := new(DeleteServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.DeleteServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.DeleteServiceApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteServiceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteServiceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteServiceApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
*
Esx client configuration.
*/
type EsxClientConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  X.509 certificate.
	*/
	Certificate *string `json:"certificate,omitempty"`
	/**
	  VCenter extension Key.
	*/
	ExtensionKey *string `json:"extensionKey,omitempty"`
	/**
	  For insecure tls connection.
	*/
	IsInsecure *bool `json:"isInsecure,omitempty"`
	/**
	  client key.
	*/
	Key *string `json:"key,omitempty"`
	/**
	  client password for the username.
	*/
	Password *string `json:"password,omitempty"`
	/**
	  TLS certificate thumbprint.
	*/
	ThumbPrint *string `json:"thumbPrint,omitempty"`
	/**
	  client username.
	*/
	Username *string `json:"username,omitempty"`

	VcenterIp *import1.IPv4Address `json:"vcenterIp,omitempty"`
	/**
	  Vcenter Port.
	*/
	VcenterPort *int `json:"vcenterPort,omitempty"`
}

func NewEsxClientConfig() *EsxClientConfig {
	p := new(EsxClientConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.EsxClientConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.EsxClientConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Network resources for Esx.
*/
type EsxNetworkResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Cluster Id of the Prism Element which hosts the Kubernetes cluster.
	*/
	ClusterId *string `json:"clusterId,omitempty"`

	Location *VcenterLocation `json:"location,omitempty"`
	/**
	  Network id to be used for this vm config.
	*/
	Network *string `json:"network"`
}

func (p *EsxNetworkResourceConfig) MarshalJSON() ([]byte, error) {
	type EsxNetworkResourceConfigProxy EsxNetworkResourceConfig
	return json.Marshal(struct {
		*EsxNetworkResourceConfigProxy
		Network *string `json:"network,omitempty"`
	}{
		EsxNetworkResourceConfigProxy: (*EsxNetworkResourceConfigProxy)(p),
		Network:                       p.Network,
	})
}

func NewEsxNetworkResourceConfig() *EsxNetworkResourceConfig {
	p := new(EsxNetworkResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.EsxNetworkResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.EsxNetworkResourceConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
ESX platform specific config.
*/
type EsxResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NetworkResource *EsxNetworkResourceConfig `json:"networkResource"`

	VmResource *VmResourceConfig `json:"vmResource"`
}

func (p *EsxResourceConfig) MarshalJSON() ([]byte, error) {
	type EsxResourceConfigProxy EsxResourceConfig
	return json.Marshal(struct {
		*EsxResourceConfigProxy
		NetworkResource *EsxNetworkResourceConfig `json:"networkResource,omitempty"`
		VmResource      *VmResourceConfig         `json:"vmResource,omitempty"`
	}{
		EsxResourceConfigProxy: (*EsxResourceConfigProxy)(p),
		NetworkResource:        p.NetworkResource,
		VmResource:             p.VmResource,
	})
}

func NewEsxResourceConfig() *EsxResourceConfig {
	p := new(EsxResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.EsxResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.EsxResourceConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/applications Get operation
*/
type GetAllAppApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAllAppApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetAllAppApiResponse() *GetAllAppApiResponse {
	p := new(GetAllAppApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.GetAllAppApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.GetAllAppApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAllAppApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAllAppApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAllAppApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/services Get operation
*/
type GetAllServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAllServiceApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetAllServiceApiResponse() *GetAllServiceApiResponse {
	p := new(GetAllServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.GetAllServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.GetAllServiceApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAllServiceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAllServiceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAllServiceApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/tasks Get operation
*/
type GetAllTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAllTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetAllTaskApiResponse() *GetAllTaskApiResponse {
	p := new(GetAllTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.GetAllTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.GetAllTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAllTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAllTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAllTaskApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/applications/{appUUID} Get operation
*/
type GetAppApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetAppApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetAppApiResponse() *GetAppApiResponse {
	p := new(GetAppApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.GetAppApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.GetAppApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetAppApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetAppApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetAppApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/history Get operation
*/
type GetHistoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetHistoryApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetHistoryApiResponse() *GetHistoryApiResponse {
	p := new(GetHistoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.GetHistoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.GetHistoryApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetHistoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetHistoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetHistoryApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/services/{serviceUUID} Get operation
*/
type GetServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetServiceApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetServiceApiResponse() *GetServiceApiResponse {
	p := new(GetServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.GetServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.GetServiceApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetServiceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetServiceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetServiceApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/tasks/{taskUUID} Get operation
*/
type GetTaskApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetTaskApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewGetTaskApiResponse() *GetTaskApiResponse {
	p := new(GetTaskApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.GetTaskApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.GetTaskApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetTaskApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetTaskApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetTaskApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

type Health struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	HealthComponents []HealthComponent `json:"healthComponents,omitempty"`

	Status *Status `json:"status,omitempty"`
}

func NewHealth() *Health {
	p := new(Health)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.Health"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.Health"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/health Get operation
*/
type HealthApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfHealthApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewHealthApiResponse() *HealthApiResponse {
	p := new(HealthApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.HealthApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.HealthApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *HealthApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *HealthApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfHealthApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
*
Health of the Component.
*/
type HealthComponent struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Details *ComponentDetails `json:"details,omitempty"`
	/**
	  Name of the component.
	*/
	Name *string `json:"name,omitempty"`

	Status *Status `json:"status,omitempty"`
}

func NewHealthComponent() *HealthComponent {
	p := new(HealthComponent)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.HealthComponent"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.HealthComponent"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type History struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	HistoryList []AppOrServiceHistory `json:"historyList,omitempty"`
}

func NewHistory() *History {
	p := new(History)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.History"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.History"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Install struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Name(or UUID) of the cluster. Optional. If not specified, the default cluster(CMSP) will be used.
	*/
	ClusterID *string `json:"clusterID,omitempty"`

	CustomValues []CustomValue `json:"customValues,omitempty"`

	Name *string `json:"name"`

	Version *string `json:"version"`
}

func (p *Install) MarshalJSON() ([]byte, error) {
	type InstallProxy Install
	return json.Marshal(struct {
		*InstallProxy
		Name    *string `json:"name,omitempty"`
		Version *string `json:"version,omitempty"`
	}{
		InstallProxy: (*InstallProxy)(p),
		Name:         p.Name,
		Version:      p.Version,
	})
}

func NewInstall() *Install {
	p := new(Install)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.Install"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.Install"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Ip ranges from start to end.
*/
type IpRange struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	EndIp *import1.IPv4Address `json:"endIp"`

	StartIp *import1.IPv4Address `json:"startIp"`
}

func (p *IpRange) MarshalJSON() ([]byte, error) {
	type IpRangeProxy IpRange
	return json.Marshal(struct {
		*IpRangeProxy
		EndIp   *import1.IPv4Address `json:"endIp,omitempty"`
		StartIp *import1.IPv4Address `json:"startIp,omitempty"`
	}{
		IpRangeProxy: (*IpRangeProxy)(p),
		EndIp:        p.EndIp,
		StartIp:      p.StartIp,
	})
}

func NewIpRange() *IpRange {
	p := new(IpRange)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.IpRange"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.IpRange"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Load balancer representing id and network config for it.
*/
type LoadBalancer struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Identifier for each load balancer. Can be as simple as an incrementing counter.
	*/
	Id *string `json:"id,omitempty"`

	NetworkConfig *LoadBalancerNetworkConfig `json:"networkConfig,omitempty"`
}

func NewLoadBalancer() *LoadBalancer {
	p := new(LoadBalancer)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.LoadBalancer"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.LoadBalancer"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
An access interface is a user's notion of a named load balancer network. It abstracts
out the network details from developers and service deployers (dev-ops).
The name should be something that defines the positioning of the load balancer
in the data center. e.g. "corp", "wan", "dc", "it", "it-hr". Currenlty in a given
subdomain - the interface should be unique - which greatly simplifies what
the service deployer has to specifiy to expose the service.
A single load balancer defines one or more interfaces's at load balancer creation times
*/
type LoadBalancerAccessInterface struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  load balancer CIDR (Classless Inter-Domain Routing) address.
	*/
	Cidr *string `json:"cidr,omitempty"`
	/**
	  load balancer access interface ips.
	*/
	Ips []import1.IPv4Address `json:"ips,omitempty"`
	/**
	  load balancer frontent network id.
	*/
	Network *string `json:"network"`
	/**
	  load balancer frontent subnet.
	*/
	Subnet *string `json:"subnet"`
}

func (p *LoadBalancerAccessInterface) MarshalJSON() ([]byte, error) {
	type LoadBalancerAccessInterfaceProxy LoadBalancerAccessInterface
	return json.Marshal(struct {
		*LoadBalancerAccessInterfaceProxy
		Network *string `json:"network,omitempty"`
		Subnet  *string `json:"subnet,omitempty"`
	}{
		LoadBalancerAccessInterfaceProxy: (*LoadBalancerAccessInterfaceProxy)(p),
		Network:                          p.Network,
		Subnet:                           p.Subnet,
	})
}

func NewLoadBalancerAccessInterface() *LoadBalancerAccessInterface {
	p := new(LoadBalancerAccessInterface)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.LoadBalancerAccessInterface"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.LoadBalancerAccessInterface"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Load balancer configuration.
*/
type LoadBalancerConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ConfigType *LoadBalancerConfigType `json:"configType,omitempty"`

	Instances []LoadBalancer `json:"instances"`
	/**
	  UUID of the MSP wanting to use this LB.
	*/
	MspUuid *string `json:"mspUuid,omitempty"`
	/**
	  Name of the Load Balancer config like envoy-1, etc. This needs to be unique per deployment.
	*/
	Name *string `json:"name"`

	ResourceConfig *LoadBalancerResourceConfig `json:"resourceConfig,omitempty"`

	Type *LoadBalancerType `json:"type,omitempty"`
}

func (p *LoadBalancerConfig) MarshalJSON() ([]byte, error) {
	type LoadBalancerConfigProxy LoadBalancerConfig
	return json.Marshal(struct {
		*LoadBalancerConfigProxy
		Instances []LoadBalancer `json:"instances,omitempty"`
		Name      *string        `json:"name,omitempty"`
	}{
		LoadBalancerConfigProxy: (*LoadBalancerConfigProxy)(p),
		Instances:               p.Instances,
		Name:                    p.Name,
	})
}

func NewLoadBalancerConfig() *LoadBalancerConfig {
	p := new(LoadBalancerConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.LoadBalancerConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.LoadBalancerConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Defines the type of configuration, static, dynamic etc.
*/
type LoadBalancerConfigType int

const (
	LOADBALANCERCONFIGTYPE_UNKNOWN  LoadBalancerConfigType = 0
	LOADBALANCERCONFIGTYPE_REDACTED LoadBalancerConfigType = 1
	LOADBALANCERCONFIGTYPE_STATIC   LoadBalancerConfigType = 2
	LOADBALANCERCONFIGTYPE_DYNAMIC  LoadBalancerConfigType = 3
)

// returns the name of the enum given an ordinal number
func (e *LoadBalancerConfigType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STATIC",
		"DYNAMIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *LoadBalancerConfigType) index(name string) LoadBalancerConfigType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"STATIC",
		"DYNAMIC",
	}
	for idx := range names {
		if names[idx] == name {
			return LoadBalancerConfigType(idx)
		}
	}
	return LOADBALANCERCONFIGTYPE_UNKNOWN
}

func (e *LoadBalancerConfigType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LoadBalancerConfigType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LoadBalancerConfigType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LoadBalancerConfigType) Ref() *LoadBalancerConfigType {
	return &e
}

/*
*
Network configuration of load balancer.
*/
type LoadBalancerNetworkConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AccessInterfaces []LoadBalancerAccessInterface `json:"accessInterfaces,omitempty"`

	AhvNetworkConfig *AhvNetworkResourceConfig `json:"ahvNetworkConfig,omitempty"`
	/**
	  Client for load balancer.
	*/
	Client *string `json:"client,omitempty"`

	EsxNetworkConfig *EsxNetworkResourceConfig `json:"esxNetworkConfig,omitempty"`
}

func NewLoadBalancerNetworkConfig() *LoadBalancerNetworkConfig {
	p := new(LoadBalancerNetworkConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.LoadBalancerNetworkConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.LoadBalancerNetworkConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
load balancer vm resource configuration.
*/
type LoadBalancerResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Number of VCPUs to be allocated to the vm.
	*/
	Cpu *int `json:"cpu,omitempty"`
	/**
	  Memory requirements of VM.
	*/
	DiskMib *int `json:"diskMib,omitempty"`
	/**
	  Gold image name/UUID.
	*/
	GoldImageRef *string `json:"goldImageRef,omitempty"`
	/**
	  OS version of the gold image.
	*/
	GoldImageVersion *string `json:"goldImageVersion,omitempty"`
	/**
	  Labels for this resource configuration. Its a key value pair of strings.
	*/
	Labels []import1.KVStringPair `json:"labels,omitempty"`
	/**
	  Memory requirements of VM.
	*/
	MemoryMib *int `json:"memoryMib,omitempty"`
}

func NewLoadBalancerResourceConfig() *LoadBalancerResourceConfig {
	p := new(LoadBalancerResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.LoadBalancerResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.LoadBalancerResourceConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Type of load balancer such as VM, hardware load balancer, container, etc.
*/
type LoadBalancerType int

const (
	LOADBALANCERTYPE_UNKNOWN           LoadBalancerType = 0
	LOADBALANCERTYPE_REDACTED          LoadBalancerType = 1
	LOADBALANCERTYPE_VM                LoadBalancerType = 2
	LOADBALANCERTYPE_CONTAINER_MASTERS LoadBalancerType = 3
)

// returns the name of the enum given an ordinal number
func (e *LoadBalancerType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"CONTAINER_MASTERS",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *LoadBalancerType) index(name string) LoadBalancerType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"VM",
		"CONTAINER_MASTERS",
	}
	for idx := range names {
		if names[idx] == name {
			return LoadBalancerType(idx)
		}
	}
	return LOADBALANCERTYPE_UNKNOWN
}

func (e *LoadBalancerType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for LoadBalancerType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *LoadBalancerType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e LoadBalancerType) Ref() *LoadBalancerType {
	return &e
}

/*
*
Nic type to be configured on the msp vm. Accepted values are NORMAL_NIC(default)/DIRECT_NIC.
NORMAL_NIC indicates nics are inserted in br0.local, where as DIRECT_NIC indicates
nics are inserted in br0 on AHV hypervisor.
*/
type NicType int

const (
	NICTYPE_UNKNOWN    NicType = 0
	NICTYPE_REDACTED   NicType = 1
	NICTYPE_NORMAL_NIC NicType = 2
	NICTYPE_DIRECT_NIC NicType = 3
)

// returns the name of the enum given an ordinal number
func (e *NicType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL_NIC",
		"DIRECT_NIC",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *NicType) index(name string) NicType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"NORMAL_NIC",
		"DIRECT_NIC",
	}
	for idx := range names {
		if names[idx] == name {
			return NicType(idx)
		}
	}
	return NICTYPE_UNKNOWN
}

func (e *NicType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for NicType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *NicType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e NicType) Ref() *NicType {
	return &e
}

/*
*
Persistent volume config to be attached to worker VM.
*/
type PersistentVolumeConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Mount path of persistent volume on worker VM.
	*/
	MountPath *string `json:"mountPath"`
	/**
	  Size of the Volume group in MB.
	*/
	SizeMib *int `json:"sizeMib"`
	/**
	  Name of the storage config.
	*/
	StorageConfigName *string `json:"storageConfigName"`
}

func (p *PersistentVolumeConfig) MarshalJSON() ([]byte, error) {
	type PersistentVolumeConfigProxy PersistentVolumeConfig
	return json.Marshal(struct {
		*PersistentVolumeConfigProxy
		MountPath         *string `json:"mountPath,omitempty"`
		SizeMib           *int    `json:"sizeMib,omitempty"`
		StorageConfigName *string `json:"storageConfigName,omitempty"`
	}{
		PersistentVolumeConfigProxy: (*PersistentVolumeConfigProxy)(p),
		MountPath:                   p.MountPath,
		SizeMib:                     p.SizeMib,
		StorageConfigName:           p.StorageConfigName,
	})
}

func NewPersistentVolumeConfig() *PersistentVolumeConfig {
	p := new(PersistentVolumeConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.PersistentVolumeConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.PersistentVolumeConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/applications Post operation
*/
type PostAppApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfPostAppApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewPostAppApiResponse() *PostAppApiResponse {
	p := new(PostAppApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.PostAppApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.PostAppApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PostAppApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *PostAppApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfPostAppApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

type PostAppResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	AppUUID *string `json:"appUUID,omitempty"`

	TaskUuid *string `json:"taskUuid,omitempty"`
}

func NewPostAppResponse() *PostAppResponse {
	p := new(PostAppResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.PostAppResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.PostAppResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/clusters Post operation
*/
type PostClusterApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfPostClusterApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewPostClusterApiResponse() *PostClusterApiResponse {
	p := new(PostClusterApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.PostClusterApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.PostClusterApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PostClusterApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *PostClusterApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfPostClusterApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
*
REST Response for all response codes in api path /msp/v1.0.a1/svcmgr/services Post operation
*/
type PostServiceApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfPostServiceApiResponseData `json:"data,omitempty"`

	Metadata *import3.ApiResponseMetadata `json:"metadata,omitempty"`
}

func NewPostServiceApiResponse() *PostServiceApiResponse {
	p := new(PostServiceApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.PostServiceApiResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.PostServiceApiResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *PostServiceApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *PostServiceApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfPostServiceApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

type PostServiceResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ServiceUUID *string `json:"serviceUUID,omitempty"`

	TaskUuid *string `json:"taskUuid,omitempty"`
}

func NewPostServiceResponse() *PostServiceResponse {
	p := new(PostServiceResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.PostServiceResponse"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.PostServiceResponse"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Service struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *string `json:"action,omitempty"`

	Apptype *string `json:"apptype,omitempty"`

	CreatedTimestamp *time.Time `json:"createdTimestamp,omitempty"`

	LastUpdatedTimestamp *time.Time `json:"lastUpdatedTimestamp,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	Version *string `json:"version,omitempty"`
}

func NewService() *Service {
	p := new(Service)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.Service"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.Service"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Services struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	ServiceList []Service `json:"serviceList,omitempty"`
}

func NewServices() *Services {
	p := new(Services)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.Services"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.Services"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Status of the service.
*/
type Status int

const (
	STATUS_UP      Status = 0
	STATUS_DOWN    Status = 1
	STATUS_UNKNOWN Status = 2
)

// returns the name of the enum given an ordinal number
func (e *Status) name(index int) string {
	names := [...]string{
		"UP",
		"DOWN",
		"$UNKNOWN",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// returns the enum type given a string value
func (e *Status) index(name string) Status {
	names := [...]string{
		"UP",
		"DOWN",
		"$UNKNOWN",
	}
	for idx := range names {
		if names[idx] == name {
			return Status(idx)
		}
	}
	return STATUS_UNKNOWN
}

func (e *Status) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for Status:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *Status) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e Status) Ref() *Status {
	return &e
}

type Task struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	TaskUuid *string `json:"taskUuid,omitempty"`
}

func NewTask() *Task {
	p := new(Task)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.Task"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.Task"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type TaskComponents struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	Action *string `json:"action,omitempty"`

	Apptype *string `json:"apptype,omitempty"`

	ChildTasks []TaskComponents `json:"childTasks,omitempty"`

	CompletedTimestamp *time.Time `json:"completedTimestamp,omitempty"`

	CreatedTimestamp *time.Time `json:"createdTimestamp,omitempty"`

	LastUpdatedTimestamp *time.Time `json:"lastUpdatedTimestamp,omitempty"`

	Message *string `json:"message,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	TaskUuid *string `json:"taskUuid,omitempty"`

	Uuid *string `json:"uuid,omitempty"`

	Version *string `json:"version,omitempty"`
}

func NewTaskComponents() *TaskComponents {
	p := new(TaskComponents)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.TaskComponents"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.TaskComponents"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type Tasks struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	TaskList []TaskComponents `json:"taskList,omitempty"`
}

func NewTasks() *Tasks {
	p := new(Tasks)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.Tasks"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.Tasks"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Vcenter location
*/
type VcenterLocation struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Cluster name in Vcenter.
	*/
	ClusterName *string `json:"clusterName,omitempty"`
	/**
	  Data center name in Vcenter.
	*/
	Datacenter *string `json:"datacenter"`

	VcenterIp *import1.IPv4Address `json:"vcenterIp,omitempty"`
	/**
	  Vcenter Port.
	*/
	VcenterPort *int `json:"vcenterPort,omitempty"`
}

func (p *VcenterLocation) MarshalJSON() ([]byte, error) {
	type VcenterLocationProxy VcenterLocation
	return json.Marshal(struct {
		*VcenterLocationProxy
		Datacenter *string `json:"datacenter,omitempty"`
	}{
		VcenterLocationProxy: (*VcenterLocationProxy)(p),
		Datacenter:           p.Datacenter,
	})
}

func NewVcenterLocation() *VcenterLocation {
	p := new(VcenterLocation)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.VcenterLocation"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.VcenterLocation"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Network resources for Vm.
*/
type VmNetworkResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Cluster Id of the Prism Element which hosts the Kubernetes cluster.
	*/
	ClusterId *string `json:"clusterId,omitempty"`
	/**
	  Network id to be used for this vm config.
	*/
	Network *string `json:"network"`
}

func (p *VmNetworkResourceConfig) MarshalJSON() ([]byte, error) {
	type VmNetworkResourceConfigProxy VmNetworkResourceConfig
	return json.Marshal(struct {
		*VmNetworkResourceConfigProxy
		Network *string `json:"network,omitempty"`
	}{
		VmNetworkResourceConfigProxy: (*VmNetworkResourceConfigProxy)(p),
		Network:                      p.Network,
	})
}

func NewVmNetworkResourceConfig() *VmNetworkResourceConfig {
	p := new(VmNetworkResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.VmNetworkResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.VmNetworkResourceConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Resource configuration for the vm.
*/
type VmResourceConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/**
	  Number of VCPUs to be allocated to the vm.
	*/
	Cpu *int `json:"cpu,omitempty"`
	/**
	  Memory requirements of VM.
	*/
	DiskMib *int `json:"diskMib,omitempty"`
	/**
	  Memory requirements of VM.
	*/
	MemoryMib *int `json:"memoryMib,omitempty"`
}

func NewVmResourceConfig() *VmResourceConfig {
	p := new(VmResourceConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.VmResourceConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.VmResourceConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
*
Configuration for the worker nodes.
*/
type WorkerConfig struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`

	NumInstances *int `json:"numInstances"`
	/**
	  This describes the name of the resource to be used.
	Actual spec is got from cluster resource configs field using this name.
	*/
	ResourceConfigName *string `json:"resourceConfigName"`
}

func (p *WorkerConfig) MarshalJSON() ([]byte, error) {
	type WorkerConfigProxy WorkerConfig
	return json.Marshal(struct {
		*WorkerConfigProxy
		NumInstances       *int    `json:"numInstances,omitempty"`
		ResourceConfigName *string `json:"resourceConfigName,omitempty"`
	}{
		WorkerConfigProxy:  (*WorkerConfigProxy)(p),
		NumInstances:       p.NumInstances,
		ResourceConfigName: p.ResourceConfigName,
	})
}

func NewWorkerConfig() *WorkerConfig {
	p := new(WorkerConfig)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "msp.v1.svcmgr.WorkerConfig"
	p.Reserved_ = map[string]interface{}{"$fqObjectType": "msp.v1.r0.a1.svcmgr.WorkerConfig"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type OneOfGetAllServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Services              `json:"-"`
}

func NewOneOfGetAllServiceApiResponseData() *OneOfGetAllServiceApiResponseData {
	p := new(OneOfGetAllServiceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAllServiceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAllServiceApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Services:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Services)
		}
		*p.oneOfType0 = v.(Services)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetAllServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetAllServiceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(Services)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.Services" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Services)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAllServiceApiResponseData"))
}

func (p *OneOfGetAllServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetAllServiceApiResponseData")
}

type OneOfPostServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *PostServiceResponse   `json:"-"`
}

func NewOneOfPostServiceApiResponseData() *OneOfPostServiceApiResponseData {
	p := new(OneOfPostServiceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPostServiceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPostServiceApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case PostServiceResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(PostServiceResponse)
		}
		*p.oneOfType0 = v.(PostServiceResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfPostServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfPostServiceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(PostServiceResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.PostServiceResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(PostServiceResponse)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPostServiceApiResponseData"))
}

func (p *OneOfPostServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfPostServiceApiResponseData")
}

type OneOfCustomValueItemValue struct {
	Discriminator *string  `json:"-"`
	ObjectType_   *string  `json:"-"`
	oneOfType1    *string  `json:"-"`
	oneOfType0    []string `json:"-"`
}

func NewOneOfCustomValueItemValue() *OneOfCustomValueItemValue {
	p := new(OneOfCustomValueItemValue)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCustomValueItemValue) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCustomValueItemValue is nil"))
	}
	switch v.(type) {
	case string:
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(string)
		}
		*p.oneOfType1 = v.(string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "String"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "String"
	case []string:
		p.oneOfType0 = v.([]string)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<String>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCustomValueItemValue) GetValue() interface{} {
	if "String" == *p.Discriminator {
		return *p.oneOfType1
	}
	if "List<String>" == *p.Discriminator {
		return p.oneOfType0
	}
	return nil
}

func (p *OneOfCustomValueItemValue) UnmarshalJSON(b []byte) error {
	vOneOfType1 := new(string)
	if err := json.Unmarshal(b, vOneOfType1); err == nil {
		if nil == p.oneOfType1 {
			p.oneOfType1 = new(string)
		}
		*p.oneOfType1 = *vOneOfType1
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "String"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "String"
		return nil
	}
	vOneOfType0 := new([]string)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		p.oneOfType0 = *vOneOfType0
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<String>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<String>"
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCustomValueItemValue"))
}

func (p *OneOfCustomValueItemValue) MarshalJSON() ([]byte, error) {
	if "String" == *p.Discriminator {
		return json.Marshal(p.oneOfType1)
	}
	if "List<String>" == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfCustomValueItemValue")
}

type OneOfGetAllAppApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Applications          `json:"-"`
}

func NewOneOfGetAllAppApiResponseData() *OneOfGetAllAppApiResponseData {
	p := new(OneOfGetAllAppApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAllAppApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAllAppApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Applications:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Applications)
		}
		*p.oneOfType0 = v.(Applications)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetAllAppApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetAllAppApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(Applications)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.Applications" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Applications)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAllAppApiResponseData"))
}

func (p *OneOfGetAllAppApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetAllAppApiResponseData")
}

type OneOfGetAllTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Tasks                 `json:"-"`
}

func NewOneOfGetAllTaskApiResponseData() *OneOfGetAllTaskApiResponseData {
	p := new(OneOfGetAllTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAllTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAllTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Tasks:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Tasks)
		}
		*p.oneOfType0 = v.(Tasks)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetAllTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetAllTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(Tasks)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.Tasks" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Tasks)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAllTaskApiResponseData"))
}

func (p *OneOfGetAllTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetAllTaskApiResponseData")
}

type OneOfGetTaskApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *TaskComponents        `json:"-"`
}

func NewOneOfGetTaskApiResponseData() *OneOfGetTaskApiResponseData {
	p := new(OneOfGetTaskApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetTaskApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetTaskApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case TaskComponents:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(TaskComponents)
		}
		*p.oneOfType0 = v.(TaskComponents)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetTaskApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetTaskApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(TaskComponents)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.TaskComponents" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(TaskComponents)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetTaskApiResponseData"))
}

func (p *OneOfGetTaskApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetTaskApiResponseData")
}

type OneOfGetHistoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *History               `json:"-"`
}

func NewOneOfGetHistoryApiResponseData() *OneOfGetHistoryApiResponseData {
	p := new(OneOfGetHistoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetHistoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetHistoryApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case History:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(History)
		}
		*p.oneOfType0 = v.(History)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetHistoryApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetHistoryApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(History)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.History" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(History)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetHistoryApiResponseData"))
}

func (p *OneOfGetHistoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetHistoryApiResponseData")
}

type OneOfGetServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Service               `json:"-"`
}

func NewOneOfGetServiceApiResponseData() *OneOfGetServiceApiResponseData {
	p := new(OneOfGetServiceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetServiceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetServiceApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Service:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Service)
		}
		*p.oneOfType0 = v.(Service)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetServiceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(Service)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.Service" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Service)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetServiceApiResponseData"))
}

func (p *OneOfGetServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetServiceApiResponseData")
}

type OneOfDeleteAppApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Task                  `json:"-"`
}

func NewOneOfDeleteAppApiResponseData() *OneOfDeleteAppApiResponseData {
	p := new(OneOfDeleteAppApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteAppApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteAppApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Task:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Task)
		}
		*p.oneOfType0 = v.(Task)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteAppApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDeleteAppApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(Task)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.Task" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Task)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteAppApiResponseData"))
}

func (p *OneOfDeleteAppApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteAppApiResponseData")
}

type OneOfDeleteServiceApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Task                  `json:"-"`
}

func NewOneOfDeleteServiceApiResponseData() *OneOfDeleteServiceApiResponseData {
	p := new(OneOfDeleteServiceApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteServiceApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteServiceApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Task:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Task)
		}
		*p.oneOfType0 = v.(Task)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteServiceApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfDeleteServiceApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(Task)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.Task" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Task)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteServiceApiResponseData"))
}

func (p *OneOfDeleteServiceApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteServiceApiResponseData")
}

type OneOfPostAppApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *PostAppResponse       `json:"-"`
}

func NewOneOfPostAppApiResponseData() *OneOfPostAppApiResponseData {
	p := new(OneOfPostAppApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPostAppApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPostAppApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case PostAppResponse:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(PostAppResponse)
		}
		*p.oneOfType0 = v.(PostAppResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfPostAppApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfPostAppApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(PostAppResponse)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.PostAppResponse" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(PostAppResponse)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPostAppApiResponseData"))
}

func (p *OneOfPostAppApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfPostAppApiResponseData")
}

type OneOfGetAppApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *Application           `json:"-"`
}

func NewOneOfGetAppApiResponseData() *OneOfGetAppApiResponseData {
	p := new(OneOfGetAppApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetAppApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetAppApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case Application:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Application)
		}
		*p.oneOfType0 = v.(Application)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetAppApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfGetAppApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(Application)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.Application" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Application)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetAppApiResponseData"))
}

func (p *OneOfGetAppApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfGetAppApiResponseData")
}

type OneOfHealthApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType0    *Health                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
}

func NewOneOfHealthApiResponseData() *OneOfHealthApiResponseData {
	p := new(OneOfHealthApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfHealthApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfHealthApiResponseData is nil"))
	}
	switch v.(type) {
	case Health:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(Health)
		}
		*p.oneOfType0 = v.(Health)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfHealthApiResponseData) GetValue() interface{} {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	return nil
}

func (p *OneOfHealthApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType0 := new(Health)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "msp.v1.svcmgr.Health" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(Health)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfHealthApiResponseData"))
}

func (p *OneOfHealthApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	return nil, errors.New("No value to marshal for OneOfHealthApiResponseData")
}

type OneOfPostClusterApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import2.ErrorResponse `json:"-"`
	oneOfType0    *import4.TaskReference `json:"-"`
}

func NewOneOfPostClusterApiResponseData() *OneOfPostClusterApiResponseData {
	p := new(OneOfPostClusterApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfPostClusterApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfPostClusterApiResponseData is nil"))
	}
	switch v.(type) {
	case import2.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import2.ErrorResponse)
		}
		*p.oneOfType400 = v.(import2.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import4.TaskReference:
		if nil == p.oneOfType0 {
			p.oneOfType0 = new(import4.TaskReference)
		}
		*p.oneOfType0 = v.(import4.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType0.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType0.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfPostClusterApiResponseData) GetValue() interface{} {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return *p.oneOfType0
	}
	return nil
}

func (p *OneOfPostClusterApiResponseData) UnmarshalJSON(b []byte) error {
	vOneOfType400 := new(import2.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if "msp.v1.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import2.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType0 := new(import4.TaskReference)
	if err := json.Unmarshal(b, vOneOfType0); err == nil {
		if "prism.v4.config.TaskReference" == *vOneOfType0.ObjectType_ {
			if nil == p.oneOfType0 {
				p.oneOfType0 = new(import4.TaskReference)
			}
			*p.oneOfType0 = *vOneOfType0
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType0.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType0.ObjectType_
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfPostClusterApiResponseData"))
}

func (p *OneOfPostClusterApiResponseData) MarshalJSON() ([]byte, error) {
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType0 != nil && *p.oneOfType0.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType0)
	}
	return nil, errors.New("No value to marshal for OneOfPostClusterApiResponseData")
}
