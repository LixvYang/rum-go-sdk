package main

import (
	"context"

	"github.com/antihax/optional"
	"github.com/lixvyang/rum-sdk-go/sdk/base"
	"github.com/lixvyang/rum-sdk-go/sdk/groups"
	"github.com/lixvyang/rum-sdk-go/sdk/model"
	"github.com/lixvyang/rum-sdk-go/service/fullnode"
)

/*
1. 创建一个测试组 CreateGroupUrl
   1.1 创建一些组  GetGroups
   1.2 加入别的组 JoinGroup
   1.3 获取组种子 Get group seed
2. 发表 11 条内容 PostToGroup
3. 查询block获取内容 GetBlock
3. 查询TrxId获取内容 GetTrx
4. 通过API查看组信息当前 GetGroupById
5. 离开组 LeaveGroup
6. 清除组数据 ClearGroupData

*/

func main() {
	cli := fullnode.New(&base.Configuration{
		BasePath: "http://127.0.0.1:8000",
	})
	ctx := context.Background()
	cli.GroupsApi.ApiV1GroupPost(ctx, &groups.GroupsApiApiV1GroupPostOpts{
		optional.NewInterface(model.HandlersCreateGroupParam{
			ConsensusType: "POA",
			
		}),
	})
}
