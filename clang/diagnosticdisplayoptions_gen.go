package clang

// #include "./clang-c/Index.h"
// #include "go-clang.h"
import "C"
import (
	"fmt"
	"strings"
)

/*
	Options to control the display of diagnostics.

	The values in this enum are meant to be combined to customize the
	behavior of clang_formatDiagnostic().
*/
type DiagnosticDisplayOptions uint32

const (
	/*
		Display the source-location information where the
		diagnostic was located.

		When set, diagnostics will be prefixed by the file, line, and
		(optionally) column to which the diagnostic refers. For example,

		\code
		test.c:28: warning: extra tokens at end of #endif directive
		\endcode

		This option corresponds to the clang flag -fshow-source-location.
	*/
	Diagnostic_DisplaySourceLocation DiagnosticDisplayOptions = C.CXDiagnostic_DisplaySourceLocation
	/*
		If displaying the source-location information of the
		diagnostic, also include the column number.

		This option corresponds to the clang flag -fshow-column.
	*/
	Diagnostic_DisplayColumn DiagnosticDisplayOptions = C.CXDiagnostic_DisplayColumn
	/*
		If displaying the source-location information of the
		diagnostic, also include information about source ranges in a
		machine-parsable format.

		This option corresponds to the clang flag
		-fdiagnostics-print-source-range-info.
	*/
	Diagnostic_DisplaySourceRanges DiagnosticDisplayOptions = C.CXDiagnostic_DisplaySourceRanges
	/*
		Display the option name associated with this diagnostic, if any.

		The option name displayed (e.g., -Wconversion) will be placed in brackets
		after the diagnostic text. This option corresponds to the clang flag
		-fdiagnostics-show-option.
	*/
	Diagnostic_DisplayOption DiagnosticDisplayOptions = C.CXDiagnostic_DisplayOption
	/*
		Display the category number associated with this diagnostic, if any.

		The category number is displayed within brackets after the diagnostic text.
		This option corresponds to the clang flag
		-fdiagnostics-show-category=id.
	*/
	Diagnostic_DisplayCategoryId DiagnosticDisplayOptions = C.CXDiagnostic_DisplayCategoryId
	/*
		Display the category name associated with this diagnostic, if any.

		The category name is displayed within brackets after the diagnostic text.
		This option corresponds to the clang flag
		-fdiagnostics-show-category=name.
	*/
	Diagnostic_DisplayCategoryName DiagnosticDisplayOptions = C.CXDiagnostic_DisplayCategoryName
)

func (ddo DiagnosticDisplayOptions) String() string {
	var r []string
	for _, t := range []struct {
		flag DiagnosticDisplayOptions
		name string
	}{
		{Diagnostic_DisplaySourceLocation, "DisplaySourceLocation"},
		{Diagnostic_DisplayColumn, "DisplayColumn"},
		{Diagnostic_DisplaySourceRanges, "DisplaySourceRanges"},
		{Diagnostic_DisplayOption, "DisplayOption"},
		{Diagnostic_DisplayCategoryId, "DisplayCategoryId"},
		{Diagnostic_DisplayCategoryName, "DisplayCategoryName"},
	} {
		if ddo&t.flag == 0 {
			continue
		}
		ddo &^= t.flag
		r = append(r, t.name)
	}
	if ddo != 0 {
		// This cast to a large intrinsic is important; it avoids recursive calls to String().
		r = append(r, fmt.Sprintf("additional-bits(%x)", uint64(ddo)))
	}
	return strings.Join(r, ",")
}
