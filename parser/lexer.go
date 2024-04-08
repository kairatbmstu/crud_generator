package parser

type Lexer struct {
	i               int
	sourceCodeRunes []rune
}

func NewLexer(sourceCode string) *Lexer {
	sourceCodeRunes := []rune(sourceCode)
	var tokenizer = Lexer{
		sourceCodeRunes: sourceCodeRunes,
	}
	return &tokenizer
}

func (t *Lexer) Tokenize() (*[]Token, error) {

	for t.i = 0; t.i < len(t.sourceCodeRunes); t.i++ {
		switch t.sourceCodeRunes[t.i] {
		case ' ':
			break
		case '\\':
			break
		case 'r':
			t.matchRelationship()
			break
		case 'p':
			t.matchPaginate()
			break
		case 'e':
			t.matchEntity()
			break
		default:
			break
		}
	}

	return nil, nil
}

func (t *Lexer) matchEntity() (*Token, bool) {
	return t.match("entity")
}

func (t *Lexer) matchRelationship() (*Token, bool) {
	return t.match("relationship")
}

func (t *Lexer) matchIdentifier() (*Token, bool) {
	return t.match("identifier")
}
func (t *Lexer) matchPaginate() (*Token, bool) {
	return t.match("paginate")
}

func (t *Lexer) matchWith() (*Token, bool) {
	return t.match("with")
}

func (t *Lexer) matchPagination() (*Token, bool) {
	return t.match("pagination")
}

func (t *Lexer) matchUUID() (*Token, bool) {
	return t.match("pagination")
}

func (t *Lexer) matchOneToOne() (*Token, bool) {
	return t.match("pagination")
}

func (t *Lexer) matchOneToMany() (*Token, bool) {
	return t.match("pagination")
}

func (t *Lexer) matchManyToOne() (*Token, bool) {
	return t.match("pagination")
}

func (t *Lexer) matchManyToMany() (*Token, bool) {
	return t.match("pagination")
}

func (t *Lexer) match(matchWord string) (*Token, bool) {
	var matchWordRune = []rune(matchWord)
	var token = Token{}
	var l int = t.i
	var r int = 0

	for r < len(matchWordRune) {
		if t.sourceCodeRunes[l] == matchWordRune[r] {
			l++
			r++
		} else {
			return nil, false
		}
	}

	return &token, true
}
