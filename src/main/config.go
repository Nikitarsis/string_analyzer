package main

type Config struct {
	shouldSaveString bool
	shouldCountCombo bool
	stopPipeline     bool
	readingFiles     []string
	outputFiles      []string
}

func GetConfig() Config {
	return Config{false, false, false, make([]string, 0), make([]string, 0)}
}

func (c *Config) SaveString() {
	c.shouldSaveString = true
}

func (c *Config) CountCombo() {
	c.shouldCountCombo = true
}

func (c *Config) DoNotPipeline() {

}

func (c *Config) SetReadingFiles(s ...string) {
	c.readingFiles = s
}

func (c *Config) SetOutputFiles(s ...string) {
	c.outputFiles = s
}

func (c Config) ShouldSaveString() bool {
	return c.shouldSaveString
}

func (c Config) ShouldCountCombo() bool {
	return c.shouldCountCombo
}

func (c Config) ShouldStopPipeline() bool {
	return c.stopPipeline
}

func (c Config) GetReadingFiles() []string {
	return c.readingFiles
}

func (c Config) GetOutputFiles() []string {
	return c.outputFiles
}
