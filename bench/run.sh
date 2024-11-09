#!/usr/bin/env bash

cd pandora || exit
pandora load-gin.yaml &
pandora load-echo.yaml &
