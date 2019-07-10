package book

import (
	"fmt"

	"github.com/rkeplin/bible-go-api/common"
	"github.com/rkeplin/bible-go-api/core"
)

type Repository struct {
	translationFactory common.TranslationFactory
}

func (r Repository) FindAll() (BookCollection, error) {
	core.WaitForDb()

	collection := BookCollection{}

	statement := "SELECT b.id, b.name, b.testament, g.id AS genreId, g.name AS genreName FROM books b LEFT JOIN genres g ON b.genreId = g.id ORDER BY b.id ASC"
	rows, err := core.DB.Query(statement)

	if err != nil {
		return collection, err
	}

	defer rows.Close()

	var b Book

	for rows.Next() {

		if err := rows.Scan(&b.ID, &b.Name, &b.Testament, &b.Genre.ID, &b.Genre.Name); err != nil {
			return collection, err
		}

		collection = append(collection, b)
	}

	return collection, nil
}

func (r Repository) FindOne(id int) (Book, error) {
	core.WaitForDb()

	var b Book

	statement := "SELECT b.id, b.name, b.testament, g.id AS genreId, g.name AS genreName FROM books b LEFT JOIN genres g ON b.genreId = g.id WHERE b.id = ?"
	err := core.DB.QueryRow(statement, id).Scan(&b.ID, &b.Name, &b.Testament, &b.Genre.ID, &b.Genre.Name)

	if err != nil {
		return b, err
	}

	return b, nil
}

func (r Repository) FindChapters(bookId int, translation string) (ChapterCollection, error) {
	core.WaitForDb()

	collection := ChapterCollection{}

	table := r.translationFactory.GetTable(translation)

	statement := fmt.Sprintf("SELECT DISTINCT(chapterId) AS id FROM %s WHERE bookId = ?", table)
	fmt.Println(statement)

	rows, err := core.DB.Query(statement, bookId)

	if err != nil {
		return collection, err
	}

	defer rows.Close()

	var c Chapter

	for rows.Next() {

		if err := rows.Scan(&c.ID); err != nil {
			return collection, err
		}

		collection = append(collection, c)
	}

	return collection, nil
}
