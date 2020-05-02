package entry

type Entry interface{
	Id()    string
	/*
	Info()  map[string]string
	Date()  time.Time
	Title() string
	Path()  string
	*/
}

type Entries []Entry
