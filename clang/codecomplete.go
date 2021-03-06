package clang

// #include "go-clang.h"
import "C"

// Flags that can be passed to clang_codeCompleteAt() to modify its behavior.
//
// The enumerators in this enumeration can be bitwise-OR'd together to provide
// multiple options to clang_codeCompleteAt().
type CodeComplete_Flags uint32

const (
	// Whether to include macros within the set of code completions returned.
	CodeComplete_IncludeMacros CodeComplete_Flags = C.CXCodeComplete_IncludeMacros

	// Whether to include code patterns for language constructs within the set of code completions, e.g., for loops.
	CodeComplete_IncludeCodePatterns CodeComplete_Flags = C.CXCodeComplete_IncludeCodePatterns

	// Whether to include brief documentation within the set of code completions returned.
	CodeComplete_IncludeBriefComments CodeComplete_Flags = C.CXCodeComplete_IncludeBriefComments

	// Whether to speed up completion by omitting top- or namespace-level entities
	// defined in the preamble. There's no guarantee any particular entity is
	// omitted. This may be useful if the headers are indexed externally.
	CodeComplete_SkipPreamble CodeComplete_Flags = C.CXCodeComplete_SkipPreamble

	// Whether to include completions with small
	// fix-its, e.g. change '.' to '->' on member access, etc.
	CodeComplete_IncludeCompletionsWithFixIts CodeComplete_Flags = C.CXCodeComplete_IncludeCompletionsWithFixIts
)
