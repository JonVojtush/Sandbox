import argparse

def greet(name="World"):
    return f"Hello {name}"

def main():
    parser = argparse.ArgumentParser(description="Greet the client.")
    
    parser.add_argument('--name', type=str, help='Name of thE client to greet')

    args = parser.parse_args()

    greeting = greet(args.name) if args.name else greet()
    print(greeting)

if __name__ == "__main__":
    main()
