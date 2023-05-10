/*
 * BitTrace
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/BitTraceProject/BitTrace-OpenAPI/server/model"
	"github.com/BitTraceProject/BitTrace-Types/pkg/structure"

	"github.com/gin-gonic/gin"
)

type IndexType string

const (
	IndexTypeBySID  IndexType = "by_snapshot_id"
	IndexTypeByTRA  IndexType = "by_timestamp_range"
	IndexTypeByBHA  IndexType = "by_block_hash"
	IndexTypeByCIN  IndexType = "by_chain_info"
	IndexTypeByPBHA IndexType = "by_parent_block_hash"
)

var (
	ErrParamInvalid = errors.New("error param invalid")
	ErrParamNotSet  = errors.New("error param not set")
	ErrQueryFailed  = errors.New("error query failed")
	ErrUIDInvalid   = errors.New("error uid invalid")
	ErrTokenInvalid = errors.New("error token invalid")
	ErrUnknown      = errors.New("error unknown")
)

// OpenapiV1Index is the index handler.
func OpenapiV1Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

// OpenapiV1SnapshotGet - Snapshot
func OpenapiV1SnapshotGet(c *gin.Context) {
	var (
		respPtr = &model.ResponseSnapshot{
			Ok:     false,
			Msg:    "",
			Number: 0,
			Data:   nil,
		}
		respCode = http.StatusOK
	)
	defer func() {
		c.JSON(respCode, *respPtr)
	}()

	token, ok := c.GetQuery("token")
	if !ok || token == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	if !checkToken(token) {
		respPtr.Ok = false
		respPtr.Msg = ErrTokenInvalid.Error()
		return
	}
	peerTag, ok := c.GetQuery("peer_tag")
	if !ok || peerTag == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	// 获取索引方式
	var indexTypes = map[IndexType]bool{
		IndexTypeBySID: false,
		IndexTypeByTRA: false,
		IndexTypeByBHA: false,
		IndexTypeByCIN: false,
	}
	if t := c.Query(string(IndexTypeBySID)); t != "" && t != "0" {
		if t == "1" {
			indexTypes[IndexTypeBySID] = true
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}
	if t := c.Query(string(IndexTypeByTRA)); t != "" && t != "0" {
		if t == "1" {
			indexTypes[IndexTypeByTRA] = true
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}
	if t := c.Query(string(IndexTypeByBHA)); t != "" && t != "0" {
		if t == "1" {
			indexTypes[IndexTypeByBHA] = true
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}
	if t := c.Query(string(IndexTypeByCIN)); t != "" && t != "0" {
		if t == "1" {
			indexTypes[IndexTypeByCIN] = true
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}

	// 选取当前索引方式
	// 优先级：by_snapshot_id > by_timestamp_range
	// > by_block_hash > by_chain_info
	var indexType = IndexTypeBySID // default BySID
	if indexTypes[IndexTypeBySID] {
		indexType = IndexTypeBySID
	} else {
		if indexTypes[IndexTypeByTRA] {
			indexType = IndexTypeByTRA
		} else {
			if indexTypes[IndexTypeByBHA] {
				indexType = IndexTypeByBHA
			} else {
				if indexTypes[IndexTypeByCIN] {
					indexType = IndexTypeByCIN
				} else {
					indexType = IndexTypeBySID // default BySID
				} // end if
			} // end if
		} // end if
	} // end if

	// 根据索引方式获取参数并且读取 snapshot 返回响应
	var (
		pageSize   = 20
		pageNumber = 1
	)
	if pSizeStr := c.Query("page_size"); pSizeStr != "" {
		if pSize, err := strconv.Atoi(pSizeStr); err != nil {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		} else {
			pageSize = pSize
		}
	}
	if pNumberStr := c.Query("page_number"); pNumberStr != "" {
		if pNumber, err := strconv.Atoi(pNumberStr); err != nil {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		} else {
			pageNumber = pNumber
		}
	}
	openapiV1SnapshotGet(c, peerTag, indexType, pageSize, pageNumber, respPtr)
}

// OpenapiV1RevisionGet - Revision
func OpenapiV1RevisionGet(c *gin.Context) {
	var (
		respPtr = &model.ResponseRevision{
			Ok:     false,
			Msg:    "",
			Number: 0,
			Data:   nil,
		}
		respCode = http.StatusOK
	)
	defer func() {
		c.JSON(respCode, *respPtr)
	}()

	token, ok := c.GetQuery("token")
	if !ok || token == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	if !checkToken(token) {
		respPtr.Ok = false
		respPtr.Msg = ErrTokenInvalid.Error()
		return
	}
	peerTag, ok := c.GetQuery("peer_tag")
	if !ok || peerTag == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	snapshotID := c.Query("snapshot_id")
	if snapshotID == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	var (
		revisionType = int(structure.RevisionTypeUnknown) // RevisionTypeUnknown 可用来代表不限制 revision type
		err          error
	)
	if rTypeStr := c.Query("revision_type"); rTypeStr != "" {
		if revisionType, err = strconv.Atoi(rTypeStr); err != nil {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}
	if revisionType < 0 || revisionType > int(structure.RevisionTypeUnknown) {
		respPtr.Ok = false
		respPtr.Msg = ErrParamInvalid.Error()
		return
	}
	initTimestamp := c.Query("init_timestamp")
	revisionData, err := QueryRevisionBySID(peerTag, snapshotID, revisionType, initTimestamp)
	if err != nil {
		respPtr.Ok = false
		respPtr.Msg = fmt.Errorf("%s:%s", ErrQueryFailed.Error(), err.Error()).Error()
		return
	}
	respPtr.Ok = true
	respPtr.Number = int32(len(revisionData))
	respPtr.Data = revisionData
}

// OpenapiV1BestStateGet - BestState
func OpenapiV1BestStateGet(c *gin.Context) {
	var (
		respPtr = &model.ResponseBestState{
			Ok:     false,
			Msg:    "",
			Number: 0,
			Data:   nil,
		}
		respCode = http.StatusOK
	)
	defer func() {
		c.JSON(respCode, *respPtr)
	}()

	token, ok := c.GetQuery("token")
	if !ok || token == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	if !checkToken(token) {
		respPtr.Ok = false
		respPtr.Msg = ErrTokenInvalid.Error()
		return
	}
	peerTag, ok := c.GetQuery("peer_tag")
	if !ok || peerTag == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	snapshotID := c.Query("snapshot_id")
	if snapshotID == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	var (
		snapshotType = int(structure.SnapshotTypeUnknown) // SnapshotTypeUnknown 可用来代表不限制 snapshot type
		err          error
	)
	if sTypeStr := c.Query("snapshot_type"); sTypeStr != "" {
		if snapshotType, err = strconv.Atoi(sTypeStr); err != nil {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}
	if snapshotType < 0 || snapshotType > int(structure.SnapshotTypeUnknown) {
		respPtr.Ok = false
		respPtr.Msg = ErrParamInvalid.Error()
		return
	}
	bestStateData, err := QueryBestStateBySID(peerTag, snapshotID, snapshotType)
	if err != nil {
		respPtr.Ok = false
		respPtr.Msg = fmt.Errorf("%s:%s", ErrQueryFailed.Error(), err.Error()).Error()
		return
	}
	respPtr.Ok = true
	respPtr.Number = int32(len(bestStateData))
	respPtr.Data = bestStateData
}

// OpenapiV1EventOrphanGet - Event_Orphan
func OpenapiV1EventOrphanGet(c *gin.Context) {
	var (
		respPtr = &model.ResponseEventOrphan{
			Ok:     false,
			Msg:    "",
			Number: 0,
			Data:   nil,
		}
		respCode = http.StatusOK
	)
	defer func() {
		c.JSON(respCode, *respPtr)
	}()

	token, ok := c.GetQuery("token")
	if !ok || token == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	if !checkToken(token) {
		respPtr.Ok = false
		respPtr.Msg = ErrTokenInvalid.Error()
		return
	}
	peerTag, ok := c.GetQuery("peer_tag")
	if !ok || peerTag == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	// 获取索引方式
	var indexTypes = map[IndexType]bool{
		IndexTypeBySID:  false,
		IndexTypeByBHA:  false,
		IndexTypeByPBHA: false,
	}
	if t := c.Query(string(IndexTypeBySID)); t != "" && t != "0" {
		if t == "1" {
			indexTypes[IndexTypeBySID] = true
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}

	if t := c.Query(string(IndexTypeByBHA)); t != "" && t != "0" {
		if t == "1" {
			indexTypes[IndexTypeByBHA] = true
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}

	if t := c.Query(string(IndexTypeByPBHA)); t != "" && t != "0" {
		if t == "1" {
			indexTypes[IndexTypeByPBHA] = true
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}

	// 选取当前索引方式
	// 优先级：by_snapshot_id > by_timestamp_range
	// > by_block_hash > by_chain_info
	var indexType = IndexTypeBySID // default BySID
	if indexTypes[IndexTypeBySID] {
		indexType = IndexTypeBySID
	} else {
		if indexTypes[IndexTypeByBHA] {
			indexType = IndexTypeByBHA
		} else {
			if indexTypes[IndexTypeByPBHA] {
				indexType = IndexTypeByPBHA
			} else {
				indexType = IndexTypeBySID // default BySID
			} // end if
		} // end if
	} // end if

	switch indexType {
	case IndexTypeBySID:
		snapshotID := c.Query("snapshot_id")
		if snapshotID == "" {
			respPtr.Ok = false
			respPtr.Msg = ErrParamNotSet.Error()
			return
		}

		var (
			eventTypeOrphan = int(structure.EventTypeUnknown) // EventTypeUnknown 可用来代表不限制 event type
			err             error
		)
		if eTypeStr := c.Query("event_type_orphan"); eTypeStr != "" {
			if eventTypeOrphan, err = strconv.Atoi(eTypeStr); err != nil {
				respPtr.Ok = false
				respPtr.Msg = ErrParamInvalid.Error()
				return
			}
		}
		if eventTypeOrphan < 0 || eventTypeOrphan > int(structure.EventTypeUnknown) {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
		eventDataOrphan, err := QueryEventOrphanBySID(peerTag, snapshotID, eventTypeOrphan)
		if err != nil {
			respPtr.Ok = false
			respPtr.Msg = fmt.Errorf("%s:%s", ErrQueryFailed.Error(), err.Error()).Error()
			return
		}
		respPtr.Ok = true
		respPtr.Number = int32(len(eventDataOrphan))
		respPtr.Data = eventDataOrphan
	case IndexTypeByBHA:
		blockHash := c.Query("block_hash")
		if blockHash == "" {
			respPtr.Ok = false
			respPtr.Msg = ErrParamNotSet.Error()
			return
		}
		eventDataOrphan, err := QueryEventOrphanByBHA(peerTag, blockHash)
		if err != nil {
			respPtr.Ok = false
			respPtr.Msg = fmt.Errorf("%s:%s", ErrQueryFailed.Error(), err.Error()).Error()
			return
		}
		respPtr.Ok = true
		respPtr.Number = int32(len(eventDataOrphan))
		respPtr.Data = eventDataOrphan
	case IndexTypeByPBHA:
		parentBlockHash := c.Query("parent_block_hash")
		if parentBlockHash == "" {
			respPtr.Ok = false
			respPtr.Msg = ErrParamNotSet.Error()
			return
		}
		eventDataOrphan, err := QueryEventOrphanByPBHA(peerTag, parentBlockHash)
		if err != nil {
			respPtr.Ok = false
			respPtr.Msg = fmt.Errorf("%s:%s", ErrQueryFailed.Error(), err.Error()).Error()
			return
		}
		respPtr.Ok = true
		respPtr.Number = int32(len(eventDataOrphan))
		respPtr.Data = eventDataOrphan
	default:
		respPtr.Ok = false
		respPtr.Msg = ErrUnknown.Error()
		return
	}
}

// OpenapiV1TimelineGet - Timeline
func OpenapiV1TimelineGet(c *gin.Context) {
	var (
		respPtr = &model.ResponseTimePoint{
			Ok:     false,
			Msg:    "",
			Number: 0,
			Data:   nil,
		}
		respCode = http.StatusOK
	)
	defer func() {
		c.JSON(respCode, *respPtr)
	}()

	token, ok := c.GetQuery("token")
	if !ok || token == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	if !checkToken(token) {
		respPtr.Ok = false
		respPtr.Msg = ErrTokenInvalid.Error()
		return
	}
	peerTagArray, ok := c.GetQueryArray("peer_tag")
	if !ok || len(peerTagArray) == 0 {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	// 获取索引方式
	var indexTypes = map[IndexType]bool{
		IndexTypeBySID: false,
		IndexTypeByTRA: false,
		IndexTypeByBHA: false,
		IndexTypeByCIN: false,
	}
	if t := c.Query(string(IndexTypeBySID)); t != "" && t != "0" {
		if t == "1" {
			indexTypes[IndexTypeBySID] = true
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}
	if t := c.Query(string(IndexTypeByTRA)); t != "" && t != "0" {
		if t == "1" {
			indexTypes[IndexTypeByTRA] = true
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}

	if t := c.Query(string(IndexTypeByBHA)); t != "" && t != "0" {
		if t == "1" {
			indexTypes[IndexTypeByBHA] = true
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}

	if t := c.Query(string(IndexTypeByCIN)); t != "" && t != "0" {
		if t == "1" {
			indexTypes[IndexTypeByCIN] = true
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
	}

	// 选取当前索引方式
	// 优先级：by_snapshot_id > by_timestamp_range
	// > by_block_hash > by_chain_info
	var indexType = IndexTypeBySID // default BySID
	if indexTypes[IndexTypeBySID] {
		indexType = IndexTypeBySID
	} else {
		if indexTypes[IndexTypeByTRA] {
			indexType = IndexTypeByTRA
		} else {
			if indexTypes[IndexTypeByBHA] {
				indexType = IndexTypeByBHA
			} else {
				if indexTypes[IndexTypeByCIN] {
					indexType = IndexTypeByCIN
				} else {
					indexType = IndexTypeBySID // default BySID
				} // end if
			} // end if
		} // end if
	} // end if

	// 根据索引方式获取参数并且读取 snapshot 返回响应
	var (
		pageSize   = 20
		pageNumber = 1
	)
	if pSizeStr := c.Query("page_size"); pSizeStr != "" {
		if pSize, err := strconv.Atoi(pSizeStr); err != nil {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		} else {
			pageSize = pSize
		}
	}
	if pNumberStr := c.Query("page_number"); pNumberStr != "" {
		if pNumber, err := strconv.Atoi(pNumberStr); err != nil {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		} else {
			pageNumber = pNumber
		}
	}
	log.Println(pageSize, pageNumber)

	// TODO 这里加上对 peer tag 是否存在的预检查，减少对数据库的调用，需要数据库中有存储 peer 元信息的支持（目前不支持）
	var (
		timePointData = []model.TimePoint{}
	)
	for _, peerTag := range peerTagArray {
		peerRespPtr := &model.ResponseSnapshot{
			Ok:     false,
			Msg:    "",
			Number: 0,
			Data:   nil,
		}
		openapiV1SnapshotGet(c, peerTag, indexType, pageSize, pageNumber, peerRespPtr)
		if !peerRespPtr.Ok {
			respPtr.Ok = false
			respPtr.Msg = peerRespPtr.Msg
			return
		}
		for _, sData := range peerRespPtr.Data {
			timePoint := model.TimePoint{
				Tag:            peerTag,
				SnapshotId:     sData.SnapshotID,
				InitTimestamp:  sData.InitTimestamp,
				FinalTimestamp: sData.FinalTimestamp,
			}
			timePointData = append(timePointData, timePoint)
		}
	}
	// 排序，倒序返回
	sort.Slice(timePointData, func(i, j int) bool {
		return timePointData[i].InitTimestamp < timePointData[j].InitTimestamp
	})

	respPtr.Ok = true
	respPtr.Number = int32(len(timePointData))
	respPtr.Data = timePointData
}

// OpenapiV1PeerGet - Peer
func OpenapiV1PeerGet(c *gin.Context) {
	var (
		respPtr = &model.ResponsePeer{
			Ok:     false,
			Msg:    "",
			Number: 0,
			Data:   nil,
		}
		respCode = http.StatusOK
	)
	defer func() {
		c.JSON(respCode, *respPtr)
	}()

	token, ok := c.GetQuery("token")
	if !ok || token == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	if !checkToken(token) {
		respPtr.Ok = false
		respPtr.Msg = ErrTokenInvalid.Error()
		return
	}
	peerData, err := QueryPeer()
	if err != nil {
		respPtr.Ok = false
		respPtr.Msg = fmt.Errorf("%s:%s", ErrQueryFailed.Error(), err.Error()).Error()
		return
	}
	respPtr.Ok = true
	respPtr.Number = int32(len(peerData))
	respPtr.Data = peerData
}

func openapiV1SnapshotGet(c *gin.Context, peerTag string, indexType IndexType, pageSize, pageNumber int, respPtr *model.ResponseSnapshot) {
	if peerTag == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}

	switch indexType {
	case IndexTypeBySID:
		snapshotID := c.Query("snapshot_id")
		if snapshotID == "" {
			respPtr.Ok = false
			respPtr.Msg = ErrParamNotSet.Error()
			return
		}
		snapshotData, err := QuerySnapshotBySID(peerTag, snapshotID, pageSize, pageNumber)
		if err != nil {
			respPtr.Ok = false
			respPtr.Msg = fmt.Errorf("%s:%s", ErrQueryFailed.Error(), err.Error()).Error()
			return
		}
		respPtr.Ok = true
		respPtr.Number = int32(len(snapshotData))
		respPtr.Data = snapshotData
	case IndexTypeByTRA:
		leftTimestamp := c.Query("left_timestamp")
		rightTimestamp := c.Query("right_timestamp")
		// precheck
		if leftTimestamp == "" {
			leftTimestamp = "0"
		}
		if rightTimestamp == "" {
			rightTimestamp = "0"
		}
		if rightTimestamp < leftTimestamp {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
		snapshotData, err := QuerySnapshotByTRA(peerTag, leftTimestamp, rightTimestamp, pageSize, pageNumber)
		if err != nil {
			respPtr.Ok = false
			respPtr.Msg = fmt.Errorf("%s:%s", ErrQueryFailed.Error(), err.Error()).Error()
			return
		}
		respPtr.Ok = true
		respPtr.Number = int32(len(snapshotData))
		respPtr.Data = snapshotData
	case IndexTypeByBHA:
		blockHash := c.Query("block_hash")
		if blockHash == "" {
			respPtr.Ok = false
			respPtr.Msg = ErrParamNotSet.Error()
			return
		}
		snapshotData, err := QuerySnapshotByBHA(peerTag, blockHash, pageSize, pageNumber)
		if err != nil {
			respPtr.Ok = false
			respPtr.Msg = fmt.Errorf("%s:%s", ErrQueryFailed.Error(), err.Error()).Error()
			return
		}
		respPtr.Ok = true
		respPtr.Number = int32(len(snapshotData))
		respPtr.Data = snapshotData
	case IndexTypeByCIN:
		var (
			targetChainID     int
			targetChainHeight int64
			err               error
		)
		if targetChainIDStr := c.Query("target_chain_id"); targetChainIDStr != "" {
			if targetChainID, err = strconv.Atoi(targetChainIDStr); err != nil {
				respPtr.Ok = false
				respPtr.Msg = ErrParamInvalid.Error()
				return
			}
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamNotSet.Error()
			return
		}
		if targetChainID < -1 {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
		if targetChainHeightStr := c.Query("target_chain_height"); targetChainHeightStr != "" {
			if targetChainHeight, err = strconv.ParseInt(targetChainHeightStr, 10, 32); err != nil {
				respPtr.Ok = false
				respPtr.Msg = ErrParamInvalid.Error()
				return
			}
		} else {
			respPtr.Ok = false
			respPtr.Msg = ErrParamNotSet.Error()
			return
		}
		if targetChainHeight < 0 {
			respPtr.Ok = false
			respPtr.Msg = ErrParamInvalid.Error()
			return
		}
		snapshotData, err := QuerySnapshotByCIN(peerTag, targetChainID, int32(targetChainHeight), pageSize, pageNumber)
		if err != nil {
			respPtr.Ok = false
			respPtr.Msg = fmt.Errorf("%s:%s", ErrQueryFailed.Error(), err.Error()).Error()
			return
		}
		respPtr.Ok = true
		respPtr.Number = int32(len(snapshotData))
		respPtr.Data = snapshotData
	default:
		respPtr.Ok = false
		respPtr.Msg = ErrUnknown.Error()
	}
	return
}

// OpenapiV1AuthRegisterPost TODO
func OpenapiV1AuthRegisterPost(c *gin.Context) {

}

// OpenapiV1AuthTokenPost TODO
func OpenapiV1AuthTokenPost(c *gin.Context) {

}

// OpenapiV1AuthTokenGet 目前暂时用这个
func OpenapiV1AuthTokenGet(c *gin.Context) {
	var (
		respPtr = &model.ResponseAuthToken{
			Ok:    false,
			Msg:   "",
			Token: "",
		}
		respCode = http.StatusOK
	)
	defer func() {
		c.JSON(respCode, *respPtr)
	}()

	// TODO 目前使用 get 方式，后面要完成 post 方式
	uid, ok := c.GetQuery("uid")
	if !ok || uid == "" {
		respPtr.Ok = false
		respPtr.Msg = ErrParamNotSet.Error()
		return
	}
	if !checkUid(uid) {
		respPtr.Ok = false
		respPtr.Msg = ErrUIDInvalid.Error()
		return
	}
	token := newToken(uid)
	respPtr.Ok = true
	respPtr.Token = token
}
