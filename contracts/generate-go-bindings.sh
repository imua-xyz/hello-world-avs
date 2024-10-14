#!/bin/bash

# Clean up previous compilation results
rm -rf out
mkdir -p out/binding

# Traverse all subdirectories in the "pre" folder
for dir in ./contracts/src/v2/; do
    # Extract the directory name (excluding the path)
    dirname=$(basename "$dir")

    # Create the corresponding "out" folder
    mkdir -p "out/$dirname"

    # Traverse all .sol files in the current directory
    for file in "$dir"*.sol; do
        # Extract the file name (excluding the path and extension)
        filename=$(basename "$file" .sol)

        # Compile the contract
        solc --bin --abi --optimize --overwrite -o "out/$dirname" "$file"
        echo "Compiled ${dirname}/${filename}.sol"
    done

    # Generate the binding.go file using abigen
    binding_dir="out/$dirname"

    # Find the .bin .abi file in the $binding_dir directory
    bin_file=$(find "$binding_dir" -maxdepth 1 -type f -name '*.bin' -exec basename {} \;)
    abi_file=$(find "$binding_dir" -maxdepth 1 -type f -name '*.abi' -exec basename {} \;)
    # Generate the binding.go file using abigen
    abigen --bin="$binding_dir/$bin_file" --abi="$binding_dir/$abi_file" --pkg="contract$dirname" --out="out/$dirname/binding.go"

    echo "Generated binding for ${dirname}"
done

echo "Compilation and binding generation completed!"