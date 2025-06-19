// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VGrammar.
type VGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VGrammar#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by VGrammar#FuncionMain.
	VisitFuncionMain(ctx *FuncionMainContext) interface{}

	// Visit a parse tree produced by VGrammar#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararInferenciaMut.
	VisitDeclararInferenciaMut(ctx *DeclararInferenciaMutContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararInferencia.
	VisitDeclararInferencia(ctx *DeclararInferenciaContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclaraTipoValor.
	VisitDeclaraTipoValor(ctx *DeclaraTipoValorContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararTipo.
	VisitDeclararTipo(ctx *DeclararTipoContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararSinMutValor.
	VisitDeclararSinMutValor(ctx *DeclararSinMutValorContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararSlice.
	VisitDeclararSlice(ctx *DeclararSliceContext) interface{}

	// Visit a parse tree produced by VGrammar#AsignacionSliceItem.
	VisitAsignacionSliceItem(ctx *AsignacionSliceItemContext) interface{}

	// Visit a parse tree produced by VGrammar#AsignacionDirecta.
	VisitAsignacionDirecta(ctx *AsignacionDirectaContext) interface{}

	// Visit a parse tree produced by VGrammar#AsignacionAritmetica.
	VisitAsignacionAritmetica(ctx *AsignacionAritmeticaContext) interface{}

	// Visit a parse tree produced by VGrammar#AsignacionSlice.
	VisitAsignacionSlice(ctx *AsignacionSliceContext) interface{}

	// Visit a parse tree produced by VGrammar#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#BreakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#ContinueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#IFstmt.
	VisitIFstmt(ctx *IFstmtContext) interface{}

	// Visit a parse tree produced by VGrammar#IFcadena.
	VisitIFcadena(ctx *IFcadenaContext) interface{}

	// Visit a parse tree produced by VGrammar#ElseStmt.
	VisitElseStmt(ctx *ElseStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#SwitchStmt.
	VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#SwitchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by VGrammar#DefaultCase.
	VisitDefaultCase(ctx *DefaultCaseContext) interface{}

	// Visit a parse tree produced by VGrammar#WhileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#ForClasicoStmt.
	VisitForClasicoStmt(ctx *ForClasicoStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#ForStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#RangoNum.
	VisitRangoNum(ctx *RangoNumContext) interface{}

	// Visit a parse tree produced by VGrammar#FuncionDeclerada.
	VisitFuncionDeclerada(ctx *FuncionDecleradaContext) interface{}

	// Visit a parse tree produced by VGrammar#LlamarFuncion.
	VisitLlamarFuncion(ctx *LlamarFuncionContext) interface{}

	// Visit a parse tree produced by VGrammar#ListaParametros.
	VisitListaParametros(ctx *ListaParametrosContext) interface{}

	// Visit a parse tree produced by VGrammar#ParametroFun.
	VisitParametroFun(ctx *ParametroFunContext) interface{}

	// Visit a parse tree produced by VGrammar#ListaArgumentos.
	VisitListaArgumentos(ctx *ListaArgumentosContext) interface{}

	// Visit a parse tree produced by VGrammar#FuncionArg.
	VisitFuncionArg(ctx *FuncionArgContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararStruct.
	VisitDeclararStruct(ctx *DeclararStructContext) interface{}

	// Visit a parse tree produced by VGrammar#PropiedadStruct.
	VisitPropiedadStruct(ctx *PropiedadStructContext) interface{}

	// Visit a parse tree produced by VGrammar#CrearStruct.
	VisitCrearStruct(ctx *CrearStructContext) interface{}

	// Visit a parse tree produced by VGrammar#ListaParametrosInit.
	VisitListaParametrosInit(ctx *ListaParametrosInitContext) interface{}

	// Visit a parse tree produced by VGrammar#ParametrosInitStruct.
	VisitParametrosInitStruct(ctx *ParametrosInitStructContext) interface{}

	// Visit a parse tree produced by VGrammar#ListaSlice.
	VisitListaSlice(ctx *ListaSliceContext) interface{}

	// Visit a parse tree produced by VGrammar#ItemSlice.
	VisitItemSlice(ctx *ItemSliceContext) interface{}

	// Visit a parse tree produced by VGrammar#PropSlice.
	VisitPropSlice(ctx *PropSliceContext) interface{}

	// Visit a parse tree produced by VGrammar#FuncionSlice.
	VisitFuncionSlice(ctx *FuncionSliceContext) interface{}

	// Visit a parse tree produced by VGrammar#tipo.
	VisitTipo(ctx *TipoContext) interface{}

	// Visit a parse tree produced by VGrammar#VectorSimple.
	VisitVectorSimple(ctx *VectorSimpleContext) interface{}

	// Visit a parse tree produced by VGrammar#MatrizDoble.
	VisitMatrizDoble(ctx *MatrizDobleContext) interface{}

	// Visit a parse tree produced by VGrammar#LiteralExp.
	VisitLiteralExp(ctx *LiteralExpContext) interface{}

	// Visit a parse tree produced by VGrammar#IdExp.
	VisitIdExp(ctx *IdExpContext) interface{}

	// Visit a parse tree produced by VGrammar#CrearStructExp.
	VisitCrearStructExp(ctx *CrearStructExpContext) interface{}

	// Visit a parse tree produced by VGrammar#UnarioExp.
	VisitUnarioExp(ctx *UnarioExpContext) interface{}

	// Visit a parse tree produced by VGrammar#ItemSliceExp.
	VisitItemSliceExp(ctx *ItemSliceExpContext) interface{}

	// Visit a parse tree produced by VGrammar#VectorFuncExp.
	VisitVectorFuncExp(ctx *VectorFuncExpContext) interface{}

	// Visit a parse tree produced by VGrammar#SliceExp.
	VisitSliceExp(ctx *SliceExpContext) interface{}

	// Visit a parse tree produced by VGrammar#PropSliceExp.
	VisitPropSliceExp(ctx *PropSliceExpContext) interface{}

	// Visit a parse tree produced by VGrammar#LlamarFuncionExp.
	VisitLlamarFuncionExp(ctx *LlamarFuncionExpContext) interface{}

	// Visit a parse tree produced by VGrammar#BinarioExp.
	VisitBinarioExp(ctx *BinarioExpContext) interface{}

	// Visit a parse tree produced by VGrammar#ParentecisExp.
	VisitParentecisExp(ctx *ParentecisExpContext) interface{}

	// Visit a parse tree produced by VGrammar#IntLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by VGrammar#FloatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by VGrammar#StringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by VGrammar#BoolLiteral.
	VisitBoolLiteral(ctx *BoolLiteralContext) interface{}

	// Visit a parse tree produced by VGrammar#NilLiteral.
	VisitNilLiteral(ctx *NilLiteralContext) interface{}

	// Visit a parse tree produced by VGrammar#ID_Patron.
	VisitID_Patron(ctx *ID_PatronContext) interface{}
}
