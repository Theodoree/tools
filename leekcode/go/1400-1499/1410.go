package _400_1499


func entityParser(text string) string {
	return strings.NewReplacer(`&quot;`, `"`, `&apos;`, `'`, `&gt;`, `>`, `&lt;`, `<`, `&frasl;`, `/`, `&amp;`, `&`).Replace(text)
}
