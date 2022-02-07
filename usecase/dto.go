package usecase

// Serve para transferir dados de um lado para outro

// input para criação de um curso

type CreateCourseInputDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateCourseOutputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
