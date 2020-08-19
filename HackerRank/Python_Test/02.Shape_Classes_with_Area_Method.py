# https://www.hackerrank.com/test/ch1k00qp1cn/questions/7tqd3k6jaf9
import math
import os
import random
import re
import sys

class Rectangle:
    def __init__(self, a, b):
        self.a = a
        self.b = b
        
    def area(self):
        return self.a * self.b
    
class Circle:
    def __init__(self, radius):
        self.radius = radius
        
    def area(self):
        return self.radius**2 * math.pi


if __name__ == '__main__':  
    fptr = open(os.environ['OUTPUT_PATH'], 'w')
    q = int(input())
    queries = []
    for _ in range(q):
        args = input().split()
        shape_name, params = args[0], tuple(map(int, args[1:]))
        if shape_name == "rectangle":
            a, b = params[0], params[1]
            shape = Rectangle(a, b)
        elif shape_name == "circle":
            r = params[0]
            shape = Circle(r)
        else:
            raise ValueError("invalid shape type")
        fptr.write("%.2f\n" % shape.area())
    fptr.close()
