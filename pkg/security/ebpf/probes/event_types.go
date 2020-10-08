// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build linux

package probes

import "github.com/DataDog/ebpf/manager"

// SelectorsPerEventType is the list of probes that should be activated for each event
var SelectorsPerEventType = map[string][]manager.ProbesSelector{

	// The following events will always be activated, regardless of the rules loaded
	"*": {
		// Exec probes
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "tracepoint/sched/sched_process_fork"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/do_exit"}},
		}},
		&manager.OneOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/cgroup_procs_write"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/cgroup1_procs_write"}},
		}},
		&manager.OneOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/cgroup_tasks_write"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/cgroup1_tasks_write"}},
		}},
		&manager.OneOf{Selectors: []manager.ProbesSelector{
			&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
				manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "execve"}, Entry),
			},
			&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
				manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "execveat"}, Entry),
			},
		}},

		// Mount probes
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/attach_recursive_mnt"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/propagate_mnt"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/security_sb_umount"}},
		}},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "mount"}, EntryAndExit, true),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "umount"}, EntryAndExit),
		},
	},

	// List of probes required to capture chmod events
	"chmod": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/security_inode_setattr"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write"}},
		}},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "chmod"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "fchmod"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "fchmodat"}, EntryAndExit),
		},
	},

	// List of probes required to capture chown events
	"chown": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/security_inode_setattr"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write"}},
		}},
		&manager.OneOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write_file"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write_file_path"}},
		}},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "chown"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "chown16"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "fchown"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "fchown16"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "fchownat"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "lchown"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "lchown16"}, EntryAndExit),
		},
	},

	// List of probes required to capture link events
	"link": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/vfs_link"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/filename_create"}},
		}},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "link"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "linkat"}, EntryAndExit),
		},
	},

	// List of probes required to capture mkdir events
	"mkdir": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/vfs_mkdir"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/filename_create"}},
		}},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "mkdir"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "mkdirat"}, EntryAndExit),
		},
	},

	// List of probes required to capture open events
	"open": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/vfs_open"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/vfs_truncate"}},
		}},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "open"}, EntryAndExit, true),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "creat"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "open_by_handle_at"}, EntryAndExit, true),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "truncate"}, EntryAndExit, true),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "openat"}, EntryAndExit, true),
		},
	},

	// List of probes required to capture ptrace events
	"ptrace": {
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "ptrace"}, EntryAndExit),
		},
	},

	// List of probes required to cqpture mmap events
	"mmap": {
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "mmap"}, EntryAndExit),
		},
	},

	// List of probes required to cqpture mprotect events
	"mprotect": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/security_file_mprotect"}},
		}},
	},

	// List of probes required to capture removexattr events
	"removexattr": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/vfs_removexattr"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write"}},
		}},
		&manager.OneOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write_file"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write_file_path"}},
		}},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "removexattr"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "fremovexattr"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "lremovexattr"}, EntryAndExit),
		},
	},

	// List of probes required to capture rename events
	"rename": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/vfs_rename"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write"}},
		}},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "rename"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "renameat"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "renameat2"}, EntryAndExit),
		},
	},

	// List of probes required to capture rmdir events
	"rmdir": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/security_inode_rmdir"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write"}},
		}},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "rmdir"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "unlinkat"}, EntryAndExit),
		},
	},

	// List of probes required to capture setxattr events
	"setxattr": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/vfs_setxattr"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write"}},
		}},
		&manager.OneOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write_file"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write_file_path"}},
		}},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "setxattr"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "fsetxattr"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "lsetxattr"}, EntryAndExit),
		},
	},

	// List of probes required to capture unlink events
	"unlink": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/vfs_unlink"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write"}},
		}},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "unlink"}, EntryAndExit),
		},
		&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
			manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "unlinkat"}, EntryAndExit),
		},
	},

	// List of probes required to capture utimes events
	"utimes": {
		&manager.AllOf{Selectors: []manager.ProbesSelector{
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/security_inode_setattr"}},
			&manager.ProbeSelector{ProbeIdentificationPair: manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "kprobe/mnt_want_write"}},
		}},
		&manager.OneOf{Selectors: []manager.ProbesSelector{
			&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
				manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "utime"}, EntryAndExit),
			},
			&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
				manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "utime32"}, EntryAndExit),
			},
		}},
		&manager.OneOf{Selectors: []manager.ProbesSelector{
			&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
				manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "utimes"}, EntryAndExit),
			},
			&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
				manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "utimes_time32"}, EntryAndExit),
			},
		}},
		&manager.OneOf{Selectors: []manager.ProbesSelector{
			&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
				manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "utimensat"}, EntryAndExit),
			},
			&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
				manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "utimensat_time32"}, EntryAndExit),
			},
		}},
		&manager.OneOf{Selectors: []manager.ProbesSelector{
			&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
				manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "futimesat"}, EntryAndExit),
			},
			&manager.AllOf{Selectors: ExpandSyscallProbesSelector(
				manager.ProbeIdentificationPair{UID: SecurityAgentUID, Section: "futimesat_time32"}, EntryAndExit),
			},
		}},
	},
}