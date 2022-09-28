package main

/* erros constantes são bons segundo dave cheney. é pq a comparacao entre 2 iguais sempre da true, acho */
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("a palavra não existe ainda no dicionário, não pode update")
)

type DictionaryErr string

/* todo type que tem método Error() e retorna string é uma interface de error, ou algo assim */
func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word] // (word -> definition, true) se 'word' existe no dic

	if !ok {
		return "", ErrNotFound //se nao tem a palavra no dic, retorna palavra vazia e o erro apropriado
	}

	return definition, nil //senao, retorna a definicao da palavra, e 'nil' (sem erro)
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound: // se a palavra nao ta no dic
		d[word] = definition // coloca no dic
	case nil: // sem erro == palavra já ta no dic
		return ErrWordExists // entao retorna o erro falando que palavra ja ta no dic...
	default:
		return err

	}
	return nil

}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound: // se a palavra nao ta no dic
		return ErrWordDoesNotExist // nao vamos update, retorna erro apropriado
	case nil: // se a palavra ja ta no dic
		d[word] = definition // atualiza a definicao
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	// aqui nao precisa tratar erro, pq delete soh deleta a chave se ela realmente existir
	delete(d, word) // funcao standard do Go que deleta chave 'word' de dicionario 'd'
}
