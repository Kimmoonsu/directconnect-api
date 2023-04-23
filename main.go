package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Connections struct {
	_ struct{} `type:"structure"`

	// The connections.
	Connections []*Connection `locationName:"connections" type:"list" json:"connections"`
}

// Information about the MAC Security (MACsec) secret key.
type MacSecKey struct {
	_ struct{} `type:"structure"`

	// The Connection Key Name (CKN) for the MAC Security secret key.
	Ckn *string `locationName:"ckn" type:"string"`

	// The Amazon Resource Name (ARN) of the MAC Security (MACsec) secret key.
	SecretARN *string `locationName:"secretARN" type:"string"`

	// The date that the MAC Security (MACsec) secret key takes effect. The value
	// is displayed in UTC format.
	StartOn *string `locationName:"startOn" type:"string"`

	// The state of the MAC Security (MACsec) secret key.
	//
	// The possible values are:
	//
	//    * associating: The MAC Security (MACsec) secret key is being validated
	//    and not yet associated with the connection or LAG.
	//
	//    * associated: The MAC Security (MACsec) secret key is validated and associated
	//    with the connection or LAG.
	//
	//    * disassociating: The MAC Security (MACsec) secret key is being disassociated
	//    from the connection or LAG
	//
	//    * disassociated: The MAC Security (MACsec) secret key is no longer associated
	//    with the connection or LAG.
	State *string `locationName:"state" type:"string"`
}

type Connection struct {
	_ struct{} `type:"structure"`

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string `locationName:"awsDevice" deprecated:"true" type:"string" json:"awsDevice,omitempty"`

	// The Direct Connect endpoint that terminates the physical connection.
	AwsDeviceV2 *string `locationName:"awsDeviceV2" type:"string" json:"awsDeviceV2,omitempty"`

	// The Direct Connect endpoint that terminates the logical connection. This
	// device might be different than the device that terminates the physical connection.
	AwsLogicalDeviceId *string `locationName:"awsLogicalDeviceId" type:"string" json:"awsLogicalDeviceId,omitempty"`

	// The bandwidth of the connection.
	Bandwidth *string `locationName:"bandwidth" type:"string" json:"bandwidth,omitempty"`

	// The ID of the connection.
	ConnectionId *string `locationName:"connectionId" type:"string" json:"connectionId,omitempty"`

	// The name of the connection.
	ConnectionName *string `locationName:"connectionName" type:"string" json:"connectionName,omitempty"`

	// The state of the connection. The following are the possible values:
	//
	//    * ordering: The initial state of a hosted connection provisioned on an
	//    interconnect. The connection stays in the ordering state until the owner
	//    of the hosted connection confirms or declines the connection order.
	//
	//    * requested: The initial state of a standard connection. The connection
	//    stays in the requested state until the Letter of Authorization (LOA) is
	//    sent to the customer.
	//
	//    * pending: The connection has been approved and is being initialized.
	//
	//    * available: The network link is up and the connection is ready for use.
	//
	//    * down: The network link is down.
	//
	//    * deleting: The connection is being deleted.
	//
	//    * deleted: The connection has been deleted.
	//
	//    * rejected: A hosted connection in the ordering state enters the rejected
	//    state if it is deleted by the customer.
	//
	//    * unknown: The state of the connection is not available.
	ConnectionState *string `locationName:"connectionState" type:"string" enum:"ConnectionState" json:"connectionState,omitempty"`

	// The MAC Security (MACsec) connection encryption mode.
	//
	// The valid values are no_encrypt, should_encrypt, and must_encrypt.
	EncryptionMode *string `locationName:"encryptionMode" type:"string" json:"encryptionMode,omitempty"`

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy *string `locationName:"hasLogicalRedundancy" type:"string" enum:"HasLogicalRedundancy" json:"hasLogicalRedundancy,omitempty"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `locationName:"jumboFrameCapable" type:"boolean" json:"jumboFrameCapable,omitempty"`

	// The ID of the LAG.
	LagId *string `locationName:"lagId" type:"string" json:"lagId,omitempty"`

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time `locationName:"loaIssueTime" type:"timestamp" json:"loaIssueTime,omitempty"`

	// The location of the connection.
	Location *string `locationName:"location" type:"string" json:"location,omitempty"`

	// Indicates whether the connection supports MAC Security (MACsec).
	MacSecCapable *bool `locationName:"macSecCapable" type:"boolean" json:"macSecCapable,omitempty"`

	// The MAC Security (MACsec) security keys associated with the connection.
	MacSecKeys []*MacSecKey `locationName:"macSecKeys" type:"list" json:"macSecKeys,omitempty"`

	// The ID of the Amazon Web Services account that owns the connection.
	OwnerAccount *string `locationName:"ownerAccount" type:"string" json:"ownerAccount,omitempty"`

	// The name of the Direct Connect service provider associated with the connection.
	PartnerName *string `locationName:"partnerName" type:"string" json:"partnerName,omitempty"`

	// The MAC Security (MACsec) port link status of the connection.
	//
	// The valid values are Encryption Up, which means that there is an active Connection
	// Key Name, or Encryption Down.
	PortEncryptionStatus *string `locationName:"portEncryptionStatus" type:"string" json:"portEncryptionsStatus,omitempty"`

	// The name of the service provider associated with the connection.
	ProviderName *string `locationName:"providerName" type:"string" json:"providerName,omitempty"`

	// The Amazon Web Services Region where the connection is located.
	Region *string `locationName:"region" type:"string" json:"region,omitempty"`

	// The tags associated with the connection.
	Tags []*Tag `locationName:"tags" min:"1" type:"list" json:"tags,omitempty"`

	// The ID of the VLAN.
	Vlan *int64 `locationName:"vlan" type:"integer" json:"vlan,omitempty"`
}

// Information about a tag.
type Tag struct {
	_ struct{} `type:"structure"`

	// The key.
	//
	// Key is a required field
	Key *string `locationName:"key" min:"1" type:"string" required:"true"`

	// The value.
	Value *string `locationName:"value" type:"string"`
}

func refString(s string) *string {
	return &s
}

type DescribeConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// A description of the connections.
	Connections []*Connection `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`
}

func main() {
	r := gin.Default()

	connection := Connection{
		AwsDevice:       refString("EqDC2-123h49s71dabc"),
		OwnerAccount:    refString("123456789012"),
		ConnectionId:    refString("dxcon-fguhmqlc"),
		LagId:           refString("dxlag-ffrz71kw"),
		ConnectionState: refString("down"),
		Bandwidth:       refString("1Gbps"),
		Location:        refString("EqDC2"),
		ConnectionName:  refString("My_Connection"),
		// LoaIssueTime:    1491568964.0,
		Region: refString("us-east-1"),
	}
	conn := []*Connection{&connection}
	connections := &Connections{}
	connections.Connections = conn

	r.POST("/", func(c *gin.Context) {
		// Request Header 출력
		fmt.Println("=============HEADER==========")
		header := c.Request.Header
		fmt.Println(header)
		fmt.Println("")

		// Request Body 출력
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			// 에러 처리
		}
		fmt.Println("==============BODY==============")
		fmt.Println(string(body))
		fmt.Println("")

		// Header 출력
		for key, value := range header {
			fmt.Printf("%s: %s\n", key, value)
		}

		targetHeader := c.GetHeader("X-Amz-Target")
		targetArr := strings.Split(targetHeader, ".")
		action := targetArr[len(targetArr)-1]
		fmt.Println("action : ", action) // "DescribeConnections"

		// Response
		c.JSONP(http.StatusOK, connections)
	})

	r.Run(":5000")
}
