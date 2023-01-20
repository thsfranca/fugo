// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package core // Fugo
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseFugoListener is a complete listener for a parse tree produced by FugoParser.
type BaseFugoListener struct{}

var _ FugoListener = &BaseFugoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFugoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFugoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFugoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFugoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile_ is called when production file_ is entered.
func (s *BaseFugoListener) EnterFile_(ctx *File_Context) {}

// ExitFile_ is called when production file_ is exited.
func (s *BaseFugoListener) ExitFile_(ctx *File_Context) {}

// EnterForm is called when production form is entered.
func (s *BaseFugoListener) EnterForm(ctx *FormContext) {}

// ExitForm is called when production form is exited.
func (s *BaseFugoListener) ExitForm(ctx *FormContext) {}

// EnterForms is called when production forms is entered.
func (s *BaseFugoListener) EnterForms(ctx *FormsContext) {}

// ExitForms is called when production forms is exited.
func (s *BaseFugoListener) ExitForms(ctx *FormsContext) {}

// EnterList_ is called when production list_ is entered.
func (s *BaseFugoListener) EnterList_(ctx *List_Context) {}

// ExitList_ is called when production list_ is exited.
func (s *BaseFugoListener) ExitList_(ctx *List_Context) {}

// EnterVector is called when production vector is entered.
func (s *BaseFugoListener) EnterVector(ctx *VectorContext) {}

// ExitVector is called when production vector is exited.
func (s *BaseFugoListener) ExitVector(ctx *VectorContext) {}

// EnterMap_ is called when production map_ is entered.
func (s *BaseFugoListener) EnterMap_(ctx *Map_Context) {}

// ExitMap_ is called when production map_ is exited.
func (s *BaseFugoListener) ExitMap_(ctx *Map_Context) {}

// EnterSet_ is called when production set_ is entered.
func (s *BaseFugoListener) EnterSet_(ctx *Set_Context) {}

// ExitSet_ is called when production set_ is exited.
func (s *BaseFugoListener) ExitSet_(ctx *Set_Context) {}

// EnterReader_macro is called when production reader_macro is entered.
func (s *BaseFugoListener) EnterReader_macro(ctx *Reader_macroContext) {}

// ExitReader_macro is called when production reader_macro is exited.
func (s *BaseFugoListener) ExitReader_macro(ctx *Reader_macroContext) {}

// EnterQuote is called when production quote is entered.
func (s *BaseFugoListener) EnterQuote(ctx *QuoteContext) {}

// ExitQuote is called when production quote is exited.
func (s *BaseFugoListener) ExitQuote(ctx *QuoteContext) {}

// EnterBacktick is called when production backtick is entered.
func (s *BaseFugoListener) EnterBacktick(ctx *BacktickContext) {}

// ExitBacktick is called when production backtick is exited.
func (s *BaseFugoListener) ExitBacktick(ctx *BacktickContext) {}

// EnterUnquote is called when production unquote is entered.
func (s *BaseFugoListener) EnterUnquote(ctx *UnquoteContext) {}

// ExitUnquote is called when production unquote is exited.
func (s *BaseFugoListener) ExitUnquote(ctx *UnquoteContext) {}

// EnterUnquote_splicing is called when production unquote_splicing is entered.
func (s *BaseFugoListener) EnterUnquote_splicing(ctx *Unquote_splicingContext) {}

// ExitUnquote_splicing is called when production unquote_splicing is exited.
func (s *BaseFugoListener) ExitUnquote_splicing(ctx *Unquote_splicingContext) {}

// EnterTag is called when production tag is entered.
func (s *BaseFugoListener) EnterTag(ctx *TagContext) {}

// ExitTag is called when production tag is exited.
func (s *BaseFugoListener) ExitTag(ctx *TagContext) {}

// EnterDeref is called when production deref is entered.
func (s *BaseFugoListener) EnterDeref(ctx *DerefContext) {}

// ExitDeref is called when production deref is exited.
func (s *BaseFugoListener) ExitDeref(ctx *DerefContext) {}

// EnterGensym is called when production gensym is entered.
func (s *BaseFugoListener) EnterGensym(ctx *GensymContext) {}

// ExitGensym is called when production gensym is exited.
func (s *BaseFugoListener) ExitGensym(ctx *GensymContext) {}

// EnterLambda_ is called when production lambda_ is entered.
func (s *BaseFugoListener) EnterLambda_(ctx *Lambda_Context) {}

// ExitLambda_ is called when production lambda_ is exited.
func (s *BaseFugoListener) ExitLambda_(ctx *Lambda_Context) {}

// EnterMeta_data is called when production meta_data is entered.
func (s *BaseFugoListener) EnterMeta_data(ctx *Meta_dataContext) {}

// ExitMeta_data is called when production meta_data is exited.
func (s *BaseFugoListener) ExitMeta_data(ctx *Meta_dataContext) {}

// EnterVar_quote is called when production var_quote is entered.
func (s *BaseFugoListener) EnterVar_quote(ctx *Var_quoteContext) {}

// ExitVar_quote is called when production var_quote is exited.
func (s *BaseFugoListener) ExitVar_quote(ctx *Var_quoteContext) {}

// EnterHost_expr is called when production host_expr is entered.
func (s *BaseFugoListener) EnterHost_expr(ctx *Host_exprContext) {}

// ExitHost_expr is called when production host_expr is exited.
func (s *BaseFugoListener) ExitHost_expr(ctx *Host_exprContext) {}

// EnterDiscard is called when production discard is entered.
func (s *BaseFugoListener) EnterDiscard(ctx *DiscardContext) {}

// ExitDiscard is called when production discard is exited.
func (s *BaseFugoListener) ExitDiscard(ctx *DiscardContext) {}

// EnterDispatch is called when production dispatch is entered.
func (s *BaseFugoListener) EnterDispatch(ctx *DispatchContext) {}

// ExitDispatch is called when production dispatch is exited.
func (s *BaseFugoListener) ExitDispatch(ctx *DispatchContext) {}

// EnterRegex is called when production regex is entered.
func (s *BaseFugoListener) EnterRegex(ctx *RegexContext) {}

// ExitRegex is called when production regex is exited.
func (s *BaseFugoListener) ExitRegex(ctx *RegexContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseFugoListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseFugoListener) ExitLiteral(ctx *LiteralContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BaseFugoListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BaseFugoListener) ExitString_(ctx *String_Context) {}

// EnterHex_ is called when production hex_ is entered.
func (s *BaseFugoListener) EnterHex_(ctx *Hex_Context) {}

// ExitHex_ is called when production hex_ is exited.
func (s *BaseFugoListener) ExitHex_(ctx *Hex_Context) {}

// EnterBin_ is called when production bin_ is entered.
func (s *BaseFugoListener) EnterBin_(ctx *Bin_Context) {}

// ExitBin_ is called when production bin_ is exited.
func (s *BaseFugoListener) ExitBin_(ctx *Bin_Context) {}

// EnterBign is called when production bign is entered.
func (s *BaseFugoListener) EnterBign(ctx *BignContext) {}

// ExitBign is called when production bign is exited.
func (s *BaseFugoListener) ExitBign(ctx *BignContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseFugoListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseFugoListener) ExitNumber(ctx *NumberContext) {}

// EnterCharacter is called when production character is entered.
func (s *BaseFugoListener) EnterCharacter(ctx *CharacterContext) {}

// ExitCharacter is called when production character is exited.
func (s *BaseFugoListener) ExitCharacter(ctx *CharacterContext) {}

// EnterNamed_char is called when production named_char is entered.
func (s *BaseFugoListener) EnterNamed_char(ctx *Named_charContext) {}

// ExitNamed_char is called when production named_char is exited.
func (s *BaseFugoListener) ExitNamed_char(ctx *Named_charContext) {}

// EnterAny_char is called when production any_char is entered.
func (s *BaseFugoListener) EnterAny_char(ctx *Any_charContext) {}

// ExitAny_char is called when production any_char is exited.
func (s *BaseFugoListener) ExitAny_char(ctx *Any_charContext) {}

// EnterU_hex_quad is called when production u_hex_quad is entered.
func (s *BaseFugoListener) EnterU_hex_quad(ctx *U_hex_quadContext) {}

// ExitU_hex_quad is called when production u_hex_quad is exited.
func (s *BaseFugoListener) ExitU_hex_quad(ctx *U_hex_quadContext) {}

// EnterNil_ is called when production nil_ is entered.
func (s *BaseFugoListener) EnterNil_(ctx *Nil_Context) {}

// ExitNil_ is called when production nil_ is exited.
func (s *BaseFugoListener) ExitNil_(ctx *Nil_Context) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseFugoListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseFugoListener) ExitKeyword(ctx *KeywordContext) {}

// EnterSimple_keyword is called when production simple_keyword is entered.
func (s *BaseFugoListener) EnterSimple_keyword(ctx *Simple_keywordContext) {}

// ExitSimple_keyword is called when production simple_keyword is exited.
func (s *BaseFugoListener) ExitSimple_keyword(ctx *Simple_keywordContext) {}

// EnterMacro_keyword is called when production macro_keyword is entered.
func (s *BaseFugoListener) EnterMacro_keyword(ctx *Macro_keywordContext) {}

// ExitMacro_keyword is called when production macro_keyword is exited.
func (s *BaseFugoListener) ExitMacro_keyword(ctx *Macro_keywordContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseFugoListener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseFugoListener) ExitSymbol(ctx *SymbolContext) {}

// EnterSimple_sym is called when production simple_sym is entered.
func (s *BaseFugoListener) EnterSimple_sym(ctx *Simple_symContext) {}

// ExitSimple_sym is called when production simple_sym is exited.
func (s *BaseFugoListener) ExitSimple_sym(ctx *Simple_symContext) {}

// EnterNs_symbol is called when production ns_symbol is entered.
func (s *BaseFugoListener) EnterNs_symbol(ctx *Ns_symbolContext) {}

// ExitNs_symbol is called when production ns_symbol is exited.
func (s *BaseFugoListener) ExitNs_symbol(ctx *Ns_symbolContext) {}

// EnterParam_name is called when production param_name is entered.
func (s *BaseFugoListener) EnterParam_name(ctx *Param_nameContext) {}

// ExitParam_name is called when production param_name is exited.
func (s *BaseFugoListener) ExitParam_name(ctx *Param_nameContext) {}
