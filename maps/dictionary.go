package maps

type Dictionary map[string]string

const (
	ErrNotFound         = dictionaryErr("could not find the word you were looking for")
	ErrWordExists       = dictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = dictionaryErr("cannot update word because it does not exist")
)

type dictionaryErr string

func (e dictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = newDefinition
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
