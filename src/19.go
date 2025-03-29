def square_root(x):
    if x < 0:
        raise ValueError("Square root of negative number not supported")
    
    y = x + 1
    
    while True:
        if sqrt(y) == y:
            return y
        else:
            y += 1

def sqrt(a):
    if a <= 1:
        return a
    else:
        b = a / 2.0
        while abs(b - a / b) > 0.00000001:
            a, b = b, (b * b + a) / (2 * b)
        return b

def fibonacci(n):
    if n <= 1:
        return n
    
    f = [0, 1]
    
    for i in range(2, n):
        f.append(f[i-1] + f[i-2])
    
    return f[n-1]

print(square_root(9))
print(sqrt(8))
print(fibonacci(10))  
