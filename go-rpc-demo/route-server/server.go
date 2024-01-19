package main

import (
	"context"
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	pb "go-rpc-demo/route"
)

type routeGuideServer struct {
	features []*pb.Feature // 假数据
	pb.UnimplementedRouteGuideServer
}

func (r *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range r.features {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return nil, nil
}

func (r *routeGuideServer) ListFeatures(rectangle *pb.Rectangle, server pb.RouteGuide_ListFeaturesServer) error {
	for _, feature := range r.features {
		left := math.Min(float64(rectangle.Lo.Longitude), float64(rectangle.Hi.Longitude))
		right := math.Max(float64(rectangle.Lo.Longitude), float64(rectangle.Hi.Longitude))
		top := math.Max(float64(rectangle.Lo.Latitude), float64(rectangle.Hi.Latitude))
		bottom := math.Min(float64(rectangle.Lo.Latitude), float64(rectangle.Hi.Latitude))
		if float64(feature.Location.Longitude) >= left &&
			float64(feature.Location.Longitude) <= right &&
			float64(feature.Location.Latitude) >= bottom &&
			float64(feature.Location.Latitude) <= top {
			if err := server.Send(feature); err != nil {
				return err
			}
		}
		// 测试：延迟1s
		time.Sleep(time.Second)
	}
	return nil
}

func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

// 计算两个点的距离
func calcDistance(p1 *pb.Point, p2 *pb.Point) int64 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int64(distance)
}

func (r *routeGuideServer) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
	startTime := time.Now()        // 初始化时间
	var pointCount, distance int64 // 初始化点的数量和距离
	var prevPoint *pb.Point        // 初始化第一个点

	for {
		point, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount:  pointCount,
				Distance:    distance,
				ElapsedTime: int64(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}

		pointCount++
		if prevPoint != nil {
			distance += calcDistance(prevPoint, point)
		}
		prevPoint = point
	}

	return nil
}

// 获取离这个点最近/最远的feature
func (r *routeGuideServer) recommendOnce(request *pb.RecommendationRequest) (*pb.Feature, error) {
	var nearest, farthest *pb.Feature           // 定义最近/最远feature
	var nearestDistance, farthestDistance int64 // 定义最近/最远距离

	for _, feature := range r.features {
		// 计算两个点的距离
		distance := calcDistance(feature.Location, request.Point)
		// 如果最近feature为空 || 当前两个点的距离小于当前最近距离
		if nearest == nil || distance < nearestDistance {
			nearestDistance = distance
			nearest = feature
		}
		if farthest == nil || distance > farthestDistance {
			farthestDistance = distance
			farthest = feature
		}
	}

	if request.Mode == pb.RecommendationMode_GetFarthest {
		return farthest, nil
	}
	return nearest, nil
}

func (r *routeGuideServer) Recommend(stream pb.RouteGuide_RecommendServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// 处理请求
		recommended, err := r.recommendOnce(request)
		if err != nil {
			return err
		}
		// 发送结果
		err = stream.Send(recommended)
		if err != nil {
			return err
		}
	}
}

// 定义一个routeGuideServer类型去实现pb.RouteGuideServer的所有方法，即函数签名一至
func newServer() pb.RouteGuideServer {
	return &routeGuideServer{
		features: []*pb.Feature{
			{Name: "上海交通大学闵行校区上海市闵行区东川路800号", Location: &pb.Point{
				Latitude:  310235000,
				Longitude: 121437403,
			}},
			{Name: "复旦大学上海市杨浦区五角场邯郸路220号", Location: &pb.Point{
				Latitude:  312978870,
				Longitude: 121503457,
			}},
			{Name: "华东理工大学上海市徐汇区梅陇路130号", Location: &pb.Point{
				Latitude:  311416130,
				Longitude: 121424904,
			}},
		},
	}
}

func main() {
	// 新建一个tcp监听服务
	listen, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
	}

	// 新建一个rpc服务
	grpcServer := grpc.NewServer()
	// 注册RouteGuideServer，即把RouteGuideServer注册到grpcServer中
	pb.RegisterRouteGuideServer(grpcServer, newServer())
	// 启动服务
	log.Fatalln(grpcServer.Serve(listen))
}
