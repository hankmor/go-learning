package main

import (
	"database/sql"
	"fmt"
	"log"

	// mysql驱动包https://github.com/go-sql-driver/mysql/
	"github.com/go-sql-driver/mysql"
)

// 定义指针变量，访问特定数据库
// 这里简单写成全局变量，实际应该再封装
var db *sql.DB

func main() {
	// mysql数据库配置
	cfg := mysql.Config{
		User:   "root",
		Passwd: "12345678",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	// 获取数据库处理变量指针
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// 确保数据库可以正常连接
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// CRUD

	fmt.Println("insertOne ...")
	id := insertOne()

	fmt.Println("deleteById by id ...")
	deleteById(id)

	fmt.Println("query all ...")
	queryAll()

	fmt.Println("query by name ...")
	queryByName("John")

	fmt.Println("query by id ...")
	queryById(2)
}

func queryAll() {
	// 按名称模糊查询数据
	albums, err := albumsAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}

func queryByName(name string) {
	// 按名称模糊查询数据
	albums, err := albumsByArtist(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}

func queryById(id int64) {
	// 按ID查询
	alb, err := albumByID(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)
}

func insertOne() int64 {
	// 查询数据
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
	return albID
}

func deleteById(id int64) {
	r, err := deleteAlbumById(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted row count: %v\n", r)
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func albumsAll() ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("albumsAll: %v", err)
	}
	defer rows.Close() // 执行完成后关闭Rows,释放资源
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() { // 遍历游标
		var alb Album
		// 按查询结果字段顺序赋值
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsAll: %v", err)
		}
		albums = append(albums, alb) // 向slice添加值
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsAll: %v", err)
	}
	return albums, nil
}

// albumsByArtist queries for albums that have the specified artist name
// 按艺术家名称模糊查询
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	// Rows表示查询结果为多行
	rows, err := db.Query("SELECT * FROM album WHERE artist like ?", "%"+name+"%")
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close() // 执行完成后关闭Rows,释放资源
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() { // 遍历游标
		var alb Album
		// 按查询结果字段顺序赋值
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb) // 向slice添加值
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID.
// 按ID精确查询
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	// Row为单行数据
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	// 获取查询成功后的id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func deleteAlbumById(id int64) (int64, error) {
	result, err := db.Exec("DELETE FROM album where id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("deleteAlbumById: %v", err)
	}
	r, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteAlbumById: %v", err)
	}
	return r, nil
}
