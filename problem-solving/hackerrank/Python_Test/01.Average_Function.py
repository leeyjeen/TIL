# https://www.hackerrank.com/test/ch1k00qp1cn/questions/6tr5fq7is1e
import math
import os
import random
import re
import sys

def avg(*args):
    return sum(args)/len(args)

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')
    
    nums = list(map(int, input().split()))
    res = avg(*nums)
    
    fptr.write('%.2f' % res + '\n')

    fptr.close()