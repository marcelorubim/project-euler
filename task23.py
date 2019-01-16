def findDivisors(n):
    """creates a list of a number's proper divisors"""
    divisors=[1]
    for x in range(2, n/2 + 1):
        if n % x == 0:
            divisors.append(x)
    return divisors

abundantNums = [12]
for x in range(13, 28123):
    if sum(findDivisors(x)) > x:
        abundantNums.append(x)
        
NonAbSums = [x for x in range(28124)]        

for i in range(len(abundantNums)):
    for j in range(i, len(abundantNums)):
        currentSum = abundantNums[i] + abundantNums[j]
        if currentSum > 28123:
            break
        else:
            if currentSum in NonAbSums:
                NonAbSums.remove(currentSum)
                
print sum(NonAbSums)