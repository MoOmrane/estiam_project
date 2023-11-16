package dictionary

import "fmt"

type Entry struct {
}

func (e Entry) String() string {

	return ""
}

type Dictionary struct {
	entries map[string]Entry
}

func New() *Dictionary {
    dict := make(map[string]Entry)
    return &Dictionary{entries: dict}
}


func (d *Dictionary) Add(word string, definition string) {
	entry := Entry{}
	d.entries[word] = entry
}

func (d *Dictionary) Get(word string) (Entry, error) {
    entry, found := d.entries[word]
    if !found {
        return Entry{}, fmt.Errorf("mot non trouvé dans le dictionnaire")
    }
    return entry, nil
}


func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
}

func (d *Dictionary) List() ([]string, map[string]Entry) {
    words := make([]string, 0, len(d.entries))
    for word := range d.entries {
        words = append(words, word)
    }
    return words, d.entries
}
