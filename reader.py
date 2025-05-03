import os

def read_and_write_files(root_dir, output_file):
    """
    Reads all files in a directory and its subdirectories, then writes
    the full path and content of each file to an output file.

    Args:
        root_dir (str): The root directory to start the search from.
        output_file (str): The path to the output file.
    """

    try:
        with open(output_file, 'w', encoding='utf-8') as outfile:  # Added encoding for broader compatibility
            for root, _, files in os.walk(root_dir):
                for file in files:
                    if file != "reader.py" and file != "prompt":
                        file_path = os.path.join(root, file)
                        try:
                            with open(file_path, 'r', encoding='utf-8') as infile:  # Added encoding for broader compatibility
                                content = infile.read()
                            outfile.write(f"Path: {file_path}\n")
                            outfile.write(f"Content:\n{content}\n")
                            outfile.write("-" * 40 + "\n")  # Separator for readability
                        except Exception as e:
                            print(f"Error reading or writing file {file_path}: {e}")

    except Exception as e:
        print(f"Error opening or writing to output file: {e}")


if __name__ == "__main__":
    root_directory = input("Enter the root directory to scan: ")
    output_filename = input("Enter the name of the output file: ")

    read_and_write_files(root_directory, output_filename)
    print(f"Files read and written to {output_filename}")