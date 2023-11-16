package dictionary

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	Definition string
}

func (e Entry) String() string {
	return e.Definition
}

type Dictionary struct {
	filePath string
}

func New(filePath string) *Dictionary {
	return &Dictionary{filePath: filePath}
}

func (d *Dictionary) Add(word string, definition string) error {
	file, err := os.OpenFile(d.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier : %s", err)
	}
	defer file.Close()

	entry := fmt.Sprintf("%s:%s\n", word, definition)
	if _, err := file.WriteString(entry); err != nil {
		return fmt.Errorf("erreur lors de l'écriture dans le fichier : %s", err)
	}

	return nil
}

func (d *Dictionary) Get(word string) (Entry, error) {
	file, err := os.Open(d.filePath)
	if err != nil {
		return Entry{}, fmt.Errorf("erreur lors de l'ouverture du fichier : %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) >= 2 && parts[0] == word {
			return Entry{Definition: parts[1]}, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return Entry{}, fmt.Errorf("erreur lors de la lecture du fichier : %s", err)
	}

	return Entry{}, fmt.Errorf("mot non trouvé dans le dictionnaire")
}

func (d *Dictionary) Remove(word string) error {
	file, err := os.Open(d.filePath)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier : %s", err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) >= 1 && parts[0] != word {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("erreur lors de la lecture du fichier : %s", err)
	}

	file, err = os.Create(d.filePath)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du fichier : %s", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

func (d *Dictionary) List() ([]string, map[string]Entry, error) {
	file, err := os.Open(d.filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("erreur lors de l'ouverture du fichier : %s", err)
	}
	defer file.Close()

	words := make([]string, 0)
	entries := make(map[string]Entry)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) >= 2 {
			word := parts[0]
			definition := parts[1]
			words = append(words, word)
			entries[word] = Entry{Definition: definition}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("erreur lors de la lecture du fichier : %s", err)
	}

	return words, entries, nil
}
