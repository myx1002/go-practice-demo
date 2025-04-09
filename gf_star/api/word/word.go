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
}
