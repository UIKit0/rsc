package x86

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

const (
	fmtLong = 1 << iota
)

var Anames8 = []string{
	"XXX",
	"AAA",
	"AAD",
	"AAM",
	"AAS",
	"ADCB",
	"ADCL",
	"ADCW",
	"ADDB",
	"ADDL",
	"ADDW",
	"ADJSP",
	"ANDB",
	"ANDL",
	"ANDW",
	"ARPL",
	"BOUNDL",
	"BOUNDW",
	"BSFL",
	"BSFW",
	"BSRL",
	"BSRW",
	"BTL",
	"BTW",
	"BTCL",
	"BTCW",
	"BTRL",
	"BTRW",
	"BTSL",
	"BTSW",
	"BYTE",
	"CALL",
	"CLC",
	"CLD",
	"CLI",
	"CLTS",
	"CMC",
	"CMPB",
	"CMPL",
	"CMPW",
	"CMPSB",
	"CMPSL",
	"CMPSW",
	"DAA",
	"DAS",
	"DATA",
	"DECB",
	"DECL",
	"DECW",
	"DIVB",
	"DIVL",
	"DIVW",
	"ENTER",
	"GLOBL",
	"GOK",
	"HISTORY",
	"HLT",
	"IDIVB",
	"IDIVL",
	"IDIVW",
	"IMULB",
	"IMULL",
	"IMULW",
	"INB",
	"INL",
	"INW",
	"INCB",
	"INCL",
	"INCW",
	"INSB",
	"INSL",
	"INSW",
	"INT",
	"INTO",
	"IRETL",
	"IRETW",
	"JCC",
	"JCS",
	"JCXZL",
	"JCXZW",
	"JEQ",
	"JGE",
	"JGT",
	"JHI",
	"JLE",
	"JLS",
	"JLT",
	"JMI",
	"JMP",
	"JNE",
	"JOC",
	"JOS",
	"JPC",
	"JPL",
	"JPS",
	"LAHF",
	"LARL",
	"LARW",
	"LEAL",
	"LEAW",
	"LEAVEL",
	"LEAVEW",
	"LOCK",
	"LODSB",
	"LODSL",
	"LODSW",
	"LONG",
	"LOOP",
	"LOOPEQ",
	"LOOPNE",
	"LSLL",
	"LSLW",
	"MOVB",
	"MOVL",
	"MOVW",
	"MOVQ",
	"MOVBLSX",
	"MOVBLZX",
	"MOVBWSX",
	"MOVBWZX",
	"MOVWLSX",
	"MOVWLZX",
	"MOVSB",
	"MOVSL",
	"MOVSW",
	"MULB",
	"MULL",
	"MULW",
	"NAME",
	"NEGB",
	"NEGL",
	"NEGW",
	"NOP",
	"NOTB",
	"NOTL",
	"NOTW",
	"ORB",
	"ORL",
	"ORW",
	"OUTB",
	"OUTL",
	"OUTW",
	"OUTSB",
	"OUTSL",
	"OUTSW",
	"PAUSE",
	"POPAL",
	"POPAW",
	"POPFL",
	"POPFW",
	"POPL",
	"POPW",
	"PUSHAL",
	"PUSHAW",
	"PUSHFL",
	"PUSHFW",
	"PUSHL",
	"PUSHW",
	"RCLB",
	"RCLL",
	"RCLW",
	"RCRB",
	"RCRL",
	"RCRW",
	"REP",
	"REPN",
	"RET",
	"ROLB",
	"ROLL",
	"ROLW",
	"RORB",
	"RORL",
	"RORW",
	"SAHF",
	"SALB",
	"SALL",
	"SALW",
	"SARB",
	"SARL",
	"SARW",
	"SBBB",
	"SBBL",
	"SBBW",
	"SCASB",
	"SCASL",
	"SCASW",
	"SETCC",
	"SETCS",
	"SETEQ",
	"SETGE",
	"SETGT",
	"SETHI",
	"SETLE",
	"SETLS",
	"SETLT",
	"SETMI",
	"SETNE",
	"SETOC",
	"SETOS",
	"SETPC",
	"SETPL",
	"SETPS",
	"CDQ",
	"CWD",
	"SHLB",
	"SHLL",
	"SHLW",
	"SHRB",
	"SHRL",
	"SHRW",
	"STC",
	"STD",
	"STI",
	"STOSB",
	"STOSL",
	"STOSW",
	"SUBB",
	"SUBL",
	"SUBW",
	"SYSCALL",
	"TESTB",
	"TESTL",
	"TESTW",
	"TEXT",
	"VERR",
	"VERW",
	"WAIT",
	"WORD",
	"XCHGB",
	"XCHGL",
	"XCHGW",
	"XLAT",
	"XORB",
	"XORL",
	"XORW",
	"FMOVB",
	"FMOVBP",
	"FMOVD",
	"FMOVDP",
	"FMOVF",
	"FMOVFP",
	"FMOVL",
	"FMOVLP",
	"FMOVV",
	"FMOVVP",
	"FMOVW",
	"FMOVWP",
	"FMOVX",
	"FMOVXP",
	"FCOMB",
	"FCOMBP",
	"FCOMD",
	"FCOMDP",
	"FCOMDPP",
	"FCOMF",
	"FCOMFP",
	"FCOMI",
	"FCOMIP",
	"FCOML",
	"FCOMLP",
	"FCOMW",
	"FCOMWP",
	"FUCOM",
	"FUCOMI",
	"FUCOMIP",
	"FUCOMP",
	"FUCOMPP",
	"FADDDP",
	"FADDW",
	"FADDL",
	"FADDF",
	"FADDD",
	"FMULDP",
	"FMULW",
	"FMULL",
	"FMULF",
	"FMULD",
	"FSUBDP",
	"FSUBW",
	"FSUBL",
	"FSUBF",
	"FSUBD",
	"FSUBRDP",
	"FSUBRW",
	"FSUBRL",
	"FSUBRF",
	"FSUBRD",
	"FDIVDP",
	"FDIVW",
	"FDIVL",
	"FDIVF",
	"FDIVD",
	"FDIVRDP",
	"FDIVRW",
	"FDIVRL",
	"FDIVRF",
	"FDIVRD",
	"FXCHD",
	"FFREE",
	"FLDCW",
	"FLDENV",
	"FRSTOR",
	"FSAVE",
	"FSTCW",
	"FSTENV",
	"FSTSW",
	"F2XM1",
	"FABS",
	"FCHS",
	"FCLEX",
	"FCOS",
	"FDECSTP",
	"FINCSTP",
	"FINIT",
	"FLD1",
	"FLDL2E",
	"FLDL2T",
	"FLDLG2",
	"FLDLN2",
	"FLDPI",
	"FLDZ",
	"FNOP",
	"FPATAN",
	"FPREM",
	"FPREM1",
	"FPTAN",
	"FRNDINT",
	"FSCALE",
	"FSIN",
	"FSINCOS",
	"FSQRT",
	"FTST",
	"FXAM",
	"FXTRACT",
	"FYL2X",
	"FYL2XP1",
	"END",
	"DYNT_",
	"INIT_",
	"SIGNAME",
	"CMPXCHGB",
	"CMPXCHGL",
	"CMPXCHGW",
	"CMPXCHG8B",
	"CPUID",
	"RDTSC",
	"XADDB",
	"XADDL",
	"XADDW",
	"CMOVLCC",
	"CMOVLCS",
	"CMOVLEQ",
	"CMOVLGE",
	"CMOVLGT",
	"CMOVLHI",
	"CMOVLLE",
	"CMOVLLS",
	"CMOVLLT",
	"CMOVLMI",
	"CMOVLNE",
	"CMOVLOC",
	"CMOVLOS",
	"CMOVLPC",
	"CMOVLPL",
	"CMOVLPS",
	"CMOVWCC",
	"CMOVWCS",
	"CMOVWEQ",
	"CMOVWGE",
	"CMOVWGT",
	"CMOVWHI",
	"CMOVWLE",
	"CMOVWLS",
	"CMOVWLT",
	"CMOVWMI",
	"CMOVWNE",
	"CMOVWOC",
	"CMOVWOS",
	"CMOVWPC",
	"CMOVWPL",
	"CMOVWPS",
	"FCMOVCC",
	"FCMOVCS",
	"FCMOVEQ",
	"FCMOVHI",
	"FCMOVLS",
	"FCMOVNE",
	"FCMOVNU",
	"FCMOVUN",
	"LFENCE",
	"MFENCE",
	"SFENCE",
	"EMMS",
	"PREFETCHT0",
	"PREFETCHT1",
	"PREFETCHT2",
	"PREFETCHNTA",
	"BSWAPL",
	"UNDEF",
	"ADDPD",
	"ADDPS",
	"ADDSD",
	"ADDSS",
	"ANDNPD",
	"ANDNPS",
	"ANDPD",
	"ANDPS",
	"CMPPD",
	"CMPPS",
	"CMPSD",
	"CMPSS",
	"COMISD",
	"COMISS",
	"CVTPL2PD",
	"CVTPL2PS",
	"CVTPD2PL",
	"CVTPD2PS",
	"CVTPS2PL",
	"CVTPS2PD",
	"CVTSD2SL",
	"CVTSD2SS",
	"CVTSL2SD",
	"CVTSL2SS",
	"CVTSS2SD",
	"CVTSS2SL",
	"CVTTPD2PL",
	"CVTTPS2PL",
	"CVTTSD2SL",
	"CVTTSS2SL",
	"DIVPD",
	"DIVPS",
	"DIVSD",
	"DIVSS",
	"MASKMOVOU",
	"MAXPD",
	"MAXPS",
	"MAXSD",
	"MAXSS",
	"MINPD",
	"MINPS",
	"MINSD",
	"MINSS",
	"MOVAPD",
	"MOVAPS",
	"MOVO",
	"MOVOU",
	"MOVHLPS",
	"MOVHPD",
	"MOVHPS",
	"MOVLHPS",
	"MOVLPD",
	"MOVLPS",
	"MOVMSKPD",
	"MOVMSKPS",
	"MOVNTO",
	"MOVNTPD",
	"MOVNTPS",
	"MOVSD",
	"MOVSS",
	"MOVUPD",
	"MOVUPS",
	"MULPD",
	"MULPS",
	"MULSD",
	"MULSS",
	"ORPD",
	"ORPS",
	"PADDQ",
	"PAND",
	"PCMPEQB",
	"PMAXSW",
	"PMAXUB",
	"PMINSW",
	"PMINUB",
	"PMOVMSKB",
	"PSADBW",
	"PSUBB",
	"PSUBL",
	"PSUBQ",
	"PSUBSB",
	"PSUBSW",
	"PSUBUSB",
	"PSUBUSW",
	"PSUBW",
	"PUNPCKHQDQ",
	"PUNPCKLQDQ",
	"PXOR",
	"RCPPS",
	"RCPSS",
	"RSQRTPS",
	"RSQRTSS",
	"SQRTPD",
	"SQRTPS",
	"SQRTSD",
	"SQRTSS",
	"SUBPD",
	"SUBPS",
	"SUBSD",
	"SUBSS",
	"UCOMISD",
	"UCOMISS",
	"UNPCKHPD",
	"UNPCKHPS",
	"UNPCKLPD",
	"UNPCKLPS",
	"XORPD",
	"XORPS",
	"AESENC",
	"PINSRD",
	"PSHUFB",
	"USEFIELD",
	"TYPE",
	"FUNCDATA",
	"PCDATA",
	"CHECKNIL",
	"VARDEF",
	"VARKILL",
	"DUFFCOPY",
	"DUFFZERO",
	"LAST",
}
