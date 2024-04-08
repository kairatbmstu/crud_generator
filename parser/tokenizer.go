package parser

type Tokenizer struct {
	i               int
	sourceCodeRunes []rune
}

func NewTokenizer(sourceCode string) *Tokenizer {
	sourceCodeRunes := []rune(sourceCode)
	var tokenizer = Tokenizer{
		sourceCodeRunes: sourceCodeRunes,
	}
	return &tokenizer
}

func (t *Tokenizer) tokenize() (*[]Token, error) {

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

func (t *Tokenizer) matchEntity() (*Token, bool) {
	return t.match("entity")
}

func (t *Tokenizer) matchRelationship() (*Token, bool) {
	return t.match("relationship")
}

func (t *Tokenizer) matchIdentifier() (*Token, bool) {
	return t.match("identifier")
}
func (t *Tokenizer) matchPaginate() (*Token, bool) {
	return t.match("paginate")
}

func (t *Tokenizer) match(matchWord string) (*Token, bool) {
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
