/*
 * BitTrace
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

import (
	"github.com/BitTraceProject/BitTrace-Types/pkg/database"
)

type (
	ResponseBestState struct {

		// OK，是否查询成功
		Ok bool `json:"ok"`

		// 消息，返回的消息
		Msg string `json:"msg"`

		// 数目，返回结果的数目
		Number int32 `json:"number"`

		// 数据，返回的数据
		Data []BestState `json:"data"`
	}
	// BestState - 某一时间点的区块链系统状态
	BestState database.TableState
)