package repository

import (
	"github.com/AlmasOrazgaliev/halyk-life-task2/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewDB(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetAuthors() (*[]models.Author, error) {
	rows, err := r.DB.Find(&models.Author{}).Rows()
	if err != nil {
		return nil, err
	}
	var authors []models.Author
	for rows.Next() {
		var author models.Author
		err := rows.Scan(&author.ID, &author.FullName, &author.Alias, &author.Specialization)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return &authors, nil
}

func (r *Repository) GetBooks() (*[]models.Book, error) {
	rows, err := r.DB.Find(&models.Book{}).Rows()
	if err != nil {
		return nil, err
	}
	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Name, &book.Category, &book.ISBNCode)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return &books, nil
}

func (r *Repository) GetMembers() (*[]models.Member, error) {
	rows, err := r.DB.Find(&models.Member{}).Rows()
	if err != nil {
		return nil, err
	}
	var members []models.Member
	for rows.Next() {
		var member models.Member
		err := rows.Scan(&member.ID, &member.FullName)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return &members, nil
}

func (r *Repository) GetAuthor(id uint) (*models.Author, error) {
	res := r.DB.First(&models.Author{}, id)
	if res.Error != nil {
		return nil, res.Error
	}
	var author models.Author
	res.Scan(&author)
	return &author, nil
}

func (r *Repository) GetBook(id uint) (*models.Book, error) {
	res := r.DB.First(&models.Book{}, id)
	if res.Error != nil {
		return nil, res.Error
	}
	var book models.Book
	res.Scan(&book)
	return &book, nil
}

func (r *Repository) GetMember(id uint) (*models.Member, error) {
	res := r.DB.First(&models.Member{}, id)
	if res.Error != nil {
		return nil, res.Error
	}
	var member models.Member
	res.Scan(&member)
	return &member, nil
}

func (r *Repository) GetAuthorBooks(id uint) (*[]models.Book, error) {
	rows, err := r.DB.Model(models.Book{AuthorId: id}).Find(&models.Book{}).Rows()
	if err != nil {
		return nil, err
	}
	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Name, &book.Category, &book.ISBNCode, &book.AuthorId)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return &books, nil
}

func (r *Repository) GetMemberBooks(id uint) (*[]models.Book, error) {
	rows, err := r.DB.Find(&models.BookMember{}, "member_id = ?", id).Rows()
	if err != nil {
		return nil, err
	}
	var books []models.Book
	for rows.Next() {
		var memberBooks models.BookMember
		var book models.Book
		err = rows.Scan(&memberBooks.BookId, &memberBooks.MemberId)
		r.DB.Model(models.Book{ID: memberBooks.BookId}).First(&book)
		books = append(books, book)
	}
	return &books, nil
}

func (r *Repository) CreateAuthor(author *models.Author) error {
	res := r.DB.Model(&models.Author{}).Create(author)
	return res.Error
}

func (r *Repository) CreateBook(book *models.Book) error {
	res := r.DB.Model(&models.Book{}).Create(book)
	return res.Error
}

func (r *Repository) CreateMember(member *models.Member) error {
	res := r.DB.Model(&models.Member{}).Create(member)
	return res.Error
}

func (r *Repository) CreateSubscribe(subscribe *models.BookMember) error {
	res := r.DB.Model(&models.BookMember{}).Create(subscribe)
	return res.Error
}

func (r *Repository) UpdateAuthor(author, updatedAuthor *models.Author) error {
	res := r.DB.Model(&author).Updates(updatedAuthor)
	res.Save(&author)
	return res.Error
}

func (r *Repository) UpdateBook(book, updatedBook *models.Book) error {
	res := r.DB.Model(&book).Updates(updatedBook)
	res.Save(&book)
	return res.Error
}

func (r *Repository) UpdateMember(member, updatedMember *models.Member) error {
	res := r.DB.Model(&member).Updates(updatedMember)
	res.Save(&member)
	return res.Error
}

func (r *Repository) DeleteAuthor(author *models.Author) error {
	res := r.DB.Delete(&author, author.ID)
	return res.Error
}

func (r *Repository) DeleteBook(book *models.Book) error {
	res := r.DB.Delete(&book, book.ID)
	return res.Error
}

func (r *Repository) DeleteMember(member *models.Member) error {
	res := r.DB.Delete(&member, member.ID)
	return res.Error
}
