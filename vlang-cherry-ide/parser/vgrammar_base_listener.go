// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseVGrammarListener is a complete listener for a parse tree produced by VGrammar.
type BaseVGrammarListener struct{}

var _ VGrammarListener = &BaseVGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseVGrammarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseVGrammarListener) ExitProgram(ctx *ProgramContext) {}

// EnterFuncionMain is called when production FuncionMain is entered.
func (s *BaseVGrammarListener) EnterFuncionMain(ctx *FuncionMainContext) {}

// ExitFuncionMain is called when production FuncionMain is exited.
func (s *BaseVGrammarListener) ExitFuncionMain(ctx *FuncionMainContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseVGrammarListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseVGrammarListener) ExitStmt(ctx *StmtContext) {}

// EnterDeclararInferenciaMut is called when production DeclararInferenciaMut is entered.
func (s *BaseVGrammarListener) EnterDeclararInferenciaMut(ctx *DeclararInferenciaMutContext) {}

// ExitDeclararInferenciaMut is called when production DeclararInferenciaMut is exited.
func (s *BaseVGrammarListener) ExitDeclararInferenciaMut(ctx *DeclararInferenciaMutContext) {}

// EnterDeclararInferencia is called when production DeclararInferencia is entered.
func (s *BaseVGrammarListener) EnterDeclararInferencia(ctx *DeclararInferenciaContext) {}

// ExitDeclararInferencia is called when production DeclararInferencia is exited.
func (s *BaseVGrammarListener) ExitDeclararInferencia(ctx *DeclararInferenciaContext) {}

// EnterDeclaraTipoValor is called when production DeclaraTipoValor is entered.
func (s *BaseVGrammarListener) EnterDeclaraTipoValor(ctx *DeclaraTipoValorContext) {}

// ExitDeclaraTipoValor is called when production DeclaraTipoValor is exited.
func (s *BaseVGrammarListener) ExitDeclaraTipoValor(ctx *DeclaraTipoValorContext) {}

// EnterDeclararTipo is called when production DeclararTipo is entered.
func (s *BaseVGrammarListener) EnterDeclararTipo(ctx *DeclararTipoContext) {}

// ExitDeclararTipo is called when production DeclararTipo is exited.
func (s *BaseVGrammarListener) ExitDeclararTipo(ctx *DeclararTipoContext) {}

// EnterDeclararSinMutValor is called when production DeclararSinMutValor is entered.
func (s *BaseVGrammarListener) EnterDeclararSinMutValor(ctx *DeclararSinMutValorContext) {}

// ExitDeclararSinMutValor is called when production DeclararSinMutValor is exited.
func (s *BaseVGrammarListener) ExitDeclararSinMutValor(ctx *DeclararSinMutValorContext) {}

// EnterDeclararSlice is called when production DeclararSlice is entered.
func (s *BaseVGrammarListener) EnterDeclararSlice(ctx *DeclararSliceContext) {}

// ExitDeclararSlice is called when production DeclararSlice is exited.
func (s *BaseVGrammarListener) ExitDeclararSlice(ctx *DeclararSliceContext) {}

// EnterAsignacionSliceItem is called when production AsignacionSliceItem is entered.
func (s *BaseVGrammarListener) EnterAsignacionSliceItem(ctx *AsignacionSliceItemContext) {}

// ExitAsignacionSliceItem is called when production AsignacionSliceItem is exited.
func (s *BaseVGrammarListener) ExitAsignacionSliceItem(ctx *AsignacionSliceItemContext) {}

// EnterAsignacionDirecta is called when production AsignacionDirecta is entered.
func (s *BaseVGrammarListener) EnterAsignacionDirecta(ctx *AsignacionDirectaContext) {}

// ExitAsignacionDirecta is called when production AsignacionDirecta is exited.
func (s *BaseVGrammarListener) ExitAsignacionDirecta(ctx *AsignacionDirectaContext) {}

// EnterAsignacionAritmetica is called when production AsignacionAritmetica is entered.
func (s *BaseVGrammarListener) EnterAsignacionAritmetica(ctx *AsignacionAritmeticaContext) {}

// ExitAsignacionAritmetica is called when production AsignacionAritmetica is exited.
func (s *BaseVGrammarListener) ExitAsignacionAritmetica(ctx *AsignacionAritmeticaContext) {}

// EnterAsignacionSlice is called when production AsignacionSlice is entered.
func (s *BaseVGrammarListener) EnterAsignacionSlice(ctx *AsignacionSliceContext) {}

// ExitAsignacionSlice is called when production AsignacionSlice is exited.
func (s *BaseVGrammarListener) ExitAsignacionSlice(ctx *AsignacionSliceContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BaseVGrammarListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BaseVGrammarListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterBreakStmt is called when production BreakStmt is entered.
func (s *BaseVGrammarListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production BreakStmt is exited.
func (s *BaseVGrammarListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production ContinueStmt is entered.
func (s *BaseVGrammarListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production ContinueStmt is exited.
func (s *BaseVGrammarListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterIFstmt is called when production IFstmt is entered.
func (s *BaseVGrammarListener) EnterIFstmt(ctx *IFstmtContext) {}

// ExitIFstmt is called when production IFstmt is exited.
func (s *BaseVGrammarListener) ExitIFstmt(ctx *IFstmtContext) {}

// EnterIFcadena is called when production IFcadena is entered.
func (s *BaseVGrammarListener) EnterIFcadena(ctx *IFcadenaContext) {}

// ExitIFcadena is called when production IFcadena is exited.
func (s *BaseVGrammarListener) ExitIFcadena(ctx *IFcadenaContext) {}

// EnterElseStmt is called when production ElseStmt is entered.
func (s *BaseVGrammarListener) EnterElseStmt(ctx *ElseStmtContext) {}

// ExitElseStmt is called when production ElseStmt is exited.
func (s *BaseVGrammarListener) ExitElseStmt(ctx *ElseStmtContext) {}

// EnterSwitchStmt is called when production SwitchStmt is entered.
func (s *BaseVGrammarListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production SwitchStmt is exited.
func (s *BaseVGrammarListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterSwitchCase is called when production SwitchCase is entered.
func (s *BaseVGrammarListener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production SwitchCase is exited.
func (s *BaseVGrammarListener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterDefaultCase is called when production DefaultCase is entered.
func (s *BaseVGrammarListener) EnterDefaultCase(ctx *DefaultCaseContext) {}

// ExitDefaultCase is called when production DefaultCase is exited.
func (s *BaseVGrammarListener) ExitDefaultCase(ctx *DefaultCaseContext) {}

// EnterWhileStmt is called when production WhileStmt is entered.
func (s *BaseVGrammarListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production WhileStmt is exited.
func (s *BaseVGrammarListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterForClasicoStmt is called when production ForClasicoStmt is entered.
func (s *BaseVGrammarListener) EnterForClasicoStmt(ctx *ForClasicoStmtContext) {}

// ExitForClasicoStmt is called when production ForClasicoStmt is exited.
func (s *BaseVGrammarListener) ExitForClasicoStmt(ctx *ForClasicoStmtContext) {}

// EnterForStmt is called when production ForStmt is entered.
func (s *BaseVGrammarListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production ForStmt is exited.
func (s *BaseVGrammarListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterRangoNum is called when production RangoNum is entered.
func (s *BaseVGrammarListener) EnterRangoNum(ctx *RangoNumContext) {}

// ExitRangoNum is called when production RangoNum is exited.
func (s *BaseVGrammarListener) ExitRangoNum(ctx *RangoNumContext) {}

// EnterFuncionDeclerada is called when production FuncionDeclerada is entered.
func (s *BaseVGrammarListener) EnterFuncionDeclerada(ctx *FuncionDecleradaContext) {}

// ExitFuncionDeclerada is called when production FuncionDeclerada is exited.
func (s *BaseVGrammarListener) ExitFuncionDeclerada(ctx *FuncionDecleradaContext) {}

// EnterLlamarFuncion is called when production LlamarFuncion is entered.
func (s *BaseVGrammarListener) EnterLlamarFuncion(ctx *LlamarFuncionContext) {}

// ExitLlamarFuncion is called when production LlamarFuncion is exited.
func (s *BaseVGrammarListener) ExitLlamarFuncion(ctx *LlamarFuncionContext) {}

// EnterListaParametros is called when production ListaParametros is entered.
func (s *BaseVGrammarListener) EnterListaParametros(ctx *ListaParametrosContext) {}

// ExitListaParametros is called when production ListaParametros is exited.
func (s *BaseVGrammarListener) ExitListaParametros(ctx *ListaParametrosContext) {}

// EnterParametroFun is called when production ParametroFun is entered.
func (s *BaseVGrammarListener) EnterParametroFun(ctx *ParametroFunContext) {}

// ExitParametroFun is called when production ParametroFun is exited.
func (s *BaseVGrammarListener) ExitParametroFun(ctx *ParametroFunContext) {}

// EnterListaArgumentos is called when production ListaArgumentos is entered.
func (s *BaseVGrammarListener) EnterListaArgumentos(ctx *ListaArgumentosContext) {}

// ExitListaArgumentos is called when production ListaArgumentos is exited.
func (s *BaseVGrammarListener) ExitListaArgumentos(ctx *ListaArgumentosContext) {}

// EnterFuncionArg is called when production FuncionArg is entered.
func (s *BaseVGrammarListener) EnterFuncionArg(ctx *FuncionArgContext) {}

// ExitFuncionArg is called when production FuncionArg is exited.
func (s *BaseVGrammarListener) ExitFuncionArg(ctx *FuncionArgContext) {}

// EnterDeclararStruct is called when production DeclararStruct is entered.
func (s *BaseVGrammarListener) EnterDeclararStruct(ctx *DeclararStructContext) {}

// ExitDeclararStruct is called when production DeclararStruct is exited.
func (s *BaseVGrammarListener) ExitDeclararStruct(ctx *DeclararStructContext) {}

// EnterPropiedadStruct is called when production PropiedadStruct is entered.
func (s *BaseVGrammarListener) EnterPropiedadStruct(ctx *PropiedadStructContext) {}

// ExitPropiedadStruct is called when production PropiedadStruct is exited.
func (s *BaseVGrammarListener) ExitPropiedadStruct(ctx *PropiedadStructContext) {}

// EnterCrearStruct is called when production CrearStruct is entered.
func (s *BaseVGrammarListener) EnterCrearStruct(ctx *CrearStructContext) {}

// ExitCrearStruct is called when production CrearStruct is exited.
func (s *BaseVGrammarListener) ExitCrearStruct(ctx *CrearStructContext) {}

// EnterListaParametrosInit is called when production ListaParametrosInit is entered.
func (s *BaseVGrammarListener) EnterListaParametrosInit(ctx *ListaParametrosInitContext) {}

// ExitListaParametrosInit is called when production ListaParametrosInit is exited.
func (s *BaseVGrammarListener) ExitListaParametrosInit(ctx *ListaParametrosInitContext) {}

// EnterParametrosInitStruct is called when production ParametrosInitStruct is entered.
func (s *BaseVGrammarListener) EnterParametrosInitStruct(ctx *ParametrosInitStructContext) {}

// ExitParametrosInitStruct is called when production ParametrosInitStruct is exited.
func (s *BaseVGrammarListener) ExitParametrosInitStruct(ctx *ParametrosInitStructContext) {}

// EnterListaSlice is called when production ListaSlice is entered.
func (s *BaseVGrammarListener) EnterListaSlice(ctx *ListaSliceContext) {}

// ExitListaSlice is called when production ListaSlice is exited.
func (s *BaseVGrammarListener) ExitListaSlice(ctx *ListaSliceContext) {}

// EnterItemSlice is called when production ItemSlice is entered.
func (s *BaseVGrammarListener) EnterItemSlice(ctx *ItemSliceContext) {}

// ExitItemSlice is called when production ItemSlice is exited.
func (s *BaseVGrammarListener) ExitItemSlice(ctx *ItemSliceContext) {}

// EnterPropSlice is called when production PropSlice is entered.
func (s *BaseVGrammarListener) EnterPropSlice(ctx *PropSliceContext) {}

// ExitPropSlice is called when production PropSlice is exited.
func (s *BaseVGrammarListener) ExitPropSlice(ctx *PropSliceContext) {}

// EnterFuncionSlice is called when production FuncionSlice is entered.
func (s *BaseVGrammarListener) EnterFuncionSlice(ctx *FuncionSliceContext) {}

// ExitFuncionSlice is called when production FuncionSlice is exited.
func (s *BaseVGrammarListener) ExitFuncionSlice(ctx *FuncionSliceContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseVGrammarListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseVGrammarListener) ExitTipo(ctx *TipoContext) {}

// EnterVectorSimple is called when production VectorSimple is entered.
func (s *BaseVGrammarListener) EnterVectorSimple(ctx *VectorSimpleContext) {}

// ExitVectorSimple is called when production VectorSimple is exited.
func (s *BaseVGrammarListener) ExitVectorSimple(ctx *VectorSimpleContext) {}

// EnterMatrizDoble is called when production MatrizDoble is entered.
func (s *BaseVGrammarListener) EnterMatrizDoble(ctx *MatrizDobleContext) {}

// ExitMatrizDoble is called when production MatrizDoble is exited.
func (s *BaseVGrammarListener) ExitMatrizDoble(ctx *MatrizDobleContext) {}

// EnterLiteralExp is called when production LiteralExp is entered.
func (s *BaseVGrammarListener) EnterLiteralExp(ctx *LiteralExpContext) {}

// ExitLiteralExp is called when production LiteralExp is exited.
func (s *BaseVGrammarListener) ExitLiteralExp(ctx *LiteralExpContext) {}

// EnterIdExp is called when production IdExp is entered.
func (s *BaseVGrammarListener) EnterIdExp(ctx *IdExpContext) {}

// ExitIdExp is called when production IdExp is exited.
func (s *BaseVGrammarListener) ExitIdExp(ctx *IdExpContext) {}

// EnterCrearStructExp is called when production CrearStructExp is entered.
func (s *BaseVGrammarListener) EnterCrearStructExp(ctx *CrearStructExpContext) {}

// ExitCrearStructExp is called when production CrearStructExp is exited.
func (s *BaseVGrammarListener) ExitCrearStructExp(ctx *CrearStructExpContext) {}

// EnterUnarioExp is called when production UnarioExp is entered.
func (s *BaseVGrammarListener) EnterUnarioExp(ctx *UnarioExpContext) {}

// ExitUnarioExp is called when production UnarioExp is exited.
func (s *BaseVGrammarListener) ExitUnarioExp(ctx *UnarioExpContext) {}

// EnterItemSliceExp is called when production ItemSliceExp is entered.
func (s *BaseVGrammarListener) EnterItemSliceExp(ctx *ItemSliceExpContext) {}

// ExitItemSliceExp is called when production ItemSliceExp is exited.
func (s *BaseVGrammarListener) ExitItemSliceExp(ctx *ItemSliceExpContext) {}

// EnterVectorFuncExp is called when production VectorFuncExp is entered.
func (s *BaseVGrammarListener) EnterVectorFuncExp(ctx *VectorFuncExpContext) {}

// ExitVectorFuncExp is called when production VectorFuncExp is exited.
func (s *BaseVGrammarListener) ExitVectorFuncExp(ctx *VectorFuncExpContext) {}

// EnterSliceExp is called when production SliceExp is entered.
func (s *BaseVGrammarListener) EnterSliceExp(ctx *SliceExpContext) {}

// ExitSliceExp is called when production SliceExp is exited.
func (s *BaseVGrammarListener) ExitSliceExp(ctx *SliceExpContext) {}

// EnterPropSliceExp is called when production PropSliceExp is entered.
func (s *BaseVGrammarListener) EnterPropSliceExp(ctx *PropSliceExpContext) {}

// ExitPropSliceExp is called when production PropSliceExp is exited.
func (s *BaseVGrammarListener) ExitPropSliceExp(ctx *PropSliceExpContext) {}

// EnterLlamarFuncionExp is called when production LlamarFuncionExp is entered.
func (s *BaseVGrammarListener) EnterLlamarFuncionExp(ctx *LlamarFuncionExpContext) {}

// ExitLlamarFuncionExp is called when production LlamarFuncionExp is exited.
func (s *BaseVGrammarListener) ExitLlamarFuncionExp(ctx *LlamarFuncionExpContext) {}

// EnterBinarioExp is called when production BinarioExp is entered.
func (s *BaseVGrammarListener) EnterBinarioExp(ctx *BinarioExpContext) {}

// ExitBinarioExp is called when production BinarioExp is exited.
func (s *BaseVGrammarListener) ExitBinarioExp(ctx *BinarioExpContext) {}

// EnterParentecisExp is called when production ParentecisExp is entered.
func (s *BaseVGrammarListener) EnterParentecisExp(ctx *ParentecisExpContext) {}

// ExitParentecisExp is called when production ParentecisExp is exited.
func (s *BaseVGrammarListener) ExitParentecisExp(ctx *ParentecisExpContext) {}

// EnterIntLiteral is called when production IntLiteral is entered.
func (s *BaseVGrammarListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production IntLiteral is exited.
func (s *BaseVGrammarListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFloatLiteral is called when production FloatLiteral is entered.
func (s *BaseVGrammarListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production FloatLiteral is exited.
func (s *BaseVGrammarListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseVGrammarListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseVGrammarListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBoolLiteral is called when production BoolLiteral is entered.
func (s *BaseVGrammarListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production BoolLiteral is exited.
func (s *BaseVGrammarListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterNilLiteral is called when production NilLiteral is entered.
func (s *BaseVGrammarListener) EnterNilLiteral(ctx *NilLiteralContext) {}

// ExitNilLiteral is called when production NilLiteral is exited.
func (s *BaseVGrammarListener) ExitNilLiteral(ctx *NilLiteralContext) {}

// EnterID_Patron is called when production ID_Patron is entered.
func (s *BaseVGrammarListener) EnterID_Patron(ctx *ID_PatronContext) {}

// ExitID_Patron is called when production ID_Patron is exited.
func (s *BaseVGrammarListener) ExitID_Patron(ctx *ID_PatronContext) {}
