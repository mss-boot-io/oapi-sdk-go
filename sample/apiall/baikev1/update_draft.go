// Package baike code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"context"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/baike/v1"
)

// HTTP PATH: /open-apis/baike/v1/drafts/:draft_id"
func main() {
	// 创建 Client
	client := lark.NewClient("appID", "appSecret")
	// 创建请求对象
	req := larkbaike.NewUpdateDraftReqBuilder().
		DraftId("42322").
		UserIdType("user_id").
		Entity(larkbaike.NewEntityBuilder().
			Id("enterprise_40217521").
			MainKeys([]*larkbaike.Term{larkbaike.NewTermBuilder().Build()}).
			Aliases([]*larkbaike.Term{larkbaike.NewTermBuilder().Build()}).
			Description("企业百科是飞书提供的一款知识管理工具，通过企业百科可以帮助企业将分散的知识信息进行聚合，并通过UGC的方式，促进企业知识的保鲜和流通").
			RelatedMeta(larkbaike.NewRelatedMetaBuilder().Build()).
			RichText("").
			Build()).
		Build()
	// 发起请求
	resp, err := client.Baike.Draft.Update(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理处理
	fmt.Println(larkcore.Prettify(resp))
}
