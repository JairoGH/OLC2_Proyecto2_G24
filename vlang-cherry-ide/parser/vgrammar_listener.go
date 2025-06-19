// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VGrammar
import "github.com/antlr4-go/antlr/v4"

// VGrammarListener is a complete listener for a parse tree produced by VGrammar.
type VGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterFuncionMain is called when entering the FuncionMain production.
	EnterFuncionMain(c *FuncionMainContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterDeclararInferenciaMut is called when entering the DeclararInferenciaMut production.
	EnterDeclararInferenciaMut(c *DeclararInferenciaMutContext)

	// EnterDeclararInferencia is called when entering the DeclararInferencia production.
	EnterDeclararInferencia(c *DeclararInferenciaContext)

	// EnterDeclaraTipoValor is called when entering the DeclaraTipoValor production.
	EnterDeclaraTipoValor(c *DeclaraTipoValorContext)

	// EnterDeclararTipo is called when entering the DeclararTipo production.
	EnterDeclararTipo(c *DeclararTipoContext)

	// EnterDeclararSinMutValor is called when entering the DeclararSinMutValor production.
	EnterDeclararSinMutValor(c *DeclararSinMutValorContext)

	// EnterDeclararSlice is called when entering the DeclararSlice production.
	EnterDeclararSlice(c *DeclararSliceContext)

	// EnterAsignacionSliceItem is called when entering the AsignacionSliceItem production.
	EnterAsignacionSliceItem(c *AsignacionSliceItemContext)

	// EnterAsignacionDirecta is called when entering the AsignacionDirecta production.
	EnterAsignacionDirecta(c *AsignacionDirectaContext)

	// EnterAsignacionAritmetica is called when entering the AsignacionAritmetica production.
	EnterAsignacionAritmetica(c *AsignacionAritmeticaContext)

	// EnterAsignacionSlice is called when entering the AsignacionSlice production.
	EnterAsignacionSlice(c *AsignacionSliceContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterBreakStmt is called when entering the BreakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the ContinueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterIFstmt is called when entering the IFstmt production.
	EnterIFstmt(c *IFstmtContext)

	// EnterIFcadena is called when entering the IFcadena production.
	EnterIFcadena(c *IFcadenaContext)

	// EnterElseStmt is called when entering the ElseStmt production.
	EnterElseStmt(c *ElseStmtContext)

	// EnterSwitchStmt is called when entering the SwitchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterSwitchCase is called when entering the SwitchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterDefaultCase is called when entering the DefaultCase production.
	EnterDefaultCase(c *DefaultCaseContext)

	// EnterWhileStmt is called when entering the WhileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterForClasicoStmt is called when entering the ForClasicoStmt production.
	EnterForClasicoStmt(c *ForClasicoStmtContext)

	// EnterForStmt is called when entering the ForStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterRangoNum is called when entering the RangoNum production.
	EnterRangoNum(c *RangoNumContext)

	// EnterFuncionDeclerada is called when entering the FuncionDeclerada production.
	EnterFuncionDeclerada(c *FuncionDecleradaContext)

	// EnterLlamarFuncion is called when entering the LlamarFuncion production.
	EnterLlamarFuncion(c *LlamarFuncionContext)

	// EnterListaParametros is called when entering the ListaParametros production.
	EnterListaParametros(c *ListaParametrosContext)

	// EnterParametroFun is called when entering the ParametroFun production.
	EnterParametroFun(c *ParametroFunContext)

	// EnterListaArgumentos is called when entering the ListaArgumentos production.
	EnterListaArgumentos(c *ListaArgumentosContext)

	// EnterFuncionArg is called when entering the FuncionArg production.
	EnterFuncionArg(c *FuncionArgContext)

	// EnterDeclararStruct is called when entering the DeclararStruct production.
	EnterDeclararStruct(c *DeclararStructContext)

	// EnterPropiedadStruct is called when entering the PropiedadStruct production.
	EnterPropiedadStruct(c *PropiedadStructContext)

	// EnterCrearStruct is called when entering the CrearStruct production.
	EnterCrearStruct(c *CrearStructContext)

	// EnterListaParametrosInit is called when entering the ListaParametrosInit production.
	EnterListaParametrosInit(c *ListaParametrosInitContext)

	// EnterParametrosInitStruct is called when entering the ParametrosInitStruct production.
	EnterParametrosInitStruct(c *ParametrosInitStructContext)

	// EnterListaSlice is called when entering the ListaSlice production.
	EnterListaSlice(c *ListaSliceContext)

	// EnterItemSlice is called when entering the ItemSlice production.
	EnterItemSlice(c *ItemSliceContext)

	// EnterPropSlice is called when entering the PropSlice production.
	EnterPropSlice(c *PropSliceContext)

	// EnterFuncionSlice is called when entering the FuncionSlice production.
	EnterFuncionSlice(c *FuncionSliceContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterVectorSimple is called when entering the VectorSimple production.
	EnterVectorSimple(c *VectorSimpleContext)

	// EnterMatrizDoble is called when entering the MatrizDoble production.
	EnterMatrizDoble(c *MatrizDobleContext)

	// EnterLiteralExp is called when entering the LiteralExp production.
	EnterLiteralExp(c *LiteralExpContext)

	// EnterIdExp is called when entering the IdExp production.
	EnterIdExp(c *IdExpContext)

	// EnterCrearStructExp is called when entering the CrearStructExp production.
	EnterCrearStructExp(c *CrearStructExpContext)

	// EnterUnarioExp is called when entering the UnarioExp production.
	EnterUnarioExp(c *UnarioExpContext)

	// EnterItemSliceExp is called when entering the ItemSliceExp production.
	EnterItemSliceExp(c *ItemSliceExpContext)

	// EnterVectorFuncExp is called when entering the VectorFuncExp production.
	EnterVectorFuncExp(c *VectorFuncExpContext)

	// EnterSliceExp is called when entering the SliceExp production.
	EnterSliceExp(c *SliceExpContext)

	// EnterPropSliceExp is called when entering the PropSliceExp production.
	EnterPropSliceExp(c *PropSliceExpContext)

	// EnterLlamarFuncionExp is called when entering the LlamarFuncionExp production.
	EnterLlamarFuncionExp(c *LlamarFuncionExpContext)

	// EnterBinarioExp is called when entering the BinarioExp production.
	EnterBinarioExp(c *BinarioExpContext)

	// EnterParentecisExp is called when entering the ParentecisExp production.
	EnterParentecisExp(c *ParentecisExpContext)

	// EnterIntLiteral is called when entering the IntLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterFloatLiteral is called when entering the FloatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterBoolLiteral is called when entering the BoolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterNilLiteral is called when entering the NilLiteral production.
	EnterNilLiteral(c *NilLiteralContext)

	// EnterID_Patron is called when entering the ID_Patron production.
	EnterID_Patron(c *ID_PatronContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitFuncionMain is called when exiting the FuncionMain production.
	ExitFuncionMain(c *FuncionMainContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitDeclararInferenciaMut is called when exiting the DeclararInferenciaMut production.
	ExitDeclararInferenciaMut(c *DeclararInferenciaMutContext)

	// ExitDeclararInferencia is called when exiting the DeclararInferencia production.
	ExitDeclararInferencia(c *DeclararInferenciaContext)

	// ExitDeclaraTipoValor is called when exiting the DeclaraTipoValor production.
	ExitDeclaraTipoValor(c *DeclaraTipoValorContext)

	// ExitDeclararTipo is called when exiting the DeclararTipo production.
	ExitDeclararTipo(c *DeclararTipoContext)

	// ExitDeclararSinMutValor is called when exiting the DeclararSinMutValor production.
	ExitDeclararSinMutValor(c *DeclararSinMutValorContext)

	// ExitDeclararSlice is called when exiting the DeclararSlice production.
	ExitDeclararSlice(c *DeclararSliceContext)

	// ExitAsignacionSliceItem is called when exiting the AsignacionSliceItem production.
	ExitAsignacionSliceItem(c *AsignacionSliceItemContext)

	// ExitAsignacionDirecta is called when exiting the AsignacionDirecta production.
	ExitAsignacionDirecta(c *AsignacionDirectaContext)

	// ExitAsignacionAritmetica is called when exiting the AsignacionAritmetica production.
	ExitAsignacionAritmetica(c *AsignacionAritmeticaContext)

	// ExitAsignacionSlice is called when exiting the AsignacionSlice production.
	ExitAsignacionSlice(c *AsignacionSliceContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitBreakStmt is called when exiting the BreakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the ContinueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitIFstmt is called when exiting the IFstmt production.
	ExitIFstmt(c *IFstmtContext)

	// ExitIFcadena is called when exiting the IFcadena production.
	ExitIFcadena(c *IFcadenaContext)

	// ExitElseStmt is called when exiting the ElseStmt production.
	ExitElseStmt(c *ElseStmtContext)

	// ExitSwitchStmt is called when exiting the SwitchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitSwitchCase is called when exiting the SwitchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitDefaultCase is called when exiting the DefaultCase production.
	ExitDefaultCase(c *DefaultCaseContext)

	// ExitWhileStmt is called when exiting the WhileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitForClasicoStmt is called when exiting the ForClasicoStmt production.
	ExitForClasicoStmt(c *ForClasicoStmtContext)

	// ExitForStmt is called when exiting the ForStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitRangoNum is called when exiting the RangoNum production.
	ExitRangoNum(c *RangoNumContext)

	// ExitFuncionDeclerada is called when exiting the FuncionDeclerada production.
	ExitFuncionDeclerada(c *FuncionDecleradaContext)

	// ExitLlamarFuncion is called when exiting the LlamarFuncion production.
	ExitLlamarFuncion(c *LlamarFuncionContext)

	// ExitListaParametros is called when exiting the ListaParametros production.
	ExitListaParametros(c *ListaParametrosContext)

	// ExitParametroFun is called when exiting the ParametroFun production.
	ExitParametroFun(c *ParametroFunContext)

	// ExitListaArgumentos is called when exiting the ListaArgumentos production.
	ExitListaArgumentos(c *ListaArgumentosContext)

	// ExitFuncionArg is called when exiting the FuncionArg production.
	ExitFuncionArg(c *FuncionArgContext)

	// ExitDeclararStruct is called when exiting the DeclararStruct production.
	ExitDeclararStruct(c *DeclararStructContext)

	// ExitPropiedadStruct is called when exiting the PropiedadStruct production.
	ExitPropiedadStruct(c *PropiedadStructContext)

	// ExitCrearStruct is called when exiting the CrearStruct production.
	ExitCrearStruct(c *CrearStructContext)

	// ExitListaParametrosInit is called when exiting the ListaParametrosInit production.
	ExitListaParametrosInit(c *ListaParametrosInitContext)

	// ExitParametrosInitStruct is called when exiting the ParametrosInitStruct production.
	ExitParametrosInitStruct(c *ParametrosInitStructContext)

	// ExitListaSlice is called when exiting the ListaSlice production.
	ExitListaSlice(c *ListaSliceContext)

	// ExitItemSlice is called when exiting the ItemSlice production.
	ExitItemSlice(c *ItemSliceContext)

	// ExitPropSlice is called when exiting the PropSlice production.
	ExitPropSlice(c *PropSliceContext)

	// ExitFuncionSlice is called when exiting the FuncionSlice production.
	ExitFuncionSlice(c *FuncionSliceContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitVectorSimple is called when exiting the VectorSimple production.
	ExitVectorSimple(c *VectorSimpleContext)

	// ExitMatrizDoble is called when exiting the MatrizDoble production.
	ExitMatrizDoble(c *MatrizDobleContext)

	// ExitLiteralExp is called when exiting the LiteralExp production.
	ExitLiteralExp(c *LiteralExpContext)

	// ExitIdExp is called when exiting the IdExp production.
	ExitIdExp(c *IdExpContext)

	// ExitCrearStructExp is called when exiting the CrearStructExp production.
	ExitCrearStructExp(c *CrearStructExpContext)

	// ExitUnarioExp is called when exiting the UnarioExp production.
	ExitUnarioExp(c *UnarioExpContext)

	// ExitItemSliceExp is called when exiting the ItemSliceExp production.
	ExitItemSliceExp(c *ItemSliceExpContext)

	// ExitVectorFuncExp is called when exiting the VectorFuncExp production.
	ExitVectorFuncExp(c *VectorFuncExpContext)

	// ExitSliceExp is called when exiting the SliceExp production.
	ExitSliceExp(c *SliceExpContext)

	// ExitPropSliceExp is called when exiting the PropSliceExp production.
	ExitPropSliceExp(c *PropSliceExpContext)

	// ExitLlamarFuncionExp is called when exiting the LlamarFuncionExp production.
	ExitLlamarFuncionExp(c *LlamarFuncionExpContext)

	// ExitBinarioExp is called when exiting the BinarioExp production.
	ExitBinarioExp(c *BinarioExpContext)

	// ExitParentecisExp is called when exiting the ParentecisExp production.
	ExitParentecisExp(c *ParentecisExpContext)

	// ExitIntLiteral is called when exiting the IntLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitFloatLiteral is called when exiting the FloatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBoolLiteral is called when exiting the BoolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitNilLiteral is called when exiting the NilLiteral production.
	ExitNilLiteral(c *NilLiteralContext)

	// ExitID_Patron is called when exiting the ID_Patron production.
	ExitID_Patron(c *ID_PatronContext)
}
