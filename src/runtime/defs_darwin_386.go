// created by cgo -cdefs and then converted to Go
// cgo -cdefs defs_darwin.go

package runtime

import "unsafe"

const (
	_EINTR  = 0x4
	_EFAULT = 0xe

	_PROT_NONE  = 0x0
	_PROT_READ  = 0x1
	_PROT_WRITE = 0x2
	_PROT_EXEC  = 0x4

	_MAP_ANON    = 0x1000
	_MAP_PRIVATE = 0x2
	_MAP_FIXED   = 0x10

	_MADV_DONTNEED = 0x4
	_MADV_FREE     = 0x5

	_MACH_MSG_TYPE_MOVE_RECEIVE   = 0x10
	_MACH_MSG_TYPE_MOVE_SEND      = 0x11
	_MACH_MSG_TYPE_MOVE_SEND_ONCE = 0x12
	_MACH_MSG_TYPE_COPY_SEND      = 0x13
	_MACH_MSG_TYPE_MAKE_SEND      = 0x14
	_MACH_MSG_TYPE_MAKE_SEND_ONCE = 0x15
	_MACH_MSG_TYPE_COPY_RECEIVE   = 0x16

	_MACH_MSG_PORT_DESCRIPTOR         = 0x0
	_MACH_MSG_OOL_DESCRIPTOR          = 0x1
	_MACH_MSG_OOL_PORTS_DESCRIPTOR    = 0x2
	_MACH_MSG_OOL_VOLATILE_DESCRIPTOR = 0x3

	_MACH_MSGH_BITS_COMPLEX = 0x80000000

	_MACH_SEND_MSG  = 0x1
	_MACH_RCV_MSG   = 0x2
	_MACH_RCV_LARGE = 0x4

	_MACH_SEND_TIMEOUT   = 0x10
	_MACH_SEND_INTERRUPT = 0x40
	_MACH_SEND_ALWAYS    = 0x10000
	_MACH_SEND_TRAILER   = 0x20000
	_MACH_RCV_TIMEOUT    = 0x100
	_MACH_RCV_NOTIFY     = 0x200
	_MACH_RCV_INTERRUPT  = 0x400
	_MACH_RCV_OVERWRITE  = 0x1000

	_NDR_PROTOCOL_2_0      = 0x0
	_NDR_INT_BIG_ENDIAN    = 0x0
	_NDR_INT_LITTLE_ENDIAN = 0x1
	_NDR_FLOAT_IEEE        = 0x0
	_NDR_CHAR_ASCII        = 0x0

	_SA_SIGINFO   = 0x40
	_SA_RESTART   = 0x2
	_SA_ONSTACK   = 0x1
	_SA_USERTRAMP = 0x100
	_SA_64REGSET  = 0x200

	_SIGHUP    = 0x1
	_SIGINT    = 0x2
	_SIGQUIT   = 0x3
	_SIGILL    = 0x4
	_SIGTRAP   = 0x5
	_SIGABRT   = 0x6
	_SIGEMT    = 0x7
	_SIGFPE    = 0x8
	_SIGKILL   = 0x9
	_SIGBUS    = 0xa
	_SIGSEGV   = 0xb
	_SIGSYS    = 0xc
	_SIGPIPE   = 0xd
	_SIGALRM   = 0xe
	_SIGTERM   = 0xf
	_SIGURG    = 0x10
	_SIGSTOP   = 0x11
	_SIGTSTP   = 0x12
	_SIGCONT   = 0x13
	_SIGCHLD   = 0x14
	_SIGTTIN   = 0x15
	_SIGTTOU   = 0x16
	_SIGIO     = 0x17
	_SIGXCPU   = 0x18
	_SIGXFSZ   = 0x19
	_SIGVTALRM = 0x1a
	_SIGPROF   = 0x1b
	_SIGWINCH  = 0x1c
	_SIGINFO   = 0x1d
	_SIGUSR1   = 0x1e
	_SIGUSR2   = 0x1f

	_FPE_INTDIV = 0x7
	_FPE_INTOVF = 0x8
	_FPE_FLTDIV = 0x1
	_FPE_FLTOVF = 0x2
	_FPE_FLTUND = 0x3
	_FPE_FLTRES = 0x4
	_FPE_FLTINV = 0x5
	_FPE_FLTSUB = 0x6

	_BUS_ADRALN = 0x1
	_BUS_ADRERR = 0x2
	_BUS_OBJERR = 0x3

	_SEGV_MAPERR = 0x1
	_SEGV_ACCERR = 0x2

	_ITIMER_REAL    = 0x0
	_ITIMER_VIRTUAL = 0x1
	_ITIMER_PROF    = 0x2

	_EV_ADD       = 0x1
	_EV_DELETE    = 0x2
	_EV_CLEAR     = 0x20
	_EV_RECEIPT   = 0x40
	_EV_ERROR     = 0x4000
	_EV_EOF       = 0x8000
	_EVFILT_READ  = -0x1
	_EVFILT_WRITE = -0x2

	_PTHREAD_CREATE_DETACHED = 0x2
)

type machbody struct {
	msgh_descriptor_count uint32
}

type machheader struct {
	msgh_bits        uint32
	msgh_size        uint32
	msgh_remote_port uint32
	msgh_local_port  uint32
	msgh_reserved    uint32
	msgh_id          int32
}

type machndr struct {
	mig_vers     uint8
	if_vers      uint8
	reserved1    uint8
	mig_encoding uint8
	int_rep      uint8
	char_rep     uint8
	float_rep    uint8
	reserved2    uint8
}

type machport struct {
	name        uint32
	pad1        uint32
	pad2        uint16
	disposition uint8
	_type       uint8
}

type stackt struct {
	ss_sp    *byte
	ss_size  uintptr
	ss_flags int32
}

type sigactiont struct {
	__sigaction_u [4]byte
	sa_tramp      unsafe.Pointer
	sa_mask       uint32
	sa_flags      int32
}

type usigactiont struct {
	__sigaction_u [4]byte
	sa_mask       uint32
	sa_flags      int32
}

type siginfo struct {
	si_signo  int32
	si_errno  int32
	si_code   int32
	si_pid    int32
	si_uid    uint32
	si_status int32
	si_addr   uint32
	si_value  [4]byte
	si_band   int32
	__pad     [7]uint32
}

type timeval struct {
	tv_sec  int32
	tv_usec int32
}

func (tv *timeval) set_usec(x int32) {
	tv.tv_usec = x
}

type itimerval struct {
	it_interval timeval
	it_value    timeval
}

type timespec struct {
	tv_sec  int32
	tv_nsec int32
}

type fpcontrol struct {
	pad_cgo_0 [2]byte
}

type fpstatus struct {
	pad_cgo_0 [2]byte
}

type regmmst struct {
	mmst_reg  [10]int8
	mmst_rsrv [6]int8
}

type regxmm struct {
	xmm_reg [16]int8
}

type regs64 struct {
	rax    uint64
	rbx    uint64
	rcx    uint64
	rdx    uint64
	rdi    uint64
	rsi    uint64
	rbp    uint64
	rsp    uint64
	r8     uint64
	r9     uint64
	r10    uint64
	r11    uint64
	r12    uint64
	r13    uint64
	r14    uint64
	r15    uint64
	rip    uint64
	rflags uint64
	cs     uint64
	fs     uint64
	gs     uint64
}

type floatstate64 struct {
	fpu_reserved  [2]int32
	fpu_fcw       fpcontrol
	fpu_fsw       fpstatus
	fpu_ftw       uint8
	fpu_rsrv1     uint8
	fpu_fop       uint16
	fpu_ip        uint32
	fpu_cs        uint16
	fpu_rsrv2     uint16
	fpu_dp        uint32
	fpu_ds        uint16
	fpu_rsrv3     uint16
	fpu_mxcsr     uint32
	fpu_mxcsrmask uint32
	fpu_stmm0     regmmst
	fpu_stmm1     regmmst
	fpu_stmm2     regmmst
	fpu_stmm3     regmmst
	fpu_stmm4     regmmst
	fpu_stmm5     regmmst
	fpu_stmm6     regmmst
	fpu_stmm7     regmmst
	fpu_xmm0      regxmm
	fpu_xmm1      regxmm
	fpu_xmm2      regxmm
	fpu_xmm3      regxmm
	fpu_xmm4      regxmm
	fpu_xmm5      regxmm
	fpu_xmm6      regxmm
	fpu_xmm7      regxmm
	fpu_xmm8      regxmm
	fpu_xmm9      regxmm
	fpu_xmm10     regxmm
	fpu_xmm11     regxmm
	fpu_xmm12     regxmm
	fpu_xmm13     regxmm
	fpu_xmm14     regxmm
	fpu_xmm15     regxmm
	fpu_rsrv4     [96]int8
	fpu_reserved1 int32
}

type exceptionstate64 struct {
	trapno     uint16
	cpu        uint16
	err        uint32
	faultvaddr uint64
}

type mcontext64 struct {
	es exceptionstate64
	ss regs64
	fs floatstate64
}

type regs32 struct {
	eax    uint32
	ebx    uint32
	ecx    uint32
	edx    uint32
	edi    uint32
	esi    uint32
	ebp    uint32
	esp    uint32
	ss     uint32
	eflags uint32
	eip    uint32
	cs     uint32
	ds     uint32
	es     uint32
	fs     uint32
	gs     uint32
}

type floatstate32 struct {
	fpu_reserved  [2]int32
	fpu_fcw       fpcontrol
	fpu_fsw       fpstatus
	fpu_ftw       uint8
	fpu_rsrv1     uint8
	fpu_fop       uint16
	fpu_ip        uint32
	fpu_cs        uint16
	fpu_rsrv2     uint16
	fpu_dp        uint32
	fpu_ds        uint16
	fpu_rsrv3     uint16
	fpu_mxcsr     uint32
	fpu_mxcsrmask uint32
	fpu_stmm0     regmmst
	fpu_stmm1     regmmst
	fpu_stmm2     regmmst
	fpu_stmm3     regmmst
	fpu_stmm4     regmmst
	fpu_stmm5     regmmst
	fpu_stmm6     regmmst
	fpu_stmm7     regmmst
	fpu_xmm0      regxmm
	fpu_xmm1      regxmm
	fpu_xmm2      regxmm
	fpu_xmm3      regxmm
	fpu_xmm4      regxmm
	fpu_xmm5      regxmm
	fpu_xmm6      regxmm
	fpu_xmm7      regxmm
	fpu_rsrv4     [224]int8
	fpu_reserved1 int32
}

type exceptionstate32 struct {
	trapno     uint16
	cpu        uint16
	err        uint32
	faultvaddr uint32
}

type mcontext32 struct {
	es exceptionstate32
	ss regs32
	fs floatstate32
}

type ucontext struct {
	uc_onstack  int32
	uc_sigmask  uint32
	uc_stack    stackt
	uc_link     *ucontext
	uc_mcsize   uint32
	uc_mcontext *mcontext32
}

type keventt struct {
	ident  uint32
	filter int16
	flags  uint16
	fflags uint32
	data   int32
	udata  *byte
}

type pthread uintptr
type pthreadattr struct {
	X__sig    int32
	X__opaque [36]int8
}
