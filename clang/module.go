package clang

// #include "go-clang.h"
import "C"

type Module struct {
	c C.CXModule
}

/*
	Parameter Module a module object.

	Returns the module file where the provided module object came from.
*/
func (m Module) ASTFile() File {
	return File{C.clang_Module_getASTFile(m.c)}
}

/*
	Parameter Module a module object.

	Returns the parent of a sub-module or NULL if the given module is top-level,
	e.g. for 'std.vector' it will return the 'std' module.
*/
func (m Module) Parent() Module {
	return Module{C.clang_Module_getParent(m.c)}
}

/*
	Parameter Module a module object.

	Returns the name of the module, e.g. for the 'std.vector' sub-module it
	will return "vector".
*/
func (m Module) Name() string {
	return cx2GoString(C.clang_Module_getName(m.c))
}

/*
	Parameter Module a module object.

	Returns the full name of the module, e.g. "std.vector".
*/
func (m Module) FullName() string {
	return cx2GoString(C.clang_Module_getFullName(m.c))
}

/*
	Parameter Module a module object.

	Returns non-zero if the module is a system one.
*/
func (m Module) IsSystem() bool {
	return C.clang_Module_isSystem(m.c) != 0
}
