package transport

import (
	"github.com/Lunarhalos/go-acm/pkg/raft/transport/pb"
	"github.com/hashicorp/raft"
)

func decodeAppendEntriesRequest(m *pb.AppendEntriesRequest) *raft.AppendEntriesRequest {
	return &raft.AppendEntriesRequest{
		RPCHeader:         decodeRPCHeader(m.RpcHeader),
		Term:              m.Term,
		Leader:            m.Leader,
		PrevLogEntry:      m.PrevLogEntry,
		PrevLogTerm:       m.PrevLogTerm,
		Entries:           decodeLogs(m.Entries),
		LeaderCommitIndex: m.LeaderCommitIndex,
	}
}

func decodeRPCHeader(m *pb.RPCHeader) raft.RPCHeader {
	return raft.RPCHeader{
		ProtocolVersion: raft.ProtocolVersion(m.ProtocolVersion),
	}
}

func decodeLogs(m []*pb.Log) []*raft.Log {
	ret := make([]*raft.Log, len(m))
	for i, l := range m {
		ret[i] = decodeLog(l)
	}
	return ret
}

func decodeLog(m *pb.Log) *raft.Log {
	return &raft.Log{
		Index:      m.Index,
		Term:       m.Term,
		Type:       decodeLogType(m.Type),
		Data:       m.Data,
		Extensions: m.Extensions,
		AppendedAt: m.AppendedAt.AsTime(),
	}
}

func decodeLogType(m pb.LogType) raft.LogType {
	switch m {
	case pb.LogType_LogCommand:
		return raft.LogCommand
	case pb.LogType_LogNoop:
		return raft.LogNoop
	case pb.LogType_LogAddPeerDeprecated:
		return raft.LogAddPeerDeprecated
	case pb.LogType_LogRemovePeerDeprecated:
		return raft.LogRemovePeerDeprecated
	case pb.LogType_LogBarrier:
		return raft.LogBarrier
	case pb.LogType_LogConfiguration:
		return raft.LogConfiguration
	default:
		panic("invalid LogType")
	}
}

func decodeAppendEntriesResponse(m *pb.AppendEntriesResponse) *raft.AppendEntriesResponse {
	return &raft.AppendEntriesResponse{
		RPCHeader:      decodeRPCHeader(m.RpcHeader),
		Term:           m.Term,
		LastLog:        m.LastLog,
		Success:        m.Success,
		NoRetryBackoff: m.NoRetryBackoff,
	}
}

func decodeRequestVoteRequest(m *pb.RequestVoteRequest) *raft.RequestVoteRequest {
	return &raft.RequestVoteRequest{
		RPCHeader:          decodeRPCHeader(m.RpcHeader),
		Term:               m.Term,
		Candidate:          m.Candidate,
		LastLogIndex:       m.LastLogIndex,
		LastLogTerm:        m.LastLogTerm,
		LeadershipTransfer: m.LeadershipTransfer,
	}
}

func decodeRequestVoteResponse(m *pb.RequestVoteResponse) *raft.RequestVoteResponse {
	return &raft.RequestVoteResponse{
		RPCHeader: decodeRPCHeader(m.RpcHeader),
		Term:      m.Term,
		Peers:     m.Peers,
		Granted:   m.Granted,
	}
}

func decodeInstallSnapshotRequest(m *pb.InstallSnapshotRequest) *raft.InstallSnapshotRequest {
	return &raft.InstallSnapshotRequest{
		RPCHeader:          decodeRPCHeader(m.RpcHeader),
		SnapshotVersion:    raft.SnapshotVersion(m.SnapshotVersion),
		Term:               m.Term,
		Leader:             m.Leader,
		LastLogIndex:       m.LastLogIndex,
		LastLogTerm:        m.LastLogTerm,
		Peers:              m.Peers,
		Configuration:      m.Configuration,
		ConfigurationIndex: m.ConfigurationIndex,
		Size:               m.Size,
	}
}

func decodeInstallSnapshotResponse(m *pb.InstallSnapshotResponse) *raft.InstallSnapshotResponse {
	return &raft.InstallSnapshotResponse{
		RPCHeader: decodeRPCHeader(m.RpcHeader),
		Term:      m.Term,
		Success:   m.Success,
	}
}

func decodeTimeoutNowRequest(m *pb.TimeoutNowRequest) *raft.TimeoutNowRequest {
	return &raft.TimeoutNowRequest{
		RPCHeader: decodeRPCHeader(m.RpcHeader),
	}
}

func decodeTimeoutNowResponse(m *pb.TimeoutNowResponse) *raft.TimeoutNowResponse {
	return &raft.TimeoutNowResponse{
		RPCHeader: decodeRPCHeader(m.RpcHeader),
	}
}
