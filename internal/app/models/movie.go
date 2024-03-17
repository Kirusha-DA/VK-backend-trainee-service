package models

type Movie struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        *string `gorm:"type:varchar(255)" json:"name"`
	Description *string `gorm:"type:text" json:"desctiption"`
	Rating      *uint   `gorm:"check:(rating <= 10 and rating >=1) or rating = NULL" json:"rating"`
	Actors      []Actor `gorm:"many2many:actor_movies"`
}

//func (m *Movie) ToDTO() *movie.MovieDTO {
//return &movie.MovieDTO{
//ID:          m.ID,
//Name:        *m.Description,
//Description: *m.Description,
//Rating:      *m.Rating,
//}
//}
