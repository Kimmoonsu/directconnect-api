package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Kimmoonsu/directconnect-api/internal/models"
	"github.com/Kimmoonsu/directconnect-api/internal/models/utils"
	httpresponse "github.com/Kimmoonsu/directconnect-api/internal/protocols/response"
)

func (h HttpHandlerImpl) CreateConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateConnection")
}

func (h HttpHandlerImpl) DescribeConnections(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DescribeConnections")
	// 		// Request Header 출력
	fmt.Println("=============HEADER==========")
	header := r.Header
	fmt.Println(header)
	fmt.Println("")

	// 		// Request Body 출력
	body, err := ioutil.ReadAll(r.Body)
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

	// targetHeader := r.GetHeader("X-Amz-Target")
	// targetArr := strings.Split(targetHeader, ".")
	// action := targetArr[len(targetArr)-1]
	// fmt.Println("action : ", action) // "DescribeConnections"

	//		// Response
	//		c.JSONP(http.StatusOK, connections)
	//	})

	connection := models.Connection{
		AwsDevice:       utils.RefString("EqDC2-123h49s71dabc"),
		OwnerAccount:    utils.RefString("123456789012"),
		ConnectionId:    utils.RefString("dxcon-fguhmqlc"),
		LagId:           utils.RefString("dxlag-ffrz71kw"),
		ConnectionState: utils.RefString("down"),
		Bandwidth:       utils.RefString("1Gbps"),
		Location:        utils.RefString("EqDC2"),
		ConnectionName:  utils.RefString("My_Connection"),
		// LoaIssueTime:    1491568964.0,
		Region: utils.RefString("us-east-1"),
	}
	conn := []*models.Connection{&connection}
	connections := &models.Connections{}
	connections.Connections = conn

	httpresponse.Json(w, http.StatusOK, "", connections)
}
