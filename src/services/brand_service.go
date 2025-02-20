package services

type brandService struct{}

func (bh *brandService) GetBrandService() *brandService {
	return &brandService{}
}

func (bh *brandService) GetAll() {}

func (bh *brandService) Create() {}

func (bh *brandService) Remove() {}
