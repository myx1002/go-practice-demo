package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "go-rpc-demo/route"
)

func runFirst(client pb.RouteGuideClient) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  312978870,
		Longitude: 121503457,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}

func runSecond(client pb.RouteGuideClient) {
	// 传入两个点，返回这两个点之间的Feature
	serverStream, err := client.ListFeatures(context.Background(), &pb.Rectangle{
		Lo: &pb.Point{
			Latitude:  312978870,
			Longitude: 121503457,
		},
		Hi: &pb.Point{
			Latitude:  311416130,
			Longitude: 121424904,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		feature, err := serverStream.Recv()
		// 输出完成
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(feature)
	}
}

func runThird(client pb.RouteGuideClient) {
	points := []*pb.Point{
		{Latitude: 310235000, Longitude: 121437403},
		{Latitude: 312978870, Longitude: 121503457},
		{Latitude: 311416130, Longitude: 121424904},
	}

	clientStream, err := client.RecordRoute(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// 每隔一秒上传一个点
	for _, point := range points {
		if err := clientStream.Send(point); err != nil {
			log.Fatalln(err)
		}
		// 测试：延迟1s
		time.Sleep(time.Second)
	}

	// 发送完毕后等待接收和关闭
	summary, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(summary)
}

func readIntFromCommandLine(reader *bufio.Reader, target *int64) {
	_, err := fmt.Fscanf(reader, "%d\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

func runForth(client pb.RouteGuideClient) {
	stream, err := client.Recommend(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// 启动一个goroutine异步接收服务端的流信息
	go func() {
		for {
			feature, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("Recommended", feature)
		}
	}()

	// 通过命令行形式输入mode和经纬度
	reader := bufio.NewReader(os.Stdin)

	// 客户端不断输入point流信息给服务端
	for {
		request := pb.RecommendationRequest{Point: new(pb.Point)}
		var mode int64
		fmt.Print("Enter Recommendation Mode (0 for farthest, 1 for the nearest)：")
		readIntFromCommandLine(reader, &mode)
		fmt.Print("Enter Latitude: ")
		readIntFromCommandLine(reader, &request.Point.Latitude)
		fmt.Print("Enter Longitude: ")
		readIntFromCommandLine(reader, &request.Point.Longitude)
		request.Mode = pb.RecommendationMode(mode)
		// 发送请求到服务端
		if err := stream.Send(&request); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// 相当于拨号请求，
	// 然后使用WithInsecure忽略掉TLS/SSL加密，方便测试
	// WithBlock就是拨号异常时，停止程序，不继续往下走
	dial, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc server")
	}
	defer dial.Close()

	// 新建一个客户端，传入连接请求
	client := pb.NewRouteGuideClient(dial)

	// 执行rpc第一个方法：GetFeature
	fmt.Println("执行rpc第一个方法：GetFeature")
	runFirst(client)

	// 执行rpc第二个方法：ListFeatures
	fmt.Println("执行rpc第二个方法：ListFeatures")
	runSecond(client)

	// 执行rpc第三个方法：RecordRoute
	fmt.Println("执行rpc第三个方法：RecordRoute")
	runThird(client)

	// 执行rpc第四个方法：Recommend
	fmt.Println("执行rpc第四个方法：Recommend")
	runForth(client)
}
