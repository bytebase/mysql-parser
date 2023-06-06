package parser

import "github.com/antlr4-go/antlr/v4"

type MySQLBaseLexer struct {
	*antlr.BaseLexer

	pendingTokens []antlr.Token
}

// NextToken implements antlr.TokenSource
func (l *MySQLBaseLexer) NextToken() antlr.Token {
	// First respond with pending tokens to the next token request, if there are any.
	if len(l.pendingTokens) != 0 {
		pending := l.pendingTokens[0]
		l.pendingTokens = l.pendingTokens[1:]
		return pending
	}

	// Let the main lexer class run the next token recognition.
	// This might create additional tokens again.
	next := l.BaseLexer.NextToken()
	if len(l.pendingTokens) != 0 {
		pending := l.pendingTokens[0]
		l.pendingTokens = l.pendingTokens[1:]
		l.pendingTokens = append(l.pendingTokens, next)
		return pending
	}
	return next
}

// EmitDot puts a DOT token onto the pending token list.
func (l *MySQLBaseLexer) EmitDot() {
	// TODO(rebelice): may use DOT_SYMBOL instead of l.GetType()
	dot := l.GetTokenFactory().Create(l.GetTokenSourceCharStreamPair(), l.GetType(), l.GetText(), antlr.TokenDefaultChannel, l.TokenStartCharIndex, l.TokenStartCharIndex, l.TokenStartLine, l.TokenStartColumn)
	l.pendingTokens = append(l.pendingTokens, dot)
	l.TokenStartCharIndex++
}
