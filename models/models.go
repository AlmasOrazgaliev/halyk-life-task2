package models

type Author struct {
	ID             uint   `json:"id"`
	FullName       string `json:"full-name"`
	Alias          string `json:"alias"`
	Specialization string `json:"specialization"`
}

type Book struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	ISBNCode string `json:"isbn-code"`
	AuthorId uint   `json:"author-id"`
	Author   Author `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Member struct {
	ID       uint   `json:"id"`
	FullName string `json:"full-name"`
}

type BookMember struct {
	BookId   uint   `json:"book-id"`
	Book     Book   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MemberId uint   `json:"member-id"`
	Member   Member `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
