# https://programmers.co.kr/learn/courses/30/lessons/42583

class Queue:
    def __init__(self):
        self.queue = []
    
    def Enqueue(self, x):
        self.queue.append(x)

    def Dequeue(self):
        return self.queue.pop(0) if len(self.queue) != 0 else False
    
    def GetSize(self):
        return len(self.queue)


def solution(bridge_length, weight, truck_weights):
    q = Queue()
    q.queue = [0]*bridge_length
    time_required = bridge_length
    total_queue = 0
    for i, v in enumerate(truck_weights):
        while total_queue + v > weight:
            if q.GetSize() < bridge_length:
                q.Enqueue(0)
                time_required += 1
            else:
                total_queue -= q.Dequeue()
        q.Enqueue(v)
        total_queue += v
        time_required += 1
    return time_required

# if __name__ == "__main__":
#     print(solution(2, 10, [7,4,5,6]))
#     print(solution(100, 100, [10]))
#     print(solution(100, 100, [10,10,10,10,10,10,10,10,10,10]))