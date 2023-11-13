package cmd

type InputFile struct {
	Filename  string `yaml:"filename"`
	Substring string `yaml:"substring"`
}

type Config struct {
	Files []InputFile `yaml:"files"`
}
