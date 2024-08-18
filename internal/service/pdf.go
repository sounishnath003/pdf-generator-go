package service

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

const LOCAL_GOOG_CHROME_PATH = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"

// Attrib: https://github.com/bitcrowd/chromic_pdf/blob/main/lib/chromic_pdf/pdf/chrome_runner.ex
var chromeFlags = []string{
	"--headless",
	"--disable-accelerated-2d-canvas",
	"--disable-gpu",
	"--allow-pre-commit-input",
	"--disable-background-networking",
	"--disable-background-timer-throttling",
	"--disable-backgrounding-occluded-windows",
	"--disable-breakpad",
	"--disable-client-side-phishing-detection",
	"--disable-component-extensions-with-background-pages",
	"--disable-component-update",
	"--disable-default-apps",
	"--disable-extensions",
	"--disable-features=Translate,BackForwardCache,AcceptCHFrame,MediaRouter,OptimizationHints",
	"--disable-hang-monitor",
	"--disable-ipc-flooding-protection",
	"--disable-popup-blocking",
	"--disable-prompt-on-repost",
	"--disable-renderer-backgrounding",
	"--disable-sync",
	"--enable-automation",
	"--enable-features=NetworkServiceInProcess2",
	"--export-tagged-pdf",
	"--force-color-profile=srgb",
	"--hide-scrollbars",
	"--metrics-recording-only",
	"--no-default-browser-check",
	"--no-first-run",
	"--no-service-autorun",
	"--password-store=basic",
	"--use-mock-keychain",
	"--no-sandbox",
}

// GeneratePdfFromSource helps to generate the PDF and Html file as per the requiremenet and bytes of data.
func GeneratePdfFromSource(source []byte) (string, error) {
	// tmpDir := os.TempDir()
	tmpDir, err := os.Getwd()
	if err := os.MkdirAll(tmpDir+"temp", 0o755); err != nil {
		return "", err
	}
	tmpFileId := generateIdForFileName()

	htmlFile := path.Join(tmpDir, "temp", tmpFileId+".html")
	pdfFile := path.Join(tmpDir, "temp", tmpFileId+".pdf")
	publicPdfFile := fmt.Sprintf("/export/%s.pdf", tmpFileId)

	os.WriteFile(htmlFile, source, 0o644)
	log.Printf("generated.htmlFile.name: %s", htmlFile)
	log.Printf("generated.pdfFile.name: %s", pdfFile)

	// Setting up the necessary runtime args for the google chrome
	args := append(chromeFlags, fmt.Sprintf(`--print-to-pdf=%s`, pdfFile), htmlFile)
	log.Printf("runtime.chrome.flags: %s\n", args)

	outBuf := bytes.NewBuffer([]byte{})

	cmd := exec.Command(LOCAL_GOOG_CHROME_PATH, args...)
	cmd.Stderr = outBuf
	cmd.Stdout = outBuf

	err = cmd.Run()
	if err != nil {
		log.Printf("an error occured %v\n", err)
		return "", errors.New(outBuf.String())
	}

	return publicPdfFile, nil
}

func generateIdForFileName() string {
	bytes := make([]byte, 6)
	rand.Read(bytes)
	return fmt.Sprintf("%#x", bytes)[2:]
}
