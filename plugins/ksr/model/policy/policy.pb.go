// Code generated by protoc-gen-go. DO NOT EDIT.
// source: policy.proto

/*
Package policy is a generated protocol buffer package.

Package policy defines data model for Kubernetes Network Policy.

It is generated from these files:
	policy.proto

It has these top-level messages:
	Policy
*/
package policy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// PolicyType selects the rule types that the network policy relates to.
// By default, rule types are determined based on the existence of Ingress or
// Egress rules: policies that contain an Egress section are assumed to affect
// Egress, and all policies (whether or not they contain an Ingress section)
// are assumed to affect Ingress.
// For example, policies are egress-only if and only if policyType is set
// to EGRESS.
// Likewise, policies blocking all egress traffic are either EGRESS
// or INGRESS_AND_EGRESS as they do not include an Egress section and would
// otherwise default to just INGRESS.
// This field is beta-level in Kubernetes 1.8.
// +optional
type Policy_PolicyType int32

const (
	Policy_DEFAULT            Policy_PolicyType = 0
	Policy_INGRESS            Policy_PolicyType = 1
	Policy_EGRESS             Policy_PolicyType = 2
	Policy_INGRESS_AND_EGRESS Policy_PolicyType = 3
)

var Policy_PolicyType_name = map[int32]string{
	0: "DEFAULT",
	1: "INGRESS",
	2: "EGRESS",
	3: "INGRESS_AND_EGRESS",
}
var Policy_PolicyType_value = map[string]int32{
	"DEFAULT":            0,
	"INGRESS":            1,
	"EGRESS":             2,
	"INGRESS_AND_EGRESS": 3,
}

func (x Policy_PolicyType) String() string {
	return proto.EnumName(Policy_PolicyType_name, int32(x))
}
func (Policy_PolicyType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Operator represents a key's relationship to a set of values.
type Policy_LabelSelector_LabelExpression_Operator int32

const (
	Policy_LabelSelector_LabelExpression_IN             Policy_LabelSelector_LabelExpression_Operator = 0
	Policy_LabelSelector_LabelExpression_NOT_IN         Policy_LabelSelector_LabelExpression_Operator = 1
	Policy_LabelSelector_LabelExpression_EXISTS         Policy_LabelSelector_LabelExpression_Operator = 2
	Policy_LabelSelector_LabelExpression_DOES_NOT_EXIST Policy_LabelSelector_LabelExpression_Operator = 3
)

var Policy_LabelSelector_LabelExpression_Operator_name = map[int32]string{
	0: "IN",
	1: "NOT_IN",
	2: "EXISTS",
	3: "DOES_NOT_EXIST",
}
var Policy_LabelSelector_LabelExpression_Operator_value = map[string]int32{
	"IN":             0,
	"NOT_IN":         1,
	"EXISTS":         2,
	"DOES_NOT_EXIST": 3,
}

func (x Policy_LabelSelector_LabelExpression_Operator) String() string {
	return proto.EnumName(Policy_LabelSelector_LabelExpression_Operator_name, int32(x))
}
func (Policy_LabelSelector_LabelExpression_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 0, 0}
}

// The protocol (TCP or UDP) which traffic must match.
// If not specified, this field defaults to TCP.
// +optional
type Policy_Port_Protocol int32

const (
	Policy_Port_TCP Policy_Port_Protocol = 0
	Policy_Port_UDP Policy_Port_Protocol = 1
)

var Policy_Port_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}
var Policy_Port_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x Policy_Port_Protocol) String() string {
	return proto.EnumName(Policy_Port_Protocol_name, int32(x))
}
func (Policy_Port_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2, 0} }

// Port reference type.
type Policy_Port_PortNameOrNumber_Type int32

const (
	Policy_Port_PortNameOrNumber_NUMBER Policy_Port_PortNameOrNumber_Type = 0
	Policy_Port_PortNameOrNumber_NAME   Policy_Port_PortNameOrNumber_Type = 1
)

var Policy_Port_PortNameOrNumber_Type_name = map[int32]string{
	0: "NUMBER",
	1: "NAME",
}
var Policy_Port_PortNameOrNumber_Type_value = map[string]int32{
	"NUMBER": 0,
	"NAME":   1,
}

func (x Policy_Port_PortNameOrNumber_Type) String() string {
	return proto.EnumName(Policy_Port_PortNameOrNumber_Type_name, int32(x))
}
func (Policy_Port_PortNameOrNumber_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2, 0, 0}
}

// Policy describes what network traffic is allowed for a set of Pods.
type Policy struct {
	// Name of the policy unique within the namespace.
	// Cannot be updated.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Namespace the policy is inserted into.
	// An empty namespace is equivalent to the "default" namespace, but "default"
	// is the canonical representation used in the key for a key-value store.
	// Cannot be updated.
	Namespace string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	// A list of labels attached to this policy.
	// +optional
	Label []*Policy_Label `protobuf:"bytes,3,rep,name=label" json:"label,omitempty"`
	// Pods to which this policy applies. The array of ingress rules is applied
	// to all pods selected by this field. Multiple network policies can select
	// the same set of pods. In such case, the ingress rules for each are combined
	// additively.
	// This field is NOT optional and follows standard label selector semantics.
	// An empty selector matches all pods in this namespace.
	Pods       *Policy_LabelSelector `protobuf:"bytes,4,opt,name=pods" json:"pods,omitempty"`
	PolicyType Policy_PolicyType     `protobuf:"varint,5,opt,name=policy_type,json=policyType,enum=policy.Policy_PolicyType" json:"policy_type,omitempty"`
	// List of ingress rules applied to the selected pods.
	// Traffic is allowed to a pod if there are no network policies selecting the pod
	// OR if the traffic source is the pod's local node,
	// OR if the traffic matches at least one ingress rule across all of the network
	// policies applied to the pod.
	// If there are no ingress rules then this network policy does not allow
	// any traffic (and serves solely to ensure that the selected pods are isolated
	// by default).
	// +optional
	IngressRule []*Policy_IngressRule `protobuf:"bytes,6,rep,name=ingress_rule,json=ingressRule" json:"ingress_rule,omitempty"`
	// List of egress rules to be applied to the selected pods.
	// Outgoing traffic is allowed if there are no network policies selecting
	// the pod OR if the traffic matches at least one egress rule across
	// all of the network policies applied to the pod.
	// If there are no egress rules then this network policy does not allow
	// any outgoing traffic (and serves solely to ensure that the selected pods
	// are isolated by default).
	// This field is beta-level in Kubernetes 1.8.
	// +optional
	EgressRule []*Policy_EgressRule `protobuf:"bytes,7,rep,name=egress_rule,json=egressRule" json:"egress_rule,omitempty"`
}

func (m *Policy) Reset()                    { *m = Policy{} }
func (m *Policy) String() string            { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()               {}
func (*Policy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Policy) GetLabel() []*Policy_Label {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *Policy) GetPods() *Policy_LabelSelector {
	if m != nil {
		return m.Pods
	}
	return nil
}

func (m *Policy) GetPolicyType() Policy_PolicyType {
	if m != nil {
		return m.PolicyType
	}
	return Policy_DEFAULT
}

func (m *Policy) GetIngressRule() []*Policy_IngressRule {
	if m != nil {
		return m.IngressRule
	}
	return nil
}

func (m *Policy) GetEgressRule() []*Policy_EgressRule {
	if m != nil {
		return m.EgressRule
	}
	return nil
}

// Label is a key/value pair attached to an object (namespace in this case).
// Labels are used to organize and to select subsets of objects.
type Policy_Label struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Policy_Label) Reset()                    { *m = Policy_Label{} }
func (m *Policy_Label) String() string            { return proto.CompactTextString(m) }
func (*Policy_Label) ProtoMessage()               {}
func (*Policy_Label) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Policy_Label) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Policy_Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// A label selector is a label query over a set of resources.
// The result of match_label-s and match_expression-s are ANDed.
// An empty label selector matches all objects. A null label selector matches
// no objects.
type Policy_LabelSelector struct {
	// A list of labels that a resource needs to have attached in order to get
	// selected.
	// +optional
	MatchLabel []*Policy_Label `protobuf:"bytes,1,rep,name=match_label,json=matchLabel" json:"match_label,omitempty"`
	// A list of key-value expressions applied to labels.
	// For a given resource and its labels, all expressions must evaluate
	// to TRUE for the resource to get selected.
	MatchExpression []*Policy_LabelSelector_LabelExpression `protobuf:"bytes,2,rep,name=match_expression,json=matchExpression" json:"match_expression,omitempty"`
}

func (m *Policy_LabelSelector) Reset()                    { *m = Policy_LabelSelector{} }
func (m *Policy_LabelSelector) String() string            { return proto.CompactTextString(m) }
func (*Policy_LabelSelector) ProtoMessage()               {}
func (*Policy_LabelSelector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *Policy_LabelSelector) GetMatchLabel() []*Policy_Label {
	if m != nil {
		return m.MatchLabel
	}
	return nil
}

func (m *Policy_LabelSelector) GetMatchExpression() []*Policy_LabelSelector_LabelExpression {
	if m != nil {
		return m.MatchExpression
	}
	return nil
}

// An expression that contains values, a label key, and an operator that
// relates the key and values.
type Policy_LabelSelector_LabelExpression struct {
	// Key is the label key that the expression applies to.
	Key      string                                        `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Operator Policy_LabelSelector_LabelExpression_Operator `protobuf:"varint,2,opt,name=operator,enum=policy.Policy_LabelSelector_LabelExpression_Operator" json:"operator,omitempty"`
	// An array of string values.
	// If the operator is IN or NOT_IN, the values array must be non-empty.
	// If the operator is EXISTS or DOES_NOT_EXIST, the values array
	// must be empty.
	// +optional
	Value []string `protobuf:"bytes,3,rep,name=value" json:"value,omitempty"`
}

func (m *Policy_LabelSelector_LabelExpression) Reset()         { *m = Policy_LabelSelector_LabelExpression{} }
func (m *Policy_LabelSelector_LabelExpression) String() string { return proto.CompactTextString(m) }
func (*Policy_LabelSelector_LabelExpression) ProtoMessage()    {}
func (*Policy_LabelSelector_LabelExpression) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1, 0}
}

func (m *Policy_LabelSelector_LabelExpression) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Policy_LabelSelector_LabelExpression) GetOperator() Policy_LabelSelector_LabelExpression_Operator {
	if m != nil {
		return m.Operator
	}
	return Policy_LabelSelector_LabelExpression_IN
}

func (m *Policy_LabelSelector_LabelExpression) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

// A port selector.
type Policy_Port struct {
	Protocol Policy_Port_Protocol `protobuf:"varint,3,opt,name=protocol,enum=policy.Policy_Port_Protocol" json:"protocol,omitempty"`
	// If specified, the port on the given protocol.
	// This can either be a numerical or named port on a pod.
	// If this field is not provided, the rule matches all port names and
	// numbers.
	// If present, only traffic on the specified protocol AND port
	// will be matched.
	// +optional
	Port *Policy_Port_PortNameOrNumber `protobuf:"bytes,1,opt,name=port" json:"port,omitempty"`
}

func (m *Policy_Port) Reset()                    { *m = Policy_Port{} }
func (m *Policy_Port) String() string            { return proto.CompactTextString(m) }
func (*Policy_Port) ProtoMessage()               {}
func (*Policy_Port) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *Policy_Port) GetProtocol() Policy_Port_Protocol {
	if m != nil {
		return m.Protocol
	}
	return Policy_Port_TCP
}

func (m *Policy_Port) GetPort() *Policy_Port_PortNameOrNumber {
	if m != nil {
		return m.Port
	}
	return nil
}

// Numerical or named port.
type Policy_Port_PortNameOrNumber struct {
	Type Policy_Port_PortNameOrNumber_Type `protobuf:"varint,1,opt,name=type,enum=policy.Policy_Port_PortNameOrNumber_Type" json:"type,omitempty"`
	// Port number from the range: 0 < x < 65536.
	Number int32 `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	// Port name as defined by containers in the pod.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Policy_Port_PortNameOrNumber) Reset()         { *m = Policy_Port_PortNameOrNumber{} }
func (m *Policy_Port_PortNameOrNumber) String() string { return proto.CompactTextString(m) }
func (*Policy_Port_PortNameOrNumber) ProtoMessage()    {}
func (*Policy_Port_PortNameOrNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 2, 0}
}

func (m *Policy_Port_PortNameOrNumber) GetType() Policy_Port_PortNameOrNumber_Type {
	if m != nil {
		return m.Type
	}
	return Policy_Port_PortNameOrNumber_NUMBER
}

func (m *Policy_Port_PortNameOrNumber) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Policy_Port_PortNameOrNumber) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A selector for a set of pods.
type Policy_Peer struct {
	// This is a label selector which selects Pods in this namespace.
	// If present but empty, this selector selects all pods in this namespace.
	// +optional
	Pods *Policy_LabelSelector `protobuf:"bytes,1,opt,name=pods" json:"pods,omitempty"`
	// Selects namespaces using cluster scoped-labels.
	// This matches all pods in all namespaces selected by this label selector.
	// If present but empty, this selector selects all namespaces.
	// +optional
	Namespaces *Policy_LabelSelector `protobuf:"bytes,2,opt,name=namespaces" json:"namespaces,omitempty"`
	IpBlock    *Policy_Peer_IPBlock  `protobuf:"bytes,3,opt,name=ip_block,json=ipBlock" json:"ip_block,omitempty"`
}

func (m *Policy_Peer) Reset()                    { *m = Policy_Peer{} }
func (m *Policy_Peer) String() string            { return proto.CompactTextString(m) }
func (*Policy_Peer) ProtoMessage()               {}
func (*Policy_Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

func (m *Policy_Peer) GetPods() *Policy_LabelSelector {
	if m != nil {
		return m.Pods
	}
	return nil
}

func (m *Policy_Peer) GetNamespaces() *Policy_LabelSelector {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *Policy_Peer) GetIpBlock() *Policy_Peer_IPBlock {
	if m != nil {
		return m.IpBlock
	}
	return nil
}

// IPBlock describes a particular CIDR (Ex. "192.168.1.1/24") that is allowed
// to/from the pods selected for this network policy. The except entries
// describe CIDRs that should not be included within this rule.
type Policy_Peer_IPBlock struct {
	// CIDR is a string representing the IP Block.
	// Valid examples are "192.168.1.1/24".
	Cidr string `protobuf:"bytes,1,opt,name=cidr" json:"cidr,omitempty"`
	// Except is a slice of CIDRs that should not be included within an IP Block
	// Valid examples are "192.168.1.1/24".
	// Except values are inside the CIDR range.
	// +optional
	Except []string `protobuf:"bytes,2,rep,name=except" json:"except,omitempty"`
}

func (m *Policy_Peer_IPBlock) Reset()                    { *m = Policy_Peer_IPBlock{} }
func (m *Policy_Peer_IPBlock) String() string            { return proto.CompactTextString(m) }
func (*Policy_Peer_IPBlock) ProtoMessage()               {}
func (*Policy_Peer_IPBlock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3, 0} }

func (m *Policy_Peer_IPBlock) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *Policy_Peer_IPBlock) GetExcept() []string {
	if m != nil {
		return m.Except
	}
	return nil
}

// Ingress rule matches traffic if and only if the traffic matches both port-s
// AND from.
type Policy_IngressRule struct {
	// List of ports made accessible on the pods selected for this policy.
	// Each item in this list is combined using a logical OR.
	// If the array is empty or null, then this ingress rule matches all ports
	// (traffic not restricted by port).
	// If the array is non-empty, then this ingress rule allows traffic
	// only if the traffic matches at least one port in the list.
	// +optional
	Port []*Policy_Port `protobuf:"bytes,1,rep,name=port" json:"port,omitempty"`
	// List of sources which are able to access the pods selected for this
	// policy.
	// Items in this list are combined using a logical OR operation.
	// If the array is empty or null, then this ingress rule matches all sources
	// (traffic not restricted by source).
	// If the array is non-empty, then this ingress rule allows traffic only
	// if the traffic matches at least one item in the from list.
	// +optional
	From []*Policy_Peer `protobuf:"bytes,2,rep,name=from" json:"from,omitempty"`
}

func (m *Policy_IngressRule) Reset()                    { *m = Policy_IngressRule{} }
func (m *Policy_IngressRule) String() string            { return proto.CompactTextString(m) }
func (*Policy_IngressRule) ProtoMessage()               {}
func (*Policy_IngressRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 4} }

func (m *Policy_IngressRule) GetPort() []*Policy_Port {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *Policy_IngressRule) GetFrom() []*Policy_Peer {
	if m != nil {
		return m.From
	}
	return nil
}

// Egress rule matches traffic if and only if the traffic matches both port-s
// AND to.
// This field is beta-level in Kubernetes 1.8.
type Policy_EgressRule struct {
	// List of destination ports for outgoing traffic.
	// Each item in this list is combined using a logical OR.
	// If the array is empty or null, then this egress rule matches all ports
	// (traffic not restricted by port).
	// If the array is non-empty, then this egress rule allows traffic
	// only if the traffic matches at least one port in the list.
	// +optional
	Port []*Policy_Port `protobuf:"bytes,1,rep,name=port" json:"port,omitempty"`
	// List of destinations for outgoing traffic of pods selected for this policy.
	// Items in this list are combined using a logical OR operation.
	// If the array is empty or null, this egress rule matches all destinations
	// (traffic not restricted by destination).
	// If the array is non-empty, then this egress rule allows traffic only
	// if the traffic matches at least one item in the to list.
	// +optional
	To []*Policy_Peer `protobuf:"bytes,2,rep,name=to" json:"to,omitempty"`
}

func (m *Policy_EgressRule) Reset()                    { *m = Policy_EgressRule{} }
func (m *Policy_EgressRule) String() string            { return proto.CompactTextString(m) }
func (*Policy_EgressRule) ProtoMessage()               {}
func (*Policy_EgressRule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 5} }

func (m *Policy_EgressRule) GetPort() []*Policy_Port {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *Policy_EgressRule) GetTo() []*Policy_Peer {
	if m != nil {
		return m.To
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "policy.Policy")
	proto.RegisterType((*Policy_Label)(nil), "policy.Policy.Label")
	proto.RegisterType((*Policy_LabelSelector)(nil), "policy.Policy.LabelSelector")
	proto.RegisterType((*Policy_LabelSelector_LabelExpression)(nil), "policy.Policy.LabelSelector.LabelExpression")
	proto.RegisterType((*Policy_Port)(nil), "policy.Policy.Port")
	proto.RegisterType((*Policy_Port_PortNameOrNumber)(nil), "policy.Policy.Port.PortNameOrNumber")
	proto.RegisterType((*Policy_Peer)(nil), "policy.Policy.Peer")
	proto.RegisterType((*Policy_Peer_IPBlock)(nil), "policy.Policy.Peer.IPBlock")
	proto.RegisterType((*Policy_IngressRule)(nil), "policy.Policy.IngressRule")
	proto.RegisterType((*Policy_EgressRule)(nil), "policy.Policy.EgressRule")
	proto.RegisterEnum("policy.Policy_PolicyType", Policy_PolicyType_name, Policy_PolicyType_value)
	proto.RegisterEnum("policy.Policy_LabelSelector_LabelExpression_Operator", Policy_LabelSelector_LabelExpression_Operator_name, Policy_LabelSelector_LabelExpression_Operator_value)
	proto.RegisterEnum("policy.Policy_Port_Protocol", Policy_Port_Protocol_name, Policy_Port_Protocol_value)
	proto.RegisterEnum("policy.Policy_Port_PortNameOrNumber_Type", Policy_Port_PortNameOrNumber_Type_name, Policy_Port_PortNameOrNumber_Type_value)
}

func init() { proto.RegisterFile("policy.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x8e, 0x1d, 0xe7, 0xa7, 0xe3, 0x9e, 0xd6, 0xda, 0x56, 0x55, 0x8e, 0x4f, 0x2e, 0xa2, 0x1c,
	0x24, 0x02, 0x42, 0x06, 0x05, 0x15, 0x55, 0x88, 0x22, 0xb5, 0xc4, 0xa0, 0xa0, 0xd6, 0x31, 0x9b,
	0x54, 0x20, 0x6e, 0x2c, 0xc7, 0x5d, 0xc0, 0xaa, 0x93, 0xb5, 0x36, 0x0e, 0x6a, 0x9e, 0xa5, 0xef,
	0xc1, 0x53, 0xf0, 0x1a, 0xdc, 0xf2, 0x0c, 0x68, 0xc7, 0x8e, 0x5d, 0x42, 0x54, 0x0a, 0x57, 0x9e,
	0x99, 0xfd, 0xbe, 0xd9, 0x99, 0xcf, 0x33, 0x0b, 0x9b, 0x31, 0x8f, 0xc2, 0x60, 0x61, 0xc5, 0x82,
	0x27, 0x9c, 0x54, 0x53, 0xaf, 0x7d, 0xb5, 0x09, 0x55, 0x17, 0x4d, 0x42, 0x40, 0x9b, 0xfa, 0x13,
	0xd6, 0x50, 0x5a, 0x4a, 0x67, 0x83, 0xa2, 0x4d, 0x9a, 0xb0, 0x21, 0xbf, 0xb3, 0xd8, 0x0f, 0x58,
	0x43, 0xc5, 0x83, 0x22, 0x40, 0xee, 0x43, 0x25, 0xf2, 0xc7, 0x2c, 0x6a, 0x94, 0x5b, 0xe5, 0x8e,
	0xde, 0xdd, 0xb5, 0xb2, 0x2b, 0xd2, 0x84, 0xd6, 0x89, 0x3c, 0xa3, 0x29, 0x84, 0x3c, 0x02, 0x2d,
	0xe6, 0xe7, 0xb3, 0x86, 0xd6, 0x52, 0x3a, 0x7a, 0xb7, 0xb9, 0x0e, 0x3a, 0x64, 0x11, 0x0b, 0x12,
	0x2e, 0x28, 0x22, 0xc9, 0x53, 0xd0, 0x53, 0x90, 0x97, 0x2c, 0x62, 0xd6, 0xa8, 0xb4, 0x94, 0xce,
	0x56, 0xf7, 0xdf, 0x15, 0x62, 0xfa, 0x19, 0x2d, 0x62, 0x46, 0x21, 0xce, 0x6d, 0x72, 0x08, 0x9b,
	0xe1, 0xf4, 0xa3, 0x60, 0xb3, 0x99, 0x27, 0xe6, 0x11, 0x6b, 0x54, 0xb1, 0x40, 0x73, 0x85, 0xdc,
	0x4f, 0x21, 0x74, 0x1e, 0x31, 0xaa, 0x87, 0x85, 0x23, 0xaf, 0x66, 0xd7, 0xd8, 0x35, 0x64, 0xaf,
	0x5e, 0x6d, 0x17, 0x64, 0x60, 0xb9, 0x6d, 0x3e, 0x84, 0x0a, 0x76, 0x43, 0x0c, 0x28, 0x5f, 0xb0,
	0x45, 0x26, 0xa7, 0x34, 0xc9, 0x2e, 0x54, 0x3e, 0xfb, 0xd1, 0x7c, 0xa9, 0x64, 0xea, 0x98, 0xdf,
	0x55, 0xf8, 0xe7, 0xa7, 0xfe, 0xc9, 0x3e, 0xe8, 0x13, 0x3f, 0x09, 0x3e, 0x79, 0xa9, 0xba, 0xca,
	0x0d, 0xea, 0x02, 0x02, 0xd3, 0x0b, 0xdf, 0x82, 0x91, 0xd2, 0xd8, 0x65, 0x2c, 0xcb, 0x09, 0xf9,
	0xb4, 0xa1, 0x22, 0xf7, 0xc1, 0x4d, 0x72, 0xa7, 0x9e, 0x9d, 0x73, 0xe8, 0x36, 0x66, 0x29, 0x02,
	0xe6, 0x57, 0x05, 0xb6, 0x57, 0x40, 0x6b, 0xba, 0x7b, 0x03, 0x75, 0x1e, 0x33, 0xe1, 0x27, 0x5c,
	0x60, 0x83, 0x5b, 0xdd, 0xfd, 0x3f, 0xb9, 0xd6, 0x1a, 0x64, 0x64, 0x9a, 0xa7, 0x29, 0x04, 0x93,
	0x03, 0xb6, 0x14, 0xac, 0xfd, 0x1c, 0xea, 0x4b, 0x2c, 0xa9, 0x82, 0xda, 0x77, 0x8c, 0x12, 0x01,
	0xa8, 0x3a, 0x83, 0x91, 0xd7, 0x77, 0x0c, 0x45, 0xda, 0xf6, 0xbb, 0xfe, 0x70, 0x34, 0x34, 0x54,
	0x42, 0x60, 0xab, 0x37, 0xb0, 0x87, 0x9e, 0x3c, 0xc4, 0xa0, 0x51, 0x36, 0xbf, 0xa8, 0xa0, 0xb9,
	0x5c, 0x24, 0xe4, 0x00, 0xea, 0xb8, 0x0d, 0x01, 0x97, 0x23, 0x2c, 0x2b, 0x6e, 0xfe, 0x32, 0x5e,
	0x22, 0xb1, 0xdc, 0x0c, 0x43, 0x73, 0x34, 0x39, 0x90, 0xd3, 0x2c, 0x12, 0x6c, 0x5f, 0xef, 0xde,
	0x59, 0xcb, 0xe2, 0x22, 0x71, 0xfc, 0x09, 0x1b, 0x08, 0x67, 0x3e, 0x19, 0x33, 0x9c, 0x6a, 0x91,
	0x98, 0x57, 0x0a, 0x18, 0xab, 0x47, 0xe4, 0x10, 0x34, 0x9c, 0x71, 0x05, 0x8b, 0xb8, 0x77, 0x9b,
	0x74, 0x16, 0xce, 0x3c, 0xd2, 0xc8, 0x1e, 0x54, 0xa7, 0x18, 0x44, 0xdd, 0x2b, 0x34, 0xf3, 0xf2,
	0x8d, 0x2e, 0x17, 0x1b, 0xdd, 0x6e, 0x82, 0x86, 0x1b, 0x22, 0x05, 0x3b, 0x3b, 0x3d, 0xb6, 0xa9,
	0x51, 0x22, 0x75, 0xd0, 0x9c, 0xa3, 0x53, 0xdb, 0x50, 0xda, 0x4d, 0xa8, 0x2f, 0xbb, 0x25, 0x35,
	0x28, 0x8f, 0x5e, 0xb8, 0x46, 0x49, 0x1a, 0x67, 0x3d, 0xd7, 0x50, 0xcc, 0x6f, 0x0a, 0x68, 0x2e,
	0x63, 0x22, 0x5f, 0x66, 0xe5, 0xd6, 0xcb, 0xfc, 0x0c, 0x20, 0x7f, 0x37, 0x66, 0x58, 0xe6, 0xef,
	0x78, 0xd7, 0xf0, 0xe4, 0x09, 0xd4, 0xc3, 0xd8, 0x1b, 0x47, 0x3c, 0xb8, 0xc0, 0x66, 0xf4, 0xee,
	0x7f, 0xab, 0x1a, 0x31, 0x26, 0xac, 0xbe, 0x7b, 0x2c, 0x21, 0xb4, 0x16, 0xc6, 0x68, 0x98, 0xfb,
	0x50, 0xcb, 0x62, 0x52, 0x8b, 0x20, 0x3c, 0x17, 0xcb, 0xd7, 0x4d, 0xda, 0x52, 0x37, 0x76, 0x19,
	0xb0, 0x38, 0xc1, 0x35, 0xd9, 0xa0, 0x99, 0x67, 0x7a, 0xa0, 0x5f, 0x7b, 0x1a, 0xc8, 0xdd, 0xfc,
	0x67, 0xcb, 0x5d, 0xda, 0x59, 0xf3, 0x77, 0xd2, 0x7f, 0x2b, 0x81, 0x1f, 0x04, 0x9f, 0x64, 0x4b,
	0xb7, 0xb3, 0xa6, 0x44, 0x8a, 0x00, 0xf3, 0x3d, 0x80, 0xfd, 0x17, 0xf9, 0xff, 0x07, 0x35, 0xe1,
	0x37, 0x65, 0x57, 0x13, 0xde, 0x7e, 0x0d, 0x50, 0x3c, 0x8a, 0x44, 0x87, 0x5a, 0xcf, 0x7e, 0x79,
	0x74, 0x76, 0x32, 0x32, 0x4a, 0xd2, 0xe9, 0x3b, 0xaf, 0xa8, 0x3d, 0x1c, 0x66, 0x5b, 0x92, 0xda,
	0x2a, 0xd9, 0x03, 0x92, 0x1d, 0x78, 0x47, 0x4e, 0xcf, 0xcb, 0xe2, 0xe5, 0x71, 0x15, 0x07, 0xfe,
	0xf1, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0x91, 0x36, 0x48, 0x3c, 0x06, 0x00, 0x00,
}
