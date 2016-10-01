package kiss

type Framework struct {
	*Resource
}

func (f *Framework) NewHandler(hf func(*Context)) *Handler {
	return &Handler{f.Resource, hf}
}

func New(r *Resource) *Framework {
	return &Framework{r}
}
