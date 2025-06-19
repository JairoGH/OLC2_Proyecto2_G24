// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseVGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVGrammarVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFuncionMain(ctx *FuncionMainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararInferenciaMut(ctx *DeclararInferenciaMutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararInferencia(ctx *DeclararInferenciaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclaraTipoValor(ctx *DeclaraTipoValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararTipo(ctx *DeclararTipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararSinMutValor(ctx *DeclararSinMutValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararSlice(ctx *DeclararSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAsignacionSliceItem(ctx *AsignacionSliceItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAsignacionDirecta(ctx *AsignacionDirectaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAsignacionAritmetica(ctx *AsignacionAritmeticaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAsignacionSlice(ctx *AsignacionSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitIFstmt(ctx *IFstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitIFcadena(ctx *IFcadenaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitElseStmt(ctx *ElseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitSwitchCase(ctx *SwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDefaultCase(ctx *DefaultCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitForClasicoStmt(ctx *ForClasicoStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitRangoNum(ctx *RangoNumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFuncionDeclerada(ctx *FuncionDecleradaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLlamarFuncion(ctx *LlamarFuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitListaParametros(ctx *ListaParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitParametroFun(ctx *ParametroFunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitListaArgumentos(ctx *ListaArgumentosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFuncionArg(ctx *FuncionArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararStruct(ctx *DeclararStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitPropiedadStruct(ctx *PropiedadStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitCrearStruct(ctx *CrearStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitListaParametrosInit(ctx *ListaParametrosInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitParametrosInitStruct(ctx *ParametrosInitStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitListaSlice(ctx *ListaSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitItemSlice(ctx *ItemSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitPropSlice(ctx *PropSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFuncionSlice(ctx *FuncionSliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitTipo(ctx *TipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitVectorSimple(ctx *VectorSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitMatrizDoble(ctx *MatrizDobleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLiteralExp(ctx *LiteralExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitIdExp(ctx *IdExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitCrearStructExp(ctx *CrearStructExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitUnarioExp(ctx *UnarioExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitItemSliceExp(ctx *ItemSliceExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitVectorFuncExp(ctx *VectorFuncExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitSliceExp(ctx *SliceExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitPropSliceExp(ctx *PropSliceExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLlamarFuncionExp(ctx *LlamarFuncionExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitBinarioExp(ctx *BinarioExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitParentecisExp(ctx *ParentecisExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitNilLiteral(ctx *NilLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitID_Patron(ctx *ID_PatronContext) interface{} {
	return v.VisitChildren(ctx)
}
