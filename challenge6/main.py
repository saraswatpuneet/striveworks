def binary_period(n):
# when edge case is considered int Max this will be fail due to 30, make it 31 
    d = [0] * 31
    l = 0
    while (n > 0):
        d[l] = n % 2
        n //= 2
        l += 1
    # as per the definition of P , P<Q/2 
    for p in range(1, 1 + l//2):
        ok = True
        for i in range(l - p):
            if d[i] != d[i + p]:
                ok = False
                break
        if ok:
            return p
    return -1

# main function
if __name__ == "__main__":
    print(binary_period(100))