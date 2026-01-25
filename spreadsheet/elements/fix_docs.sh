#!/bin/bash
ROOT="/home/connerohnesorge/Documents/001Repos/goffice"
GEN_LINES="//go:generate gomarkdoc -u -o CLAUDE.md .\n\n//go:generate gomarkdoc -u -o AGENTS.md .\n\n"

# Filter out directories
find "$ROOT" -name "*.go" -not -path "*/.*" -not -path "*/vendor/*" -not -path "*/examples/*" -not -path "*/testdata/*" -not -path "*/Open-XML-SDK/*" -printf "%h\n" | sort -u | while read dir; do
    pkg_name=$(grep -h "^package " "$dir"/*.go | head -n 1 | sed 's/.*package //')
    if [ -z "$pkg_name" ]; then continue; fi
    # Skip main packages in cmd/ or elsewhere if they are meant to be binaries
    if [[ "$pkg_name" == "main" ]]; then continue; fi

    doc_file="$dir/doc.go"
    if [ -f "$doc_file" ]; then
        if ! grep -q "gomarkdoc" "$doc_file"; then
            echo "Updating $doc_file"
            tmp_doc=$(mktemp)
            printf "$GEN_LINES" > "$tmp_doc"
            cat "$doc_file" >> "$tmp_doc"
            mv "$tmp_doc" "$doc_file"
        fi
    else
        echo "Creating $doc_file"
        # Find package comment
        source_file=$(grep -l "^// Package $pkg_name" "$dir"/*.go | head -n 1)
        if [ -n "$source_file" ]; then
            # Extract comment
            line_num=$(grep -n "^package $pkg_name" "$source_file" | cut -d: -f1)
            comment=$(head -n $((line_num - 1)) "$source_file" | grep "^//" | grep -v "Copyright")
            
            printf "$GEN_LINES" > "$doc_file"
            if [ -n "$comment" ]; then
                echo "$comment" >> "$doc_file"
            else
                echo "// Package $pkg_name provides support for Office Open XML." >> "$doc_file"
            fi
            echo "package $pkg_name" >> "$doc_file"
            
            # Remove comment from source file
            sed -i "/^\[// Package $pkg_name/d" "$source_file"
        else
            printf "$GEN_LINES" > "$doc_file"
            echo "// Package $pkg_name provides support for Office Open XML." >> "$doc_file"
            echo "package $pkg_name" >> "$doc_file"
        fi
    fi
done
