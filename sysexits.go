package sysexits

import (
	"fmt"
	"os"
)

type SysExit int

const (
	ExOk           SysExit = 0
	ExUsage        SysExit = 64
	ExDataErr      SysExit = 65
	ExNoInput      SysExit = 66
	ExInterr       SysExit = 75
	ExInval        SysExit = 77
	ExIO           SysExit = 78
	ExNoMem        SysExit = 79
	ExPerm         SysExit = 80
	ExNoFile       SysExit = 81
	ExTooBig       SysExit = 82
	ExNoSpace      SysExit = 83
	ExCmdRun       SysExit = 84
	ExCmdErr       SysExit = 85
	ExSystem       SysExit = 86
	ExCantCp       SysExit = 87
	ExCantCreate   SysExit = 88
	ExCantOpen     SysExit = 89
	ExCantClose    SysExit = 90
	ExCantRead     SysExit = 91
	ExCantWrite    SysExit = 92
	ExCantExec     SysExit = 93
	ExFileExists   SysExit = 94
	ExFileInUse    SysExit = 95
	ExFileNotFound SysExit = 96
	ExFileTooLarge SysExit = 97
	ExNoExec       SysExit = 98
	ExNoLogin      SysExit = 99
	ExNoUser       SysExit = 100
	ExNoGroup      SysExit = 101
	ExNoPerm       SysExit = 102
	ExNoDir        SysExit = 103
	ExNoFileSys    SysExit = 104
	ExNoPath       SysExit = 105
	ExNoLink       SysExit = 106
	ExNoMemSys     SysExit = 107
	ExNoProc       SysExit = 108
	ExNoSocket     SysExit = 109
	ExNoSockSys    SysExit = 110
	ExNoSys        SysExit = 111
)

func (se SysExit) Do() {
	os.Exit(int(se))
}

func (se SysExit) Error() string {
	switch se {
	case ExOk:
		return "ok"
	case ExUsage:
		return "command line usage error"
	case ExDataErr:
		return "data format error"
	case ExNoInput:
		return "input error"
	case ExInterr:
		return "interrupted system call"
	case ExInval:
		return "invalid argument"
	case ExIO:
		return "i/o error"
	case ExNoMem:
		return "insufficient memory"
	case ExPerm:
		return "permission denied"
	case ExNoFile:
		return "file does not exist"
	case ExTooBig:
		return "file too large"
	case ExNoSpace:
		return "no space left on device"
	case ExCmdRun:
		return "cannot run program"
	case ExCmdErr:
		return "command error"
	case ExSystem:
		return "system error"
	case ExCantCp:
		return "cannot copy"
	case ExCantCreate:
		return "cannot create"
	case ExCantOpen:
		return "cannot open"
	case ExCantClose:
		return "cannot close"
	case ExCantRead:
		return "cannot read"
	case ExCantWrite:
		return "cannot write"
	case ExCantExec:
		return "cannot execute"
	case ExFileExists:
		return "file exists"
	case ExFileInUse:
		return "file in use"
	case ExFileNotFound:
		return "file not found"
	case ExFileTooLarge:
		return "file too large"
	case ExNoExec:
		return "cannot execute"
	case ExNoLogin:
		return "no login"
	case ExNoUser:
		return "no such user"
	case ExNoGroup:
		return "no such group"
	case ExNoPerm:
		return "no permission"
	case ExNoDir:
		return "no such directory"
	case ExNoFileSys:
		return "no file system"
	case ExNoPath:
		return "no such path"
	case ExNoLink:
		return "no such link"
	case ExNoMemSys:
		return "no memory"
	case ExNoProc:
		return "no process"
	case ExNoSocket:
		return "no socket"
	case ExNoSockSys:
		return "no socket system"
	case ExNoSys:
		return "no system"
	default:
		return "unknown error (" + fmt.Sprint(se) + ")"
	}
}
