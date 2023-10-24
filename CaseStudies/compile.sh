#!/bin/bash

DIR=$1

plantuml -tsvg $DIR/*.puml -output ./Compile/SVG
plantuml -tpng $DIR/*.puml -output ./Compile/PNG