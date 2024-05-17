#!/bin/bash

template='<!doctype html>
<html lang="en">
<head>
<meta charset="UTF-8" />
<title>%s</title>
</head>
<body></body>
</html>'

printf "$template" "Post #$1" >./blog/$2.html
