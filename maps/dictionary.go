package maps

//creating a dictionary type, which acts as a thin wrapper around map.

type Dictionary map[string]string

// var ErrNotFound = errors.New("could not find the word you were looking for")
// var ErrWordExists = errors.New("cannot add word because it already exists")

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operatio on word as it does not exist")
)

type DictionaryErr string

// we made out error constant, by creating our own type which implements error interface
func (e DictionaryErr) Error() string {
	return string(e)
}

//const value must be determined at compile time. error.New() returns value at
//runtime. Thus, can't initialize this value as const. error.new() returns interface.

func (d Dictionary) Search(word string) (string, error) {
	defintition, ok := d[word] //map lookup can return 2 values.
	if !ok {
		return "", ErrNotFound
	}
	return defintition, nil //returning multiple values. multiple return types
}

func (d Dictionary) Add(word, defintion string) error {
	_, err := d.Search(word)
	// if err == nil {
	// 	return ErrWordExists
	// }

	// d[word] = defintion
	// return nil

	switch err {
	case ErrNotFound:
		d[word] = defintion
	case nil:
		return ErrWordExists
	default:
		return err
	}
	//switch case covers the case where, some other error is sent
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
