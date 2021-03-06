// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

type mOS struct {
	machport uint32
	waitsema uint32
}

const (
	_CTL_HW      = 6
	_HW_NCPU     = 3
	_HW_PAGESIZE = 7
)

const _DebugMach = false

// Mach RPC (MIG)
const (
	_MinMachMsg = 48
	_MachReply  = 100
)

type codemsg struct {
	h    machheader
	ndr  machndr
	code int32
}

const (
	tmach_semcreate = 3418
	rmach_semcreate = tmach_semcreate + _MachReply

	tmach_semdestroy = 3419
	rmach_semdestroy = tmach_semdestroy + _MachReply

	_KERN_ABORTED             = 14
	_KERN_OPERATION_TIMED_OUT = 49
)

type tmach_semcreatemsg struct {
	h      machheader
	ndr    machndr
	policy int32
	value  int32
}

type rmach_semcreatemsg struct {
	h         machheader
	body      machbody
	semaphore machport
}

type tmach_semdestroymsg struct {
	h         machheader
	body      machbody
	semaphore machport
}

const (
	_NSIG        = 32
	_SI_USER     = 0
	_SIG_BLOCK   = 1
	_SIG_UNBLOCK = 2
	_SIG_SETMASK = 3
	_SS_DISABLE  = 4
)

type sigset uint32
