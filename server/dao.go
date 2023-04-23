package server

import (
	"fmt"
	"log"
	"strings"

	"github.com/BitTraceProject/BitTrace-OpenAPI/server/model"
	"github.com/BitTraceProject/BitTrace-Types/pkg/constants"
	"github.com/BitTraceProject/BitTrace-Types/pkg/structure"
)

// 都按照默认顺序返回

func QuerySnapshotBySID(peerTag string, snapshotID string, pageSize, pageNumber int) ([]model.Snapshot, error) {
	var (
		snapshotTableName = strings.Join([]string{constants.TABLE_SNAPSHOT_DATA_PREFIX, peerTag}, constants.DEFAULT_SEP_SYMBOL)
		sql               = fmt.Sprintf("select * from `%s` where snapshot_id='%s' limit %d offset %d", snapshotTableName, snapshotID, pageSize, (pageNumber-1)*pageSize)
		data              = make([]model.Snapshot, 0)
	)
	db, err := oas.Raw(sql)
	if err != nil {
		log.Printf("[QuerySnapshotBySID]raw got error:%v", err)
		return nil, ErrQueryFailed
	}
	db = db.Scan(&data)
	if db.Error != nil {
		log.Printf("[QuerySnapshotBySID]scan got error:%v", db.Error)
		return nil, ErrQueryFailed
	}
	return data, nil
}

func QuerySnapshotByTRA(peerTag string, leftTimestamp, rightTimestamp string, pageSize, pageNumber int) ([]model.Snapshot, error) {
	var (
		snapshotTableName = strings.Join([]string{constants.TABLE_SNAPSHOT_DATA_PREFIX, peerTag}, constants.DEFAULT_SEP_SYMBOL)
		sql               = fmt.Sprintf("select * from `%s` where init_timestamp between '%s' and '%s' limit %d offset %d", snapshotTableName, leftTimestamp, rightTimestamp, pageSize, (pageNumber-1)*pageSize)
		data              = make([]model.Snapshot, 0)
	)
	db, err := oas.Raw(sql)
	if err != nil {
		log.Printf("[QuerySnapshotByTRA]raw got error:%v", err)
		return nil, ErrQueryFailed
	}
	db = db.Scan(&data)
	if db.Error != nil {
		log.Printf("[QuerySnapshotByTRA]scan got error:%v", db.Error)
		return nil, ErrQueryFailed
	}
	return data, nil
}

func QuerySnapshotByBHA(peerTag string, blockHash string, pageSize, pageNumber int) ([]model.Snapshot, error) {
	var (
		snapshotTableName = strings.Join([]string{constants.TABLE_SNAPSHOT_DATA_PREFIX, peerTag}, constants.DEFAULT_SEP_SYMBOL)
		sql               = fmt.Sprintf("select * from `%s` where block_hash='%s' limit %d offset %d", snapshotTableName, blockHash, pageSize, (pageNumber-1)*pageSize)
		data              = make([]model.Snapshot, 0)
	)
	db, err := oas.Raw(sql)
	if err != nil {
		log.Printf("[QuerySnapshotByBHA]raw got error:%v", err)
		return nil, ErrQueryFailed
	}
	db = db.Scan(&data)
	if db.Error != nil {
		log.Printf("[QuerySnapshotByBHA]scan got error:%v", db.Error)
		return nil, ErrQueryFailed
	}
	return data, nil
}

func QuerySnapshotByCIN(peerTag string, targetChainID int, targetChainHeight int32, pageSize, pageNumber int) ([]model.Snapshot, error) {
	var (
		snapshotTableName = strings.Join([]string{constants.TABLE_SNAPSHOT_DATA_PREFIX, peerTag}, constants.DEFAULT_SEP_SYMBOL)
		sql               = fmt.Sprintf("select * from `%s` where target_chain_id='%d' and target_chain_height='%d' limit %d offset %d", snapshotTableName, targetChainID, targetChainHeight, pageSize, (pageNumber-1)*pageSize)
		data              = make([]model.Snapshot, 0)
	)
	db, err := oas.Raw(sql)
	if err != nil {
		log.Printf("[QuerySnapshotByCIN]raw got error:%v", err)
		return nil, ErrQueryFailed
	}
	db = db.Scan(&data)
	if db.Error != nil {
		log.Printf("[QuerySnapshotByCIN]scan got error:%v", db.Error)
		return nil, ErrQueryFailed
	}
	return data, nil
}

func QueryRevisionBySID(peerTag string, snapshotID string, revisionType int, initTimestamp string) ([]model.Revision, error) {
	var (
		revisionTableName = strings.Join([]string{constants.TABLE_REVISION_PREFIX, peerTag}, constants.DEFAULT_SEP_SYMBOL)
		sql               = fmt.Sprintf("select * from `%s` where snapshot_id='%s'", revisionTableName, snapshotID)
		data              = make([]model.Revision, 0)
	)
	if revisionType != int(structure.RevisionTypeUnknown) {
		sql = fmt.Sprintf("select * from `%s` where snapshot_id='%s' and revision_type='%d'", revisionTableName, snapshotID, revisionType)
	}
	if initTimestamp != "" {
		sql = fmt.Sprintf("select * from `%s` where snapshot_id='%s' and revision_type='%d' and init_timestamp>'%s'", revisionTableName, snapshotID, revisionType, initTimestamp)
	}
	db, err := oas.Raw(sql)
	if err != nil {
		log.Printf("[QueryRevisionBySID]raw got error:%v", err)
		return nil, ErrQueryFailed
	}
	db = db.Scan(&data)
	if db.Error != nil {
		log.Printf("[QueryRevisionBySID]scan got error:%v", db.Error)
		return nil, ErrQueryFailed
	}
	return data, nil
}

func QueryBestStateBySID(peerTag string, snapshotID string, snapshotType int) ([]model.BestState, error) {
	var (
		bestStatusTableName = strings.Join([]string{constants.TABLE_STATE_PREFIX, peerTag}, constants.DEFAULT_SEP_SYMBOL)
		sql                 = fmt.Sprintf("select * from `%s` where snapshot_id='%s'", bestStatusTableName, snapshotID)
		data                = make([]model.BestState, 0)
	)
	if snapshotType != int(structure.SnapshotTypeUnknown) {
		sql = fmt.Sprintf("select * from `%s` where snapshot_id='%s' and snapshot_type='%d'", bestStatusTableName, snapshotID, snapshotType)
	}
	db, err := oas.Raw(sql)
	if err != nil {
		log.Printf("[QueryBestStateBySID]raw got error:%v", err)
		return nil, ErrQueryFailed
	}
	db = db.Scan(&data)
	if db.Error != nil {
		log.Printf("[QueryBestStateBySID]scan got error:%v", db.Error)
		return nil, ErrQueryFailed
	}
	return data, nil
}

func QueryEventOrphanBySID(peerTag string, snapshotID string, eventOrphanType int) ([]model.EventOrphan, error) {
	var (
		eventOrphanTableName = strings.Join([]string{constants.TABLE_EVENT_ORPHAN_PREFIX, peerTag}, constants.DEFAULT_SEP_SYMBOL)
		sql                  = fmt.Sprintf("select * from `%s` where snapshot_id='%s'", eventOrphanTableName, snapshotID)
		data                 = make([]model.EventOrphan, 0)
	)
	if eventOrphanType != int(structure.EventTypeUnknown) {
		sql = fmt.Sprintf("select * from `%s` where snapshot_id='%s' and event_type_orphan='%d'", eventOrphanTableName, snapshotID, eventOrphanType)
	}
	db, err := oas.Raw(sql)
	if err != nil {
		log.Printf("[QueryEventOrphanBySID]raw got error:%v", err)
		return nil, ErrQueryFailed
	}
	db = db.Scan(&data)
	if db.Error != nil {
		log.Printf("[QueryEventOrphanBySID]scan got error:%v", db.Error)
		return nil, ErrQueryFailed
	}
	return data, nil
}

func QueryEventOrphanByBHA(peerTag string, blockHash string) ([]model.EventOrphan, error) {
	var (
		eventOrphanTableName = strings.Join([]string{constants.TABLE_EVENT_ORPHAN_PREFIX, peerTag}, constants.DEFAULT_SEP_SYMBOL)
		sql                  = fmt.Sprintf("select * from `%s` where orphan_block_hash='%s'", eventOrphanTableName, blockHash)
		data                 = make([]model.EventOrphan, 0)
	)
	db, err := oas.Raw(sql)
	if err != nil {
		log.Printf("[QueryEventOrphanByBHA]raw got error:%v", err)
		return nil, ErrQueryFailed
	}
	db = db.Scan(&data)
	if db.Error != nil {
		log.Printf("[QueryEventOrphanByBHA]scan got error:%v", db.Error)
		return nil, ErrQueryFailed
	}
	return data, nil
}

func QueryEventOrphanByPBHA(peerTag string, parentBlockHash string) ([]model.EventOrphan, error) {
	var (
		eventOrphanTableName = strings.Join([]string{constants.TABLE_EVENT_ORPHAN_PREFIX, peerTag}, constants.DEFAULT_SEP_SYMBOL)
		sql                  = fmt.Sprintf("select * from `%s` where orphan_parent_block_hash='%s'", eventOrphanTableName, parentBlockHash)
		data                 = make([]model.EventOrphan, 0)
	)
	db, err := oas.Raw(sql)
	if err != nil {
		log.Printf("[QueryEventOrphanByPBHA]raw got error:%v", err)
		return nil, ErrQueryFailed
	}
	db = db.Scan(&data)
	if db.Error != nil {
		log.Printf("[QueryEventOrphanByPBHA]scan got error:%v", db.Error)
		return nil, ErrQueryFailed
	}
	return data, nil
}

func QueryPeer() ([]string, error) {
	// 如果 peerTag 为空，则查询全部 peer（按照 peer 分组），如果 peerTag 不为空，则查询该 peer 的全部信息
	var (
		data   = []string{}
		peers  = map[string]struct{}{}
		sql    = "SHOW TABLES;"
		tables = []string{}
	)
	db, err := oas.Raw(sql)
	if err != nil {
		log.Printf("[QueryPeer]raw got error:%v", err)
		return nil, ErrQueryFailed
	}
	db = db.Scan(&tables)
	if db.Error != nil {
		log.Printf("[QueryPeer]scan got error:%v", db.Error)
		return nil, ErrQueryFailed
	}
	for _, table := range tables {
		// 获取当前 peerTag
		segs := strings.Split(table, constants.DEFAULT_SEP_SYMBOL)
		currentPeerTag := strings.Join(segs[len(segs)-2:], constants.DEFAULT_SEP_SYMBOL)
		if _, ok := peers[currentPeerTag]; !ok {
			peers[currentPeerTag] = struct{}{}
		}
	}
	for peer := range peers {
		data = append(data, peer)
	}
	return data, nil
}
