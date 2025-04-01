import random

def generate_random_string(length):
    """
    Generate a random string of specified length.
    Example usage: generate_random_string(5) returns 'aBcD' and generate_random_string(10) returns 'xYzP'.
    """
    return ''.join(random.choice('abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ') for _ in range(length))

# Test the function
print(generate_random_string(5))  # Example output: "abcA"
