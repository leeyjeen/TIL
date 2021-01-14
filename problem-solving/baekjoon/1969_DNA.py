# https://www.acmicpc.net/problem/1969
import sys

if __name__ == "__main__":
    n, m = list(map(int, sys.stdin.readline().split()))
    counts = [{} for i in range(m)]
    
    for i in range(n):
        dna = sys.stdin.readline()
        nucleotide = list(dna)
        for j in range(m):
            if counts[j].get(nucleotide[j]):
                counts[j][nucleotide[j]] += 1
            else:
                counts[j][nucleotide[j]] = 1
    
    answer_dna = ""
    answer_count = 0
    for i in range(len(counts)):
        count = counts[i]
        answer = sorted(count.items(), key=lambda item: (-item[1], item[0]))[0]
        answer_dna += answer[0]
        answer_count += n - answer[1]
        
    print("{}\n{}".format(answer_dna, answer_count))
