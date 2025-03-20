def fibonacci(n):
    if n <= 1:
        return n
    else:
        return fibonacci(n - 1) + fibonacci(n - 2)

for _ in range(10):  # Generate 10 random numbers using the fibonacci sequence
    print(fibonacci(random.randint(0, 99)))
