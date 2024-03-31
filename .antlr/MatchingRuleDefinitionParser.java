// Generated from /Users/praveen/my-workspace/pact-avro-plugin/MatchingRuleDefinition.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class MatchingRuleDefinitionParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, INTEGER_LITERAL=19, DECIMAL_LITERAL=20, LEFT_BRACKET=21, RIGHT_BRACKET=22, 
		STRING_LITERAL=23, BOOLEAN_LITERAL=24, COMMA=25, DOLLAR=26, WS=27;
	public static final int
		RULE_matchingDefinition = 0, RULE_matchingDefinitionExp = 1, RULE_matchingRule = 2, 
		RULE_primitiveValue = 3, RULE_string = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"matchingDefinition", "matchingDefinitionExp", "matchingRule", "primitiveValue", 
			"string"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'matching'", "'notEmpty'", "'eachKey'", "'eachValue'", "'equalTo'", 
			"'type'", "'number'", "'integer'", "'decimal'", "'datetime'", "'date'", 
			"'time'", "'regex'", "'include'", "'boolean'", "'semver'", "'contentType'", 
			"'null'", null, null, "'('", "')'", null, null, "','", "'$'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, "INTEGER_LITERAL", "DECIMAL_LITERAL", 
			"LEFT_BRACKET", "RIGHT_BRACKET", "STRING_LITERAL", "BOOLEAN_LITERAL", 
			"COMMA", "DOLLAR", "WS"
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
	public String getGrammarFileName() { return "MatchingRuleDefinition.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public MatchingRuleDefinitionParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MatchingDefinitionContext extends ParserRuleContext {
		public List<MatchingDefinitionExpContext> matchingDefinitionExp() {
			return getRuleContexts(MatchingDefinitionExpContext.class);
		}
		public MatchingDefinitionExpContext matchingDefinitionExp(int i) {
			return getRuleContext(MatchingDefinitionExpContext.class,i);
		}
		public TerminalNode EOF() { return getToken(MatchingRuleDefinitionParser.EOF, 0); }
		public List<TerminalNode> COMMA() { return getTokens(MatchingRuleDefinitionParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MatchingRuleDefinitionParser.COMMA, i);
		}
		public MatchingDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchingDefinition; }
	}

	public final MatchingDefinitionContext matchingDefinition() throws RecognitionException {
		MatchingDefinitionContext _localctx = new MatchingDefinitionContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_matchingDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(10);
			matchingDefinitionExp();
			setState(15);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(11);
				match(COMMA);
				setState(12);
				matchingDefinitionExp();
				}
				}
				setState(17);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(18);
			match(EOF);
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
	public static class MatchingDefinitionExpContext extends ParserRuleContext {
		public TerminalNode LEFT_BRACKET() { return getToken(MatchingRuleDefinitionParser.LEFT_BRACKET, 0); }
		public MatchingRuleContext matchingRule() {
			return getRuleContext(MatchingRuleContext.class,0);
		}
		public TerminalNode RIGHT_BRACKET() { return getToken(MatchingRuleDefinitionParser.RIGHT_BRACKET, 0); }
		public PrimitiveValueContext primitiveValue() {
			return getRuleContext(PrimitiveValueContext.class,0);
		}
		public MatchingDefinitionExpContext matchingDefinitionExp() {
			return getRuleContext(MatchingDefinitionExpContext.class,0);
		}
		public MatchingDefinitionExpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchingDefinitionExp; }
	}

	public final MatchingDefinitionExpContext matchingDefinitionExp() throws RecognitionException {
		MatchingDefinitionExpContext _localctx = new MatchingDefinitionExpContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_matchingDefinitionExp);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(40);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				{
				setState(20);
				match(T__0);
				setState(21);
				match(LEFT_BRACKET);
				setState(22);
				matchingRule();
				setState(23);
				match(RIGHT_BRACKET);
				}
				break;
			case T__1:
				{
				setState(25);
				match(T__1);
				setState(26);
				match(LEFT_BRACKET);
				setState(27);
				primitiveValue();
				setState(28);
				match(RIGHT_BRACKET);
				}
				break;
			case T__2:
				{
				setState(30);
				match(T__2);
				setState(31);
				match(LEFT_BRACKET);
				setState(32);
				matchingDefinitionExp();
				setState(33);
				match(RIGHT_BRACKET);
				}
				break;
			case T__3:
				{
				setState(35);
				match(T__3);
				setState(36);
				match(LEFT_BRACKET);
				setState(37);
				matchingDefinitionExp();
				setState(38);
				match(RIGHT_BRACKET);
				}
				break;
			default:
				throw new NoViableAltException(this);
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
	public static class MatchingRuleContext extends ParserRuleContext {
		public Token matcherType;
		public List<TerminalNode> COMMA() { return getTokens(MatchingRuleDefinitionParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(MatchingRuleDefinitionParser.COMMA, i);
		}
		public PrimitiveValueContext primitiveValue() {
			return getRuleContext(PrimitiveValueContext.class,0);
		}
		public TerminalNode DECIMAL_LITERAL() { return getToken(MatchingRuleDefinitionParser.DECIMAL_LITERAL, 0); }
		public TerminalNode INTEGER_LITERAL() { return getToken(MatchingRuleDefinitionParser.INTEGER_LITERAL, 0); }
		public List<StringContext> string() {
			return getRuleContexts(StringContext.class);
		}
		public StringContext string(int i) {
			return getRuleContext(StringContext.class,i);
		}
		public TerminalNode BOOLEAN_LITERAL() { return getToken(MatchingRuleDefinitionParser.BOOLEAN_LITERAL, 0); }
		public TerminalNode DOLLAR() { return getToken(MatchingRuleDefinitionParser.DOLLAR, 0); }
		public MatchingRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matchingRule; }
	}

	public final MatchingRuleContext matchingRule() throws RecognitionException {
		MatchingRuleContext _localctx = new MatchingRuleContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_matchingRule);
		int _la;
		try {
			setState(83);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__4:
			case T__5:
				enterOuterAlt(_localctx, 1);
				{
				{
				setState(42);
				_la = _input.LA(1);
				if ( !(_la==T__4 || _la==T__5) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(43);
				match(COMMA);
				setState(44);
				primitiveValue();
				}
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 2);
				{
				setState(45);
				match(T__6);
				setState(46);
				match(COMMA);
				setState(47);
				_la = _input.LA(1);
				if ( !(_la==INTEGER_LITERAL || _la==DECIMAL_LITERAL) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case T__7:
				enterOuterAlt(_localctx, 3);
				{
				setState(48);
				match(T__7);
				setState(49);
				match(COMMA);
				setState(50);
				match(INTEGER_LITERAL);
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 4);
				{
				setState(51);
				match(T__8);
				setState(52);
				match(COMMA);
				setState(53);
				match(DECIMAL_LITERAL);
				}
				break;
			case T__9:
			case T__10:
			case T__11:
				enterOuterAlt(_localctx, 5);
				{
				setState(54);
				((MatchingRuleContext)_localctx).matcherType = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 7168L) != 0)) ) {
					((MatchingRuleContext)_localctx).matcherType = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(55);
				match(COMMA);
				setState(56);
				string();
				setState(57);
				match(COMMA);
				setState(58);
				string();
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 6);
				{
				setState(60);
				match(T__12);
				setState(61);
				match(COMMA);
				setState(62);
				string();
				setState(63);
				match(COMMA);
				setState(64);
				string();
				}
				break;
			case T__13:
				enterOuterAlt(_localctx, 7);
				{
				setState(66);
				match(T__13);
				setState(67);
				match(COMMA);
				setState(68);
				string();
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 8);
				{
				setState(69);
				match(T__14);
				setState(70);
				match(COMMA);
				setState(71);
				match(BOOLEAN_LITERAL);
				}
				break;
			case T__15:
				enterOuterAlt(_localctx, 9);
				{
				setState(72);
				match(T__15);
				setState(73);
				match(COMMA);
				setState(74);
				string();
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 10);
				{
				setState(75);
				match(T__16);
				setState(76);
				match(COMMA);
				setState(77);
				string();
				setState(78);
				match(COMMA);
				setState(79);
				string();
				}
				break;
			case DOLLAR:
				enterOuterAlt(_localctx, 11);
				{
				setState(81);
				match(DOLLAR);
				setState(82);
				string();
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
	public static class PrimitiveValueContext extends ParserRuleContext {
		public StringContext string() {
			return getRuleContext(StringContext.class,0);
		}
		public TerminalNode DECIMAL_LITERAL() { return getToken(MatchingRuleDefinitionParser.DECIMAL_LITERAL, 0); }
		public TerminalNode INTEGER_LITERAL() { return getToken(MatchingRuleDefinitionParser.INTEGER_LITERAL, 0); }
		public TerminalNode BOOLEAN_LITERAL() { return getToken(MatchingRuleDefinitionParser.BOOLEAN_LITERAL, 0); }
		public PrimitiveValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitiveValue; }
	}

	public final PrimitiveValueContext primitiveValue() throws RecognitionException {
		PrimitiveValueContext _localctx = new PrimitiveValueContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_primitiveValue);
		try {
			setState(89);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__17:
			case STRING_LITERAL:
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				string();
				}
				break;
			case DECIMAL_LITERAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(86);
				match(DECIMAL_LITERAL);
				}
				break;
			case INTEGER_LITERAL:
				enterOuterAlt(_localctx, 3);
				{
				setState(87);
				match(INTEGER_LITERAL);
				}
				break;
			case BOOLEAN_LITERAL:
				enterOuterAlt(_localctx, 4);
				{
				setState(88);
				match(BOOLEAN_LITERAL);
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
	public static class StringContext extends ParserRuleContext {
		public TerminalNode STRING_LITERAL() { return getToken(MatchingRuleDefinitionParser.STRING_LITERAL, 0); }
		public StringContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string; }
	}

	public final StringContext string() throws RecognitionException {
		StringContext _localctx = new StringContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_string);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(91);
			_la = _input.LA(1);
			if ( !(_la==T__17 || _la==STRING_LITERAL) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	public static final String _serializedATN =
		"\u0004\u0001\u001b^\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0005\u0000\u000e\b\u0000\n\u0000\f\u0000"+
		"\u0011\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001"+
		")\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0003\u0002T\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0003\u0003Z\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0000\u0000"+
		"\u0005\u0000\u0002\u0004\u0006\b\u0000\u0004\u0001\u0000\u0005\u0006\u0001"+
		"\u0000\u0013\u0014\u0001\u0000\n\f\u0002\u0000\u0012\u0012\u0017\u0017"+
		"i\u0000\n\u0001\u0000\u0000\u0000\u0002(\u0001\u0000\u0000\u0000\u0004"+
		"S\u0001\u0000\u0000\u0000\u0006Y\u0001\u0000\u0000\u0000\b[\u0001\u0000"+
		"\u0000\u0000\n\u000f\u0003\u0002\u0001\u0000\u000b\f\u0005\u0019\u0000"+
		"\u0000\f\u000e\u0003\u0002\u0001\u0000\r\u000b\u0001\u0000\u0000\u0000"+
		"\u000e\u0011\u0001\u0000\u0000\u0000\u000f\r\u0001\u0000\u0000\u0000\u000f"+
		"\u0010\u0001\u0000\u0000\u0000\u0010\u0012\u0001\u0000\u0000\u0000\u0011"+
		"\u000f\u0001\u0000\u0000\u0000\u0012\u0013\u0005\u0000\u0000\u0001\u0013"+
		"\u0001\u0001\u0000\u0000\u0000\u0014\u0015\u0005\u0001\u0000\u0000\u0015"+
		"\u0016\u0005\u0015\u0000\u0000\u0016\u0017\u0003\u0004\u0002\u0000\u0017"+
		"\u0018\u0005\u0016\u0000\u0000\u0018)\u0001\u0000\u0000\u0000\u0019\u001a"+
		"\u0005\u0002\u0000\u0000\u001a\u001b\u0005\u0015\u0000\u0000\u001b\u001c"+
		"\u0003\u0006\u0003\u0000\u001c\u001d\u0005\u0016\u0000\u0000\u001d)\u0001"+
		"\u0000\u0000\u0000\u001e\u001f\u0005\u0003\u0000\u0000\u001f \u0005\u0015"+
		"\u0000\u0000 !\u0003\u0002\u0001\u0000!\"\u0005\u0016\u0000\u0000\")\u0001"+
		"\u0000\u0000\u0000#$\u0005\u0004\u0000\u0000$%\u0005\u0015\u0000\u0000"+
		"%&\u0003\u0002\u0001\u0000&\'\u0005\u0016\u0000\u0000\')\u0001\u0000\u0000"+
		"\u0000(\u0014\u0001\u0000\u0000\u0000(\u0019\u0001\u0000\u0000\u0000("+
		"\u001e\u0001\u0000\u0000\u0000(#\u0001\u0000\u0000\u0000)\u0003\u0001"+
		"\u0000\u0000\u0000*+\u0007\u0000\u0000\u0000+,\u0005\u0019\u0000\u0000"+
		",T\u0003\u0006\u0003\u0000-.\u0005\u0007\u0000\u0000./\u0005\u0019\u0000"+
		"\u0000/T\u0007\u0001\u0000\u000001\u0005\b\u0000\u000012\u0005\u0019\u0000"+
		"\u00002T\u0005\u0013\u0000\u000034\u0005\t\u0000\u000045\u0005\u0019\u0000"+
		"\u00005T\u0005\u0014\u0000\u000067\u0007\u0002\u0000\u000078\u0005\u0019"+
		"\u0000\u000089\u0003\b\u0004\u00009:\u0005\u0019\u0000\u0000:;\u0003\b"+
		"\u0004\u0000;T\u0001\u0000\u0000\u0000<=\u0005\r\u0000\u0000=>\u0005\u0019"+
		"\u0000\u0000>?\u0003\b\u0004\u0000?@\u0005\u0019\u0000\u0000@A\u0003\b"+
		"\u0004\u0000AT\u0001\u0000\u0000\u0000BC\u0005\u000e\u0000\u0000CD\u0005"+
		"\u0019\u0000\u0000DT\u0003\b\u0004\u0000EF\u0005\u000f\u0000\u0000FG\u0005"+
		"\u0019\u0000\u0000GT\u0005\u0018\u0000\u0000HI\u0005\u0010\u0000\u0000"+
		"IJ\u0005\u0019\u0000\u0000JT\u0003\b\u0004\u0000KL\u0005\u0011\u0000\u0000"+
		"LM\u0005\u0019\u0000\u0000MN\u0003\b\u0004\u0000NO\u0005\u0019\u0000\u0000"+
		"OP\u0003\b\u0004\u0000PT\u0001\u0000\u0000\u0000QR\u0005\u001a\u0000\u0000"+
		"RT\u0003\b\u0004\u0000S*\u0001\u0000\u0000\u0000S-\u0001\u0000\u0000\u0000"+
		"S0\u0001\u0000\u0000\u0000S3\u0001\u0000\u0000\u0000S6\u0001\u0000\u0000"+
		"\u0000S<\u0001\u0000\u0000\u0000SB\u0001\u0000\u0000\u0000SE\u0001\u0000"+
		"\u0000\u0000SH\u0001\u0000\u0000\u0000SK\u0001\u0000\u0000\u0000SQ\u0001"+
		"\u0000\u0000\u0000T\u0005\u0001\u0000\u0000\u0000UZ\u0003\b\u0004\u0000"+
		"VZ\u0005\u0014\u0000\u0000WZ\u0005\u0013\u0000\u0000XZ\u0005\u0018\u0000"+
		"\u0000YU\u0001\u0000\u0000\u0000YV\u0001\u0000\u0000\u0000YW\u0001\u0000"+
		"\u0000\u0000YX\u0001\u0000\u0000\u0000Z\u0007\u0001\u0000\u0000\u0000"+
		"[\\\u0007\u0003\u0000\u0000\\\t\u0001\u0000\u0000\u0000\u0004\u000f(S"+
		"Y";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}