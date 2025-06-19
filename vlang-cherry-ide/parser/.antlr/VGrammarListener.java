// Generated from /home/daaniieel/Escritorio/PROYECTO1/vlang-cherry-ide/parser/VGrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link VGrammar}.
 */
public interface VGrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link VGrammar#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(VGrammar.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(VGrammar.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncionMain}
	 * labeled alternative in {@link VGrammar#main_func}.
	 * @param ctx the parse tree
	 */
	void enterFuncionMain(VGrammar.FuncionMainContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncionMain}
	 * labeled alternative in {@link VGrammar#main_func}.
	 * @param ctx the parse tree
	 */
	void exitFuncionMain(VGrammar.FuncionMainContext ctx);
	/**
	 * Enter a parse tree produced by {@link VGrammar#stmt}.
	 * @param ctx the parse tree
	 */
	void enterStmt(VGrammar.StmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#stmt}.
	 * @param ctx the parse tree
	 */
	void exitStmt(VGrammar.StmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararInferenciaMut}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDeclararInferenciaMut(VGrammar.DeclararInferenciaMutContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararInferenciaMut}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDeclararInferenciaMut(VGrammar.DeclararInferenciaMutContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararInferencia}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDeclararInferencia(VGrammar.DeclararInferenciaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararInferencia}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDeclararInferencia(VGrammar.DeclararInferenciaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclaraTipoValor}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDeclaraTipoValor(VGrammar.DeclaraTipoValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclaraTipoValor}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDeclaraTipoValor(VGrammar.DeclaraTipoValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararTipo}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDeclararTipo(VGrammar.DeclararTipoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararTipo}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDeclararTipo(VGrammar.DeclararTipoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararSinMutValor}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDeclararSinMutValor(VGrammar.DeclararSinMutValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararSinMutValor}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDeclararSinMutValor(VGrammar.DeclararSinMutValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararSlice}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void enterDeclararSlice(VGrammar.DeclararSliceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararSlice}
	 * labeled alternative in {@link VGrammar#stmt_declaracion}.
	 * @param ctx the parse tree
	 */
	void exitDeclararSlice(VGrammar.DeclararSliceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignacionSliceItem}
	 * labeled alternative in {@link VGrammar#stmt_asignar}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionSliceItem(VGrammar.AsignacionSliceItemContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignacionSliceItem}
	 * labeled alternative in {@link VGrammar#stmt_asignar}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionSliceItem(VGrammar.AsignacionSliceItemContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignacionDirecta}
	 * labeled alternative in {@link VGrammar#stmt_asignar}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionDirecta(VGrammar.AsignacionDirectaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignacionDirecta}
	 * labeled alternative in {@link VGrammar#stmt_asignar}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionDirecta(VGrammar.AsignacionDirectaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignacionAritmetica}
	 * labeled alternative in {@link VGrammar#stmt_asignar}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionAritmetica(VGrammar.AsignacionAritmeticaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignacionAritmetica}
	 * labeled alternative in {@link VGrammar#stmt_asignar}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionAritmetica(VGrammar.AsignacionAritmeticaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignacionSlice}
	 * labeled alternative in {@link VGrammar#stmt_asignar}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionSlice(VGrammar.AsignacionSliceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignacionSlice}
	 * labeled alternative in {@link VGrammar#stmt_asignar}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionSlice(VGrammar.AsignacionSliceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ReturnStmt}
	 * labeled alternative in {@link VGrammar#stmt_transferencia}.
	 * @param ctx the parse tree
	 */
	void enterReturnStmt(VGrammar.ReturnStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ReturnStmt}
	 * labeled alternative in {@link VGrammar#stmt_transferencia}.
	 * @param ctx the parse tree
	 */
	void exitReturnStmt(VGrammar.ReturnStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BreakStmt}
	 * labeled alternative in {@link VGrammar#stmt_transferencia}.
	 * @param ctx the parse tree
	 */
	void enterBreakStmt(VGrammar.BreakStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BreakStmt}
	 * labeled alternative in {@link VGrammar#stmt_transferencia}.
	 * @param ctx the parse tree
	 */
	void exitBreakStmt(VGrammar.BreakStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ContinueStmt}
	 * labeled alternative in {@link VGrammar#stmt_transferencia}.
	 * @param ctx the parse tree
	 */
	void enterContinueStmt(VGrammar.ContinueStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ContinueStmt}
	 * labeled alternative in {@link VGrammar#stmt_transferencia}.
	 * @param ctx the parse tree
	 */
	void exitContinueStmt(VGrammar.ContinueStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IFstmt}
	 * labeled alternative in {@link VGrammar#if_stmt}.
	 * @param ctx the parse tree
	 */
	void enterIFstmt(VGrammar.IFstmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IFstmt}
	 * labeled alternative in {@link VGrammar#if_stmt}.
	 * @param ctx the parse tree
	 */
	void exitIFstmt(VGrammar.IFstmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IFcadena}
	 * labeled alternative in {@link VGrammar#if_chain}.
	 * @param ctx the parse tree
	 */
	void enterIFcadena(VGrammar.IFcadenaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IFcadena}
	 * labeled alternative in {@link VGrammar#if_chain}.
	 * @param ctx the parse tree
	 */
	void exitIFcadena(VGrammar.IFcadenaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ElseStmt}
	 * labeled alternative in {@link VGrammar#else_stmt}.
	 * @param ctx the parse tree
	 */
	void enterElseStmt(VGrammar.ElseStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ElseStmt}
	 * labeled alternative in {@link VGrammar#else_stmt}.
	 * @param ctx the parse tree
	 */
	void exitElseStmt(VGrammar.ElseStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SwitchStmt}
	 * labeled alternative in {@link VGrammar#switch_stmt}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStmt(VGrammar.SwitchStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SwitchStmt}
	 * labeled alternative in {@link VGrammar#switch_stmt}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStmt(VGrammar.SwitchStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SwitchCase}
	 * labeled alternative in {@link VGrammar#switch_case}.
	 * @param ctx the parse tree
	 */
	void enterSwitchCase(VGrammar.SwitchCaseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SwitchCase}
	 * labeled alternative in {@link VGrammar#switch_case}.
	 * @param ctx the parse tree
	 */
	void exitSwitchCase(VGrammar.SwitchCaseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DefaultCase}
	 * labeled alternative in {@link VGrammar#default_case}.
	 * @param ctx the parse tree
	 */
	void enterDefaultCase(VGrammar.DefaultCaseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DefaultCase}
	 * labeled alternative in {@link VGrammar#default_case}.
	 * @param ctx the parse tree
	 */
	void exitDefaultCase(VGrammar.DefaultCaseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code WhileStmt}
	 * labeled alternative in {@link VGrammar#while_stmt}.
	 * @param ctx the parse tree
	 */
	void enterWhileStmt(VGrammar.WhileStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code WhileStmt}
	 * labeled alternative in {@link VGrammar#while_stmt}.
	 * @param ctx the parse tree
	 */
	void exitWhileStmt(VGrammar.WhileStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForStmt}
	 * labeled alternative in {@link VGrammar#for_stmt}.
	 * @param ctx the parse tree
	 */
	void enterForStmt(VGrammar.ForStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForStmt}
	 * labeled alternative in {@link VGrammar#for_stmt}.
	 * @param ctx the parse tree
	 */
	void exitForStmt(VGrammar.ForStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code RangoNum}
	 * labeled alternative in {@link VGrammar#range}.
	 * @param ctx the parse tree
	 */
	void enterRangoNum(VGrammar.RangoNumContext ctx);
	/**
	 * Exit a parse tree produced by the {@code RangoNum}
	 * labeled alternative in {@link VGrammar#range}.
	 * @param ctx the parse tree
	 */
	void exitRangoNum(VGrammar.RangoNumContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncionDeclerada}
	 * labeled alternative in {@link VGrammar#declarar_funcion}.
	 * @param ctx the parse tree
	 */
	void enterFuncionDeclerada(VGrammar.FuncionDecleradaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncionDeclerada}
	 * labeled alternative in {@link VGrammar#declarar_funcion}.
	 * @param ctx the parse tree
	 */
	void exitFuncionDeclerada(VGrammar.FuncionDecleradaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code LlamarFuncion}
	 * labeled alternative in {@link VGrammar#llamar_funcion}.
	 * @param ctx the parse tree
	 */
	void enterLlamarFuncion(VGrammar.LlamarFuncionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code LlamarFuncion}
	 * labeled alternative in {@link VGrammar#llamar_funcion}.
	 * @param ctx the parse tree
	 */
	void exitLlamarFuncion(VGrammar.LlamarFuncionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ListaParametros}
	 * labeled alternative in {@link VGrammar#lista_parametros}.
	 * @param ctx the parse tree
	 */
	void enterListaParametros(VGrammar.ListaParametrosContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ListaParametros}
	 * labeled alternative in {@link VGrammar#lista_parametros}.
	 * @param ctx the parse tree
	 */
	void exitListaParametros(VGrammar.ListaParametrosContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ParametroFun}
	 * labeled alternative in {@link VGrammar#parametro_fun}.
	 * @param ctx the parse tree
	 */
	void enterParametroFun(VGrammar.ParametroFunContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ParametroFun}
	 * labeled alternative in {@link VGrammar#parametro_fun}.
	 * @param ctx the parse tree
	 */
	void exitParametroFun(VGrammar.ParametroFunContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ListaArgumentos}
	 * labeled alternative in {@link VGrammar#lista_argumentos}.
	 * @param ctx the parse tree
	 */
	void enterListaArgumentos(VGrammar.ListaArgumentosContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ListaArgumentos}
	 * labeled alternative in {@link VGrammar#lista_argumentos}.
	 * @param ctx the parse tree
	 */
	void exitListaArgumentos(VGrammar.ListaArgumentosContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncionArg}
	 * labeled alternative in {@link VGrammar#argumento_fun}.
	 * @param ctx the parse tree
	 */
	void enterFuncionArg(VGrammar.FuncionArgContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncionArg}
	 * labeled alternative in {@link VGrammar#argumento_fun}.
	 * @param ctx the parse tree
	 */
	void exitFuncionArg(VGrammar.FuncionArgContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararStruct}
	 * labeled alternative in {@link VGrammar#declarar_struct}.
	 * @param ctx the parse tree
	 */
	void enterDeclararStruct(VGrammar.DeclararStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararStruct}
	 * labeled alternative in {@link VGrammar#declarar_struct}.
	 * @param ctx the parse tree
	 */
	void exitDeclararStruct(VGrammar.DeclararStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PropiedadStruct}
	 * labeled alternative in {@link VGrammar#propiedad_struct}.
	 * @param ctx the parse tree
	 */
	void enterPropiedadStruct(VGrammar.PropiedadStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PropiedadStruct}
	 * labeled alternative in {@link VGrammar#propiedad_struct}.
	 * @param ctx the parse tree
	 */
	void exitPropiedadStruct(VGrammar.PropiedadStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CrearStruct}
	 * labeled alternative in {@link VGrammar#crear_struct}.
	 * @param ctx the parse tree
	 */
	void enterCrearStruct(VGrammar.CrearStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CrearStruct}
	 * labeled alternative in {@link VGrammar#crear_struct}.
	 * @param ctx the parse tree
	 */
	void exitCrearStruct(VGrammar.CrearStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ListaParametrosInit}
	 * labeled alternative in {@link VGrammar#lista_parametros_init}.
	 * @param ctx the parse tree
	 */
	void enterListaParametrosInit(VGrammar.ListaParametrosInitContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ListaParametrosInit}
	 * labeled alternative in {@link VGrammar#lista_parametros_init}.
	 * @param ctx the parse tree
	 */
	void exitListaParametrosInit(VGrammar.ListaParametrosInitContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ParametrosInitStruct}
	 * labeled alternative in {@link VGrammar#parametros_init_struct}.
	 * @param ctx the parse tree
	 */
	void enterParametrosInitStruct(VGrammar.ParametrosInitStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ParametrosInitStruct}
	 * labeled alternative in {@link VGrammar#parametros_init_struct}.
	 * @param ctx the parse tree
	 */
	void exitParametrosInitStruct(VGrammar.ParametrosInitStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ListaSlice}
	 * labeled alternative in {@link VGrammar#lista_slice}.
	 * @param ctx the parse tree
	 */
	void enterListaSlice(VGrammar.ListaSliceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ListaSlice}
	 * labeled alternative in {@link VGrammar#lista_slice}.
	 * @param ctx the parse tree
	 */
	void exitListaSlice(VGrammar.ListaSliceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ItemSlice}
	 * labeled alternative in {@link VGrammar#item_slice}.
	 * @param ctx the parse tree
	 */
	void enterItemSlice(VGrammar.ItemSliceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ItemSlice}
	 * labeled alternative in {@link VGrammar#item_slice}.
	 * @param ctx the parse tree
	 */
	void exitItemSlice(VGrammar.ItemSliceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PropSlice}
	 * labeled alternative in {@link VGrammar#prop_slice}.
	 * @param ctx the parse tree
	 */
	void enterPropSlice(VGrammar.PropSliceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PropSlice}
	 * labeled alternative in {@link VGrammar#prop_slice}.
	 * @param ctx the parse tree
	 */
	void exitPropSlice(VGrammar.PropSliceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncionSlice}
	 * labeled alternative in {@link VGrammar#fun_slice}.
	 * @param ctx the parse tree
	 */
	void enterFuncionSlice(VGrammar.FuncionSliceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncionSlice}
	 * labeled alternative in {@link VGrammar#fun_slice}.
	 * @param ctx the parse tree
	 */
	void exitFuncionSlice(VGrammar.FuncionSliceContext ctx);
	/**
	 * Enter a parse tree produced by {@link VGrammar#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo(VGrammar.TipoContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo(VGrammar.TipoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorSimple}
	 * labeled alternative in {@link VGrammar#tipos_slices}.
	 * @param ctx the parse tree
	 */
	void enterVectorSimple(VGrammar.VectorSimpleContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorSimple}
	 * labeled alternative in {@link VGrammar#tipos_slices}.
	 * @param ctx the parse tree
	 */
	void exitVectorSimple(VGrammar.VectorSimpleContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MatrizDoble}
	 * labeled alternative in {@link VGrammar#tipos_slices}.
	 * @param ctx the parse tree
	 */
	void enterMatrizDoble(VGrammar.MatrizDobleContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MatrizDoble}
	 * labeled alternative in {@link VGrammar#tipos_slices}.
	 * @param ctx the parse tree
	 */
	void exitMatrizDoble(VGrammar.MatrizDobleContext ctx);
	/**
	 * Enter a parse tree produced by the {@code LiteralExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterLiteralExp(VGrammar.LiteralExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code LiteralExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitLiteralExp(VGrammar.LiteralExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IdExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterIdExp(VGrammar.IdExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IdExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitIdExp(VGrammar.IdExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CrearStructExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterCrearStructExp(VGrammar.CrearStructExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CrearStructExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitCrearStructExp(VGrammar.CrearStructExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code UnarioExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterUnarioExp(VGrammar.UnarioExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code UnarioExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitUnarioExp(VGrammar.UnarioExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ItemSliceExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterItemSliceExp(VGrammar.ItemSliceExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ItemSliceExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitItemSliceExp(VGrammar.ItemSliceExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorFuncExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterVectorFuncExp(VGrammar.VectorFuncExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorFuncExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitVectorFuncExp(VGrammar.VectorFuncExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SliceExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterSliceExp(VGrammar.SliceExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SliceExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitSliceExp(VGrammar.SliceExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PropSliceExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterPropSliceExp(VGrammar.PropSliceExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PropSliceExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitPropSliceExp(VGrammar.PropSliceExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code LlamarFuncionExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterLlamarFuncionExp(VGrammar.LlamarFuncionExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code LlamarFuncionExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitLlamarFuncionExp(VGrammar.LlamarFuncionExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BinarioExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterBinarioExp(VGrammar.BinarioExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BinarioExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitBinarioExp(VGrammar.BinarioExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ParentecisExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterParentecisExp(VGrammar.ParentecisExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ParentecisExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitParentecisExp(VGrammar.ParentecisExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IntLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void enterIntLiteral(VGrammar.IntLiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IntLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void exitIntLiteral(VGrammar.IntLiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FloatLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void enterFloatLiteral(VGrammar.FloatLiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FloatLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void exitFloatLiteral(VGrammar.FloatLiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StringLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void enterStringLiteral(VGrammar.StringLiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StringLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void exitStringLiteral(VGrammar.StringLiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BoolLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void enterBoolLiteral(VGrammar.BoolLiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BoolLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void exitBoolLiteral(VGrammar.BoolLiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code NilLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void enterNilLiteral(VGrammar.NilLiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code NilLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void exitNilLiteral(VGrammar.NilLiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ID_Patron}
	 * labeled alternative in {@link VGrammar#patronId}.
	 * @param ctx the parse tree
	 */
	void enterID_Patron(VGrammar.ID_PatronContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ID_Patron}
	 * labeled alternative in {@link VGrammar#patronId}.
	 * @param ctx the parse tree
	 */
	void exitID_Patron(VGrammar.ID_PatronContext ctx);
}