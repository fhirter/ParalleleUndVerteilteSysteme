#!/usr/bin/env bash

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <input_directory>"
  exit 1
fi

# Check if the input is a directory
input_directory=$1
if [ ! -d "$input_directory" ]; then
    echo "The provided input is not a directory. Please provide a valid directory."
    exit 1
fi

output_directory=$input_directory

for markdown_file in "$input_directory"/*.md; do
    if [ -f "$markdown_file" ]; then
        # Extract the file name without the extension
        filename=$(basename -- "$markdown_file")
        filename_noext="${filename%.*}"

        # Convert Markdown to PDF using Pandoc
        pandoc "$markdown_file" \
          -o "$output_directory/$filename_noext.pdf" \
          --pdf-engine=xelatex \
          -H header.sty \

        echo "Converted $markdown_file to $output_directory/$filename_noext.pdf"
    fi
done