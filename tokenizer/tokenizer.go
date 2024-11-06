package tokenizer

import (
	"strings"
	"unicode"
	"pati/patistructs"
)

// Tokenize function takes the content of a BASIC program and returns a list of tokens
func Tokenize(content string) []*patistructs.Token {
	var tokens []*patistructs.Token
	lines := strings.Split(content, "\n")
	lineNumber := 1

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			lineNumber++
			continue
		}

		pos := 0
		for pos < len(line) {
			ch := line[pos]

			// Skip whitespace
			if unicode.IsSpace(rune(ch)) {
				pos++
				continue
			}

			switch ch {
			case '+':
				tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_PLUS, lineNumber, pos, string(ch)))
				pos++
			case '-':
				tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_MINUS, lineNumber, pos, string(ch)))
				pos++
			case '*':
				tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_MULTIPLY, lineNumber, pos, string(ch)))
				pos++
			case '/':
				tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_DIVIDE, lineNumber, pos, string(ch)))
				pos++
			case '=':
				tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_EQUAL, lineNumber, pos, string(ch)))
				pos++
			case '(':
				tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_LEFT_PARENTHESIS, lineNumber, pos, string(ch)))
				pos++
			case ')':
				tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_RIGHT_PARENTHESIS, lineNumber, pos, string(ch)))
				pos++
			case '{':
				tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_LEFT_BRACE, lineNumber, pos, string(ch)))
				pos++
			case '}':
				tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_RIGHT_BRACE, lineNumber, pos, string(ch)))
				pos++
			case ',':
				tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_COMMA, lineNumber, pos, string(ch)))
				pos++
			case '"':
				// Parse string literal
				startPos := pos
				pos++
				for pos < len(line) && line[pos] != '"' {
					pos++
				}
				if pos < len(line) {
					pos++ // Include the closing quote
				}
				tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_STRING, lineNumber, startPos, line[startPos:pos]))
			default:
				if unicode.IsLetter(rune(ch)) {
					// Parse identifier or keyword
					startPos := pos
					for pos < len(line) && (unicode.IsLetter(rune(line[pos])) || unicode.IsDigit(rune(line[pos])) || line[pos] == '$') {
						pos++
					}
					word := line[startPos:pos]
					switch word {
					case "LET":
						tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_LET, lineNumber, startPos, word))
					case "IF":
						tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_IF, lineNumber, startPos, word))
					case "THEN":
						tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_THEN, lineNumber, startPos, word))
					case "RETURN":
						tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_RETURN, lineNumber, startPos, word))
					case "END":
						tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_END, lineNumber, startPos, word))
					case "PRINT":
						tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_PRINT, lineNumber, startPos, word))
					case "INPUT":
						tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_INPUT, lineNumber, startPos, word))
					case "PROC":
						tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_WORD, lineNumber, startPos, word))
					default:
						tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_VARIABLE, lineNumber, startPos, word))
					}
				} else if unicode.IsDigit(rune(ch)) {
					// Parse number
					startPos := pos
					for pos < len(line) && unicode.IsDigit(rune(line[pos])) {
						pos++
					}
					tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_NUMBER, lineNumber, startPos, line[startPos:pos]))
				} else {
					// Unknown token
					tokens = append(tokens, patistructs.NewTokenWithValues(patistructs.TOKEN_ILLEGAL, lineNumber, pos, string(ch)))
					pos++
				}
			}
		}
		lineNumber++
	}

	return tokens
}
