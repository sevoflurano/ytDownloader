package main

import (
	"fmt"
	"os"
	"os/exec"
)

func menu() uint8 {
	var escolher uint8
	for {
		fmt.Println("1. Vídeo completo\n2. Só audio (mp3)\nEscolha: ")
		if _, err := fmt.Scan(&escolher); err != nil {
			fmt.Println("Por favor, digite apenas números!")
			var descartar string
			fmt.Scanln(&descartar)
			continue
		}
		if escolher == 1 || escolher == 2 {
			return escolher
		}
		fmt.Println("Opção inválida! Escolha 1 ou 2.")
	}
}

func validar_pasta(pasta string) error {
	info, err := os.Stat(pasta)
	if err != nil {
		return fmt.Errorf("Pasta não existe: %v", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s Não é uma pasta", pasta)
	}
	return nil
}

func menuAudio() uint8 {
	var escolha uint8
	for {
		fmt.Println("1. MP3 (mais compatível)\n2. OPUS/m4a (melhor qualidade)\nEscolha: ")
		if _, err := fmt.Scan(&escolha); err != nil {
			fmt.Println("Por favor, digite apenas números!")
			var descartar string
			fmt.Scanln(&descartar)
			continue
		}
		if escolha == 1 || escolha == 2 {
			return escolha
		}
		fmt.Println("Opção inválida! Escolha 1 ou 2.")
	}
}

func downloader(escolher uint8) {
	url := os.Args[1]
	pasta := os.Args[2]
	fmt.Println("baixando:", url)
	fmt.Println("pasta:", pasta)

	var cmd *exec.Cmd
	switch escolher {
	case 1:
		cmd = exec.Command("yt-dlp", "-o", pasta+"/%(title)s.%(ext)s", url)
	case 2:
		switch menuAudio() {
		case 1:
			cmd = exec.Command("yt-dlp", "-x", "--audio-format", "mp3", "-o", pasta+"/%(title)s.%(ext)s", url)
		case 2:
			cmd = exec.Command("yt-dlp", "-x", "-f", "bestaudio", "-o", pasta+"/%(title)s.%(ext)s", url)
		}
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("erro:", err)
		return
	}
	fmt.Println("concluído!")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: ./downloader <URL> <PASTA>")
		os.Exit(1)
	}
	escolha := menu()
	downloader(escolha)
}
