package entity

// a interface fica dentro do serviço de domínio

type CourseRepository interface {
	Insert(course Course) error
}
