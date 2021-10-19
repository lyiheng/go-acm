package transport

import (
	"github.com/Lunarhalos/go-acm/pkg/raft/transport/pb"
	"github.com/hashicorp/raft"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func encodeAppendEntriesRequest(s *raft.AppendEntriesRequest) *pb.AppendEntriesRequest {
	return &pb.AppendEntriesRequest{
		RpcHeader:         encodeRPCHeader(s.RPCHeader),
		Term:              s.Term,
		Leader:            s.Leader,
		PrevLogEntry:      s.PrevLogEntry,
		PrevLogTerm:       s.PrevLogTerm,
		Entries:           encodeLogs(s.Entries),
		LeaderCommitIndex: s.LeaderCommitIndex,
	}
}

func encodeRPCHeader(s raft.RPCHeader) *pb.RPCHeader {
	return &pb.RPCHeader{
		ProtocolVersion: int64(s.ProtocolVersion),
	}
}

func encodeLogs(s []*raft.Log) []*pb.Log {
	ret := make([]*pb.Log, len(s))
	for i, l := range s {
		ret[i] = encodeLog(l)
	}
	return ret
}

func encodeLog(s *raft.Log) *pb.Log {
	return &pb.Log{
		Index:      s.Index,
		Term:       s.Term,
		Type:       encodeLogType(s.Type),
		Data:       s.Data,
		Extensions: s.Extensions,
		AppendedAt: timestamppb.New(s.AppendedAt),
	}
}

func encodeLogType(s raft.LogType) pb.LogType {
	switch s {
	case raft.LogCommand:
		return pb.LogType_LogCommand
	case raft.LogNoop:
		return pb.LogType_LogNoop
	case raft.LogAddPeerDeprecated:
		return pb.LogType_LogAddPeerDeprecated
	case raft.LogRemovePeerDeprecated:
		return pb.LogType_LogRemovePeerDeprecated
	case raft.LogBarrier:
		return pb.LogType_LogBarrier
	case raft.LogConfiguration:
		return pb.LogType_LogConfiguration
	default:
		panic("invalid LogType")
	}
}

func encodeAppendEntriesResponse(s *raft.AppendEntriesResponse) *pb.AppendEntriesResponse {
	return &pb.AppendEntriesResponse{
		RpcHeader:      encodeRPCHeader(s.RPCHeader),
		Term:           s.Term,
		LastLog:        s.LastLog,
		Success:        s.Success,
		NoRetryBackoff: s.NoRetryBackoff,
	}
}

func encodeRequestVoteRequest(s *raft.RequestVoteRequest) *pb.RequestVoteRequest {
	return &pb.RequestVoteRequest{
		RpcHeader:          encodeRPCHeader(s.RPCHeader),
		Term:               s.Term,
		Candidate:          s.Candidate,
		LastLogIndex:       s.LastLogIndex,
		LastLogTerm:        s.LastLogTerm,
		LeadershipTransfer: s.LeadershipTransfer,
	}
}

func encodeRequestVoteResponse(s *raft.RequestVoteResponse) *pb.RequestVoteResponse {
	return &pb.RequestVoteResponse{
		RpcHeader: encodeRPCHeader(s.RPCHeader),
		Term:      s.Term,
		Peers:     s.Peers,
		Granted:   s.Granted,
	}
}

func encodeInstallSnapshotRequest(s *raft.InstallSnapshotRequest) *pb.InstallSnapshotRequest {
	return &pb.InstallSnapshotRequest{
		RpcHeader:          encodeRPCHeader(s.RPCHeader),
		SnapshotVersion:    int64(s.SnapshotVersion),
		Term:               s.Term,
		Leader:             s.Leader,
		LastLogIndex:       s.LastLogIndex,
		LastLogTerm:        s.LastLogTerm,
		Peers:              s.Peers,
		Configuration:      s.Configuration,
		ConfigurationIndex: s.ConfigurationIndex,
		Size:               s.Size,
	}
}

func encodeInstallSnapshotResponse(s *raft.InstallSnapshotResponse) *pb.InstallSnapshotResponse {
	return &pb.InstallSnapshotResponse{
		RpcHeader: encodeRPCHeader(s.RPCHeader),
		Term:      s.Term,
		Success:   s.Success,
	}
}

func encodeTimeoutNowRequest(s *raft.TimeoutNowRequest) *pb.TimeoutNowRequest {
	return &pb.TimeoutNowRequest{
		RpcHeader: encodeRPCHeader(s.RPCHeader),
	}
}

func encodeTimeoutNowResponse(s *raft.TimeoutNowResponse) *pb.TimeoutNowResponse {
	return &pb.TimeoutNowResponse{
		RpcHeader: encodeRPCHeader(s.RPCHeader),
	}
}
