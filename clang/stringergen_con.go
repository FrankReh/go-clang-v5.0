// generated by stringer -output=stringergen_con.go -type=AccessSpecifier,AvailabilityKind,CallingConv,ChildVisitResult,CommentInlineCommandRenderKind,CommentKind,CommentParamPassDirection,CompletionChunkKind,CursorKind,ExceptionSpecification,DiagnosticSeverity,EvalResultKind,IdxAttrKind,IdxDeclInfoFlags,IdxEntityCXXTemplateKind,IdxEntityKind,IdxEntityLanguage,IdxEntityRefKind,IdxObjCContainerKind,IndexOptFlags,LanguageKind,LinkageKind,NameRefFlags,RefQualifierKind,Reparse_Flags,Result,SaveTranslationUnit_Flags,StorageClass,TemplateArgumentKind,TokenKind,TUResourceUsageKind,TypeKind,VisibilityKind,VisitorResult; DO NOT EDIT

package clang

import "strconv"

const _AccessSpecifier_name = "AccessSpecifier_InvalidAccessSpecifier_PublicAccessSpecifier_ProtectedAccessSpecifier_Private"

var _AccessSpecifier_index = [...]uint8{0, 23, 45, 70, 93}

func (i AccessSpecifier) String() string {
	if i >= AccessSpecifier(len(_AccessSpecifier_index)-1) {
		return "AccessSpecifier(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AccessSpecifier_name[_AccessSpecifier_index[i]:_AccessSpecifier_index[i+1]]
}

const _AvailabilityKind_name = "Availability_AvailableAvailability_DeprecatedAvailability_NotAvailableAvailability_NotAccessible"

var _AvailabilityKind_index = [...]uint8{0, 22, 45, 70, 96}

func (i AvailabilityKind) String() string {
	if i >= AvailabilityKind(len(_AvailabilityKind_index)-1) {
		return "AvailabilityKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _AvailabilityKind_name[_AvailabilityKind_index[i]:_AvailabilityKind_index[i+1]]
}

const (
	_CallingConv_name_0 = "CallingConv_DefaultCallingConv_CCallingConv_X86StdCallCallingConv_X86FastCallCallingConv_X86ThisCallCallingConv_X86PascalCallingConv_AAPCSCallingConv_AAPCS_VFPCallingConv_X86RegCallCallingConv_IntelOclBiccCallingConv_Win64CallingConv_X86_64SysVCallingConv_X86VectorCallCallingConv_SwiftCallingConv_PreserveMostCallingConv_PreserveAll"
	_CallingConv_name_1 = "CallingConv_Invalid"
	_CallingConv_name_2 = "CallingConv_Unexposed"
)

var (
	_CallingConv_index_0 = [...]uint16{0, 19, 32, 54, 77, 100, 121, 138, 159, 181, 205, 222, 244, 269, 286, 310, 333}
)

func (i CallingConv) String() string {
	switch {
	case 0 <= i && i <= 15:
		return _CallingConv_name_0[_CallingConv_index_0[i]:_CallingConv_index_0[i+1]]
	case i == 100:
		return _CallingConv_name_1
	case i == 200:
		return _CallingConv_name_2
	default:
		return "CallingConv(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

const _ChildVisitResult_name = "ChildVisit_BreakChildVisit_ContinueChildVisit_Recurse"

var _ChildVisitResult_index = [...]uint8{0, 16, 35, 53}

func (i ChildVisitResult) String() string {
	if i >= ChildVisitResult(len(_ChildVisitResult_index)-1) {
		return "ChildVisitResult(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ChildVisitResult_name[_ChildVisitResult_index[i]:_ChildVisitResult_index[i+1]]
}

const _CommentInlineCommandRenderKind_name = "CommentInlineCommandRenderKind_NormalCommentInlineCommandRenderKind_BoldCommentInlineCommandRenderKind_MonospacedCommentInlineCommandRenderKind_Emphasized"

var _CommentInlineCommandRenderKind_index = [...]uint8{0, 37, 72, 113, 154}

func (i CommentInlineCommandRenderKind) String() string {
	if i >= CommentInlineCommandRenderKind(len(_CommentInlineCommandRenderKind_index)-1) {
		return "CommentInlineCommandRenderKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CommentInlineCommandRenderKind_name[_CommentInlineCommandRenderKind_index[i]:_CommentInlineCommandRenderKind_index[i+1]]
}

const _CommentKind_name = "Comment_NullComment_TextComment_InlineCommandComment_HTMLStartTagComment_HTMLEndTagComment_ParagraphComment_BlockCommandComment_ParamCommandComment_TParamCommandComment_VerbatimBlockCommandComment_VerbatimBlockLineComment_VerbatimLineComment_FullComment"

var _CommentKind_index = [...]uint8{0, 12, 24, 45, 65, 83, 100, 120, 140, 161, 189, 214, 234, 253}

func (i CommentKind) String() string {
	if i >= CommentKind(len(_CommentKind_index)-1) {
		return "CommentKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CommentKind_name[_CommentKind_index[i]:_CommentKind_index[i+1]]
}

const _CommentParamPassDirection_name = "CommentParamPassDirection_InCommentParamPassDirection_OutCommentParamPassDirection_InOut"

var _CommentParamPassDirection_index = [...]uint8{0, 28, 57, 88}

func (i CommentParamPassDirection) String() string {
	if i >= CommentParamPassDirection(len(_CommentParamPassDirection_index)-1) {
		return "CommentParamPassDirection(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CommentParamPassDirection_name[_CommentParamPassDirection_index[i]:_CommentParamPassDirection_index[i+1]]
}

const _CompletionChunkKind_name = "CompletionChunk_OptionalCompletionChunk_TypedTextCompletionChunk_TextCompletionChunk_PlaceholderCompletionChunk_InformativeCompletionChunk_CurrentParameterCompletionChunk_LeftParenCompletionChunk_RightParenCompletionChunk_LeftBracketCompletionChunk_RightBracketCompletionChunk_LeftBraceCompletionChunk_RightBraceCompletionChunk_LeftAngleCompletionChunk_RightAngleCompletionChunk_CommaCompletionChunk_ResultTypeCompletionChunk_ColonCompletionChunk_SemiColonCompletionChunk_EqualCompletionChunk_HorizontalSpaceCompletionChunk_VerticalSpace"

var _CompletionChunkKind_index = [...]uint16{0, 24, 49, 69, 96, 123, 155, 180, 206, 233, 261, 286, 312, 337, 363, 384, 410, 431, 456, 477, 508, 537}

func (i CompletionChunkKind) String() string {
	if i >= CompletionChunkKind(len(_CompletionChunkKind_index)-1) {
		return "CompletionChunkKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CompletionChunkKind_name[_CompletionChunkKind_index[i]:_CompletionChunkKind_index[i+1]]
}

const (
	_CursorKind_name_0 = "Cursor_UnexposedDeclCursor_StructDeclCursor_UnionDeclCursor_ClassDeclCursor_EnumDeclCursor_FieldDeclCursor_EnumConstantDeclCursor_FunctionDeclCursor_VarDeclCursor_ParmDeclCursor_ObjCInterfaceDeclCursor_ObjCCategoryDeclCursor_ObjCProtocolDeclCursor_ObjCPropertyDeclCursor_ObjCIvarDeclCursor_ObjCInstanceMethodDeclCursor_ObjCClassMethodDeclCursor_ObjCImplementationDeclCursor_ObjCCategoryImplDeclCursor_TypedefDeclCursor_CXXMethodCursor_NamespaceCursor_LinkageSpecCursor_ConstructorCursor_DestructorCursor_ConversionFunctionCursor_TemplateTypeParameterCursor_NonTypeTemplateParameterCursor_TemplateTemplateParameterCursor_FunctionTemplateCursor_ClassTemplateCursor_ClassTemplatePartialSpecializationCursor_NamespaceAliasCursor_UsingDirectiveCursor_UsingDeclarationCursor_TypeAliasDeclCursor_ObjCSynthesizeDeclCursor_ObjCDynamicDeclCursor_CXXAccessSpecifierCursor_ObjCSuperClassRefCursor_ObjCProtocolRefCursor_ObjCClassRefCursor_TypeRefCursor_CXXBaseSpecifierCursor_TemplateRefCursor_NamespaceRefCursor_MemberRefCursor_LabelRefCursor_OverloadedDeclRefCursor_VariableRef"
	_CursorKind_name_1 = "Cursor_InvalidFileCursor_NoDeclFoundCursor_NotImplementedCursor_InvalidCode"
	_CursorKind_name_2 = "Cursor_UnexposedExprCursor_DeclRefExprCursor_MemberRefExprCursor_CallExprCursor_ObjCMessageExprCursor_BlockExprCursor_IntegerLiteralCursor_FloatingLiteralCursor_ImaginaryLiteralCursor_StringLiteralCursor_CharacterLiteralCursor_ParenExprCursor_UnaryOperatorCursor_ArraySubscriptExprCursor_BinaryOperatorCursor_CompoundAssignOperatorCursor_ConditionalOperatorCursor_CStyleCastExprCursor_CompoundLiteralExprCursor_InitListExprCursor_AddrLabelExprCursor_StmtExprCursor_GenericSelectionExprCursor_GNUNullExprCursor_CXXStaticCastExprCursor_CXXDynamicCastExprCursor_CXXReinterpretCastExprCursor_CXXConstCastExprCursor_CXXFunctionalCastExprCursor_CXXTypeidExprCursor_CXXBoolLiteralExprCursor_CXXNullPtrLiteralExprCursor_CXXThisExprCursor_CXXThrowExprCursor_CXXNewExprCursor_CXXDeleteExprCursor_UnaryExprCursor_ObjCStringLiteralCursor_ObjCEncodeExprCursor_ObjCSelectorExprCursor_ObjCProtocolExprCursor_ObjCBridgedCastExprCursor_PackExpansionExprCursor_SizeOfPackExprCursor_LambdaExprCursor_ObjCBoolLiteralExprCursor_ObjCSelfExprCursor_OMPArraySectionExprCursor_ObjCAvailabilityCheckExpr"
	_CursorKind_name_3 = "Cursor_UnexposedStmtCursor_LabelStmtCursor_CompoundStmtCursor_CaseStmtCursor_DefaultStmtCursor_IfStmtCursor_SwitchStmtCursor_WhileStmtCursor_DoStmtCursor_ForStmtCursor_GotoStmtCursor_IndirectGotoStmtCursor_ContinueStmtCursor_BreakStmtCursor_ReturnStmtCursor_GCCAsmStmtCursor_ObjCAtTryStmtCursor_ObjCAtCatchStmtCursor_ObjCAtFinallyStmtCursor_ObjCAtThrowStmtCursor_ObjCAtSynchronizedStmtCursor_ObjCAutoreleasePoolStmtCursor_ObjCForCollectionStmtCursor_CXXCatchStmtCursor_CXXTryStmtCursor_CXXForRangeStmtCursor_SEHTryStmtCursor_SEHExceptStmtCursor_SEHFinallyStmtCursor_MSAsmStmtCursor_NullStmtCursor_DeclStmtCursor_OMPParallelDirectiveCursor_OMPSimdDirectiveCursor_OMPForDirectiveCursor_OMPSectionsDirectiveCursor_OMPSectionDirectiveCursor_OMPSingleDirectiveCursor_OMPParallelForDirectiveCursor_OMPParallelSectionsDirectiveCursor_OMPTaskDirectiveCursor_OMPMasterDirectiveCursor_OMPCriticalDirectiveCursor_OMPTaskyieldDirectiveCursor_OMPBarrierDirectiveCursor_OMPTaskwaitDirectiveCursor_OMPFlushDirectiveCursor_SEHLeaveStmtCursor_OMPOrderedDirectiveCursor_OMPAtomicDirectiveCursor_OMPForSimdDirectiveCursor_OMPParallelForSimdDirectiveCursor_OMPTargetDirectiveCursor_OMPTeamsDirectiveCursor_OMPTaskgroupDirectiveCursor_OMPCancellationPointDirectiveCursor_OMPCancelDirectiveCursor_OMPTargetDataDirectiveCursor_OMPTaskLoopDirectiveCursor_OMPTaskLoopSimdDirectiveCursor_OMPDistributeDirectiveCursor_OMPTargetEnterDataDirectiveCursor_OMPTargetExitDataDirectiveCursor_OMPTargetParallelDirectiveCursor_OMPTargetParallelForDirectiveCursor_OMPTargetUpdateDirectiveCursor_OMPDistributeParallelForDirectiveCursor_OMPDistributeParallelForSimdDirectiveCursor_OMPDistributeSimdDirectiveCursor_OMPTargetParallelForSimdDirectiveCursor_OMPTargetSimdDirectiveCursor_OMPTeamsDistributeDirectiveCursor_OMPTeamsDistributeSimdDirectiveCursor_OMPTeamsDistributeParallelForSimdDirectiveCursor_OMPTeamsDistributeParallelForDirectiveCursor_OMPTargetTeamsDirectiveCursor_OMPTargetTeamsDistributeDirectiveCursor_OMPTargetTeamsDistributeParallelForDirectiveCursor_OMPTargetTeamsDistributeParallelForSimdDirectiveCursor_OMPTargetTeamsDistributeSimdDirective"
	_CursorKind_name_4 = "Cursor_TranslationUnit"
	_CursorKind_name_5 = "Cursor_UnexposedAttrCursor_IBActionAttrCursor_IBOutletAttrCursor_IBOutletCollectionAttrCursor_CXXFinalAttrCursor_CXXOverrideAttrCursor_AnnotateAttrCursor_AsmLabelAttrCursor_PackedAttrCursor_PureAttrCursor_ConstAttrCursor_NoDuplicateAttrCursor_CUDAConstantAttrCursor_CUDADeviceAttrCursor_CUDAGlobalAttrCursor_CUDAHostAttrCursor_CUDASharedAttrCursor_VisibilityAttrCursor_DLLExportCursor_DLLImport"
	_CursorKind_name_6 = "Cursor_PreprocessingDirectiveCursor_MacroDefinitionCursor_MacroExpansionCursor_InclusionDirective"
	_CursorKind_name_7 = "Cursor_ModuleImportDeclCursor_TypeAliasTemplateDeclCursor_StaticAssertCursor_FriendDecl"
	_CursorKind_name_8 = "Cursor_OverloadCandidate"
)

var (
	_CursorKind_index_0 = [...]uint16{0, 20, 37, 53, 69, 84, 100, 123, 142, 156, 171, 195, 218, 241, 264, 283, 312, 338, 367, 394, 412, 428, 444, 462, 480, 497, 522, 550, 581, 613, 636, 656, 697, 718, 739, 762, 782, 807, 829, 854, 878, 900, 919, 933, 956, 974, 993, 1009, 1024, 1048, 1066}
	_CursorKind_index_1 = [...]uint8{0, 18, 36, 57, 75}
	_CursorKind_index_2 = [...]uint16{0, 20, 38, 58, 73, 95, 111, 132, 154, 177, 197, 220, 236, 256, 281, 302, 331, 357, 378, 404, 423, 443, 458, 485, 503, 527, 552, 581, 604, 632, 652, 677, 705, 723, 742, 759, 779, 795, 819, 840, 863, 886, 912, 936, 957, 974, 1000, 1019, 1045, 1077}
	_CursorKind_index_3 = [...]uint16{0, 20, 36, 55, 70, 88, 101, 118, 134, 147, 161, 176, 199, 218, 234, 251, 268, 288, 310, 334, 356, 385, 415, 443, 462, 479, 501, 518, 538, 559, 575, 590, 605, 632, 655, 677, 704, 730, 755, 785, 820, 843, 868, 895, 923, 949, 976, 1000, 1019, 1045, 1070, 1096, 1130, 1155, 1179, 1207, 1243, 1268, 1297, 1324, 1355, 1384, 1418, 1451, 1484, 1520, 1551, 1591, 1635, 1668, 1708, 1737, 1771, 1809, 1858, 1903, 1933, 1973, 2024, 2079, 2123}
	_CursorKind_index_5 = [...]uint16{0, 20, 39, 58, 87, 106, 128, 147, 166, 183, 198, 214, 236, 259, 280, 301, 320, 341, 362, 378, 394}
	_CursorKind_index_6 = [...]uint8{0, 29, 51, 72, 97}
	_CursorKind_index_7 = [...]uint8{0, 23, 51, 70, 87}
)

func (i CursorKind) String() string {
	switch {
	case 1 <= i && i <= 50:
		i -= 1
		return _CursorKind_name_0[_CursorKind_index_0[i]:_CursorKind_index_0[i+1]]
	case 70 <= i && i <= 73:
		i -= 70
		return _CursorKind_name_1[_CursorKind_index_1[i]:_CursorKind_index_1[i+1]]
	case 100 <= i && i <= 148:
		i -= 100
		return _CursorKind_name_2[_CursorKind_index_2[i]:_CursorKind_index_2[i+1]]
	case 200 <= i && i <= 279:
		i -= 200
		return _CursorKind_name_3[_CursorKind_index_3[i]:_CursorKind_index_3[i+1]]
	case i == 300:
		return _CursorKind_name_4
	case 400 <= i && i <= 419:
		i -= 400
		return _CursorKind_name_5[_CursorKind_index_5[i]:_CursorKind_index_5[i+1]]
	case 500 <= i && i <= 503:
		i -= 500
		return _CursorKind_name_6[_CursorKind_index_6[i]:_CursorKind_index_6[i+1]]
	case 600 <= i && i <= 603:
		i -= 600
		return _CursorKind_name_7[_CursorKind_index_7[i]:_CursorKind_index_7[i+1]]
	case i == 700:
		return _CursorKind_name_8
	default:
		return "CursorKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

const _ExceptionSpecification_name = "ExceptionSpecification_NonFunctionExceptionSpecification_NoneExceptionSpecification_DynamicNoneExceptionSpecification_DynamicExceptionSpecification_MSAnyExceptionSpecification_BasicNoexceptExceptionSpecification_ComputedNoexceptExceptionSpecification_UnevaluatedExceptionSpecification_UninstantiatedExceptionSpecification_Unparsed"

var _ExceptionSpecification_index = [...]uint16{0, 34, 61, 95, 125, 153, 189, 228, 262, 299, 330}

func (i ExceptionSpecification) String() string {
	i -= -1
	if i < 0 || i >= ExceptionSpecification(len(_ExceptionSpecification_index)-1) {
		return "ExceptionSpecification(" + strconv.FormatInt(int64(i+-1), 10) + ")"
	}
	return _ExceptionSpecification_name[_ExceptionSpecification_index[i]:_ExceptionSpecification_index[i+1]]
}

const _DiagnosticSeverity_name = "Diagnostic_IgnoredDiagnostic_NoteDiagnostic_WarningDiagnostic_ErrorDiagnostic_Fatal"

var _DiagnosticSeverity_index = [...]uint8{0, 18, 33, 51, 67, 83}

func (i DiagnosticSeverity) String() string {
	if i >= DiagnosticSeverity(len(_DiagnosticSeverity_index)-1) {
		return "DiagnosticSeverity(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DiagnosticSeverity_name[_DiagnosticSeverity_index[i]:_DiagnosticSeverity_index[i+1]]
}

const _EvalResultKind_name = "Eval_UnExposedEval_IntEval_FloatEval_ObjCStrLiteralEval_StrLiteralEval_CFStrEval_Other"

var _EvalResultKind_index = [...]uint8{0, 14, 22, 32, 51, 66, 76, 86}

func (i EvalResultKind) String() string {
	if i >= EvalResultKind(len(_EvalResultKind_index)-1) {
		return "EvalResultKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EvalResultKind_name[_EvalResultKind_index[i]:_EvalResultKind_index[i+1]]
}

const _IdxAttrKind_name = "IdxAttr_UnexposedIdxAttr_IBActionIdxAttr_IBOutletIdxAttr_IBOutletCollection"

var _IdxAttrKind_index = [...]uint8{0, 17, 33, 49, 75}

func (i IdxAttrKind) String() string {
	if i >= IdxAttrKind(len(_IdxAttrKind_index)-1) {
		return "IdxAttrKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdxAttrKind_name[_IdxAttrKind_index[i]:_IdxAttrKind_index[i+1]]
}

const _IdxDeclInfoFlags_name = "IdxDeclFlag_Skipped"

var _IdxDeclInfoFlags_index = [...]uint8{0, 19}

func (i IdxDeclInfoFlags) String() string {
	i -= 1
	if i >= IdxDeclInfoFlags(len(_IdxDeclInfoFlags_index)-1) {
		return "IdxDeclInfoFlags(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _IdxDeclInfoFlags_name[_IdxDeclInfoFlags_index[i]:_IdxDeclInfoFlags_index[i+1]]
}

const _IdxEntityCXXTemplateKind_name = "IdxEntity_NonTemplateIdxEntity_TemplateIdxEntity_TemplatePartialSpecializationIdxEntity_TemplateSpecialization"

var _IdxEntityCXXTemplateKind_index = [...]uint8{0, 21, 39, 78, 110}

func (i IdxEntityCXXTemplateKind) String() string {
	if i >= IdxEntityCXXTemplateKind(len(_IdxEntityCXXTemplateKind_index)-1) {
		return "IdxEntityCXXTemplateKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdxEntityCXXTemplateKind_name[_IdxEntityCXXTemplateKind_index[i]:_IdxEntityCXXTemplateKind_index[i+1]]
}

const _IdxEntityKind_name = "IdxEntity_UnexposedIdxEntity_TypedefIdxEntity_FunctionIdxEntity_VariableIdxEntity_FieldIdxEntity_EnumConstantIdxEntity_ObjCClassIdxEntity_ObjCProtocolIdxEntity_ObjCCategoryIdxEntity_ObjCInstanceMethodIdxEntity_ObjCClassMethodIdxEntity_ObjCPropertyIdxEntity_ObjCIvarIdxEntity_EnumIdxEntity_StructIdxEntity_UnionIdxEntity_CXXClassIdxEntity_CXXNamespaceIdxEntity_CXXNamespaceAliasIdxEntity_CXXStaticVariableIdxEntity_CXXStaticMethodIdxEntity_CXXInstanceMethodIdxEntity_CXXConstructorIdxEntity_CXXDestructorIdxEntity_CXXConversionFunctionIdxEntity_CXXTypeAliasIdxEntity_CXXInterface"

var _IdxEntityKind_index = [...]uint16{0, 19, 36, 54, 72, 87, 109, 128, 150, 172, 200, 225, 247, 265, 279, 295, 310, 328, 350, 377, 404, 429, 456, 480, 503, 534, 556, 578}

func (i IdxEntityKind) String() string {
	if i >= IdxEntityKind(len(_IdxEntityKind_index)-1) {
		return "IdxEntityKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdxEntityKind_name[_IdxEntityKind_index[i]:_IdxEntityKind_index[i+1]]
}

const _IdxEntityLanguage_name = "IdxEntityLang_NoneIdxEntityLang_CIdxEntityLang_ObjCIdxEntityLang_CXXIdxEntityLang_Swift"

var _IdxEntityLanguage_index = [...]uint8{0, 18, 33, 51, 68, 87}

func (i IdxEntityLanguage) String() string {
	if i >= IdxEntityLanguage(len(_IdxEntityLanguage_index)-1) {
		return "IdxEntityLanguage(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdxEntityLanguage_name[_IdxEntityLanguage_index[i]:_IdxEntityLanguage_index[i+1]]
}

const _IdxEntityRefKind_name = "IdxEntityRef_DirectIdxEntityRef_Implicit"

var _IdxEntityRefKind_index = [...]uint8{0, 19, 40}

func (i IdxEntityRefKind) String() string {
	i -= 1
	if i >= IdxEntityRefKind(len(_IdxEntityRefKind_index)-1) {
		return "IdxEntityRefKind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _IdxEntityRefKind_name[_IdxEntityRefKind_index[i]:_IdxEntityRefKind_index[i+1]]
}

const _IdxObjCContainerKind_name = "IdxObjCContainer_ForwardRefIdxObjCContainer_InterfaceIdxObjCContainer_Implementation"

var _IdxObjCContainerKind_index = [...]uint8{0, 27, 53, 84}

func (i IdxObjCContainerKind) String() string {
	if i >= IdxObjCContainerKind(len(_IdxObjCContainerKind_index)-1) {
		return "IdxObjCContainerKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IdxObjCContainerKind_name[_IdxObjCContainerKind_index[i]:_IdxObjCContainerKind_index[i+1]]
}

const (
	_IndexOptFlags_name_0 = "IndexOpt_NoneIndexOpt_SuppressRedundantRefsIndexOpt_IndexFunctionLocalSymbols"
	_IndexOptFlags_name_1 = "IndexOpt_IndexImplicitTemplateInstantiations"
	_IndexOptFlags_name_2 = "IndexOpt_SuppressWarnings"
	_IndexOptFlags_name_3 = "IndexOpt_SkipParsedBodiesInSession"
)

var (
	_IndexOptFlags_index_0 = [...]uint8{0, 13, 43, 77}
)

func (i IndexOptFlags) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _IndexOptFlags_name_0[_IndexOptFlags_index_0[i]:_IndexOptFlags_index_0[i+1]]
	case i == 4:
		return _IndexOptFlags_name_1
	case i == 8:
		return _IndexOptFlags_name_2
	case i == 16:
		return _IndexOptFlags_name_3
	default:
		return "IndexOptFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

const _LanguageKind_name = "Language_InvalidLanguage_CLanguage_ObjCLanguage_CPlusPlus"

var _LanguageKind_index = [...]uint8{0, 16, 26, 39, 57}

func (i LanguageKind) String() string {
	if i >= LanguageKind(len(_LanguageKind_index)-1) {
		return "LanguageKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LanguageKind_name[_LanguageKind_index[i]:_LanguageKind_index[i+1]]
}

const _LinkageKind_name = "Linkage_InvalidLinkage_NoLinkageLinkage_InternalLinkage_UniqueExternalLinkage_External"

var _LinkageKind_index = [...]uint8{0, 15, 32, 48, 70, 86}

func (i LinkageKind) String() string {
	if i >= LinkageKind(len(_LinkageKind_index)-1) {
		return "LinkageKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LinkageKind_name[_LinkageKind_index[i]:_LinkageKind_index[i+1]]
}

const (
	_NameRefFlags_name_0 = "NameRange_WantQualifierNameRange_WantTemplateArgs"
	_NameRefFlags_name_1 = "NameRange_WantSinglePiece"
)

var (
	_NameRefFlags_index_0 = [...]uint8{0, 23, 49}
)

func (i NameRefFlags) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _NameRefFlags_name_0[_NameRefFlags_index_0[i]:_NameRefFlags_index_0[i+1]]
	case i == 4:
		return _NameRefFlags_name_1
	default:
		return "NameRefFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

const _RefQualifierKind_name = "RefQualifier_NoneRefQualifier_LValueRefQualifier_RValue"

var _RefQualifierKind_index = [...]uint8{0, 17, 36, 55}

func (i RefQualifierKind) String() string {
	if i >= RefQualifierKind(len(_RefQualifierKind_index)-1) {
		return "RefQualifierKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RefQualifierKind_name[_RefQualifierKind_index[i]:_RefQualifierKind_index[i+1]]
}

const _Reparse_Flags_name = "Reparse_None"

var _Reparse_Flags_index = [...]uint8{0, 12}

func (i Reparse_Flags) String() string {
	if i >= Reparse_Flags(len(_Reparse_Flags_index)-1) {
		return "Reparse_Flags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Reparse_Flags_name[_Reparse_Flags_index[i]:_Reparse_Flags_index[i+1]]
}

const _Result_name = "Result_SuccessResult_InvalidResult_VisitBreak"

var _Result_index = [...]uint8{0, 14, 28, 45}

func (i Result) String() string {
	if i >= Result(len(_Result_index)-1) {
		return "Result(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Result_name[_Result_index[i]:_Result_index[i+1]]
}

const _SaveTranslationUnit_Flags_name = "SaveTranslationUnit_None"

var _SaveTranslationUnit_Flags_index = [...]uint8{0, 24}

func (i SaveTranslationUnit_Flags) String() string {
	if i >= SaveTranslationUnit_Flags(len(_SaveTranslationUnit_Flags_index)-1) {
		return "SaveTranslationUnit_Flags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SaveTranslationUnit_Flags_name[_SaveTranslationUnit_Flags_index[i]:_SaveTranslationUnit_Flags_index[i+1]]
}

const _StorageClass_name = "SC_InvalidSC_NoneSC_ExternSC_StaticSC_PrivateExternSC_OpenCLWorkGroupLocalSC_AutoSC_Register"

var _StorageClass_index = [...]uint8{0, 10, 17, 26, 35, 51, 74, 81, 92}

func (i StorageClass) String() string {
	if i >= StorageClass(len(_StorageClass_index)-1) {
		return "StorageClass(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _StorageClass_name[_StorageClass_index[i]:_StorageClass_index[i+1]]
}

const _TemplateArgumentKind_name = "TemplateArgumentKind_NullTemplateArgumentKind_TypeTemplateArgumentKind_DeclarationTemplateArgumentKind_NullPtrTemplateArgumentKind_IntegralTemplateArgumentKind_TemplateTemplateArgumentKind_TemplateExpansionTemplateArgumentKind_ExpressionTemplateArgumentKind_PackTemplateArgumentKind_Invalid"

var _TemplateArgumentKind_index = [...]uint16{0, 25, 50, 82, 110, 139, 168, 206, 237, 262, 290}

func (i TemplateArgumentKind) String() string {
	if i >= TemplateArgumentKind(len(_TemplateArgumentKind_index)-1) {
		return "TemplateArgumentKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TemplateArgumentKind_name[_TemplateArgumentKind_index[i]:_TemplateArgumentKind_index[i+1]]
}

const _TokenKind_name = "Token_PunctuationToken_KeywordToken_IdentifierToken_LiteralToken_Comment"

var _TokenKind_index = [...]uint8{0, 17, 30, 46, 59, 72}

func (i TokenKind) String() string {
	if i >= TokenKind(len(_TokenKind_index)-1) {
		return "TokenKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenKind_name[_TokenKind_index[i]:_TokenKind_index[i+1]]
}

const _TUResourceUsageKind_name = "TUResourceUsage_ASTTUResourceUsage_IdentifiersTUResourceUsage_SelectorsTUResourceUsage_GlobalCompletionResultsTUResourceUsage_SourceManagerContentCacheTUResourceUsage_AST_SideTablesTUResourceUsage_SourceManager_Membuffer_MallocTUResourceUsage_SourceManager_Membuffer_MMapTUResourceUsage_ExternalASTSource_Membuffer_MallocTUResourceUsage_ExternalASTSource_Membuffer_MMapTUResourceUsage_PreprocessorTUResourceUsage_PreprocessingRecordTUResourceUsage_SourceManager_DataStructuresTUResourceUsage_Preprocessor_HeaderSearch"

var _TUResourceUsageKind_index = [...]uint16{0, 19, 46, 71, 110, 151, 181, 227, 271, 321, 369, 397, 432, 476, 517}

func (i TUResourceUsageKind) String() string {
	i -= 1
	if i >= TUResourceUsageKind(len(_TUResourceUsageKind_index)-1) {
		return "TUResourceUsageKind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _TUResourceUsageKind_name[_TUResourceUsageKind_index[i]:_TUResourceUsageKind_index[i+1]]
}

const (
	_TypeKind_name_0 = "Type_InvalidType_UnexposedType_VoidType_BoolType_Char_UType_UCharType_Char16Type_Char32Type_UShortType_UIntType_ULongType_ULongLongType_UInt128Type_Char_SType_SCharType_WCharType_ShortType_IntType_LongType_LongLongType_Int128Type_FloatType_DoubleType_LongDoubleType_NullPtrType_OverloadType_DependentType_ObjCIdType_ObjCClassType_ObjCSelType_Float128Type_Half"
	_TypeKind_name_1 = "Type_ComplexType_PointerType_BlockPointerType_LValueReferenceType_RValueReferenceType_RecordType_EnumType_TypedefType_ObjCInterfaceType_ObjCObjectPointerType_FunctionNoProtoType_FunctionProtoType_ConstantArrayType_VectorType_IncompleteArrayType_VariableArrayType_DependentSizedArrayType_MemberPointerType_AutoType_ElaboratedType_PipeType_OCLImage1dROType_OCLImage1dArrayROType_OCLImage1dBufferROType_OCLImage2dROType_OCLImage2dArrayROType_OCLImage2dDepthROType_OCLImage2dArrayDepthROType_OCLImage2dMSAAROType_OCLImage2dArrayMSAAROType_OCLImage2dMSAADepthROType_OCLImage2dArrayMSAADepthROType_OCLImage3dROType_OCLImage1dWOType_OCLImage1dArrayWOType_OCLImage1dBufferWOType_OCLImage2dWOType_OCLImage2dArrayWOType_OCLImage2dDepthWOType_OCLImage2dArrayDepthWOType_OCLImage2dMSAAWOType_OCLImage2dArrayMSAAWOType_OCLImage2dMSAADepthWOType_OCLImage2dArrayMSAADepthWOType_OCLImage3dWOType_OCLImage1dRWType_OCLImage1dArrayRWType_OCLImage1dBufferRWType_OCLImage2dRWType_OCLImage2dArrayRWType_OCLImage2dDepthRWType_OCLImage2dArrayDepthRWType_OCLImage2dMSAARWType_OCLImage2dArrayMSAARWType_OCLImage2dMSAADepthRWType_OCLImage2dArrayMSAADepthRWType_OCLImage3dRWType_OCLSamplerType_OCLEventType_OCLQueueType_OCLReserveID"
)

var (
	_TypeKind_index_0 = [...]uint16{0, 12, 26, 35, 44, 55, 65, 76, 87, 98, 107, 117, 131, 143, 154, 164, 174, 184, 192, 201, 214, 225, 235, 246, 261, 273, 286, 300, 311, 325, 337, 350, 359}
	_TypeKind_index_1 = [...]uint16{0, 12, 24, 41, 61, 81, 92, 101, 113, 131, 153, 173, 191, 209, 220, 240, 258, 282, 300, 309, 324, 333, 350, 372, 395, 412, 434, 456, 483, 504, 530, 556, 587, 604, 621, 643, 666, 683, 705, 727, 754, 775, 801, 827, 858, 875, 892, 914, 937, 954, 976, 998, 1025, 1046, 1072, 1098, 1129, 1146, 1161, 1174, 1187, 1204}
)

func (i TypeKind) String() string {
	switch {
	case 0 <= i && i <= 31:
		return _TypeKind_name_0[_TypeKind_index_0[i]:_TypeKind_index_0[i+1]]
	case 100 <= i && i <= 160:
		i -= 100
		return _TypeKind_name_1[_TypeKind_index_1[i]:_TypeKind_index_1[i+1]]
	default:
		return "TypeKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

const _VisibilityKind_name = "Visibility_InvalidVisibility_HiddenVisibility_ProtectedVisibility_Default"

var _VisibilityKind_index = [...]uint8{0, 18, 35, 55, 73}

func (i VisibilityKind) String() string {
	if i >= VisibilityKind(len(_VisibilityKind_index)-1) {
		return "VisibilityKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VisibilityKind_name[_VisibilityKind_index[i]:_VisibilityKind_index[i+1]]
}

const _VisitorResult_name = "Visit_BreakVisit_Continue"

var _VisitorResult_index = [...]uint8{0, 11, 25}

func (i VisitorResult) String() string {
	if i >= VisitorResult(len(_VisitorResult_index)-1) {
		return "VisitorResult(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _VisitorResult_name[_VisitorResult_index[i]:_VisitorResult_index[i+1]]
}
