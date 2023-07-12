package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/muriloabranches/Go-Expert-SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	fmt.Println("Connecting to Database...")
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()
	fmt.Println("Connected...")

	queries := db.New(dbConn)

	fmt.Println("Creating new category...")
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend description", Valid: true},
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("New category created...")

	fmt.Println("Listing all categories...")
	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	categoryID := categories[len(categories)-1].ID

	fmt.Println("Updating last category...")
	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          categoryID,
		Name:        "Backend updated",
		Description: sql.NullString{String: "Backend description updated", Valid: true},
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("Last category updated...")

	fmt.Println("Listing all categories...")
	categories, err = queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	fmt.Println("Deleting first category...")
	err = queries.DeleteCategory(ctx, categories[0].ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("First category deleted...")

	fmt.Println("Listing all categories...")
	categories, err = queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
