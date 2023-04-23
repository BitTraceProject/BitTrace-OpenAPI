/*
 * BitTrace
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

type (
	ResponseTimePoint struct {

		// OK，是否查询成功
		Ok bool `json:"ok"`

		// 消息，返回的消息
		Msg string `json:"msg"`

		// 数目，返回结果的数目
		Number int32 `json:"number"`

		// 数据，返回的数据
		Data []TimePoint `json:"data"`
	}
	// TimePoint - 某一区块同步过程所在的时间点
	TimePoint struct {

		// 所属对等节点的唯一标识
		Tag string `json:"tag"`

		// Snapshot的唯一ID
		SnapshotId string `json:"snapshot_id"`

		// Snapshot表示的区块同步过程的起始时间
		InitTimestamp string `json:"init_timestamp"`

		// Snapshot表示的区块同步过程的结束时间
		FinalTimestamp string `json:"final_timestamp"`
	}
)