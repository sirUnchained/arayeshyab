package handlers

type brandHandler struct{}

func (bh *brandHandler) GetBrandHandler() *brandHandler {
	return &brandHandler{}
}

func (bh *brandHandler) GetAll() {}

func (bh *brandHandler) Create() {}

func (bh *brandHandler) Remove() {}
