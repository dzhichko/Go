package cmd

type InputFile struct {
	filename  string `yaml:"filename"`
	substring string `yaml:"substring"`
}

type Config struct {
	files []InputFile `yaml:"files"`
}
