package cmd

import (
	"bytes"
	"regexp"
	"strings"
	"testing"
)

func removeEmojis(s string) string {
	emojiRegex := regexp.MustCompile(`[\x{1F600}-\x{1F6FF}\x{2700}-\x{27BF}\x{1F300}-\x{1F5FF}\x{1F1E0}-\x{1F1FF}]`)
	return emojiRegex.ReplaceAllString(s, "")
}

func TestGreetCommand(t *testing.T) {
	output := new(bytes.Buffer)
	rootCmd.SetOut(output)

	rootCmd.SetArgs([]string{"hello", "greet", "Robert"})

	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Erro ao executar comando: %v", err)
	}

	result := removeEmojis(output.String())
	t.Log("Saída do comando:", result)
	if !strings.Contains(result, "Olá, Robert") {
		t.Errorf("Esperado 'Olá, Robert' na saída, mas obtido: %s", result)
	}
}
