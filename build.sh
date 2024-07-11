#!/bin/bash

xk6 build v0.52.0 \
  --replace google.golang.org/genproto@v0.0.0-20210226172003-ab064af71705=google.golang.org/genproto@v0.0.0-20230526161137-0005af68ea54 \
  --replace google.golang.org/genproto@v0.0.0-20230410155749-daa745c078e1=google.golang.org/genproto@v0.0.0-20230526161137-0005af68ea54 \
  --with github.com/avitalique/xk6-file@latest \
  --with github.com/ezeeb/xk6-mqtt=.