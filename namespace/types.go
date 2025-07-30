package namespace

type tagRelation struct {
	Name     string `json:"name"`
	Instance string `json:"instance"`
	Pattern  string `json:"pattern"`
	Args     []Arg  `json:"args"`
}

type tagRetrieval struct {
	Name string `json:"name"`
	Args []Arg  `json:"args"`
}

type tagThing struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Author string `json:"author"`
}

type tagLink struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Thing1 string `json:"thing1"`
	Thing2 string `json:"thing2"`
	Author string `json:"author"`
}
