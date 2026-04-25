package reset

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func resetWindows(h string, products []string) {
	cleanDir(filepath.Join(h, "AppData", "Roaming", "JetBrains"), products)
	cleanDir(filepath.Join(h, "AppData", "Local", "JetBrains"), products)
	cleanRegistry(products)
}

func cleanRegistry(products []string) {
	const basePath = `HKCU\Software\JavaSoft\Prefs\jetbrains`

	for _, product := range products {
		keyPath := basePath + `\` + strings.ToLower(product)
		err := exec.Command("reg", "delete", keyPath, "/f").Run()
		if err == nil {
			fmt.Printf("🧹 Registry: cleaned %s\n", product)
		}
		// Ключ не существует — молча пропускаем
	}
}
