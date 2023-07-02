package main

import (
	"github.com/Kimmoonsu/directconnect-api/internal/logger"
)

// func main() {
// 	r := gin.Default()

// 	connection := Connection{
// 		AwsDevice:       refString("EqDC2-123h49s71dabc"),
// 		OwnerAccount:    refString("123456789012"),
// 		ConnectionId:    refString("dxcon-fguhmqlc"),
// 		LagId:           refString("dxlag-ffrz71kw"),
// 		ConnectionState: refString("down"),
// 		Bandwidth:       refString("1Gbps"),
// 		Location:        refString("EqDC2"),
// 		ConnectionName:  refString("My_Connection"),
// 		// LoaIssueTime:    1491568964.0,
// 		Region: refString("us-east-1"),
// 	}
// 	conn := []*Connection{&connection}
// 	connections := &Connections{}
// 	connections.Connections = conn

// 	r.POST("/", func(c *gin.Context) {
// 		// Request Header 출력
// 		fmt.Println("=============HEADER==========")
// 		header := c.Request.Header
// 		fmt.Println(header)
// 		fmt.Println("")

// 		// Request Body 출력
// 		body, err := ioutil.ReadAll(c.Request.Body)
// 		if err != nil {
// 			// 에러 처리
// 		}
// 		fmt.Println("==============BODY==============")
// 		fmt.Println(string(body))
// 		fmt.Println("")

// 		// Header 출력
// 		for key, value := range header {
// 			fmt.Printf("%s: %s\n", key, value)
// 		}

// 		targetHeader := c.GetHeader("X-Amz-Target")
// 		targetArr := strings.Split(targetHeader, ".")
// 		action := targetArr[len(targetArr)-1]
// 		fmt.Println("action : ", action) // "DescribeConnections"

// 		// Response
// 		c.JSONP(http.StatusOK, connections)
// 	})

// 	r.Run(":5000")
// }

func main() {
	logger.InitLogger()

	initProtocol := InitHttpProtocol()

	initProtocol.Listen()

}
