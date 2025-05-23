// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package word

import (
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"proxima/app/gateway/api/word"
	"proxima/app/gateway/utility"
	v1 "proxima/app/word/api/words/v1"
)

type ControllerV1 struct {
	WordClient v1.WordsClient
}

func NewV1() word.IWordV1 {
	return &ControllerV1{
		WordClient: v1.NewWordsClient(grpcx.Client.MustNewGrpcClientConn("word", grpcx.Client.ChainUnary(
			utility.GrpcClientTimeout,
		))),
	}
}
