package boostrap

import "kumemori/internal/adapters/repository"

func InitApp() {
	repository.InitDb()
}
