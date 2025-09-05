from typing import AnyStr, Any

def main():
    data = read_file("test-input.txt")
    print(data)

def process_input(data: AnyStr) -> Any:
    data.split("\n")

def read_file(filename: AnyStr) -> AnyStr:
    with open(filename, "r+") as f:
        data = f.read()
    return data

main()