#!/bin/bash

# Define the output directory
OUTPUT_DIR="output"
SRC_DIR="src"

# Remove everything in the output directory
rm -rf "$OUTPUT_DIR/*"

# Create the output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Build the main application
echo "Building the main application..."
go build -o "$OUTPUT_DIR/pdp" "$SRC_DIR/main.go"

# Navigate to the cmd directory
cd "$SRC_DIR/cmd" || exit

# Build the command plugins dynamically
echo "Building command plugins..."
for cmd_file in *.go; do
    # Check if there are any .go files
    if [ ! -e "$cmd_file" ]; then
        echo "No command files found in $SRC_DIR/cmd."
        exit 1
    fi

    # Get the base name of the command file (without the .go extension)
    cmd_name=$(basename "$cmd_file" .go)

    # Compile the command into a shared object file
    go build -buildmode=plugin -o "../../$OUTPUT_DIR/cmd/$cmd_name.so" "$cmd_file"

    # Check if the build was successful
    if [ $? -ne 0 ]; then
        echo "Error compiling $cmd_name. Skipping..."
    else
        echo "Compiled $cmd_name successfully."
    fi
done

# Navigate back to the root directory
cd ../../

# Check if the init directory exists before copying
if [ -d "$SRC_DIR/init" ]; then
    # Create the init directory in the output directory if it doesn't exist
    mkdir -p "$OUTPUT_DIR/init"
    
    # Copy everything from the init directory in src to the output directory
    cp -r "$SRC_DIR/init/"* "$OUTPUT_DIR/init/"
else
    echo "Warning: The init directory does not exist. No files copied."
fi

echo "Build completed. All output files are in the '$OUTPUT_DIR' directory."