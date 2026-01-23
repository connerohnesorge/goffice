#!/usr/bin/env bash

cd /home/connerohnesorge/Documents/001Repos/goffice/spreadsheet/elements

# Find all TestNew functions and check their content
for file in *_test.go; do
    grep -h "^func TestNew" "$file" | while read line; do
        func_name=$(echo "$line" | sed 's/func \([^(]*\).*/\1/')

        echo "=== $func_name in $file ==="

        # Extract the function body
        awk "/^func $func_name/,/^}/" "$file" > /tmp/testfunc.tmp

        # Check if it only has nil check
        nil_check=$(grep -c "if.*== nil" /tmp/testfunc.tmp)
        localname_check=$(grep -c "LocalName()" /tmp/testfunc.tmp)
        namespace_check=$(grep -c "NamespaceURI()" /tmp/testfunc.tmp)
        ref_check=$(grep -c "Ref()" /tmp/testfunc.tmp)

        if [ "$nil_check" -gt 0 ] && [ "$localname_check" -eq 0 ] && [ "$namespace_check" -eq 0 ] && [ "$ref_check" -eq 0 ]; then
            echo "PURE NIL CHECK ONLY - Can be consolidated"
            echo "$file:$func_name" >> /tmp/nil_only_tests.txt
        elif [ "$nil_check" -gt 0 ]; then
            echo "Has nil check plus other tests - Keep separate"
        else
            echo "No nil check - Keep separate"
        fi
        echo

        rm -f /tmp/testfunc.tmp
    done
done