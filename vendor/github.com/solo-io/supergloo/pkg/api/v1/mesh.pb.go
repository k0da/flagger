// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/supergloo/api/v1/mesh.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//Meshes represent a currently registered service mesh.
type Mesh struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by supergloo during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	// Types that are valid to be assigned to MeshType:
	//	*Mesh_Istio
	//	*Mesh_AwsAppMesh
	//	*Mesh_LinkerdMesh
	MeshType isMesh_MeshType `protobuf_oneof:"mesh_type"`
	// mtls config specifies configuration options for enabling mutual
	// tls between pods in this mesh
	MtlsConfig *MtlsConfig `protobuf:"bytes,2,opt,name=mtls_config,json=mtlsConfig,proto3" json:"mtls_config,omitempty"`
	// configuration for propagating stats and metrics from
	// mesh controllers and sidecars to a centralized datastore
	// such as prometheus
	MonitoringConfig *MonitoringConfig `protobuf:"bytes,3,opt,name=monitoring_config,json=monitoringConfig,proto3" json:"monitoring_config,omitempty"`
	// object which represents the data mesh discovery finds about a given mesh
	DiscoveryMetadata    *DiscoveryMetadata `protobuf:"bytes,6,opt,name=discovery_metadata,json=discoveryMetadata,proto3" json:"discovery_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Mesh) Reset()         { *m = Mesh{} }
func (m *Mesh) String() string { return proto.CompactTextString(m) }
func (*Mesh) ProtoMessage()    {}
func (*Mesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{0}
}
func (m *Mesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mesh.Unmarshal(m, b)
}
func (m *Mesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mesh.Marshal(b, m, deterministic)
}
func (m *Mesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesh.Merge(m, src)
}
func (m *Mesh) XXX_Size() int {
	return xxx_messageInfo_Mesh.Size(m)
}
func (m *Mesh) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesh.DiscardUnknown(m)
}

var xxx_messageInfo_Mesh proto.InternalMessageInfo

type isMesh_MeshType interface {
	isMesh_MeshType()
	Equal(interface{}) bool
}

type Mesh_Istio struct {
	Istio *IstioMesh `protobuf:"bytes,1,opt,name=istio,proto3,oneof"`
}
type Mesh_AwsAppMesh struct {
	AwsAppMesh *AwsAppMesh `protobuf:"bytes,4,opt,name=aws_app_mesh,json=awsAppMesh,proto3,oneof"`
}
type Mesh_LinkerdMesh struct {
	LinkerdMesh *LinkerdMesh `protobuf:"bytes,5,opt,name=linkerd_mesh,json=linkerdMesh,proto3,oneof"`
}

func (*Mesh_Istio) isMesh_MeshType()       {}
func (*Mesh_AwsAppMesh) isMesh_MeshType()  {}
func (*Mesh_LinkerdMesh) isMesh_MeshType() {}

func (m *Mesh) GetMeshType() isMesh_MeshType {
	if m != nil {
		return m.MeshType
	}
	return nil
}

func (m *Mesh) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Mesh) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Mesh) GetIstio() *IstioMesh {
	if x, ok := m.GetMeshType().(*Mesh_Istio); ok {
		return x.Istio
	}
	return nil
}

func (m *Mesh) GetAwsAppMesh() *AwsAppMesh {
	if x, ok := m.GetMeshType().(*Mesh_AwsAppMesh); ok {
		return x.AwsAppMesh
	}
	return nil
}

func (m *Mesh) GetLinkerdMesh() *LinkerdMesh {
	if x, ok := m.GetMeshType().(*Mesh_LinkerdMesh); ok {
		return x.LinkerdMesh
	}
	return nil
}

func (m *Mesh) GetMtlsConfig() *MtlsConfig {
	if m != nil {
		return m.MtlsConfig
	}
	return nil
}

func (m *Mesh) GetMonitoringConfig() *MonitoringConfig {
	if m != nil {
		return m.MonitoringConfig
	}
	return nil
}

func (m *Mesh) GetDiscoveryMetadata() *DiscoveryMetadata {
	if m != nil {
		return m.DiscoveryMetadata
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Mesh) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Mesh_OneofMarshaler, _Mesh_OneofUnmarshaler, _Mesh_OneofSizer, []interface{}{
		(*Mesh_Istio)(nil),
		(*Mesh_AwsAppMesh)(nil),
		(*Mesh_LinkerdMesh)(nil),
	}
}

func _Mesh_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Mesh)
	// mesh_type
	switch x := m.MeshType.(type) {
	case *Mesh_Istio:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Istio); err != nil {
			return err
		}
	case *Mesh_AwsAppMesh:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AwsAppMesh); err != nil {
			return err
		}
	case *Mesh_LinkerdMesh:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LinkerdMesh); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Mesh.MeshType has unexpected type %T", x)
	}
	return nil
}

func _Mesh_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Mesh)
	switch tag {
	case 1: // mesh_type.istio
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(IstioMesh)
		err := b.DecodeMessage(msg)
		m.MeshType = &Mesh_Istio{msg}
		return true, err
	case 4: // mesh_type.aws_app_mesh
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AwsAppMesh)
		err := b.DecodeMessage(msg)
		m.MeshType = &Mesh_AwsAppMesh{msg}
		return true, err
	case 5: // mesh_type.linkerd_mesh
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LinkerdMesh)
		err := b.DecodeMessage(msg)
		m.MeshType = &Mesh_LinkerdMesh{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Mesh_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Mesh)
	// mesh_type
	switch x := m.MeshType.(type) {
	case *Mesh_Istio:
		s := proto.Size(x.Istio)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mesh_AwsAppMesh:
		s := proto.Size(x.AwsAppMesh)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Mesh_LinkerdMesh:
		s := proto.Size(x.LinkerdMesh)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Generic discovery data shared between different meshes
type DiscoveryMetadata struct {
	// list of namespaces which we know are being injected by a given mesh
	InjectedNamespaceLabel string `protobuf:"bytes,1,opt,name=injected_namespace_label,json=injectedNamespaceLabel,proto3" json:"injected_namespace_label,omitempty"`
	// Whether or not auto-injection is enabled for a given mesh
	EnableAutoInject bool `protobuf:"varint,2,opt,name=enable_auto_inject,json=enableAutoInject,proto3" json:"enable_auto_inject,omitempty"`
	// version of the mesh which is installed
	MeshVersion string `protobuf:"bytes,3,opt,name=mesh_version,json=meshVersion,proto3" json:"mesh_version,omitempty"`
	// namespace which the mesh is installed into
	InstallationNamespace string `protobuf:"bytes,4,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// discovered mtls config of the given mesh
	MtlsConfig           *MtlsConfig `protobuf:"bytes,5,opt,name=mtls_config,json=mtlsConfig,proto3" json:"mtls_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DiscoveryMetadata) Reset()         { *m = DiscoveryMetadata{} }
func (m *DiscoveryMetadata) String() string { return proto.CompactTextString(m) }
func (*DiscoveryMetadata) ProtoMessage()    {}
func (*DiscoveryMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{1}
}
func (m *DiscoveryMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryMetadata.Unmarshal(m, b)
}
func (m *DiscoveryMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryMetadata.Marshal(b, m, deterministic)
}
func (m *DiscoveryMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryMetadata.Merge(m, src)
}
func (m *DiscoveryMetadata) XXX_Size() int {
	return xxx_messageInfo_DiscoveryMetadata.Size(m)
}
func (m *DiscoveryMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryMetadata proto.InternalMessageInfo

func (m *DiscoveryMetadata) GetInjectedNamespaceLabel() string {
	if m != nil {
		return m.InjectedNamespaceLabel
	}
	return ""
}

func (m *DiscoveryMetadata) GetEnableAutoInject() bool {
	if m != nil {
		return m.EnableAutoInject
	}
	return false
}

func (m *DiscoveryMetadata) GetMeshVersion() string {
	if m != nil {
		return m.MeshVersion
	}
	return ""
}

func (m *DiscoveryMetadata) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *DiscoveryMetadata) GetMtlsConfig() *MtlsConfig {
	if m != nil {
		return m.MtlsConfig
	}
	return nil
}

// Mesh object representing an installed Istio control plane
type IstioMesh struct {
	// where the istio control plane has been installed
	InstallationNamespace string `protobuf:"bytes,1,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// version of istio which has been installed
	IstioVersion         string   `protobuf:"bytes,2,opt,name=istio_version,json=istioVersion,proto3" json:"istio_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioMesh) Reset()         { *m = IstioMesh{} }
func (m *IstioMesh) String() string { return proto.CompactTextString(m) }
func (*IstioMesh) ProtoMessage()    {}
func (*IstioMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{2}
}
func (m *IstioMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioMesh.Unmarshal(m, b)
}
func (m *IstioMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioMesh.Marshal(b, m, deterministic)
}
func (m *IstioMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioMesh.Merge(m, src)
}
func (m *IstioMesh) XXX_Size() int {
	return xxx_messageInfo_IstioMesh.Size(m)
}
func (m *IstioMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioMesh.DiscardUnknown(m)
}

var xxx_messageInfo_IstioMesh proto.InternalMessageInfo

func (m *IstioMesh) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *IstioMesh) GetIstioVersion() string {
	if m != nil {
		return m.IstioVersion
	}
	return ""
}

// Mesh object representing AWS App Mesh
type AwsAppMesh struct {
	// Reference to the secret that holds the AWS credentials that will be used to access the AWS App Mesh service.
	AwsSecret *core.ResourceRef `protobuf:"bytes,1,opt,name=aws_secret,json=awsSecret,proto3" json:"aws_secret,omitempty"`
	// The AWS region the AWS App Mesh control plane resources (Virtual Nodes, Virtual Routers, etc.) will be created in.
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// Determines whether pods will be automatically injected with the AWS App Mesh Envoy sidecar proxy.
	//
	// If set to true, supergloo will ensure that a MutatingAdmissionWebhook server with the injection logic is deployed
	// to the cluster and that it has been registered with the Kubernetes API server via a MutatingWebhookConfiguration.
	// This will cause the webhook to be invoked on each pod creation event.
	EnableAutoInject bool `protobuf:"varint,3,opt,name=enable_auto_inject,json=enableAutoInject,proto3" json:"enable_auto_inject,omitempty"`
	// Pods matching this selector will be injected with the sidecar proxy at creation time.
	//
	// NOTE: the sidecar injector webhook currently supports only the NamespaceSelector and LabelSelector
	InjectionSelector *PodSelector `protobuf:"bytes,4,opt,name=injection_selector,json=injectionSelector,proto3" json:"injection_selector,omitempty"`
	// If auto-injection is enabled, the value of the pod label with this key will be used to calculate the value of
	// APPMESH_VIRTUAL_NODE_NAME environment variable that is set on the injected sidecar proxy container.
	VirtualNodeLabel string `protobuf:"bytes,5,opt,name=virtual_node_label,json=virtualNodeLabel,proto3" json:"virtual_node_label,omitempty"`
	// Reference to the config map that contains the patch that will be applied to the spec of the pods matching the
	// injection_selector.
	SidecarPatchConfigMap *core.ResourceRef `protobuf:"bytes,6,opt,name=sidecar_patch_config_map,json=sidecarPatchConfigMap,proto3" json:"sidecar_patch_config_map,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *AwsAppMesh) Reset()         { *m = AwsAppMesh{} }
func (m *AwsAppMesh) String() string { return proto.CompactTextString(m) }
func (*AwsAppMesh) ProtoMessage()    {}
func (*AwsAppMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{3}
}
func (m *AwsAppMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsAppMesh.Unmarshal(m, b)
}
func (m *AwsAppMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsAppMesh.Marshal(b, m, deterministic)
}
func (m *AwsAppMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsAppMesh.Merge(m, src)
}
func (m *AwsAppMesh) XXX_Size() int {
	return xxx_messageInfo_AwsAppMesh.Size(m)
}
func (m *AwsAppMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsAppMesh.DiscardUnknown(m)
}

var xxx_messageInfo_AwsAppMesh proto.InternalMessageInfo

func (m *AwsAppMesh) GetAwsSecret() *core.ResourceRef {
	if m != nil {
		return m.AwsSecret
	}
	return nil
}

func (m *AwsAppMesh) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *AwsAppMesh) GetEnableAutoInject() bool {
	if m != nil {
		return m.EnableAutoInject
	}
	return false
}

func (m *AwsAppMesh) GetInjectionSelector() *PodSelector {
	if m != nil {
		return m.InjectionSelector
	}
	return nil
}

func (m *AwsAppMesh) GetVirtualNodeLabel() string {
	if m != nil {
		return m.VirtualNodeLabel
	}
	return ""
}

func (m *AwsAppMesh) GetSidecarPatchConfigMap() *core.ResourceRef {
	if m != nil {
		return m.SidecarPatchConfigMap
	}
	return nil
}

// Mesh object representing an installed Linkerd control plane
type LinkerdMesh struct {
	// where the Linkerd control plane has been installed
	InstallationNamespace string `protobuf:"bytes,1,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// version of istio which has been installed
	LinkerdVersion       string   `protobuf:"bytes,2,opt,name=linkerd_version,json=linkerdVersion,proto3" json:"linkerd_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkerdMesh) Reset()         { *m = LinkerdMesh{} }
func (m *LinkerdMesh) String() string { return proto.CompactTextString(m) }
func (*LinkerdMesh) ProtoMessage()    {}
func (*LinkerdMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{4}
}
func (m *LinkerdMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkerdMesh.Unmarshal(m, b)
}
func (m *LinkerdMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkerdMesh.Marshal(b, m, deterministic)
}
func (m *LinkerdMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkerdMesh.Merge(m, src)
}
func (m *LinkerdMesh) XXX_Size() int {
	return xxx_messageInfo_LinkerdMesh.Size(m)
}
func (m *LinkerdMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkerdMesh.DiscardUnknown(m)
}

var xxx_messageInfo_LinkerdMesh proto.InternalMessageInfo

func (m *LinkerdMesh) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *LinkerdMesh) GetLinkerdVersion() string {
	if m != nil {
		return m.LinkerdVersion
	}
	return ""
}

// the encryption configuration that will be applied by the role
type MtlsConfig struct {
	// whether or not mutual TLS should be enabled between pods in this mesh
	MtlsEnabled bool `protobuf:"varint,1,opt,name=mtls_enabled,json=mtlsEnabled,proto3" json:"mtls_enabled,omitempty"`
	// if set, rootCertificate will override the root certificate used by the mesh
	// to encrypt mtls connections.
	//
	// The structure of the secret must be a standard kubernetes TLS secret
	// such as can be created via `kubectl create secret tls`
	//
	// if mtlsEnabled is false, this field is ignored
	// If deploying to Consul, Consul Connect requires that the cert and key are generated using ec, not rsa.
	RootCertificate      *core.ResourceRef `protobuf:"bytes,3,opt,name=root_certificate,json=rootCertificate,proto3" json:"root_certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MtlsConfig) Reset()         { *m = MtlsConfig{} }
func (m *MtlsConfig) String() string { return proto.CompactTextString(m) }
func (*MtlsConfig) ProtoMessage()    {}
func (*MtlsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{5}
}
func (m *MtlsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MtlsConfig.Unmarshal(m, b)
}
func (m *MtlsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MtlsConfig.Marshal(b, m, deterministic)
}
func (m *MtlsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MtlsConfig.Merge(m, src)
}
func (m *MtlsConfig) XXX_Size() int {
	return xxx_messageInfo_MtlsConfig.Size(m)
}
func (m *MtlsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MtlsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MtlsConfig proto.InternalMessageInfo

func (m *MtlsConfig) GetMtlsEnabled() bool {
	if m != nil {
		return m.MtlsEnabled
	}
	return false
}

func (m *MtlsConfig) GetRootCertificate() *core.ResourceRef {
	if m != nil {
		return m.RootCertificate
	}
	return nil
}

// Contains configuration options for monitoring a mesh
// Currently MonitoringConfig only contains options for configuring
// an in-cluster Prometheus instance to scrape a mesh for metrics
type MonitoringConfig struct {
	// indicates to supergloo that metrics should be propagated to one or more instances of prometheus.
	// add a [`core.solo.io.ResourceRef`](../../../../solo-kit/api/v1/ref.proto.sk#ResourceRef) for each
	// NAMESPACE.NAME of the configmap used to configure each prometheus instance.
	// assumes that the configmap contains a key named `prometheus.yml` whose value
	// is the prometheus yaml config as an inline string
	PrometheusConfigmaps []core.ResourceRef `protobuf:"bytes,1,rep,name=prometheus_configmaps,json=prometheusConfigmaps,proto3" json:"prometheus_configmaps"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MonitoringConfig) Reset()         { *m = MonitoringConfig{} }
func (m *MonitoringConfig) String() string { return proto.CompactTextString(m) }
func (*MonitoringConfig) ProtoMessage()    {}
func (*MonitoringConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{6}
}
func (m *MonitoringConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitoringConfig.Unmarshal(m, b)
}
func (m *MonitoringConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitoringConfig.Marshal(b, m, deterministic)
}
func (m *MonitoringConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoringConfig.Merge(m, src)
}
func (m *MonitoringConfig) XXX_Size() int {
	return xxx_messageInfo_MonitoringConfig.Size(m)
}
func (m *MonitoringConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoringConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoringConfig proto.InternalMessageInfo

func (m *MonitoringConfig) GetPrometheusConfigmaps() []core.ResourceRef {
	if m != nil {
		return m.PrometheusConfigmaps
	}
	return nil
}

type MeshGroup struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by supergloo during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	// the meshes contained in this group
	Meshes               []*core.ResourceRef `protobuf:"bytes,3,rep,name=meshes,proto3" json:"meshes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MeshGroup) Reset()         { *m = MeshGroup{} }
func (m *MeshGroup) String() string { return proto.CompactTextString(m) }
func (*MeshGroup) ProtoMessage()    {}
func (*MeshGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_713281dd1a237b0d, []int{7}
}
func (m *MeshGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshGroup.Unmarshal(m, b)
}
func (m *MeshGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshGroup.Marshal(b, m, deterministic)
}
func (m *MeshGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshGroup.Merge(m, src)
}
func (m *MeshGroup) XXX_Size() int {
	return xxx_messageInfo_MeshGroup.Size(m)
}
func (m *MeshGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MeshGroup proto.InternalMessageInfo

func (m *MeshGroup) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *MeshGroup) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *MeshGroup) GetMeshes() []*core.ResourceRef {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func init() {
	proto.RegisterType((*Mesh)(nil), "supergloo.solo.io.Mesh")
	proto.RegisterType((*DiscoveryMetadata)(nil), "supergloo.solo.io.DiscoveryMetadata")
	proto.RegisterType((*IstioMesh)(nil), "supergloo.solo.io.IstioMesh")
	proto.RegisterType((*AwsAppMesh)(nil), "supergloo.solo.io.AwsAppMesh")
	proto.RegisterType((*LinkerdMesh)(nil), "supergloo.solo.io.LinkerdMesh")
	proto.RegisterType((*MtlsConfig)(nil), "supergloo.solo.io.MtlsConfig")
	proto.RegisterType((*MonitoringConfig)(nil), "supergloo.solo.io.MonitoringConfig")
	proto.RegisterType((*MeshGroup)(nil), "supergloo.solo.io.MeshGroup")
}

func init() {
	proto.RegisterFile("github.com/solo-io/supergloo/api/v1/mesh.proto", fileDescriptor_713281dd1a237b0d)
}

var fileDescriptor_713281dd1a237b0d = []byte{
	// 902 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xc1, 0x6e, 0xdb, 0x46,
	0x10, 0x8d, 0x2c, 0x59, 0xb0, 0x46, 0x6a, 0x23, 0x6d, 0x6d, 0x83, 0x09, 0x5a, 0x27, 0x55, 0x0a,
	0xa4, 0x87, 0x84, 0x84, 0xd3, 0x16, 0x30, 0x7c, 0x28, 0x60, 0x3b, 0x45, 0x1a, 0x20, 0x0a, 0x0c,
	0xba, 0xe8, 0xa1, 0x17, 0x62, 0x45, 0x8e, 0xa8, 0xad, 0x49, 0xee, 0x62, 0x77, 0x69, 0x23, 0x57,
	0xff, 0x43, 0xff, 0xa1, 0x9f, 0xd2, 0x63, 0xef, 0x45, 0x73, 0xe8, 0x1f, 0xb8, 0x5f, 0x50, 0xec,
	0x72, 0x49, 0xc5, 0x8e, 0xac, 0x3a, 0x3d, 0xf5, 0x44, 0xee, 0xcc, 0x7b, 0x6f, 0x67, 0x67, 0x66,
	0x67, 0xc1, 0x4f, 0x99, 0x9e, 0x97, 0x53, 0x3f, 0xe6, 0x79, 0xa0, 0x78, 0xc6, 0x9f, 0x32, 0x1e,
	0xa8, 0x52, 0xa0, 0x4c, 0x33, 0xce, 0x03, 0x2a, 0x58, 0x70, 0xb6, 0x1b, 0xe4, 0xa8, 0xe6, 0xbe,
	0x90, 0x5c, 0x73, 0x32, 0x6a, 0x9c, 0xbe, 0x81, 0xfb, 0x8c, 0xdf, 0xdf, 0x4c, 0x79, 0xca, 0xad,
	0x37, 0x30, 0x7f, 0x15, 0xf0, 0xfe, 0x4e, 0xca, 0x79, 0x9a, 0x61, 0x60, 0x57, 0xd3, 0x72, 0x16,
	0x24, 0xa5, 0xa4, 0x9a, 0xf1, 0xe2, 0x26, 0xff, 0xb9, 0xa4, 0x42, 0xa0, 0x54, 0xce, 0xbf, 0xbb,
	0x2c, 0x30, 0xf3, 0x3d, 0x65, 0x7a, 0x11, 0x97, 0xa6, 0x09, 0xd5, 0xd4, 0x51, 0x82, 0x5b, 0x50,
	0x94, 0xa6, 0xba, 0xac, 0xf7, 0x78, 0x72, 0x0b, 0x82, 0xc4, 0xd9, 0x07, 0x44, 0x54, 0xaf, 0x1d,
	0xe5, 0xd9, 0x6d, 0xb2, 0xab, 0x30, 0xc3, 0x58, 0x73, 0x59, 0x71, 0xc6, 0xbf, 0x77, 0xa0, 0x33,
	0x41, 0x35, 0x27, 0x2f, 0xa0, 0x5b, 0x45, 0xeb, 0x25, 0x0f, 0x5b, 0x5f, 0xf6, 0x9f, 0x6d, 0xfa,
	0x31, 0x97, 0x58, 0xa7, 0xdd, 0x3f, 0xb1, 0xbe, 0xc3, 0x7b, 0xbf, 0xbd, 0x7d, 0x70, 0xe7, 0xef,
	0xb7, 0x0f, 0x46, 0x1a, 0x95, 0x4e, 0xd8, 0x6c, 0xb6, 0x3f, 0x66, 0x69, 0xc1, 0x25, 0x8e, 0x43,
	0x47, 0x27, 0x7b, 0xb0, 0x51, 0x67, 0xca, 0x43, 0x2b, 0xb5, 0x7d, 0x55, 0x6a, 0xe2, 0xbc, 0x87,
	0x1d, 0x23, 0x16, 0x36, 0x68, 0xf2, 0x35, 0xac, 0x33, 0xa5, 0x19, 0xf7, 0x5a, 0x96, 0xf6, 0xa9,
	0xff, 0x5e, 0xf5, 0xfd, 0x97, 0xc6, 0x6f, 0xe2, 0xfd, 0xfe, 0x4e, 0x58, 0x81, 0xc9, 0x01, 0x0c,
	0xe8, 0xb9, 0x8a, 0xa8, 0x10, 0x91, 0xe9, 0x1c, 0xaf, 0x63, 0xc9, 0x9f, 0x2d, 0x21, 0x1f, 0x9c,
	0xab, 0x03, 0x21, 0x1c, 0x1b, 0x68, 0xb3, 0x22, 0x47, 0x30, 0xc8, 0x58, 0x71, 0x8a, 0x32, 0xa9,
	0x24, 0xd6, 0xad, 0xc4, 0xce, 0x12, 0x89, 0x57, 0x15, 0xcc, 0x69, 0xf4, 0xb3, 0xc5, 0x92, 0x7c,
	0x0b, 0xfd, 0x5c, 0x67, 0x2a, 0x8a, 0x79, 0x31, 0x63, 0xa9, 0xb7, 0x76, 0x63, 0x18, 0x13, 0x9d,
	0xa9, 0x23, 0x0b, 0x0a, 0x21, 0x6f, 0xfe, 0xc9, 0x31, 0x8c, 0x72, 0x5e, 0x30, 0xcd, 0x25, 0x2b,
	0xd2, 0x5a, 0xa5, 0x6d, 0x55, 0x1e, 0x2d, 0x53, 0x69, 0xb0, 0x4e, 0x6b, 0x98, 0x5f, 0xb3, 0x90,
	0x13, 0x20, 0x09, 0x53, 0x31, 0x3f, 0x43, 0xf9, 0x26, 0x6a, 0x6a, 0xd2, 0xb5, 0x92, 0x5f, 0x2c,
	0x91, 0x7c, 0x5e, 0x83, 0xeb, 0x0a, 0x85, 0xa3, 0xe4, 0xba, 0x69, 0xff, 0x93, 0x8b, 0xcb, 0x4e,
	0x1b, 0x5a, 0xf9, 0xc5, 0x65, 0x67, 0x83, 0x74, 0x4d, 0xba, 0x50, 0x1d, 0xf6, 0xa1, 0x67, 0xfe,
	0x22, 0xfd, 0x46, 0xe0, 0xf8, 0x97, 0x35, 0x18, 0xbd, 0x27, 0x45, 0xf6, 0xc0, 0x63, 0xc5, 0xcf,
	0x18, 0x6b, 0x4c, 0xa2, 0x82, 0xe6, 0xa8, 0x04, 0x8d, 0x31, 0xca, 0xe8, 0x14, 0x33, 0x5b, 0xef,
	0x5e, 0xb8, 0x5d, 0xfb, 0x5f, 0xd7, 0xee, 0x57, 0xc6, 0x4b, 0x9e, 0x00, 0xc1, 0x82, 0x4e, 0x33,
	0x8c, 0x68, 0xa9, 0x79, 0x54, 0xa1, 0x6c, 0x7e, 0x37, 0xc2, 0x61, 0xe5, 0x39, 0x28, 0x35, 0x7f,
	0x69, 0xed, 0xe4, 0x73, 0x18, 0xd8, 0x50, 0xce, 0x50, 0x2a, 0xc6, 0x0b, 0x9b, 0xc1, 0x5e, 0xd8,
	0x37, 0xb6, 0x1f, 0x2b, 0x13, 0xf9, 0x06, 0xb6, 0x59, 0xa1, 0x34, 0xcd, 0x32, 0x3b, 0x22, 0x16,
	0xe1, 0xd8, 0xde, 0xe9, 0x85, 0x5b, 0xef, 0x7a, 0x9b, 0x60, 0xae, 0x17, 0x78, 0xfd, 0x03, 0x0b,
	0x3c, 0x4e, 0xa1, 0xd7, 0xb4, 0xef, 0x8a, 0x18, 0x5a, 0xab, 0x62, 0x78, 0x04, 0x1f, 0xd9, 0xae,
	0x6f, 0x8e, 0xb7, 0x66, 0xd1, 0x03, 0x6b, 0x74, 0xe7, 0x1b, 0xff, 0xb9, 0x06, 0xb0, 0xe8, 0x75,
	0xb2, 0x07, 0xa6, 0xd7, 0x23, 0x85, 0xb1, 0x44, 0xed, 0xee, 0xd6, 0xbd, 0xab, 0x57, 0x32, 0x44,
	0xc5, 0x4b, 0x19, 0x63, 0x88, 0xb3, 0xb0, 0x47, 0xcf, 0xd5, 0x89, 0xc5, 0x92, 0x6d, 0xe8, 0x4a,
	0x4c, 0x17, 0xdb, 0xb8, 0xd5, 0x0d, 0x15, 0x69, 0xdf, 0x50, 0x91, 0x09, 0x90, 0x0a, 0x61, 0xce,
	0x59, 0x8f, 0x1f, 0x77, 0x4d, 0x97, 0xdd, 0xb1, 0x63, 0x9e, 0x9c, 0x38, 0x54, 0x38, 0x6a, 0x98,
	0xb5, 0xc9, 0x6c, 0x7e, 0xc6, 0xa4, 0x2e, 0x69, 0x16, 0x15, 0x3c, 0xa9, 0x5b, 0x68, 0xdd, 0x06,
	0x38, 0x74, 0x9e, 0xd7, 0x3c, 0x71, 0xcd, 0x13, 0x82, 0xa7, 0x58, 0x82, 0x31, 0x95, 0x91, 0xa0,
	0x3a, 0x9e, 0xbb, 0xea, 0x45, 0x39, 0x15, 0xee, 0x26, 0xac, 0x48, 0xc5, 0x96, 0xa3, 0x1e, 0x1b,
	0x66, 0x55, 0xc5, 0x09, 0x15, 0xe3, 0x1c, 0xfa, 0xef, 0xcc, 0x81, 0xff, 0x5a, 0xca, 0xc7, 0x70,
	0xb7, 0x1e, 0x3a, 0x57, 0x8b, 0xf9, 0xb1, 0x33, 0xd7, 0xe5, 0x2c, 0x01, 0x16, 0x1d, 0x65, 0xfb,
	0xdb, 0x74, 0x61, 0x95, 0xe6, 0xc4, 0xee, 0xb1, 0x11, 0xda, 0xce, 0xfc, 0xae, 0x32, 0x91, 0xe7,
	0x30, 0x94, 0x9c, 0xeb, 0x28, 0x46, 0xa9, 0xd9, 0x8c, 0xc5, 0x54, 0xa3, 0x1b, 0x24, 0x2b, 0xce,
	0x7a, 0xd7, 0x50, 0x8e, 0x16, 0x8c, 0xf1, 0x1c, 0x86, 0xd7, 0x67, 0x0c, 0xf9, 0x01, 0xb6, 0x84,
	0xe4, 0x39, 0xea, 0x39, 0x96, 0xf5, 0x45, 0xc8, 0xa9, 0x50, 0x5e, 0xeb, 0x61, 0x7b, 0xa5, 0xbc,
	0x9b, 0xf5, 0x9b, 0x0b, 0xf6, 0x51, 0x43, 0x1e, 0xff, 0xd1, 0x82, 0x9e, 0xc9, 0xe4, 0x0b, 0xc9,
	0x4b, 0xf1, 0x7f, 0x78, 0x88, 0x76, 0xc1, 0x0d, 0x36, 0xaf, 0xfd, 0x2f, 0xe7, 0x0a, 0x1d, 0x70,
	0xdf, 0xbb, 0xb8, 0xec, 0x74, 0x60, 0x2d, 0x4f, 0x2f, 0x2e, 0x3b, 0x03, 0x02, 0xc6, 0x9a, 0x9a,
	0xe3, 0xa8, 0xc3, 0xa7, 0xbf, 0xfe, 0xb5, 0xd3, 0xfa, 0xe9, 0xf1, 0xca, 0xb7, 0x59, 0x9c, 0xa6,
	0xee, 0x7d, 0x9e, 0x76, 0xed, 0xbb, 0xfc, 0xd5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x8a,
	0xe3, 0x3c, 0x2b, 0x09, 0x00, 0x00,
}

func (this *Mesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh)
	if !ok {
		that2, ok := that.(Mesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if that1.MeshType == nil {
		if this.MeshType != nil {
			return false
		}
	} else if this.MeshType == nil {
		return false
	} else if !this.MeshType.Equal(that1.MeshType) {
		return false
	}
	if !this.MtlsConfig.Equal(that1.MtlsConfig) {
		return false
	}
	if !this.MonitoringConfig.Equal(that1.MonitoringConfig) {
		return false
	}
	if !this.DiscoveryMetadata.Equal(that1.DiscoveryMetadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Mesh_Istio) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_Istio)
	if !ok {
		that2, ok := that.(Mesh_Istio)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Istio.Equal(that1.Istio) {
		return false
	}
	return true
}
func (this *Mesh_AwsAppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_AwsAppMesh)
	if !ok {
		that2, ok := that.(Mesh_AwsAppMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.AwsAppMesh.Equal(that1.AwsAppMesh) {
		return false
	}
	return true
}
func (this *Mesh_LinkerdMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_LinkerdMesh)
	if !ok {
		that2, ok := that.(Mesh_LinkerdMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.LinkerdMesh.Equal(that1.LinkerdMesh) {
		return false
	}
	return true
}
func (this *DiscoveryMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryMetadata)
	if !ok {
		that2, ok := that.(DiscoveryMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InjectedNamespaceLabel != that1.InjectedNamespaceLabel {
		return false
	}
	if this.EnableAutoInject != that1.EnableAutoInject {
		return false
	}
	if this.MeshVersion != that1.MeshVersion {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if !this.MtlsConfig.Equal(that1.MtlsConfig) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IstioMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IstioMesh)
	if !ok {
		that2, ok := that.(IstioMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if this.IstioVersion != that1.IstioVersion {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AwsAppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsAppMesh)
	if !ok {
		that2, ok := that.(AwsAppMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.AwsSecret.Equal(that1.AwsSecret) {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if this.EnableAutoInject != that1.EnableAutoInject {
		return false
	}
	if !this.InjectionSelector.Equal(that1.InjectionSelector) {
		return false
	}
	if this.VirtualNodeLabel != that1.VirtualNodeLabel {
		return false
	}
	if !this.SidecarPatchConfigMap.Equal(that1.SidecarPatchConfigMap) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LinkerdMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LinkerdMesh)
	if !ok {
		that2, ok := that.(LinkerdMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if this.LinkerdVersion != that1.LinkerdVersion {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MtlsConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MtlsConfig)
	if !ok {
		that2, ok := that.(MtlsConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MtlsEnabled != that1.MtlsEnabled {
		return false
	}
	if !this.RootCertificate.Equal(that1.RootCertificate) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MonitoringConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MonitoringConfig)
	if !ok {
		that2, ok := that.(MonitoringConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.PrometheusConfigmaps) != len(that1.PrometheusConfigmaps) {
		return false
	}
	for i := range this.PrometheusConfigmaps {
		if !this.PrometheusConfigmaps[i].Equal(&that1.PrometheusConfigmaps[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshGroup) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshGroup)
	if !ok {
		that2, ok := that.(MeshGroup)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
