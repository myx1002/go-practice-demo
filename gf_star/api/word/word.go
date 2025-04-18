// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package word

import (
	"context"

	"gf_star/api/word/v1"
)

type IWordV1 interface {
	WordCreate(ctx context.Context, req *v1.WordCreateReq) (res *v1.WordCreateRes, err error)
	WordUpdate(ctx context.Context, req *v1.WordUpdateReq) (res *v1.WordUpdateRes, err error)
	WordDelete(ctx context.Context, req *v1.WordDeleteReq) (res *v1.WordDeleteRes, err error)
	WordDetail(ctx context.Context, req *v1.WordDetailReq) (res *v1.WordDetailRes, err error)
	WordList(ctx context.Context, req *v1.WordListReq) (res *v1.WordListRes, err error)
	WordRandomList(ctx context.Context, req *v1.WordRandomListReq) (res *v1.WordRandomListRes, err error)
	SetLevel(ctx context.Context, req *v1.SetLevelReq) (res *v1.SetLevelRes, err error)
}
