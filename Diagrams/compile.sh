#!/bin/bash

plantuml -tsvg *.puml -output ./Compile/SVG
plantuml -tpng *.puml -output ./Compile/PNG