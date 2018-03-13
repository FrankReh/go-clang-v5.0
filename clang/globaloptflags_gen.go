package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"

type GlobalOptFlags uint32

const (
	// Used to indicate that no special CXIndex options are needed.
	GlobalOpt_None GlobalOptFlags = C.CXGlobalOpt_None
	/*
		Used to indicate that threads that libclang creates for indexing
		purposes should use background priority.

		Affects #clang_indexSourceFile, #clang_indexTranslationUnit,
		#clang_parseTranslationUnit, #clang_saveTranslationUnit.
	*/
	GlobalOpt_ThreadBackgroundPriorityForIndexing GlobalOptFlags = C.CXGlobalOpt_ThreadBackgroundPriorityForIndexing

	/*
		Used to indicate that threads that libclang creates for editing
		purposes should use background priority.

		Affects #clang_reparseTranslationUnit, #clang_codeCompleteAt,
		#clang_annotateTokens
	*/
	GlobalOpt_ThreadBackgroundPriorityForEditing GlobalOptFlags = C.CXGlobalOpt_ThreadBackgroundPriorityForEditing

	/*
		Used to indicate that all threads that libclang creates should use background priority.
		Both of the above.
	*/
	GlobalOpt_ThreadBackgroundPriorityForAll GlobalOptFlags = C.CXGlobalOpt_ThreadBackgroundPriorityForAll
)
