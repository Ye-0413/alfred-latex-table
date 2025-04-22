#!/usr/bin/env bash

go build -o alfred-latex-table
zip alfred-latex-table.alfredworkflow ./alfred-latex-table ./info.plist ./icon.png
rm ./alfred-latex-table
