package main

import (
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"
)

func moveFilesToFolder(sourcePath string, folderName string, extensions []string, wg *sync.WaitGroup) {
	defer wg.Done()

	folderPath := filepath.Join(sourcePath, folderName)
	if err := os.MkdirAll(folderPath, 0755); err != nil {
		log.Printf("ОШИБКА: ошибка создания папки %s: %v\n", folderName, err)
		return
	}

	files, err := os.ReadDir(sourcePath)
	if err != nil {
		log.Printf("ОШИБКА: ошибка чтения папки: %v\n", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileExt := strings.ToLower(filepath.Ext(file.Name()))
		for _, ext := range extensions {
			if fileExt == ext {
				subFolder := filepath.Join(folderPath, strings.TrimPrefix(ext, "."))
				if err := os.MkdirAll(subFolder, 0755); err != nil {
					log.Printf("ОШИБКА: ошибка создания подпапки %s: %v\n", subFolder, err)
					return
				}

				oldPath := filepath.Join(sourcePath, file.Name())
				newPath := filepath.Join(subFolder, file.Name())

				if err := os.Rename(oldPath, newPath); err != nil {
					log.Printf("ОШИБКА: ошибка перемещения файла %s: %v\n", file.Name(), err)
					return
				}
				log.Printf("Перемещен файл: %s в %s\n", file.Name(), subFolder)
				break
			}
		}
	}
}

func moveOthersToFolder(sourcePath string, othersFolderName string, categorizedFolders []string, wg *sync.WaitGroup) {
	defer wg.Done()

	othersPath := filepath.Join(sourcePath, othersFolderName)
	if err := os.MkdirAll(othersPath, 0755); err != nil {
		log.Printf("ОШИБКА: ошибка создания папки %s: %v\n", othersFolderName, err)
		return
	}

	files, err := os.ReadDir(sourcePath)
	if err != nil {
		log.Printf("ОШИБКА: ошибка чтения папки: %v\n", err)
		return
	}

	for _, file := range files {
		fileName := file.Name()
		filePath := filepath.Join(sourcePath, fileName)

		isCategorizedFolder := slices.Contains(categorizedFolders, fileName)
		if isCategorizedFolder {
			continue
		}

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			log.Printf("ОШИБКА: ошибка проверки элемента %s: %v\n", fileName, err)
			continue
		}

		var newPath string
		var targetSubFolder string

		if file.IsDir() {
			targetSubFolder = filepath.Join(othersPath, "Folders")
			if err := os.MkdirAll(targetSubFolder, 0755); err != nil {
				log.Printf("ОШИБКА: ошибка создания подпапки %s: %v\n", targetSubFolder, err)
				continue
			}
			newPath = filepath.Join(targetSubFolder, fileName)
			log.Printf("Перемещена папка: %s в %s\n", fileName, targetSubFolder)
		} else {
			fileExt := strings.ToLower(filepath.Ext(fileName))
			if fileExt == "" {
				fileExt = ".no_ext"
			}
			if fileExt == ".ini" {
				log.Printf("Проверка: ini пропущен: %s\n", fileName)
				continue
			}
			targetSubFolder = filepath.Join(othersPath, strings.TrimPrefix(fileExt, "."))
			if err := os.MkdirAll(targetSubFolder, 0755); err != nil {
				log.Printf("ОШИБКА: ошибка создания подпапки %s: %v\n", targetSubFolder, err)
				continue
			}
			newPath = filepath.Join(targetSubFolder, fileName)
			log.Printf("Перемещен файл (остальное): %s в %s\n", fileName, targetSubFolder)
		}

		if err := os.Rename(filePath, newPath); err != nil {
			log.Printf("ОШИБКА: ошибка перемещения %s: %v\n", fileName, err)
			continue
		}
	}
}

func main() {
	folderPath := "C:\\Users\\vyantik\\Downloads"

	folders := []string{"Executable", "Images", "Documents", "Audio", "Video", "Archives", "Code", "Others"}
	extensions := make([][]string, 0)

	extensions = append(extensions, []string{".exe", ".msi", ".dmg", ".app", ".deb", ".rpm", ".apk"})

	extensions = append(extensions, []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".tiff", ".ico", ".svg", ".psd", ".ai"})

	extensions = append(extensions, []string{".pdf", ".doc", ".docx", ".txt", ".rtf", ".xls", ".xlsx", ".ppt", ".pptx", ".odt", ".ods", ".odp"})

	extensions = append(extensions, []string{".mp3", ".wav", ".m4a", ".ogg", ".flac", ".aac", ".wma", ".m4b"})

	extensions = append(extensions, []string{".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".mpeg", ".mpg", ".m4v", ".webm"})

	extensions = append(extensions, []string{".zip", ".rar", ".7z", ".tar", ".gz", ".bz2", ".iso"})

	extensions = append(extensions, []string{
		".py", ".js", ".ts", ".go", ".java", ".cpp", ".h", ".c",
		".css", ".html", ".php", ".sql", ".json", ".xml", ".yaml", ".yml",
		".toml", ".conf", ".env", ".log", ".md",
	})

	var wg sync.WaitGroup

	log.Println("Запущена параллельная сортировка файлов...")

	timeStart := time.Now()

	for i := range extensions {
		wg.Add(1)
		go moveFilesToFolder(folderPath, folders[i], extensions[i], &wg)
	}

	wg.Wait()

	wg.Add(1)
	go moveOthersToFolder(folderPath, folders[7], folders, &wg)

	wg.Wait()

	log.Println("Время выполнения:", time.Since(timeStart))

	log.Println("Все операции сортировки завершены успешно.")
}
