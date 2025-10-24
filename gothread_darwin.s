#include "textflag.h"

TEXT libpthread_pthread_mach_thread_np_trampoline<>(SB),NOSPLIT,$0-0
	JMP	libpthread_pthread_mach_thread_np(SB)
GLOBL	·libpthread_pthread_mach_thread_np_trampoline_addr(SB), RODATA, $8
DATA	·libpthread_pthread_mach_thread_np_trampoline_addr(SB)/8, $libpthread_pthread_mach_thread_np_trampoline<>(SB)

TEXT libpthread_pthread_self_trampoline<>(SB),NOSPLIT,$0-0
	JMP	libpthread_pthread_self(SB)
GLOBL	·libpthread_pthread_self_trampoline_addr(SB), RODATA, $8
DATA	·libpthread_pthread_self_trampoline_addr(SB)/8, $libpthread_pthread_self_trampoline<>(SB)
