package page

import (
	"database/sql"
)

type PageRepository interface {
	GetPagesByChapterID(chapterID int) ([]Page, error)
	CreatePage(page *Page) error
	UpdatePage(page *Page) error
}

type pageRepository struct {
	db *sql.DB
}

func NewPageRepository(db *sql.DB) PageRepository {
	return &pageRepository{db: db}
}

func (r *pageRepository) GetPagesByChapterID(chapterID int) ([]Page, error) {
	rows, err := r.db.Query("SELECT page_id, chapter_id, page_number, image_url FROM pages WHERE chapter_id = $1", chapterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pages []Page
	for rows.Next() {
		var page Page
		if err := rows.Scan(&page.ID, &page.ChapterID, &page.PageNumber, &page.ImageURL); err != nil {
			return nil, err
		}
		pages = append(pages, page)
	}

	return pages, nil
}

func (r *pageRepository) CreatePage(page *Page) error {
	_, err := r.db.Exec("SELECT insert_page($1, $2, $3)",
		page.ChapterID, page.PageNumber, page.ImageURL)
	return err
}

func (r *pageRepository) UpdatePage(page *Page) error {
	_, err := r.db.Exec("SELECT update_page_image_url($1, $2, $3)",
		page.ChapterID, page.PageNumber, page.ImageURL)
	return err
}
