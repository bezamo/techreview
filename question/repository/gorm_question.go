package repository

import (
	"github.com/hellyab/techreview/entities"
	"github.com/hellyab/techreview/question"
	"github.com/jinzhu/gorm"
)

//QuestionGormRepo implements question.QuestionRepository interface
type QuestionGormRepo struct {
	conn *gorm.DB
}

//NewQuestionGormRepo returns new object of QuestionGormRepo
func NewQuestionGormRepo(db *gorm.DB) question.QuestionRepository {
	return &QuestionGormRepo{conn: db}
}

//Questions returns all user questions stored in the database
func (qstnRepo *QuestionGormRepo) Questions() ([]entities.Question, []error) {
	qstns := []entities.Question{}
	errs := qstnRepo.conn.Find(&qstns).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return qstns, errs
}

//Question returns a user question stored in the database which has the given id
func (qstnRepo *QuestionGormRepo) Question(id string) (*entities.Question, []error) {
	qstn := entities.Question{}
	errs := qstnRepo.conn.Where("id = ?", id).First(&qstn).GetErrors()
	// errs := qstnRepo.conn.First(&qstn, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &qstn, errs
}

//UpdateQuestion updates a given question in the database
func (qstnRepo *QuestionGormRepo) UpdateQuestion(question *entities.Question) (*entities.Question, []error) {
	qstn := question
	errs := qstnRepo.conn.Save(qstn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//DeleteQuestion deletes a question with a given id from the database
func (qstnRepo *QuestionGormRepo) DeleteQuestion(id string) (*entities.Question, []error) {
	qstn, errs := qstnRepo.Question(id)
	if len(errs) > 0 {
		return nil, errs
	}
	// errs := qstnRepo.conn.Where("id = ?", id).First(&qstn).GetErrors()
	errs = qstnRepo.conn.Delete(qstn).GetErrors()


	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//StoreQuestion stores a given question in the database
func (qstnRepo *QuestionGormRepo) StoreQuestion(question *entities.Question) (*entities.Question, []error) {
	qstn := question
	errs := qstnRepo.conn.Create(qstn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}
