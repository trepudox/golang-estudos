import time


def fibonacci(n):
    if n == 0:
        return 0
    
    if n <= 2:
        return 1

    return fibonacci(n - 1) + fibonacci(n - 2)


inicio = time.time()

print(fibonacci(50))

fim = time.time()

print(f"tempo: {(fim - inicio) * 1000}")


