package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
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

			t := pass.TypesInfo.TypeOf(sel.X)
			if !isLogReceiverType(t) {
				return true
			}

			method := sel.Sel.Name
			if !isLogMethod(method) {
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

func isLogReceiverType(t types.Type) bool {
	if t == nil {
		return false
	}

	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}

	named, ok := t.(*types.Named)
	if !ok {
		return false
	}

	obj := named.Obj()
	if obj == nil {
		return false
	}
	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}

	pkgPath := pkg.Path()
	switch pkgPath {
	case "log/slog":
		return true
	case "go.uber.org/zap":
		return true
	default:
		return false
	}
}

func isLogMethod(method string) bool {
	switch method {
	// Базовые уровни, общие для slog и zap
	case "Debug", "Info", "Warn", "Error",
		// zap.Logger дополнительные уровни
		"DPanic", "Panic", "Fatal",
		// zap.SugaredLogger printf-методы
		"Debugf", "Infof", "Warnf", "Errorf", "DPanicf", "Panicf", "Fatalf",
		// zap.SugaredLogger *w-методы
		"Debugw", "Infow", "Warnw", "Errorw", "DPanicw", "Panicw", "Fatalw":
		return true
	default:
		return false
	}
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
