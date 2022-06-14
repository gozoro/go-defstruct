# go-defstruct

Go package sets default and environment scalar values to struct fields from tags.


## Usage

```

import "github.com/gozoro/go-defstruct"

type Data struct {

	Host string `env:"DATA_HOST" default:"localhost"`
	Port int    `env:"DATA_PORT" default:"8080"`
}

func (d *Data) LoadDefault() {

	defstruct.SetDefaultFromTags(d)
}


func (d *Data) LoadFromEnv() {

	defstruct.SetEnvFromTags(d)
}

d := &Data{}
d.LoadDefault()
d.LoadFromEnv()


```
