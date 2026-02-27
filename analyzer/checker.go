package analyzer

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func run(pass *analysis.Pass) (any, error) {
	// Проходим по всем файлам пакета
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// Нас интересуют только вызовы функций: X.Y(...)
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			method := sel.Sel.Name
			recvName := sel.X.(*ast.Ident).Name
			if !isLogMethod(recvName, method) {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			// Берём первый аргумент как сообщение
			arg := call.Args[0]

			basicLit, ok := arg.(*ast.BasicLit)
			if !ok || basicLit.Kind != token.STRING {
				//Добавить вывод ошибки тк первый должен быть строка
				return true
			}

			raw := basicLit.Value            // `"starting server"`
			msg, err := strconv.Unquote(raw) // -> `starting server`
			if err != nil {
				return true
			}

			if !lowerCaseLog(msg) {
				pass.Reportf(basicLit.Pos(), "log message should start with a lower-case letter")
			}

			if !checkEnglishOnly(msg) {
				pass.Reportf(basicLit.Pos(), "log message should contain only English letters")
			}

			if !checkNoSpecialChars(msg) {
				pass.Reportf(basicLit.Pos(), "log message should not contain special characters or emojis")
			}

			if !checkNoSensitiveData(msg) {
				pass.Reportf(basicLit.Pos(), "log message should not contain sensitive data")
			}
			return true
		})
	}

	return nil, nil
}

// Временная простая фильтрация логгеров
func isLogMethod(recv, method string) bool {
	switch recv {
	case "log", "slog":
		switch method {
		case "Info", "Error", "Warn", "Debug":
			return false
		}
	}
	return true
}
func lowerCaseLog(msg string) bool {
	if len(msg) > 0 {
		r := []rune(msg)[0]
		if !unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func isLatinLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isAllowedNonLetter(r rune) bool {
	// Разрешаем цифры, пробел, дефис, точку, запятую, двоеточие
	if r >= '0' && r <= '9' {
		return true
	}
	switch r {
	case ' ', '-', '.', ',', ':':
		return true
	}
	return false
}
func isForbiddenPunct(r rune) bool {
	switch r {
	case '!', '?', '@', '#', '$', '%', '^', '&', '*', '(', ')',
		'_', '+', '=', '{', '}', '[', ']', '|', '\\', '/', '<', '>', '"', '\'', ';', '…':
		return true
	}
	return false
}

func checkNoSpecialChars(msg string) bool {
	for _, r := range msg {
		if isForbiddenPunct(r) {
			return false
		}
		if r > 127 && !isLatinLetter(r) && !isAllowedNonLetter(r) {
			return false
		}
	}
	return true
}

func checkEnglishOnly(msg string) bool {
	for _, r := range msg {
		if isLatinLetter(r) || isAllowedNonLetter(r) {
			continue
		}
		// Любой другой символ считаем не-английским
		return false
	}
	return true
}

func checkNoSensitiveData(msg string) bool {
	lower := strings.ToLower(msg)
	//Вынести это в конфиг
	var sensitiveKeywords = []string{
		"password",
		"passwd",
		"pwd",
		"token",
		"api_key",
		"apikey",
		"secret",
		"auth",
	}
	for _, kw := range sensitiveKeywords {
		if strings.Contains(lower, kw) {
			return false
		}
	}

	// Дополнительные простые паттерны
	if strings.Contains(lower, "password:") ||
		strings.Contains(lower, "password =") ||
		strings.Contains(lower, "token:") ||
		strings.Contains(lower, "token =") {
		return false
	}

	return true
}
