package categorymodel

import (
	"go-web-native-crud/config"
	"go-web-native-crud/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`SELECT * FROM categories`)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.Slug, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
		INSERT INTO categories (name, slug, created_at, updated_at) VALUES (?, ?, ?, ?)`,
		category.Name, category.Slug, category.CreatedAt, category.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LastInsertId > 0

}

func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`SELECT id, name, slug FROM categories WHERE id = ?`, id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Name,  &category.Slug); err != nil {
		panic(err.Error())
	}

	return category
}


func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`UPDATE categories SET name = ?, slug = ?, updated_at = ? WHERE id = ?`, category.Name, category.Slug,  category.UpdatedAt, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

