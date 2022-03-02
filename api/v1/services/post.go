package services

import "social-media/api/v1/repository"

type PostServiceInterfce interface {
	AddPost()
}

type PostServiceStruct struct {
	PostRepository repository.SocialMediaRepositoryInterface
}

func PostService(postRepo repository.SocialMediaRepositoryInterface) PostServiceInterfce {
	return &PostServiceStruct{PostRepository: postRepo}
}

func (ps PostServiceStruct) AddPost() {

}
