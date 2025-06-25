// Generated from /home/daaniieel/Escritorio/PRUEBA/vlang-cherry-ide/parser/VGrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class VGrammar extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WS=1, COMENTARIO=2, COMENTARIOMULT=3, RW_MUT=4, OP_ASSIGN=5, OP_INCREMENTO=6, 
		OP_DECREMENTO=7, OP_DECLARACION=8, RW_INT=9, RW_FLOAT64=10, RW_STRING=11, 
		RW_BOOL=12, INT_LITERAL=13, FLOAT_LITERAL=14, STRING_LITERAL=15, BOOL_LITERAL=16, 
		NIL_LITERAL=17, OP_SUMA=18, OP_RESTA=19, OP_MULT=20, OP_DIV=21, OP_MOD=22, 
		OP_IGUAL=23, OP_DIFERENTE=24, OP_MENORQ=25, OP_MAYORQ=26, OP_MENORIGUAL=27, 
		OP_MAYORIGUAL=28, OP_AND=29, OP_OR=30, OP_NOT=31, RW_MAIN=32, RW_FN=33, 
		RW_STRUCT=34, RW_IF=35, RW_ELSE=36, RW_SWITCH=37, RW_CASE=38, RW_DEFAULT=39, 
		RW_FOR=40, RW_WHILE=41, RW_BREAK=42, RW_CONTINUE=43, RW_RETURN=44, RW_IN=45, 
		ID=46, PAR_IZQ=47, PAR_DER=48, LLAVE_IZQ=49, LLAVE_DER=50, CORCHETE_IZQ=51, 
		CORCHETE_DER=52, PUNTO=53, COMA=54, PUNTO_Y_COMA=55, DOS_PUNTOS=56;
	public static final int
		RULE_program = 0, RULE_main_func = 1, RULE_stmt = 2, RULE_stmt_declaracion = 3, 
		RULE_stmt_asignar = 4, RULE_stmt_transferencia = 5, RULE_if_stmt = 6, 
		RULE_if_chain = 7, RULE_else_stmt = 8, RULE_switch_stmt = 9, RULE_switch_case = 10, 
		RULE_default_case = 11, RULE_while_stmt = 12, RULE_for_clasico_stmt = 13, 
		RULE_for_stmt = 14, RULE_range = 15, RULE_declarar_funcion = 16, RULE_llamar_funcion = 17, 
		RULE_lista_parametros = 18, RULE_parametro_fun = 19, RULE_lista_argumentos = 20, 
		RULE_argumento_fun = 21, RULE_declarar_struct = 22, RULE_propiedad_struct = 23, 
		RULE_crear_struct = 24, RULE_lista_parametros_init = 25, RULE_parametros_init_struct = 26, 
		RULE_lista_slice = 27, RULE_item_slice = 28, RULE_prop_slice = 29, RULE_fun_slice = 30, 
		RULE_tipo = 31, RULE_tipos_slices = 32, RULE_expr = 33, RULE_literal = 34, 
		RULE_patronId = 35;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "main_func", "stmt", "stmt_declaracion", "stmt_asignar", "stmt_transferencia", 
			"if_stmt", "if_chain", "else_stmt", "switch_stmt", "switch_case", "default_case", 
			"while_stmt", "for_clasico_stmt", "for_stmt", "range", "declarar_funcion", 
			"llamar_funcion", "lista_parametros", "parametro_fun", "lista_argumentos", 
			"argumento_fun", "declarar_struct", "propiedad_struct", "crear_struct", 
			"lista_parametros_init", "parametros_init_struct", "lista_slice", "item_slice", 
			"prop_slice", "fun_slice", "tipo", "tipos_slices", "expr", "literal", 
			"patronId"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, "'mut'", "':='", "'+='", "'-='", "'='", "'int'", 
			null, "'string'", "'bool'", null, null, null, null, "'nil'", "'+'", "'-'", 
			"'*'", "'/'", "'%'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'", 
			"'||'", "'!'", "'main'", "'fn'", "'struct'", "'if'", "'else'", "'switch'", 
			"'case'", "'default'", "'for'", "'while'", "'break'", "'continue'", "'return'", 
			"'in'", null, "'('", "')'", "'{'", "'}'", "'['", "']'", "'.'", "','", 
			"';'", "':'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WS", "COMENTARIO", "COMENTARIOMULT", "RW_MUT", "OP_ASSIGN", "OP_INCREMENTO", 
			"OP_DECREMENTO", "OP_DECLARACION", "RW_INT", "RW_FLOAT64", "RW_STRING", 
			"RW_BOOL", "INT_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL", "BOOL_LITERAL", 
			"NIL_LITERAL", "OP_SUMA", "OP_RESTA", "OP_MULT", "OP_DIV", "OP_MOD", 
			"OP_IGUAL", "OP_DIFERENTE", "OP_MENORQ", "OP_MAYORQ", "OP_MENORIGUAL", 
			"OP_MAYORIGUAL", "OP_AND", "OP_OR", "OP_NOT", "RW_MAIN", "RW_FN", "RW_STRUCT", 
			"RW_IF", "RW_ELSE", "RW_SWITCH", "RW_CASE", "RW_DEFAULT", "RW_FOR", "RW_WHILE", 
			"RW_BREAK", "RW_CONTINUE", "RW_RETURN", "RW_IN", "ID", "PAR_IZQ", "PAR_DER", 
			"LLAVE_IZQ", "LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER", "PUNTO", "COMA", 
			"PUNTO_Y_COMA", "DOS_PUNTOS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "VGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public VGrammar(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public Main_funcContext main_func() {
			return getRuleContext(Main_funcContext.class,0);
		}
		public TerminalNode EOF() { return getToken(VGrammar.EOF, 0); }
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(75);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(72);
					stmt();
					}
					} 
				}
				setState(77);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			setState(79);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RW_FN) {
				{
				setState(78);
				main_func();
				}
			}

			setState(82);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				setState(81);
				match(EOF);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Main_funcContext extends ParserRuleContext {
		public Main_funcContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_main_func; }
	 
		public Main_funcContext() { }
		public void copyFrom(Main_funcContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionMainContext extends Main_funcContext {
		public TerminalNode RW_FN() { return getToken(VGrammar.RW_FN, 0); }
		public TerminalNode RW_MAIN() { return getToken(VGrammar.RW_MAIN, 0); }
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public FuncionMainContext(Main_funcContext ctx) { copyFrom(ctx); }
	}

	public final Main_funcContext main_func() throws RecognitionException {
		Main_funcContext _localctx = new Main_funcContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_main_func);
		int _la;
		try {
			_localctx = new FuncionMainContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(84);
			match(RW_FN);
			setState(85);
			match(RW_MAIN);
			setState(86);
			match(PAR_IZQ);
			setState(87);
			match(PAR_DER);
			setState(88);
			match(LLAVE_IZQ);
			setState(92);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 102452149878800L) != 0)) {
				{
				{
				setState(89);
				stmt();
				}
				}
				setState(94);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(95);
			match(LLAVE_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StmtContext extends ParserRuleContext {
		public Stmt_declaracionContext stmt_declaracion() {
			return getRuleContext(Stmt_declaracionContext.class,0);
		}
		public Stmt_asignarContext stmt_asignar() {
			return getRuleContext(Stmt_asignarContext.class,0);
		}
		public Stmt_transferenciaContext stmt_transferencia() {
			return getRuleContext(Stmt_transferenciaContext.class,0);
		}
		public If_stmtContext if_stmt() {
			return getRuleContext(If_stmtContext.class,0);
		}
		public Switch_stmtContext switch_stmt() {
			return getRuleContext(Switch_stmtContext.class,0);
		}
		public For_clasico_stmtContext for_clasico_stmt() {
			return getRuleContext(For_clasico_stmtContext.class,0);
		}
		public While_stmtContext while_stmt() {
			return getRuleContext(While_stmtContext.class,0);
		}
		public For_stmtContext for_stmt() {
			return getRuleContext(For_stmtContext.class,0);
		}
		public Llamar_funcionContext llamar_funcion() {
			return getRuleContext(Llamar_funcionContext.class,0);
		}
		public Fun_sliceContext fun_slice() {
			return getRuleContext(Fun_sliceContext.class,0);
		}
		public Declarar_funcionContext declarar_funcion() {
			return getRuleContext(Declarar_funcionContext.class,0);
		}
		public Declarar_structContext declarar_struct() {
			return getRuleContext(Declarar_structContext.class,0);
		}
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_stmt);
		try {
			setState(109);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(97);
				stmt_declaracion();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(98);
				stmt_asignar();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(99);
				stmt_transferencia();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(100);
				if_stmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(101);
				switch_stmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(102);
				for_clasico_stmt();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(103);
				while_stmt();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(104);
				for_stmt();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(105);
				llamar_funcion();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(106);
				fun_slice();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(107);
				declarar_funcion();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(108);
				declarar_struct();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Stmt_declaracionContext extends ParserRuleContext {
		public Stmt_declaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt_declaracion; }
	 
		public Stmt_declaracionContext() { }
		public void copyFrom(Stmt_declaracionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararSliceContext extends Stmt_declaracionContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclararSliceContext(Stmt_declaracionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararInferenciaContext extends Stmt_declaracionContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_ASSIGN() { return getToken(VGrammar.OP_ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclararInferenciaContext(Stmt_declaracionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclaraTipoValorContext extends Stmt_declaracionContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclaraTipoValorContext(Stmt_declaracionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararSinMutValorContext extends Stmt_declaracionContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclararSinMutValorContext(Stmt_declaracionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararInferenciaMutContext extends Stmt_declaracionContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_ASSIGN() { return getToken(VGrammar.OP_ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclararInferenciaMutContext(Stmt_declaracionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararTipoContext extends Stmt_declaracionContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public DeclararTipoContext(Stmt_declaracionContext ctx) { copyFrom(ctx); }
	}

	public final Stmt_declaracionContext stmt_declaracion() throws RecognitionException {
		Stmt_declaracionContext _localctx = new Stmt_declaracionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_stmt_declaracion);
		int _la;
		try {
			setState(141);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				_localctx = new DeclararInferenciaMutContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(111);
				match(RW_MUT);
				setState(112);
				match(ID);
				setState(113);
				match(OP_ASSIGN);
				setState(114);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DeclararInferenciaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(115);
				match(ID);
				setState(116);
				match(OP_ASSIGN);
				setState(117);
				expr(0);
				}
				break;
			case 3:
				_localctx = new DeclaraTipoValorContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(118);
				match(RW_MUT);
				setState(119);
				match(ID);
				setState(120);
				tipo();
				setState(121);
				match(OP_DECLARACION);
				setState(122);
				expr(0);
				}
				break;
			case 4:
				_localctx = new DeclararTipoContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(124);
				match(RW_MUT);
				setState(125);
				match(ID);
				setState(126);
				tipo();
				}
				break;
			case 5:
				_localctx = new DeclararSinMutValorContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(127);
				match(ID);
				setState(128);
				tipo();
				setState(129);
				match(OP_DECLARACION);
				setState(130);
				expr(0);
				}
				break;
			case 6:
				_localctx = new DeclararSliceContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(133);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==RW_MUT) {
					{
					setState(132);
					match(RW_MUT);
					}
				}

				setState(135);
				match(ID);
				setState(136);
				match(OP_DECLARACION);
				setState(137);
				tipo();
				setState(139);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
				case 1:
					{
					setState(138);
					expr(0);
					}
					break;
				}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Stmt_asignarContext extends ParserRuleContext {
		public Stmt_asignarContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt_asignar; }
	 
		public Stmt_asignarContext() { }
		public void copyFrom(Stmt_asignarContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionSliceItemContext extends Stmt_asignarContext {
		public Item_sliceContext item_slice() {
			return getRuleContext(Item_sliceContext.class,0);
		}
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AsignacionSliceItemContext(Stmt_asignarContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionDirectaContext extends Stmt_asignarContext {
		public PatronIdContext patronId() {
			return getRuleContext(PatronIdContext.class,0);
		}
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AsignacionDirectaContext(Stmt_asignarContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionSliceContext extends Stmt_asignarContext {
		public Token op;
		public Item_sliceContext item_slice() {
			return getRuleContext(Item_sliceContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode OP_INCREMENTO() { return getToken(VGrammar.OP_INCREMENTO, 0); }
		public TerminalNode OP_DECREMENTO() { return getToken(VGrammar.OP_DECREMENTO, 0); }
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
		public AsignacionSliceContext(Stmt_asignarContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionAritmeticaContext extends Stmt_asignarContext {
		public Token op;
		public PatronIdContext patronId() {
			return getRuleContext(PatronIdContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode OP_INCREMENTO() { return getToken(VGrammar.OP_INCREMENTO, 0); }
		public TerminalNode OP_DECREMENTO() { return getToken(VGrammar.OP_DECREMENTO, 0); }
		public AsignacionAritmeticaContext(Stmt_asignarContext ctx) { copyFrom(ctx); }
	}

	public final Stmt_asignarContext stmt_asignar() throws RecognitionException {
		Stmt_asignarContext _localctx = new Stmt_asignarContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_stmt_asignar);
		int _la;
		try {
			setState(159);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				_localctx = new AsignacionSliceItemContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(143);
				item_slice();
				setState(144);
				match(OP_DECLARACION);
				setState(145);
				expr(0);
				}
				break;
			case 2:
				_localctx = new AsignacionDirectaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(147);
				patronId();
				setState(148);
				match(OP_DECLARACION);
				setState(149);
				expr(0);
				}
				break;
			case 3:
				_localctx = new AsignacionAritmeticaContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(151);
				patronId();
				setState(152);
				((AsignacionAritmeticaContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==OP_INCREMENTO || _la==OP_DECREMENTO) ) {
					((AsignacionAritmeticaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(153);
				expr(0);
				}
				break;
			case 4:
				_localctx = new AsignacionSliceContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(155);
				item_slice();
				setState(156);
				((AsignacionSliceContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 448L) != 0)) ) {
					((AsignacionSliceContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(157);
				expr(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Stmt_transferenciaContext extends ParserRuleContext {
		public Stmt_transferenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt_transferencia; }
	 
		public Stmt_transferenciaContext() { }
		public void copyFrom(Stmt_transferenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStmtContext extends Stmt_transferenciaContext {
		public TerminalNode RW_CONTINUE() { return getToken(VGrammar.RW_CONTINUE, 0); }
		public ContinueStmtContext(Stmt_transferenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStmtContext extends Stmt_transferenciaContext {
		public TerminalNode RW_BREAK() { return getToken(VGrammar.RW_BREAK, 0); }
		public BreakStmtContext(Stmt_transferenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStmtContext extends Stmt_transferenciaContext {
		public TerminalNode RW_RETURN() { return getToken(VGrammar.RW_RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ReturnStmtContext(Stmt_transferenciaContext ctx) { copyFrom(ctx); }
	}

	public final Stmt_transferenciaContext stmt_transferencia() throws RecognitionException {
		Stmt_transferenciaContext _localctx = new Stmt_transferenciaContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_stmt_transferencia);
		try {
			setState(167);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RW_RETURN:
				_localctx = new ReturnStmtContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(161);
				match(RW_RETURN);
				setState(163);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
				case 1:
					{
					setState(162);
					expr(0);
					}
					break;
				}
				}
				break;
			case RW_BREAK:
				_localctx = new BreakStmtContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(165);
				match(RW_BREAK);
				}
				break;
			case RW_CONTINUE:
				_localctx = new ContinueStmtContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(166);
				match(RW_CONTINUE);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_stmtContext extends ParserRuleContext {
		public If_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_stmt; }
	 
		public If_stmtContext() { }
		public void copyFrom(If_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IFstmtContext extends If_stmtContext {
		public List<If_chainContext> if_chain() {
			return getRuleContexts(If_chainContext.class);
		}
		public If_chainContext if_chain(int i) {
			return getRuleContext(If_chainContext.class,i);
		}
		public List<TerminalNode> RW_ELSE() { return getTokens(VGrammar.RW_ELSE); }
		public TerminalNode RW_ELSE(int i) {
			return getToken(VGrammar.RW_ELSE, i);
		}
		public Else_stmtContext else_stmt() {
			return getRuleContext(Else_stmtContext.class,0);
		}
		public IFstmtContext(If_stmtContext ctx) { copyFrom(ctx); }
	}

	public final If_stmtContext if_stmt() throws RecognitionException {
		If_stmtContext _localctx = new If_stmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_if_stmt);
		int _la;
		try {
			int _alt;
			_localctx = new IFstmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(169);
			if_chain();
			setState(174);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(170);
					match(RW_ELSE);
					setState(171);
					if_chain();
					}
					} 
				}
				setState(176);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			}
			setState(178);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RW_ELSE) {
				{
				setState(177);
				else_stmt();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_chainContext extends ParserRuleContext {
		public If_chainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_chain; }
	 
		public If_chainContext() { }
		public void copyFrom(If_chainContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IFcadenaContext extends If_chainContext {
		public TerminalNode RW_IF() { return getToken(VGrammar.RW_IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public IFcadenaContext(If_chainContext ctx) { copyFrom(ctx); }
	}

	public final If_chainContext if_chain() throws RecognitionException {
		If_chainContext _localctx = new If_chainContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_if_chain);
		int _la;
		try {
			_localctx = new IFcadenaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(180);
			match(RW_IF);
			setState(181);
			expr(0);
			setState(182);
			match(LLAVE_IZQ);
			setState(186);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 102452149878800L) != 0)) {
				{
				{
				setState(183);
				stmt();
				}
				}
				setState(188);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(189);
			match(LLAVE_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Else_stmtContext extends ParserRuleContext {
		public Else_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_else_stmt; }
	 
		public Else_stmtContext() { }
		public void copyFrom(Else_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ElseStmtContext extends Else_stmtContext {
		public TerminalNode RW_ELSE() { return getToken(VGrammar.RW_ELSE, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ElseStmtContext(Else_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Else_stmtContext else_stmt() throws RecognitionException {
		Else_stmtContext _localctx = new Else_stmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_else_stmt);
		int _la;
		try {
			_localctx = new ElseStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(191);
			match(RW_ELSE);
			setState(192);
			match(LLAVE_IZQ);
			setState(196);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 102452149878800L) != 0)) {
				{
				{
				setState(193);
				stmt();
				}
				}
				setState(198);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(199);
			match(LLAVE_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Switch_stmtContext extends ParserRuleContext {
		public Switch_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_stmt; }
	 
		public Switch_stmtContext() { }
		public void copyFrom(Switch_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStmtContext extends Switch_stmtContext {
		public TerminalNode RW_SWITCH() { return getToken(VGrammar.RW_SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public List<Switch_caseContext> switch_case() {
			return getRuleContexts(Switch_caseContext.class);
		}
		public Switch_caseContext switch_case(int i) {
			return getRuleContext(Switch_caseContext.class,i);
		}
		public Default_caseContext default_case() {
			return getRuleContext(Default_caseContext.class,0);
		}
		public SwitchStmtContext(Switch_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Switch_stmtContext switch_stmt() throws RecognitionException {
		Switch_stmtContext _localctx = new Switch_stmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_switch_stmt);
		int _la;
		try {
			_localctx = new SwitchStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(201);
			match(RW_SWITCH);
			setState(202);
			expr(0);
			setState(203);
			match(LLAVE_IZQ);
			setState(207);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==RW_CASE) {
				{
				{
				setState(204);
				switch_case();
				}
				}
				setState(209);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(211);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RW_DEFAULT) {
				{
				setState(210);
				default_case();
				}
			}

			setState(213);
			match(LLAVE_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Switch_caseContext extends ParserRuleContext {
		public Switch_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_case; }
	 
		public Switch_caseContext() { }
		public void copyFrom(Switch_caseContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchCaseContext extends Switch_caseContext {
		public TerminalNode RW_CASE() { return getToken(VGrammar.RW_CASE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode DOS_PUNTOS() { return getToken(VGrammar.DOS_PUNTOS, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public SwitchCaseContext(Switch_caseContext ctx) { copyFrom(ctx); }
	}

	public final Switch_caseContext switch_case() throws RecognitionException {
		Switch_caseContext _localctx = new Switch_caseContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_switch_case);
		int _la;
		try {
			_localctx = new SwitchCaseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(215);
			match(RW_CASE);
			setState(216);
			expr(0);
			setState(217);
			match(DOS_PUNTOS);
			setState(221);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 102452149878800L) != 0)) {
				{
				{
				setState(218);
				stmt();
				}
				}
				setState(223);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Default_caseContext extends ParserRuleContext {
		public Default_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default_case; }
	 
		public Default_caseContext() { }
		public void copyFrom(Default_caseContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DefaultCaseContext extends Default_caseContext {
		public TerminalNode RW_DEFAULT() { return getToken(VGrammar.RW_DEFAULT, 0); }
		public TerminalNode DOS_PUNTOS() { return getToken(VGrammar.DOS_PUNTOS, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public DefaultCaseContext(Default_caseContext ctx) { copyFrom(ctx); }
	}

	public final Default_caseContext default_case() throws RecognitionException {
		Default_caseContext _localctx = new Default_caseContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_default_case);
		int _la;
		try {
			_localctx = new DefaultCaseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(224);
			match(RW_DEFAULT);
			setState(225);
			match(DOS_PUNTOS);
			setState(229);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 102452149878800L) != 0)) {
				{
				{
				setState(226);
				stmt();
				}
				}
				setState(231);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class While_stmtContext extends ParserRuleContext {
		public While_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_stmt; }
	 
		public While_stmtContext() { }
		public void copyFrom(While_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class WhileStmtContext extends While_stmtContext {
		public TerminalNode RW_FOR() { return getToken(VGrammar.RW_FOR, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public WhileStmtContext(While_stmtContext ctx) { copyFrom(ctx); }
	}

	public final While_stmtContext while_stmt() throws RecognitionException {
		While_stmtContext _localctx = new While_stmtContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_while_stmt);
		int _la;
		try {
			_localctx = new WhileStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(232);
			match(RW_FOR);
			setState(233);
			expr(0);
			setState(234);
			match(LLAVE_IZQ);
			setState(238);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 102452149878800L) != 0)) {
				{
				{
				setState(235);
				stmt();
				}
				}
				setState(240);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(241);
			match(LLAVE_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class For_clasico_stmtContext extends ParserRuleContext {
		public For_clasico_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_clasico_stmt; }
	 
		public For_clasico_stmtContext() { }
		public void copyFrom(For_clasico_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForClasicoStmtContext extends For_clasico_stmtContext {
		public TerminalNode RW_FOR() { return getToken(VGrammar.RW_FOR, 0); }
		public List<TerminalNode> PUNTO_Y_COMA() { return getTokens(VGrammar.PUNTO_Y_COMA); }
		public TerminalNode PUNTO_Y_COMA(int i) {
			return getToken(VGrammar.PUNTO_Y_COMA, i);
		}
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public List<Stmt_asignarContext> stmt_asignar() {
			return getRuleContexts(Stmt_asignarContext.class);
		}
		public Stmt_asignarContext stmt_asignar(int i) {
			return getRuleContext(Stmt_asignarContext.class,i);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ForClasicoStmtContext(For_clasico_stmtContext ctx) { copyFrom(ctx); }
	}

	public final For_clasico_stmtContext for_clasico_stmt() throws RecognitionException {
		For_clasico_stmtContext _localctx = new For_clasico_stmtContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_for_clasico_stmt);
		int _la;
		try {
			_localctx = new ForClasicoStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(243);
			match(RW_FOR);
			setState(245);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(244);
				stmt_asignar();
				}
			}

			setState(247);
			match(PUNTO_Y_COMA);
			setState(249);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 774058334216192L) != 0)) {
				{
				setState(248);
				expr(0);
				}
			}

			setState(251);
			match(PUNTO_Y_COMA);
			setState(253);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(252);
				stmt_asignar();
				}
			}

			setState(255);
			match(LLAVE_IZQ);
			setState(259);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 102452149878800L) != 0)) {
				{
				{
				setState(256);
				stmt();
				}
				}
				setState(261);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(262);
			match(LLAVE_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class For_stmtContext extends ParserRuleContext {
		public For_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_stmt; }
	 
		public For_stmtContext() { }
		public void copyFrom(For_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForStmtContext extends For_stmtContext {
		public TerminalNode RW_FOR() { return getToken(VGrammar.RW_FOR, 0); }
		public List<TerminalNode> ID() { return getTokens(VGrammar.ID); }
		public TerminalNode ID(int i) {
			return getToken(VGrammar.ID, i);
		}
		public TerminalNode RW_IN() { return getToken(VGrammar.RW_IN, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public RangeContext range() {
			return getRuleContext(RangeContext.class,0);
		}
		public TerminalNode COMA() { return getToken(VGrammar.COMA, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ForStmtContext(For_stmtContext ctx) { copyFrom(ctx); }
	}

	public final For_stmtContext for_stmt() throws RecognitionException {
		For_stmtContext _localctx = new For_stmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_for_stmt);
		int _la;
		try {
			_localctx = new ForStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(264);
			match(RW_FOR);
			setState(265);
			match(ID);
			setState(268);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMA) {
				{
				setState(266);
				match(COMA);
				setState(267);
				match(ID);
				}
			}

			setState(270);
			match(RW_IN);
			setState(273);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				{
				setState(271);
				expr(0);
				}
				break;
			case 2:
				{
				setState(272);
				range();
				}
				break;
			}
			setState(275);
			match(LLAVE_IZQ);
			setState(279);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 102452149878800L) != 0)) {
				{
				{
				setState(276);
				stmt();
				}
				}
				setState(281);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(282);
			match(LLAVE_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RangeContext extends ParserRuleContext {
		public RangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_range; }
	 
		public RangeContext() { }
		public void copyFrom(RangeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RangoNumContext extends RangeContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> PUNTO() { return getTokens(VGrammar.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(VGrammar.PUNTO, i);
		}
		public RangoNumContext(RangeContext ctx) { copyFrom(ctx); }
	}

	public final RangeContext range() throws RecognitionException {
		RangeContext _localctx = new RangeContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_range);
		try {
			_localctx = new RangoNumContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(284);
			expr(0);
			setState(285);
			match(PUNTO);
			setState(286);
			match(PUNTO);
			setState(287);
			match(PUNTO);
			setState(288);
			expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Declarar_funcionContext extends ParserRuleContext {
		public Declarar_funcionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarar_funcion; }
	 
		public Declarar_funcionContext() { }
		public void copyFrom(Declarar_funcionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionDecleradaContext extends Declarar_funcionContext {
		public TerminalNode RW_FN() { return getToken(VGrammar.RW_FN, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public Lista_parametrosContext lista_parametros() {
			return getRuleContext(Lista_parametrosContext.class,0);
		}
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public FuncionDecleradaContext(Declarar_funcionContext ctx) { copyFrom(ctx); }
	}

	public final Declarar_funcionContext declarar_funcion() throws RecognitionException {
		Declarar_funcionContext _localctx = new Declarar_funcionContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_declarar_funcion);
		int _la;
		try {
			_localctx = new FuncionDecleradaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(290);
			match(RW_FN);
			setState(291);
			match(ID);
			setState(292);
			match(PAR_IZQ);
			setState(294);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(293);
				lista_parametros();
				}
			}

			setState(296);
			match(PAR_DER);
			setState(298);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2322168557870592L) != 0)) {
				{
				setState(297);
				tipo();
				}
			}

			setState(300);
			match(LLAVE_IZQ);
			setState(304);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 102452149878800L) != 0)) {
				{
				{
				setState(301);
				stmt();
				}
				}
				setState(306);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(307);
			match(LLAVE_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Llamar_funcionContext extends ParserRuleContext {
		public Llamar_funcionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_llamar_funcion; }
	 
		public Llamar_funcionContext() { }
		public void copyFrom(Llamar_funcionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LlamarFuncionContext extends Llamar_funcionContext {
		public PatronIdContext patronId() {
			return getRuleContext(PatronIdContext.class,0);
		}
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public Lista_argumentosContext lista_argumentos() {
			return getRuleContext(Lista_argumentosContext.class,0);
		}
		public LlamarFuncionContext(Llamar_funcionContext ctx) { copyFrom(ctx); }
	}

	public final Llamar_funcionContext llamar_funcion() throws RecognitionException {
		Llamar_funcionContext _localctx = new Llamar_funcionContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_llamar_funcion);
		int _la;
		try {
			_localctx = new LlamarFuncionContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(309);
			patronId();
			setState(310);
			match(PAR_IZQ);
			setState(312);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 774058334216192L) != 0)) {
				{
				setState(311);
				lista_argumentos();
				}
			}

			setState(314);
			match(PAR_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Lista_parametrosContext extends ParserRuleContext {
		public Lista_parametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista_parametros; }
	 
		public Lista_parametrosContext() { }
		public void copyFrom(Lista_parametrosContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ListaParametrosContext extends Lista_parametrosContext {
		public List<Parametro_funContext> parametro_fun() {
			return getRuleContexts(Parametro_funContext.class);
		}
		public Parametro_funContext parametro_fun(int i) {
			return getRuleContext(Parametro_funContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(VGrammar.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(VGrammar.COMA, i);
		}
		public ListaParametrosContext(Lista_parametrosContext ctx) { copyFrom(ctx); }
	}

	public final Lista_parametrosContext lista_parametros() throws RecognitionException {
		Lista_parametrosContext _localctx = new Lista_parametrosContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_lista_parametros);
		int _la;
		try {
			int _alt;
			_localctx = new ListaParametrosContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(316);
			parametro_fun();
			setState(321);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,31,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(317);
					match(COMA);
					setState(318);
					parametro_fun();
					}
					} 
				}
				setState(323);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,31,_ctx);
			}
			setState(325);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMA) {
				{
				setState(324);
				match(COMA);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Parametro_funContext extends ParserRuleContext {
		public Parametro_funContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametro_fun; }
	 
		public Parametro_funContext() { }
		public void copyFrom(Parametro_funContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParametroFunContext extends Parametro_funContext {
		public List<TerminalNode> ID() { return getTokens(VGrammar.ID); }
		public TerminalNode ID(int i) {
			return getToken(VGrammar.ID, i);
		}
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public ParametroFunContext(Parametro_funContext ctx) { copyFrom(ctx); }
	}

	public final Parametro_funContext parametro_fun() throws RecognitionException {
		Parametro_funContext _localctx = new Parametro_funContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_parametro_fun);
		try {
			_localctx = new ParametroFunContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(328);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				setState(327);
				match(ID);
				}
				break;
			}
			setState(330);
			match(ID);
			setState(331);
			tipo();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Lista_argumentosContext extends ParserRuleContext {
		public Lista_argumentosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista_argumentos; }
	 
		public Lista_argumentosContext() { }
		public void copyFrom(Lista_argumentosContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ListaArgumentosContext extends Lista_argumentosContext {
		public List<Argumento_funContext> argumento_fun() {
			return getRuleContexts(Argumento_funContext.class);
		}
		public Argumento_funContext argumento_fun(int i) {
			return getRuleContext(Argumento_funContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(VGrammar.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(VGrammar.COMA, i);
		}
		public ListaArgumentosContext(Lista_argumentosContext ctx) { copyFrom(ctx); }
	}

	public final Lista_argumentosContext lista_argumentos() throws RecognitionException {
		Lista_argumentosContext _localctx = new Lista_argumentosContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_lista_argumentos);
		int _la;
		try {
			int _alt;
			_localctx = new ListaArgumentosContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(333);
			argumento_fun();
			setState(338);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,34,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(334);
					match(COMA);
					setState(335);
					argumento_fun();
					}
					} 
				}
				setState(340);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,34,_ctx);
			}
			setState(342);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMA) {
				{
				setState(341);
				match(COMA);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Argumento_funContext extends ParserRuleContext {
		public Argumento_funContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argumento_fun; }
	 
		public Argumento_funContext() { }
		public void copyFrom(Argumento_funContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionArgContext extends Argumento_funContext {
		public PatronIdContext patronId() {
			return getRuleContext(PatronIdContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode DOS_PUNTOS() { return getToken(VGrammar.DOS_PUNTOS, 0); }
		public FuncionArgContext(Argumento_funContext ctx) { copyFrom(ctx); }
	}

	public final Argumento_funContext argumento_fun() throws RecognitionException {
		Argumento_funContext _localctx = new Argumento_funContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_argumento_fun);
		try {
			_localctx = new FuncionArgContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(346);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				{
				setState(344);
				match(ID);
				setState(345);
				match(DOS_PUNTOS);
				}
				break;
			}
			setState(350);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
			case 1:
				{
				setState(348);
				patronId();
				}
				break;
			case 2:
				{
				setState(349);
				expr(0);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Declarar_structContext extends ParserRuleContext {
		public Declarar_structContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarar_struct; }
	 
		public Declarar_structContext() { }
		public void copyFrom(Declarar_structContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararStructContext extends Declarar_structContext {
		public TerminalNode RW_STRUCT() { return getToken(VGrammar.RW_STRUCT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public List<Propiedad_structContext> propiedad_struct() {
			return getRuleContexts(Propiedad_structContext.class);
		}
		public Propiedad_structContext propiedad_struct(int i) {
			return getRuleContext(Propiedad_structContext.class,i);
		}
		public DeclararStructContext(Declarar_structContext ctx) { copyFrom(ctx); }
	}

	public final Declarar_structContext declarar_struct() throws RecognitionException {
		Declarar_structContext _localctx = new Declarar_structContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_declarar_struct);
		int _la;
		try {
			_localctx = new DeclararStructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(352);
			match(RW_STRUCT);
			setState(353);
			match(ID);
			setState(354);
			match(LLAVE_IZQ);
			setState(358);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2322168557870592L) != 0)) {
				{
				{
				setState(355);
				propiedad_struct();
				}
				}
				setState(360);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(361);
			match(LLAVE_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Propiedad_structContext extends ParserRuleContext {
		public Propiedad_structContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_propiedad_struct; }
	 
		public Propiedad_structContext() { }
		public void copyFrom(Propiedad_structContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PropiedadStructContext extends Propiedad_structContext {
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public PropiedadStructContext(Propiedad_structContext ctx) { copyFrom(ctx); }
	}

	public final Propiedad_structContext propiedad_struct() throws RecognitionException {
		Propiedad_structContext _localctx = new Propiedad_structContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_propiedad_struct);
		int _la;
		try {
			_localctx = new PropiedadStructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			tipo();
			setState(364);
			match(ID);
			setState(367);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OP_DECLARACION) {
				{
				setState(365);
				match(OP_DECLARACION);
				setState(366);
				expr(0);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Crear_structContext extends ParserRuleContext {
		public Crear_structContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_crear_struct; }
	 
		public Crear_structContext() { }
		public void copyFrom(Crear_structContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CrearStructContext extends Crear_structContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public Lista_parametros_initContext lista_parametros_init() {
			return getRuleContext(Lista_parametros_initContext.class,0);
		}
		public CrearStructContext(Crear_structContext ctx) { copyFrom(ctx); }
	}

	public final Crear_structContext crear_struct() throws RecognitionException {
		Crear_structContext _localctx = new Crear_structContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_crear_struct);
		int _la;
		try {
			_localctx = new CrearStructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(369);
			match(ID);
			setState(370);
			match(LLAVE_IZQ);
			setState(372);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(371);
				lista_parametros_init();
				}
			}

			setState(374);
			match(LLAVE_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Lista_parametros_initContext extends ParserRuleContext {
		public Lista_parametros_initContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista_parametros_init; }
	 
		public Lista_parametros_initContext() { }
		public void copyFrom(Lista_parametros_initContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ListaParametrosInitContext extends Lista_parametros_initContext {
		public List<Parametros_init_structContext> parametros_init_struct() {
			return getRuleContexts(Parametros_init_structContext.class);
		}
		public Parametros_init_structContext parametros_init_struct(int i) {
			return getRuleContext(Parametros_init_structContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(VGrammar.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(VGrammar.COMA, i);
		}
		public ListaParametrosInitContext(Lista_parametros_initContext ctx) { copyFrom(ctx); }
	}

	public final Lista_parametros_initContext lista_parametros_init() throws RecognitionException {
		Lista_parametros_initContext _localctx = new Lista_parametros_initContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_lista_parametros_init);
		int _la;
		try {
			int _alt;
			_localctx = new ListaParametrosInitContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(376);
			parametros_init_struct();
			setState(381);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,41,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(377);
					match(COMA);
					setState(378);
					parametros_init_struct();
					}
					} 
				}
				setState(383);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,41,_ctx);
			}
			setState(385);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMA) {
				{
				setState(384);
				match(COMA);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Parametros_init_structContext extends ParserRuleContext {
		public Parametros_init_structContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametros_init_struct; }
	 
		public Parametros_init_structContext() { }
		public void copyFrom(Parametros_init_structContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParametrosInitStructContext extends Parametros_init_structContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode DOS_PUNTOS() { return getToken(VGrammar.DOS_PUNTOS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ParametrosInitStructContext(Parametros_init_structContext ctx) { copyFrom(ctx); }
	}

	public final Parametros_init_structContext parametros_init_struct() throws RecognitionException {
		Parametros_init_structContext _localctx = new Parametros_init_structContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_parametros_init_struct);
		try {
			_localctx = new ParametrosInitStructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(387);
			match(ID);
			setState(388);
			match(DOS_PUNTOS);
			setState(389);
			expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Lista_sliceContext extends ParserRuleContext {
		public Lista_sliceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lista_slice; }
	 
		public Lista_sliceContext() { }
		public void copyFrom(Lista_sliceContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ListaSliceContext extends Lista_sliceContext {
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(VGrammar.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(VGrammar.COMA, i);
		}
		public ListaSliceContext(Lista_sliceContext ctx) { copyFrom(ctx); }
	}

	public final Lista_sliceContext lista_slice() throws RecognitionException {
		Lista_sliceContext _localctx = new Lista_sliceContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_lista_slice);
		int _la;
		try {
			int _alt;
			_localctx = new ListaSliceContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(391);
			match(LLAVE_IZQ);
			setState(400);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 774058334216192L) != 0)) {
				{
				setState(392);
				expr(0);
				setState(397);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(393);
						match(COMA);
						setState(394);
						expr(0);
						}
						} 
					}
					setState(399);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
				}
				}
			}

			setState(403);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMA) {
				{
				setState(402);
				match(COMA);
				}
			}

			setState(405);
			match(LLAVE_DER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Item_sliceContext extends ParserRuleContext {
		public Item_sliceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_item_slice; }
	 
		public Item_sliceContext() { }
		public void copyFrom(Item_sliceContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ItemSliceContext extends Item_sliceContext {
		public PatronIdContext patronId() {
			return getRuleContext(PatronIdContext.class,0);
		}
		public List<TerminalNode> CORCHETE_IZQ() { return getTokens(VGrammar.CORCHETE_IZQ); }
		public TerminalNode CORCHETE_IZQ(int i) {
			return getToken(VGrammar.CORCHETE_IZQ, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> CORCHETE_DER() { return getTokens(VGrammar.CORCHETE_DER); }
		public TerminalNode CORCHETE_DER(int i) {
			return getToken(VGrammar.CORCHETE_DER, i);
		}
		public ItemSliceContext(Item_sliceContext ctx) { copyFrom(ctx); }
	}

	public final Item_sliceContext item_slice() throws RecognitionException {
		Item_sliceContext _localctx = new Item_sliceContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_item_slice);
		try {
			int _alt;
			_localctx = new ItemSliceContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(407);
			patronId();
			setState(412); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(408);
					match(CORCHETE_IZQ);
					setState(409);
					expr(0);
					setState(410);
					match(CORCHETE_DER);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(414); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,46,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Prop_sliceContext extends ParserRuleContext {
		public Prop_sliceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prop_slice; }
	 
		public Prop_sliceContext() { }
		public void copyFrom(Prop_sliceContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PropSliceContext extends Prop_sliceContext {
		public Item_sliceContext item_slice() {
			return getRuleContext(Item_sliceContext.class,0);
		}
		public TerminalNode PUNTO() { return getToken(VGrammar.PUNTO, 0); }
		public PatronIdContext patronId() {
			return getRuleContext(PatronIdContext.class,0);
		}
		public PropSliceContext(Prop_sliceContext ctx) { copyFrom(ctx); }
	}

	public final Prop_sliceContext prop_slice() throws RecognitionException {
		Prop_sliceContext _localctx = new Prop_sliceContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_prop_slice);
		try {
			_localctx = new PropSliceContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(416);
			item_slice();
			setState(417);
			match(PUNTO);
			setState(418);
			patronId();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Fun_sliceContext extends ParserRuleContext {
		public Fun_sliceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fun_slice; }
	 
		public Fun_sliceContext() { }
		public void copyFrom(Fun_sliceContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionSliceContext extends Fun_sliceContext {
		public Item_sliceContext item_slice() {
			return getRuleContext(Item_sliceContext.class,0);
		}
		public TerminalNode PUNTO() { return getToken(VGrammar.PUNTO, 0); }
		public Llamar_funcionContext llamar_funcion() {
			return getRuleContext(Llamar_funcionContext.class,0);
		}
		public FuncionSliceContext(Fun_sliceContext ctx) { copyFrom(ctx); }
	}

	public final Fun_sliceContext fun_slice() throws RecognitionException {
		Fun_sliceContext _localctx = new Fun_sliceContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_fun_slice);
		try {
			_localctx = new FuncionSliceContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(420);
			item_slice();
			setState(421);
			match(PUNTO);
			setState(422);
			llamar_funcion();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TipoContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_INT() { return getToken(VGrammar.RW_INT, 0); }
		public TerminalNode RW_FLOAT64() { return getToken(VGrammar.RW_FLOAT64, 0); }
		public TerminalNode RW_STRING() { return getToken(VGrammar.RW_STRING, 0); }
		public TerminalNode RW_BOOL() { return getToken(VGrammar.RW_BOOL, 0); }
		public Tipos_slicesContext tipos_slices() {
			return getRuleContext(Tipos_slicesContext.class,0);
		}
		public TipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo; }
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_tipo);
		try {
			setState(430);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				setState(424);
				match(ID);
				}
				break;
			case RW_INT:
				enterOuterAlt(_localctx, 2);
				{
				setState(425);
				match(RW_INT);
				}
				break;
			case RW_FLOAT64:
				enterOuterAlt(_localctx, 3);
				{
				setState(426);
				match(RW_FLOAT64);
				}
				break;
			case RW_STRING:
				enterOuterAlt(_localctx, 4);
				{
				setState(427);
				match(RW_STRING);
				}
				break;
			case RW_BOOL:
				enterOuterAlt(_localctx, 5);
				{
				setState(428);
				match(RW_BOOL);
				}
				break;
			case CORCHETE_IZQ:
				enterOuterAlt(_localctx, 6);
				{
				setState(429);
				tipos_slices();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Tipos_slicesContext extends ParserRuleContext {
		public Tipos_slicesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipos_slices; }
	 
		public Tipos_slicesContext() { }
		public void copyFrom(Tipos_slicesContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorSimpleContext extends Tipos_slicesContext {
		public TerminalNode CORCHETE_IZQ() { return getToken(VGrammar.CORCHETE_IZQ, 0); }
		public TerminalNode CORCHETE_DER() { return getToken(VGrammar.CORCHETE_DER, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public VectorSimpleContext(Tipos_slicesContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MatrizDobleContext extends Tipos_slicesContext {
		public List<TerminalNode> CORCHETE_IZQ() { return getTokens(VGrammar.CORCHETE_IZQ); }
		public TerminalNode CORCHETE_IZQ(int i) {
			return getToken(VGrammar.CORCHETE_IZQ, i);
		}
		public List<TerminalNode> CORCHETE_DER() { return getTokens(VGrammar.CORCHETE_DER); }
		public TerminalNode CORCHETE_DER(int i) {
			return getToken(VGrammar.CORCHETE_DER, i);
		}
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public MatrizDobleContext(Tipos_slicesContext ctx) { copyFrom(ctx); }
	}

	public final Tipos_slicesContext tipos_slices() throws RecognitionException {
		Tipos_slicesContext _localctx = new Tipos_slicesContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_tipos_slices);
		try {
			setState(440);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,48,_ctx) ) {
			case 1:
				_localctx = new VectorSimpleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(432);
				match(CORCHETE_IZQ);
				setState(433);
				match(CORCHETE_DER);
				setState(434);
				tipo();
				}
				break;
			case 2:
				_localctx = new MatrizDobleContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(435);
				match(CORCHETE_IZQ);
				setState(436);
				match(CORCHETE_DER);
				setState(437);
				match(CORCHETE_IZQ);
				setState(438);
				match(CORCHETE_DER);
				setState(439);
				tipo();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LiteralExpContext extends ExprContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public LiteralExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdExpContext extends ExprContext {
		public PatronIdContext patronId() {
			return getRuleContext(PatronIdContext.class,0);
		}
		public IdExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CrearStructExpContext extends ExprContext {
		public Crear_structContext crear_struct() {
			return getRuleContext(Crear_structContext.class,0);
		}
		public CrearStructExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UnarioExpContext extends ExprContext {
		public Token op;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode OP_NOT() { return getToken(VGrammar.OP_NOT, 0); }
		public TerminalNode OP_RESTA() { return getToken(VGrammar.OP_RESTA, 0); }
		public UnarioExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ItemSliceExpContext extends ExprContext {
		public Item_sliceContext item_slice() {
			return getRuleContext(Item_sliceContext.class,0);
		}
		public ItemSliceExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorFuncExpContext extends ExprContext {
		public Fun_sliceContext fun_slice() {
			return getRuleContext(Fun_sliceContext.class,0);
		}
		public VectorFuncExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceExpContext extends ExprContext {
		public Lista_sliceContext lista_slice() {
			return getRuleContext(Lista_sliceContext.class,0);
		}
		public SliceExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PropSliceExpContext extends ExprContext {
		public Prop_sliceContext prop_slice() {
			return getRuleContext(Prop_sliceContext.class,0);
		}
		public PropSliceExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LlamarFuncionExpContext extends ExprContext {
		public Llamar_funcionContext llamar_funcion() {
			return getRuleContext(Llamar_funcionContext.class,0);
		}
		public LlamarFuncionExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BinarioExpContext extends ExprContext {
		public ExprContext left;
		public Token op;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_MULT() { return getToken(VGrammar.OP_MULT, 0); }
		public TerminalNode OP_DIV() { return getToken(VGrammar.OP_DIV, 0); }
		public TerminalNode OP_MOD() { return getToken(VGrammar.OP_MOD, 0); }
		public TerminalNode OP_SUMA() { return getToken(VGrammar.OP_SUMA, 0); }
		public TerminalNode OP_RESTA() { return getToken(VGrammar.OP_RESTA, 0); }
		public TerminalNode OP_MENORQ() { return getToken(VGrammar.OP_MENORQ, 0); }
		public TerminalNode OP_MENORIGUAL() { return getToken(VGrammar.OP_MENORIGUAL, 0); }
		public TerminalNode OP_MAYORQ() { return getToken(VGrammar.OP_MAYORQ, 0); }
		public TerminalNode OP_MAYORIGUAL() { return getToken(VGrammar.OP_MAYORIGUAL, 0); }
		public TerminalNode OP_IGUAL() { return getToken(VGrammar.OP_IGUAL, 0); }
		public TerminalNode OP_DIFERENTE() { return getToken(VGrammar.OP_DIFERENTE, 0); }
		public TerminalNode OP_AND() { return getToken(VGrammar.OP_AND, 0); }
		public TerminalNode OP_OR() { return getToken(VGrammar.OP_OR, 0); }
		public BinarioExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParentecisExpContext extends ExprContext {
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public ParentecisExpContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 66;
		enterRecursionRule(_localctx, 66, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(457);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,49,_ctx) ) {
			case 1:
				{
				_localctx = new ParentecisExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(443);
				match(PAR_IZQ);
				setState(444);
				expr(0);
				setState(445);
				match(PAR_DER);
				}
				break;
			case 2:
				{
				_localctx = new LlamarFuncionExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(447);
				llamar_funcion();
				}
				break;
			case 3:
				{
				_localctx = new IdExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(448);
				patronId();
				}
				break;
			case 4:
				{
				_localctx = new ItemSliceExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(449);
				item_slice();
				}
				break;
			case 5:
				{
				_localctx = new PropSliceExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(450);
				prop_slice();
				}
				break;
			case 6:
				{
				_localctx = new VectorFuncExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(451);
				fun_slice();
				}
				break;
			case 7:
				{
				_localctx = new LiteralExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(452);
				literal();
				}
				break;
			case 8:
				{
				_localctx = new SliceExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(453);
				lista_slice();
				}
				break;
			case 9:
				{
				_localctx = new CrearStructExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(454);
				crear_struct();
				}
				break;
			case 10:
				{
				_localctx = new UnarioExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(455);
				((UnarioExpContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==OP_RESTA || _la==OP_NOT) ) {
					((UnarioExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(456);
				expr(7);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(479);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,51,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(477);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,50,_ctx) ) {
					case 1:
						{
						_localctx = new BinarioExpContext(new ExprContext(_parentctx, _parentState));
						((BinarioExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(459);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(460);
						((BinarioExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 7340032L) != 0)) ) {
							((BinarioExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(461);
						((BinarioExpContext)_localctx).right = expr(7);
						}
						break;
					case 2:
						{
						_localctx = new BinarioExpContext(new ExprContext(_parentctx, _parentState));
						((BinarioExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(462);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(463);
						((BinarioExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_SUMA || _la==OP_RESTA) ) {
							((BinarioExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(464);
						((BinarioExpContext)_localctx).right = expr(6);
						}
						break;
					case 3:
						{
						_localctx = new BinarioExpContext(new ExprContext(_parentctx, _parentState));
						((BinarioExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(465);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(466);
						((BinarioExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 503316480L) != 0)) ) {
							((BinarioExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(467);
						((BinarioExpContext)_localctx).right = expr(5);
						}
						break;
					case 4:
						{
						_localctx = new BinarioExpContext(new ExprContext(_parentctx, _parentState));
						((BinarioExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(468);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(469);
						((BinarioExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_IGUAL || _la==OP_DIFERENTE) ) {
							((BinarioExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(470);
						((BinarioExpContext)_localctx).right = expr(4);
						}
						break;
					case 5:
						{
						_localctx = new BinarioExpContext(new ExprContext(_parentctx, _parentState));
						((BinarioExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(471);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(472);
						((BinarioExpContext)_localctx).op = match(OP_AND);
						setState(473);
						((BinarioExpContext)_localctx).right = expr(3);
						}
						break;
					case 6:
						{
						_localctx = new BinarioExpContext(new ExprContext(_parentctx, _parentState));
						((BinarioExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(474);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(475);
						((BinarioExpContext)_localctx).op = match(OP_OR);
						setState(476);
						((BinarioExpContext)_localctx).right = expr(2);
						}
						break;
					}
					} 
				}
				setState(481);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,51,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	 
		public LiteralContext() { }
		public void copyFrom(LiteralContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringLiteralContext extends LiteralContext {
		public TerminalNode STRING_LITERAL() { return getToken(VGrammar.STRING_LITERAL, 0); }
		public StringLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BoolLiteralContext extends LiteralContext {
		public TerminalNode BOOL_LITERAL() { return getToken(VGrammar.BOOL_LITERAL, 0); }
		public BoolLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FloatLiteralContext extends LiteralContext {
		public TerminalNode FLOAT_LITERAL() { return getToken(VGrammar.FLOAT_LITERAL, 0); }
		public FloatLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NilLiteralContext extends LiteralContext {
		public TerminalNode NIL_LITERAL() { return getToken(VGrammar.NIL_LITERAL, 0); }
		public NilLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntLiteralContext extends LiteralContext {
		public TerminalNode INT_LITERAL() { return getToken(VGrammar.INT_LITERAL, 0); }
		public IntLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_literal);
		try {
			setState(487);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_LITERAL:
				_localctx = new IntLiteralContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(482);
				match(INT_LITERAL);
				}
				break;
			case FLOAT_LITERAL:
				_localctx = new FloatLiteralContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(483);
				match(FLOAT_LITERAL);
				}
				break;
			case STRING_LITERAL:
				_localctx = new StringLiteralContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(484);
				match(STRING_LITERAL);
				}
				break;
			case BOOL_LITERAL:
				_localctx = new BoolLiteralContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(485);
				match(BOOL_LITERAL);
				}
				break;
			case NIL_LITERAL:
				_localctx = new NilLiteralContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(486);
				match(NIL_LITERAL);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PatronIdContext extends ParserRuleContext {
		public PatronIdContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_patronId; }
	 
		public PatronIdContext() { }
		public void copyFrom(PatronIdContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ID_PatronContext extends PatronIdContext {
		public List<TerminalNode> ID() { return getTokens(VGrammar.ID); }
		public TerminalNode ID(int i) {
			return getToken(VGrammar.ID, i);
		}
		public List<TerminalNode> PUNTO() { return getTokens(VGrammar.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(VGrammar.PUNTO, i);
		}
		public ID_PatronContext(PatronIdContext ctx) { copyFrom(ctx); }
	}

	public final PatronIdContext patronId() throws RecognitionException {
		PatronIdContext _localctx = new PatronIdContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_patronId);
		try {
			int _alt;
			_localctx = new ID_PatronContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(489);
			match(ID);
			setState(494);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,53,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(490);
					match(PUNTO);
					setState(491);
					match(ID);
					}
					} 
				}
				setState(496);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,53,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 33:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 6);
		case 1:
			return precpred(_ctx, 5);
		case 2:
			return precpred(_ctx, 4);
		case 3:
			return precpred(_ctx, 3);
		case 4:
			return precpred(_ctx, 2);
		case 5:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u00018\u01f2\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0001\u0000\u0005\u0000J\b\u0000\n\u0000\f\u0000M\t\u0000\u0001"+
		"\u0000\u0003\u0000P\b\u0000\u0001\u0000\u0003\u0000S\b\u0000\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0005\u0001"+
		"[\b\u0001\n\u0001\f\u0001^\t\u0001\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002"+
		"n\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003"+
		"\u0086\b\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003"+
		"\u008c\b\u0003\u0003\u0003\u008e\b\u0003\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0003\u0004\u00a0\b\u0004\u0001\u0005\u0001\u0005\u0003"+
		"\u0005\u00a4\b\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u00a8\b\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0005\u0006\u00ad\b\u0006\n\u0006"+
		"\f\u0006\u00b0\t\u0006\u0001\u0006\u0003\u0006\u00b3\b\u0006\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007\u00b9\b\u0007\n\u0007"+
		"\f\u0007\u00bc\t\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b"+
		"\u0005\b\u00c3\b\b\n\b\f\b\u00c6\t\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0005\t\u00ce\b\t\n\t\f\t\u00d1\t\t\u0001\t\u0003\t\u00d4\b"+
		"\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0005\n\u00dc\b\n\n"+
		"\n\f\n\u00df\t\n\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u00e4"+
		"\b\u000b\n\u000b\f\u000b\u00e7\t\u000b\u0001\f\u0001\f\u0001\f\u0001\f"+
		"\u0005\f\u00ed\b\f\n\f\f\f\u00f0\t\f\u0001\f\u0001\f\u0001\r\u0001\r\u0003"+
		"\r\u00f6\b\r\u0001\r\u0001\r\u0003\r\u00fa\b\r\u0001\r\u0001\r\u0003\r"+
		"\u00fe\b\r\u0001\r\u0001\r\u0005\r\u0102\b\r\n\r\f\r\u0105\t\r\u0001\r"+
		"\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u010d"+
		"\b\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u0112\b\u000e"+
		"\u0001\u000e\u0001\u000e\u0005\u000e\u0116\b\u000e\n\u000e\f\u000e\u0119"+
		"\t\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0003\u0010\u0127\b\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u012b"+
		"\b\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u012f\b\u0010\n\u0010\f\u0010"+
		"\u0132\t\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0003\u0011\u0139\b\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0005\u0012\u0140\b\u0012\n\u0012\f\u0012\u0143\t\u0012\u0001"+
		"\u0012\u0003\u0012\u0146\b\u0012\u0001\u0013\u0003\u0013\u0149\b\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0005\u0014\u0151\b\u0014\n\u0014\f\u0014\u0154\t\u0014\u0001\u0014\u0003"+
		"\u0014\u0157\b\u0014\u0001\u0015\u0001\u0015\u0003\u0015\u015b\b\u0015"+
		"\u0001\u0015\u0001\u0015\u0003\u0015\u015f\b\u0015\u0001\u0016\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0005\u0016\u0165\b\u0016\n\u0016\f\u0016\u0168"+
		"\t\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0003\u0017\u0170\b\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0003"+
		"\u0018\u0175\b\u0018\u0001\u0018\u0001\u0018\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0005\u0019\u017c\b\u0019\n\u0019\f\u0019\u017f\t\u0019\u0001\u0019"+
		"\u0003\u0019\u0182\b\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0005\u001b\u018c\b\u001b"+
		"\n\u001b\f\u001b\u018f\t\u001b\u0003\u001b\u0191\b\u001b\u0001\u001b\u0003"+
		"\u001b\u0194\b\u001b\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0004\u001c\u019d\b\u001c\u000b\u001c\f"+
		"\u001c\u019e\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001e"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001f\u0001\u001f\u0001\u001f"+
		"\u0001\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u01af\b\u001f\u0001 "+
		"\u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0003 \u01b9\b \u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0003!\u01ca\b!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001!\u0001!\u0005!\u01de\b!\n!\f!\u01e1\t!\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0003\"\u01e8\b\"\u0001#\u0001#\u0001#\u0005"+
		"#\u01ed\b#\n#\f#\u01f0\t#\u0001#\u0000\u0001B$\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,."+
		"02468:<>@BDF\u0000\u0007\u0001\u0000\u0006\u0007\u0001\u0000\u0006\b\u0002"+
		"\u0000\u0013\u0013\u001f\u001f\u0001\u0000\u0014\u0016\u0001\u0000\u0012"+
		"\u0013\u0001\u0000\u0019\u001c\u0001\u0000\u0017\u0018\u0227\u0000K\u0001"+
		"\u0000\u0000\u0000\u0002T\u0001\u0000\u0000\u0000\u0004m\u0001\u0000\u0000"+
		"\u0000\u0006\u008d\u0001\u0000\u0000\u0000\b\u009f\u0001\u0000\u0000\u0000"+
		"\n\u00a7\u0001\u0000\u0000\u0000\f\u00a9\u0001\u0000\u0000\u0000\u000e"+
		"\u00b4\u0001\u0000\u0000\u0000\u0010\u00bf\u0001\u0000\u0000\u0000\u0012"+
		"\u00c9\u0001\u0000\u0000\u0000\u0014\u00d7\u0001\u0000\u0000\u0000\u0016"+
		"\u00e0\u0001\u0000\u0000\u0000\u0018\u00e8\u0001\u0000\u0000\u0000\u001a"+
		"\u00f3\u0001\u0000\u0000\u0000\u001c\u0108\u0001\u0000\u0000\u0000\u001e"+
		"\u011c\u0001\u0000\u0000\u0000 \u0122\u0001\u0000\u0000\u0000\"\u0135"+
		"\u0001\u0000\u0000\u0000$\u013c\u0001\u0000\u0000\u0000&\u0148\u0001\u0000"+
		"\u0000\u0000(\u014d\u0001\u0000\u0000\u0000*\u015a\u0001\u0000\u0000\u0000"+
		",\u0160\u0001\u0000\u0000\u0000.\u016b\u0001\u0000\u0000\u00000\u0171"+
		"\u0001\u0000\u0000\u00002\u0178\u0001\u0000\u0000\u00004\u0183\u0001\u0000"+
		"\u0000\u00006\u0187\u0001\u0000\u0000\u00008\u0197\u0001\u0000\u0000\u0000"+
		":\u01a0\u0001\u0000\u0000\u0000<\u01a4\u0001\u0000\u0000\u0000>\u01ae"+
		"\u0001\u0000\u0000\u0000@\u01b8\u0001\u0000\u0000\u0000B\u01c9\u0001\u0000"+
		"\u0000\u0000D\u01e7\u0001\u0000\u0000\u0000F\u01e9\u0001\u0000\u0000\u0000"+
		"HJ\u0003\u0004\u0002\u0000IH\u0001\u0000\u0000\u0000JM\u0001\u0000\u0000"+
		"\u0000KI\u0001\u0000\u0000\u0000KL\u0001\u0000\u0000\u0000LO\u0001\u0000"+
		"\u0000\u0000MK\u0001\u0000\u0000\u0000NP\u0003\u0002\u0001\u0000ON\u0001"+
		"\u0000\u0000\u0000OP\u0001\u0000\u0000\u0000PR\u0001\u0000\u0000\u0000"+
		"QS\u0005\u0000\u0000\u0001RQ\u0001\u0000\u0000\u0000RS\u0001\u0000\u0000"+
		"\u0000S\u0001\u0001\u0000\u0000\u0000TU\u0005!\u0000\u0000UV\u0005 \u0000"+
		"\u0000VW\u0005/\u0000\u0000WX\u00050\u0000\u0000X\\\u00051\u0000\u0000"+
		"Y[\u0003\u0004\u0002\u0000ZY\u0001\u0000\u0000\u0000[^\u0001\u0000\u0000"+
		"\u0000\\Z\u0001\u0000\u0000\u0000\\]\u0001\u0000\u0000\u0000]_\u0001\u0000"+
		"\u0000\u0000^\\\u0001\u0000\u0000\u0000_`\u00052\u0000\u0000`\u0003\u0001"+
		"\u0000\u0000\u0000an\u0003\u0006\u0003\u0000bn\u0003\b\u0004\u0000cn\u0003"+
		"\n\u0005\u0000dn\u0003\f\u0006\u0000en\u0003\u0012\t\u0000fn\u0003\u001a"+
		"\r\u0000gn\u0003\u0018\f\u0000hn\u0003\u001c\u000e\u0000in\u0003\"\u0011"+
		"\u0000jn\u0003<\u001e\u0000kn\u0003 \u0010\u0000ln\u0003,\u0016\u0000"+
		"ma\u0001\u0000\u0000\u0000mb\u0001\u0000\u0000\u0000mc\u0001\u0000\u0000"+
		"\u0000md\u0001\u0000\u0000\u0000me\u0001\u0000\u0000\u0000mf\u0001\u0000"+
		"\u0000\u0000mg\u0001\u0000\u0000\u0000mh\u0001\u0000\u0000\u0000mi\u0001"+
		"\u0000\u0000\u0000mj\u0001\u0000\u0000\u0000mk\u0001\u0000\u0000\u0000"+
		"ml\u0001\u0000\u0000\u0000n\u0005\u0001\u0000\u0000\u0000op\u0005\u0004"+
		"\u0000\u0000pq\u0005.\u0000\u0000qr\u0005\u0005\u0000\u0000r\u008e\u0003"+
		"B!\u0000st\u0005.\u0000\u0000tu\u0005\u0005\u0000\u0000u\u008e\u0003B"+
		"!\u0000vw\u0005\u0004\u0000\u0000wx\u0005.\u0000\u0000xy\u0003>\u001f"+
		"\u0000yz\u0005\b\u0000\u0000z{\u0003B!\u0000{\u008e\u0001\u0000\u0000"+
		"\u0000|}\u0005\u0004\u0000\u0000}~\u0005.\u0000\u0000~\u008e\u0003>\u001f"+
		"\u0000\u007f\u0080\u0005.\u0000\u0000\u0080\u0081\u0003>\u001f\u0000\u0081"+
		"\u0082\u0005\b\u0000\u0000\u0082\u0083\u0003B!\u0000\u0083\u008e\u0001"+
		"\u0000\u0000\u0000\u0084\u0086\u0005\u0004\u0000\u0000\u0085\u0084\u0001"+
		"\u0000\u0000\u0000\u0085\u0086\u0001\u0000\u0000\u0000\u0086\u0087\u0001"+
		"\u0000\u0000\u0000\u0087\u0088\u0005.\u0000\u0000\u0088\u0089\u0005\b"+
		"\u0000\u0000\u0089\u008b\u0003>\u001f\u0000\u008a\u008c\u0003B!\u0000"+
		"\u008b\u008a\u0001\u0000\u0000\u0000\u008b\u008c\u0001\u0000\u0000\u0000"+
		"\u008c\u008e\u0001\u0000\u0000\u0000\u008do\u0001\u0000\u0000\u0000\u008d"+
		"s\u0001\u0000\u0000\u0000\u008dv\u0001\u0000\u0000\u0000\u008d|\u0001"+
		"\u0000\u0000\u0000\u008d\u007f\u0001\u0000\u0000\u0000\u008d\u0085\u0001"+
		"\u0000\u0000\u0000\u008e\u0007\u0001\u0000\u0000\u0000\u008f\u0090\u0003"+
		"8\u001c\u0000\u0090\u0091\u0005\b\u0000\u0000\u0091\u0092\u0003B!\u0000"+
		"\u0092\u00a0\u0001\u0000\u0000\u0000\u0093\u0094\u0003F#\u0000\u0094\u0095"+
		"\u0005\b\u0000\u0000\u0095\u0096\u0003B!\u0000\u0096\u00a0\u0001\u0000"+
		"\u0000\u0000\u0097\u0098\u0003F#\u0000\u0098\u0099\u0007\u0000\u0000\u0000"+
		"\u0099\u009a\u0003B!\u0000\u009a\u00a0\u0001\u0000\u0000\u0000\u009b\u009c"+
		"\u00038\u001c\u0000\u009c\u009d\u0007\u0001\u0000\u0000\u009d\u009e\u0003"+
		"B!\u0000\u009e\u00a0\u0001\u0000\u0000\u0000\u009f\u008f\u0001\u0000\u0000"+
		"\u0000\u009f\u0093\u0001\u0000\u0000\u0000\u009f\u0097\u0001\u0000\u0000"+
		"\u0000\u009f\u009b\u0001\u0000\u0000\u0000\u00a0\t\u0001\u0000\u0000\u0000"+
		"\u00a1\u00a3\u0005,\u0000\u0000\u00a2\u00a4\u0003B!\u0000\u00a3\u00a2"+
		"\u0001\u0000\u0000\u0000\u00a3\u00a4\u0001\u0000\u0000\u0000\u00a4\u00a8"+
		"\u0001\u0000\u0000\u0000\u00a5\u00a8\u0005*\u0000\u0000\u00a6\u00a8\u0005"+
		"+\u0000\u0000\u00a7\u00a1\u0001\u0000\u0000\u0000\u00a7\u00a5\u0001\u0000"+
		"\u0000\u0000\u00a7\u00a6\u0001\u0000\u0000\u0000\u00a8\u000b\u0001\u0000"+
		"\u0000\u0000\u00a9\u00ae\u0003\u000e\u0007\u0000\u00aa\u00ab\u0005$\u0000"+
		"\u0000\u00ab\u00ad\u0003\u000e\u0007\u0000\u00ac\u00aa\u0001\u0000\u0000"+
		"\u0000\u00ad\u00b0\u0001\u0000\u0000\u0000\u00ae\u00ac\u0001\u0000\u0000"+
		"\u0000\u00ae\u00af\u0001\u0000\u0000\u0000\u00af\u00b2\u0001\u0000\u0000"+
		"\u0000\u00b0\u00ae\u0001\u0000\u0000\u0000\u00b1\u00b3\u0003\u0010\b\u0000"+
		"\u00b2\u00b1\u0001\u0000\u0000\u0000\u00b2\u00b3\u0001\u0000\u0000\u0000"+
		"\u00b3\r\u0001\u0000\u0000\u0000\u00b4\u00b5\u0005#\u0000\u0000\u00b5"+
		"\u00b6\u0003B!\u0000\u00b6\u00ba\u00051\u0000\u0000\u00b7\u00b9\u0003"+
		"\u0004\u0002\u0000\u00b8\u00b7\u0001\u0000\u0000\u0000\u00b9\u00bc\u0001"+
		"\u0000\u0000\u0000\u00ba\u00b8\u0001\u0000\u0000\u0000\u00ba\u00bb\u0001"+
		"\u0000\u0000\u0000\u00bb\u00bd\u0001\u0000\u0000\u0000\u00bc\u00ba\u0001"+
		"\u0000\u0000\u0000\u00bd\u00be\u00052\u0000\u0000\u00be\u000f\u0001\u0000"+
		"\u0000\u0000\u00bf\u00c0\u0005$\u0000\u0000\u00c0\u00c4\u00051\u0000\u0000"+
		"\u00c1\u00c3\u0003\u0004\u0002\u0000\u00c2\u00c1\u0001\u0000\u0000\u0000"+
		"\u00c3\u00c6\u0001\u0000\u0000\u0000\u00c4\u00c2\u0001\u0000\u0000\u0000"+
		"\u00c4\u00c5\u0001\u0000\u0000\u0000\u00c5\u00c7\u0001\u0000\u0000\u0000"+
		"\u00c6\u00c4\u0001\u0000\u0000\u0000\u00c7\u00c8\u00052\u0000\u0000\u00c8"+
		"\u0011\u0001\u0000\u0000\u0000\u00c9\u00ca\u0005%\u0000\u0000\u00ca\u00cb"+
		"\u0003B!\u0000\u00cb\u00cf\u00051\u0000\u0000\u00cc\u00ce\u0003\u0014"+
		"\n\u0000\u00cd\u00cc\u0001\u0000\u0000\u0000\u00ce\u00d1\u0001\u0000\u0000"+
		"\u0000\u00cf\u00cd\u0001\u0000\u0000\u0000\u00cf\u00d0\u0001\u0000\u0000"+
		"\u0000\u00d0\u00d3\u0001\u0000\u0000\u0000\u00d1\u00cf\u0001\u0000\u0000"+
		"\u0000\u00d2\u00d4\u0003\u0016\u000b\u0000\u00d3\u00d2\u0001\u0000\u0000"+
		"\u0000\u00d3\u00d4\u0001\u0000\u0000\u0000\u00d4\u00d5\u0001\u0000\u0000"+
		"\u0000\u00d5\u00d6\u00052\u0000\u0000\u00d6\u0013\u0001\u0000\u0000\u0000"+
		"\u00d7\u00d8\u0005&\u0000\u0000\u00d8\u00d9\u0003B!\u0000\u00d9\u00dd"+
		"\u00058\u0000\u0000\u00da\u00dc\u0003\u0004\u0002\u0000\u00db\u00da\u0001"+
		"\u0000\u0000\u0000\u00dc\u00df\u0001\u0000\u0000\u0000\u00dd\u00db\u0001"+
		"\u0000\u0000\u0000\u00dd\u00de\u0001\u0000\u0000\u0000\u00de\u0015\u0001"+
		"\u0000\u0000\u0000\u00df\u00dd\u0001\u0000\u0000\u0000\u00e0\u00e1\u0005"+
		"\'\u0000\u0000\u00e1\u00e5\u00058\u0000\u0000\u00e2\u00e4\u0003\u0004"+
		"\u0002\u0000\u00e3\u00e2\u0001\u0000\u0000\u0000\u00e4\u00e7\u0001\u0000"+
		"\u0000\u0000\u00e5\u00e3\u0001\u0000\u0000\u0000\u00e5\u00e6\u0001\u0000"+
		"\u0000\u0000\u00e6\u0017\u0001\u0000\u0000\u0000\u00e7\u00e5\u0001\u0000"+
		"\u0000\u0000\u00e8\u00e9\u0005(\u0000\u0000\u00e9\u00ea\u0003B!\u0000"+
		"\u00ea\u00ee\u00051\u0000\u0000\u00eb\u00ed\u0003\u0004\u0002\u0000\u00ec"+
		"\u00eb\u0001\u0000\u0000\u0000\u00ed\u00f0\u0001\u0000\u0000\u0000\u00ee"+
		"\u00ec\u0001\u0000\u0000\u0000\u00ee\u00ef\u0001\u0000\u0000\u0000\u00ef"+
		"\u00f1\u0001\u0000\u0000\u0000\u00f0\u00ee\u0001\u0000\u0000\u0000\u00f1"+
		"\u00f2\u00052\u0000\u0000\u00f2\u0019\u0001\u0000\u0000\u0000\u00f3\u00f5"+
		"\u0005(\u0000\u0000\u00f4\u00f6\u0003\b\u0004\u0000\u00f5\u00f4\u0001"+
		"\u0000\u0000\u0000\u00f5\u00f6\u0001\u0000\u0000\u0000\u00f6\u00f7\u0001"+
		"\u0000\u0000\u0000\u00f7\u00f9\u00057\u0000\u0000\u00f8\u00fa\u0003B!"+
		"\u0000\u00f9\u00f8\u0001\u0000\u0000\u0000\u00f9\u00fa\u0001\u0000\u0000"+
		"\u0000\u00fa\u00fb\u0001\u0000\u0000\u0000\u00fb\u00fd\u00057\u0000\u0000"+
		"\u00fc\u00fe\u0003\b\u0004\u0000\u00fd\u00fc\u0001\u0000\u0000\u0000\u00fd"+
		"\u00fe\u0001\u0000\u0000\u0000\u00fe\u00ff\u0001\u0000\u0000\u0000\u00ff"+
		"\u0103\u00051\u0000\u0000\u0100\u0102\u0003\u0004\u0002\u0000\u0101\u0100"+
		"\u0001\u0000\u0000\u0000\u0102\u0105\u0001\u0000\u0000\u0000\u0103\u0101"+
		"\u0001\u0000\u0000\u0000\u0103\u0104\u0001\u0000\u0000\u0000\u0104\u0106"+
		"\u0001\u0000\u0000\u0000\u0105\u0103\u0001\u0000\u0000\u0000\u0106\u0107"+
		"\u00052\u0000\u0000\u0107\u001b\u0001\u0000\u0000\u0000\u0108\u0109\u0005"+
		"(\u0000\u0000\u0109\u010c\u0005.\u0000\u0000\u010a\u010b\u00056\u0000"+
		"\u0000\u010b\u010d\u0005.\u0000\u0000\u010c\u010a\u0001\u0000\u0000\u0000"+
		"\u010c\u010d\u0001\u0000\u0000\u0000\u010d\u010e\u0001\u0000\u0000\u0000"+
		"\u010e\u0111\u0005-\u0000\u0000\u010f\u0112\u0003B!\u0000\u0110\u0112"+
		"\u0003\u001e\u000f\u0000\u0111\u010f\u0001\u0000\u0000\u0000\u0111\u0110"+
		"\u0001\u0000\u0000\u0000\u0112\u0113\u0001\u0000\u0000\u0000\u0113\u0117"+
		"\u00051\u0000\u0000\u0114\u0116\u0003\u0004\u0002\u0000\u0115\u0114\u0001"+
		"\u0000\u0000\u0000\u0116\u0119\u0001\u0000\u0000\u0000\u0117\u0115\u0001"+
		"\u0000\u0000\u0000\u0117\u0118\u0001\u0000\u0000\u0000\u0118\u011a\u0001"+
		"\u0000\u0000\u0000\u0119\u0117\u0001\u0000\u0000\u0000\u011a\u011b\u0005"+
		"2\u0000\u0000\u011b\u001d\u0001\u0000\u0000\u0000\u011c\u011d\u0003B!"+
		"\u0000\u011d\u011e\u00055\u0000\u0000\u011e\u011f\u00055\u0000\u0000\u011f"+
		"\u0120\u00055\u0000\u0000\u0120\u0121\u0003B!\u0000\u0121\u001f\u0001"+
		"\u0000\u0000\u0000\u0122\u0123\u0005!\u0000\u0000\u0123\u0124\u0005.\u0000"+
		"\u0000\u0124\u0126\u0005/\u0000\u0000\u0125\u0127\u0003$\u0012\u0000\u0126"+
		"\u0125\u0001\u0000\u0000\u0000\u0126\u0127\u0001\u0000\u0000\u0000\u0127"+
		"\u0128\u0001\u0000\u0000\u0000\u0128\u012a\u00050\u0000\u0000\u0129\u012b"+
		"\u0003>\u001f\u0000\u012a\u0129\u0001\u0000\u0000\u0000\u012a\u012b\u0001"+
		"\u0000\u0000\u0000\u012b\u012c\u0001\u0000\u0000\u0000\u012c\u0130\u0005"+
		"1\u0000\u0000\u012d\u012f\u0003\u0004\u0002\u0000\u012e\u012d\u0001\u0000"+
		"\u0000\u0000\u012f\u0132\u0001\u0000\u0000\u0000\u0130\u012e\u0001\u0000"+
		"\u0000\u0000\u0130\u0131\u0001\u0000\u0000\u0000\u0131\u0133\u0001\u0000"+
		"\u0000\u0000\u0132\u0130\u0001\u0000\u0000\u0000\u0133\u0134\u00052\u0000"+
		"\u0000\u0134!\u0001\u0000\u0000\u0000\u0135\u0136\u0003F#\u0000\u0136"+
		"\u0138\u0005/\u0000\u0000\u0137\u0139\u0003(\u0014\u0000\u0138\u0137\u0001"+
		"\u0000\u0000\u0000\u0138\u0139\u0001\u0000\u0000\u0000\u0139\u013a\u0001"+
		"\u0000\u0000\u0000\u013a\u013b\u00050\u0000\u0000\u013b#\u0001\u0000\u0000"+
		"\u0000\u013c\u0141\u0003&\u0013\u0000\u013d\u013e\u00056\u0000\u0000\u013e"+
		"\u0140\u0003&\u0013\u0000\u013f\u013d\u0001\u0000\u0000\u0000\u0140\u0143"+
		"\u0001\u0000\u0000\u0000\u0141\u013f\u0001\u0000\u0000\u0000\u0141\u0142"+
		"\u0001\u0000\u0000\u0000\u0142\u0145\u0001\u0000\u0000\u0000\u0143\u0141"+
		"\u0001\u0000\u0000\u0000\u0144\u0146\u00056\u0000\u0000\u0145\u0144\u0001"+
		"\u0000\u0000\u0000\u0145\u0146\u0001\u0000\u0000\u0000\u0146%\u0001\u0000"+
		"\u0000\u0000\u0147\u0149\u0005.\u0000\u0000\u0148\u0147\u0001\u0000\u0000"+
		"\u0000\u0148\u0149\u0001\u0000\u0000\u0000\u0149\u014a\u0001\u0000\u0000"+
		"\u0000\u014a\u014b\u0005.\u0000\u0000\u014b\u014c\u0003>\u001f\u0000\u014c"+
		"\'\u0001\u0000\u0000\u0000\u014d\u0152\u0003*\u0015\u0000\u014e\u014f"+
		"\u00056\u0000\u0000\u014f\u0151\u0003*\u0015\u0000\u0150\u014e\u0001\u0000"+
		"\u0000\u0000\u0151\u0154\u0001\u0000\u0000\u0000\u0152\u0150\u0001\u0000"+
		"\u0000\u0000\u0152\u0153\u0001\u0000\u0000\u0000\u0153\u0156\u0001\u0000"+
		"\u0000\u0000\u0154\u0152\u0001\u0000\u0000\u0000\u0155\u0157\u00056\u0000"+
		"\u0000\u0156\u0155\u0001\u0000\u0000\u0000\u0156\u0157\u0001\u0000\u0000"+
		"\u0000\u0157)\u0001\u0000\u0000\u0000\u0158\u0159\u0005.\u0000\u0000\u0159"+
		"\u015b\u00058\u0000\u0000\u015a\u0158\u0001\u0000\u0000\u0000\u015a\u015b"+
		"\u0001\u0000\u0000\u0000\u015b\u015e\u0001\u0000\u0000\u0000\u015c\u015f"+
		"\u0003F#\u0000\u015d\u015f\u0003B!\u0000\u015e\u015c\u0001\u0000\u0000"+
		"\u0000\u015e\u015d\u0001\u0000\u0000\u0000\u015f+\u0001\u0000\u0000\u0000"+
		"\u0160\u0161\u0005\"\u0000\u0000\u0161\u0162\u0005.\u0000\u0000\u0162"+
		"\u0166\u00051\u0000\u0000\u0163\u0165\u0003.\u0017\u0000\u0164\u0163\u0001"+
		"\u0000\u0000\u0000\u0165\u0168\u0001\u0000\u0000\u0000\u0166\u0164\u0001"+
		"\u0000\u0000\u0000\u0166\u0167\u0001\u0000\u0000\u0000\u0167\u0169\u0001"+
		"\u0000\u0000\u0000\u0168\u0166\u0001\u0000\u0000\u0000\u0169\u016a\u0005"+
		"2\u0000\u0000\u016a-\u0001\u0000\u0000\u0000\u016b\u016c\u0003>\u001f"+
		"\u0000\u016c\u016f\u0005.\u0000\u0000\u016d\u016e\u0005\b\u0000\u0000"+
		"\u016e\u0170\u0003B!\u0000\u016f\u016d\u0001\u0000\u0000\u0000\u016f\u0170"+
		"\u0001\u0000\u0000\u0000\u0170/\u0001\u0000\u0000\u0000\u0171\u0172\u0005"+
		".\u0000\u0000\u0172\u0174\u00051\u0000\u0000\u0173\u0175\u00032\u0019"+
		"\u0000\u0174\u0173\u0001\u0000\u0000\u0000\u0174\u0175\u0001\u0000\u0000"+
		"\u0000\u0175\u0176\u0001\u0000\u0000\u0000\u0176\u0177\u00052\u0000\u0000"+
		"\u01771\u0001\u0000\u0000\u0000\u0178\u017d\u00034\u001a\u0000\u0179\u017a"+
		"\u00056\u0000\u0000\u017a\u017c\u00034\u001a\u0000\u017b\u0179\u0001\u0000"+
		"\u0000\u0000\u017c\u017f\u0001\u0000\u0000\u0000\u017d\u017b\u0001\u0000"+
		"\u0000\u0000\u017d\u017e\u0001\u0000\u0000\u0000\u017e\u0181\u0001\u0000"+
		"\u0000\u0000\u017f\u017d\u0001\u0000\u0000\u0000\u0180\u0182\u00056\u0000"+
		"\u0000\u0181\u0180\u0001\u0000\u0000\u0000\u0181\u0182\u0001\u0000\u0000"+
		"\u0000\u01823\u0001\u0000\u0000\u0000\u0183\u0184\u0005.\u0000\u0000\u0184"+
		"\u0185\u00058\u0000\u0000\u0185\u0186\u0003B!\u0000\u01865\u0001\u0000"+
		"\u0000\u0000\u0187\u0190\u00051\u0000\u0000\u0188\u018d\u0003B!\u0000"+
		"\u0189\u018a\u00056\u0000\u0000\u018a\u018c\u0003B!\u0000\u018b\u0189"+
		"\u0001\u0000\u0000\u0000\u018c\u018f\u0001\u0000\u0000\u0000\u018d\u018b"+
		"\u0001\u0000\u0000\u0000\u018d\u018e\u0001\u0000\u0000\u0000\u018e\u0191"+
		"\u0001\u0000\u0000\u0000\u018f\u018d\u0001\u0000\u0000\u0000\u0190\u0188"+
		"\u0001\u0000\u0000\u0000\u0190\u0191\u0001\u0000\u0000\u0000\u0191\u0193"+
		"\u0001\u0000\u0000\u0000\u0192\u0194\u00056\u0000\u0000\u0193\u0192\u0001"+
		"\u0000\u0000\u0000\u0193\u0194\u0001\u0000\u0000\u0000\u0194\u0195\u0001"+
		"\u0000\u0000\u0000\u0195\u0196\u00052\u0000\u0000\u01967\u0001\u0000\u0000"+
		"\u0000\u0197\u019c\u0003F#\u0000\u0198\u0199\u00053\u0000\u0000\u0199"+
		"\u019a\u0003B!\u0000\u019a\u019b\u00054\u0000\u0000\u019b\u019d\u0001"+
		"\u0000\u0000\u0000\u019c\u0198\u0001\u0000\u0000\u0000\u019d\u019e\u0001"+
		"\u0000\u0000\u0000\u019e\u019c\u0001\u0000\u0000\u0000\u019e\u019f\u0001"+
		"\u0000\u0000\u0000\u019f9\u0001\u0000\u0000\u0000\u01a0\u01a1\u00038\u001c"+
		"\u0000\u01a1\u01a2\u00055\u0000\u0000\u01a2\u01a3\u0003F#\u0000\u01a3"+
		";\u0001\u0000\u0000\u0000\u01a4\u01a5\u00038\u001c\u0000\u01a5\u01a6\u0005"+
		"5\u0000\u0000\u01a6\u01a7\u0003\"\u0011\u0000\u01a7=\u0001\u0000\u0000"+
		"\u0000\u01a8\u01af\u0005.\u0000\u0000\u01a9\u01af\u0005\t\u0000\u0000"+
		"\u01aa\u01af\u0005\n\u0000\u0000\u01ab\u01af\u0005\u000b\u0000\u0000\u01ac"+
		"\u01af\u0005\f\u0000\u0000\u01ad\u01af\u0003@ \u0000\u01ae\u01a8\u0001"+
		"\u0000\u0000\u0000\u01ae\u01a9\u0001\u0000\u0000\u0000\u01ae\u01aa\u0001"+
		"\u0000\u0000\u0000\u01ae\u01ab\u0001\u0000\u0000\u0000\u01ae\u01ac\u0001"+
		"\u0000\u0000\u0000\u01ae\u01ad\u0001\u0000\u0000\u0000\u01af?\u0001\u0000"+
		"\u0000\u0000\u01b0\u01b1\u00053\u0000\u0000\u01b1\u01b2\u00054\u0000\u0000"+
		"\u01b2\u01b9\u0003>\u001f\u0000\u01b3\u01b4\u00053\u0000\u0000\u01b4\u01b5"+
		"\u00054\u0000\u0000\u01b5\u01b6\u00053\u0000\u0000\u01b6\u01b7\u00054"+
		"\u0000\u0000\u01b7\u01b9\u0003>\u001f\u0000\u01b8\u01b0\u0001\u0000\u0000"+
		"\u0000\u01b8\u01b3\u0001\u0000\u0000\u0000\u01b9A\u0001\u0000\u0000\u0000"+
		"\u01ba\u01bb\u0006!\uffff\uffff\u0000\u01bb\u01bc\u0005/\u0000\u0000\u01bc"+
		"\u01bd\u0003B!\u0000\u01bd\u01be\u00050\u0000\u0000\u01be\u01ca\u0001"+
		"\u0000\u0000\u0000\u01bf\u01ca\u0003\"\u0011\u0000\u01c0\u01ca\u0003F"+
		"#\u0000\u01c1\u01ca\u00038\u001c\u0000\u01c2\u01ca\u0003:\u001d\u0000"+
		"\u01c3\u01ca\u0003<\u001e\u0000\u01c4\u01ca\u0003D\"\u0000\u01c5\u01ca"+
		"\u00036\u001b\u0000\u01c6\u01ca\u00030\u0018\u0000\u01c7\u01c8\u0007\u0002"+
		"\u0000\u0000\u01c8\u01ca\u0003B!\u0007\u01c9\u01ba\u0001\u0000\u0000\u0000"+
		"\u01c9\u01bf\u0001\u0000\u0000\u0000\u01c9\u01c0\u0001\u0000\u0000\u0000"+
		"\u01c9\u01c1\u0001\u0000\u0000\u0000\u01c9\u01c2\u0001\u0000\u0000\u0000"+
		"\u01c9\u01c3\u0001\u0000\u0000\u0000\u01c9\u01c4\u0001\u0000\u0000\u0000"+
		"\u01c9\u01c5\u0001\u0000\u0000\u0000\u01c9\u01c6\u0001\u0000\u0000\u0000"+
		"\u01c9\u01c7\u0001\u0000\u0000\u0000\u01ca\u01df\u0001\u0000\u0000\u0000"+
		"\u01cb\u01cc\n\u0006\u0000\u0000\u01cc\u01cd\u0007\u0003\u0000\u0000\u01cd"+
		"\u01de\u0003B!\u0007\u01ce\u01cf\n\u0005\u0000\u0000\u01cf\u01d0\u0007"+
		"\u0004\u0000\u0000\u01d0\u01de\u0003B!\u0006\u01d1\u01d2\n\u0004\u0000"+
		"\u0000\u01d2\u01d3\u0007\u0005\u0000\u0000\u01d3\u01de\u0003B!\u0005\u01d4"+
		"\u01d5\n\u0003\u0000\u0000\u01d5\u01d6\u0007\u0006\u0000\u0000\u01d6\u01de"+
		"\u0003B!\u0004\u01d7\u01d8\n\u0002\u0000\u0000\u01d8\u01d9\u0005\u001d"+
		"\u0000\u0000\u01d9\u01de\u0003B!\u0003\u01da\u01db\n\u0001\u0000\u0000"+
		"\u01db\u01dc\u0005\u001e\u0000\u0000\u01dc\u01de\u0003B!\u0002\u01dd\u01cb"+
		"\u0001\u0000\u0000\u0000\u01dd\u01ce\u0001\u0000\u0000\u0000\u01dd\u01d1"+
		"\u0001\u0000\u0000\u0000\u01dd\u01d4\u0001\u0000\u0000\u0000\u01dd\u01d7"+
		"\u0001\u0000\u0000\u0000\u01dd\u01da\u0001\u0000\u0000\u0000\u01de\u01e1"+
		"\u0001\u0000\u0000\u0000\u01df\u01dd\u0001\u0000\u0000\u0000\u01df\u01e0"+
		"\u0001\u0000\u0000\u0000\u01e0C\u0001\u0000\u0000\u0000\u01e1\u01df\u0001"+
		"\u0000\u0000\u0000\u01e2\u01e8\u0005\r\u0000\u0000\u01e3\u01e8\u0005\u000e"+
		"\u0000\u0000\u01e4\u01e8\u0005\u000f\u0000\u0000\u01e5\u01e8\u0005\u0010"+
		"\u0000\u0000\u01e6\u01e8\u0005\u0011\u0000\u0000\u01e7\u01e2\u0001\u0000"+
		"\u0000\u0000\u01e7\u01e3\u0001\u0000\u0000\u0000\u01e7\u01e4\u0001\u0000"+
		"\u0000\u0000\u01e7\u01e5\u0001\u0000\u0000\u0000\u01e7\u01e6\u0001\u0000"+
		"\u0000\u0000\u01e8E\u0001\u0000\u0000\u0000\u01e9\u01ee\u0005.\u0000\u0000"+
		"\u01ea\u01eb\u00055\u0000\u0000\u01eb\u01ed\u0005.\u0000\u0000\u01ec\u01ea"+
		"\u0001\u0000\u0000\u0000\u01ed\u01f0\u0001\u0000\u0000\u0000\u01ee\u01ec"+
		"\u0001\u0000\u0000\u0000\u01ee\u01ef\u0001\u0000\u0000\u0000\u01efG\u0001"+
		"\u0000\u0000\u0000\u01f0\u01ee\u0001\u0000\u0000\u00006KOR\\m\u0085\u008b"+
		"\u008d\u009f\u00a3\u00a7\u00ae\u00b2\u00ba\u00c4\u00cf\u00d3\u00dd\u00e5"+
		"\u00ee\u00f5\u00f9\u00fd\u0103\u010c\u0111\u0117\u0126\u012a\u0130\u0138"+
		"\u0141\u0145\u0148\u0152\u0156\u015a\u015e\u0166\u016f\u0174\u017d\u0181"+
		"\u018d\u0190\u0193\u019e\u01ae\u01b8\u01c9\u01dd\u01df\u01e7\u01ee";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}