// @Title
// @Description
// @Author  Wangwengang  2021/8/19 下午3:36
// @Update  Wangwengang  2021/8/19 下午3:36
package proto

//go:generate protoc -I. --go_out=plugins=rpcx:. --go_opt=paths=source_relative pigeon/room.proto
//go:generate protoc -I. --go_out=plugins=rpcx:. --go_opt=paths=source_relative pigeon/pigeon.proto
//go:generate protoc -I. --go_out=plugins=rpcx:. --go_opt=paths=source_relative pigeon/protocol.proto
//go:generate protoc -I. --go_out=plugins=rpcx:. --go_opt=paths=source_relative identity/identity.proto
//go:generate protoc -I. --go_out=plugins=rpcx:. --go_opt=paths=source_relative common/common.proto
