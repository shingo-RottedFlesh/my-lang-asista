package token

type TokenKind string

type Token struct {
	Kind  TokenKind
	Value string
}

const (
	TokenKindLet               TokenKind = "定数定義キーワード"
	TokenKindPlus              TokenKind = "+"
	TokenKindMinus             TokenKind = "-"
	TokenKindAssignment        TokenKind = "="
	TokenKindSemicolon         TokenKind = ";"
	TokenKindRoundBracketStart TokenKind = "("
	TokenKindRoundBracketEnd   TokenKind = ")"
	TokenKindCurlyBracketStart TokenKind = "{"
	TokenKindCurlyBracketEnd   TokenKind = "}"
	TokenKindComma             TokenKind = ","
	TokenKindEndOfFile         TokenKind = "入力の終わり"
	TokenKindStringLiteral     TokenKind = "文字列リテラル"
	TokenKindNumberLiteral     TokenKind = "数値リテラル"
	TokenKindIdentifier        TokenKind = "識別子"
	TokenKindIllegal           TokenKind = "不正なトークン"
)
