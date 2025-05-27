def fib(max):
    a, b = 0, 1
    while a < max:
        yield a
        a, b = b, a + b

fibonacci_numbers = list(fib(100))
print(fibonacci_numbers)
