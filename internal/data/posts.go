package data

import (
	"context"
	"database/sql"
	"log"

	"github.com/Caps1d/Spill/internal/construct"
	_ "github.com/lib/pq"
)

type PostModel struct {
	DB *sql.DB
}

func (pm *PostModel) All() (*[]construct.Post, error) {
	posts := []construct.Post{}

	rows, err := pm.DB.Query("SELECT * FROM post;")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post construct.Post

		if err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.UserId, &post.CreatedAt); err != nil {
			log.Println(err)
			continue
		}
		//append that struct into the posts array
		posts = append(posts, post)
	}
	rows.Close()
	return &posts, nil
}

func (pm *PostModel) Get(id int64) (*construct.Post, error) {
	var post construct.Post

	row, err := pm.DB.QueryContext(context.Background(), "SELECT * FROM post WHERE id = $1;", id)

	if err != nil {
		log.Printf("DB query error: %s", err)
		return nil, err
	}

	if err := row.Scan(&post.Id, &post.Title, &post.Content, &post.UserId, &post.CreatedAt); err != nil {
		return nil, err
	}
	row.Close()

	return &post, nil
}

func (pm *PostModel) Create(p *construct.Post) (int64, error) {

	query := "INSERT INTO post (title, content, userid) VALUES ($1, $2, $3) Returning id;"

	row := pm.DB.QueryRowContext(context.Background(), query, p.Title, p.Content, p.UserId)

	var insertedID int64
	if err := row.Scan(&insertedID); err != nil {
		log.Printf("No row was inserted, err: %s", err)
		return insertedID, err
	}

	return insertedID, nil

}

func (pm *PostModel) Delete(id int64) error {

	query := "DELETE FROM post WHERE id = $1;"

	row := pm.DB.QueryRowContext(context.Background(), query, id)

	if err := row.Scan(); err != nil {
		return err
	}

	return nil
}
