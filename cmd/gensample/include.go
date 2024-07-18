package main

import (
	_ "github.com/taylor-swanson/gensample/internal/emitter/template"

	_ "github.com/taylor-swanson/gensample/internal/generator/date"
	_ "github.com/taylor-swanson/gensample/internal/generator/names"
	_ "github.com/taylor-swanson/gensample/internal/generator/net"
	_ "github.com/taylor-swanson/gensample/internal/generator/number"
	_ "github.com/taylor-swanson/gensample/internal/generator/strings"
	_ "github.com/taylor-swanson/gensample/internal/generator/uuid"
	_ "github.com/taylor-swanson/gensample/internal/generator/web"

	_ "github.com/taylor-swanson/gensample/internal/output/file"
	_ "github.com/taylor-swanson/gensample/internal/output/stdout"
	_ "github.com/taylor-swanson/gensample/internal/output/tcp"
	_ "github.com/taylor-swanson/gensample/internal/output/udp"
)
