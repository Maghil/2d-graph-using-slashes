def populate_graph(data):
    graph=[]
    max_y = sum(data)
    max_x = -1             #to keep track of max x-axis of list #starting from -1 since no x axis exists as at start
    x,y = 0,0
    x_is_zero=False
    for i in range(0,len(data)):
        for _ in range(0,data[i]):
            if not x_is_zero:
                if i % 2==0:
                    if x > max_x:                       #creating a empty list filled with spaces so later we can fill with slashes
                        max_x+=1
                        graph.append([" "] * max_y)
                    graph[x][y]="/"
                    x+=1
                else:
                    if x==0:
                        print("the graph is getting pushed to -x axis, we currently don't a support for that \n")
                        x_is_zero = True
                        print("PRINTING GRAPH BEFORE REACHING -X AXIS")
                        break
                    x-=1
                    graph[x][y]="\\"
                y+=1
    return(graph[::-1])

def graph_printer(graph):
    for row in graph:
        print("".join(row))

if __name__ == "__main__":
    data = [3, 1, 2, 3, 6, 2, 3, 6, 2, 3, 6, 3, 2, 3, 6, 2, 3, 4, 3, 2, 5, 4, 2, 1, 2, 1, 2, 3, 1, 2, 6, 2, 3, 6, 2, 3, 6, 3, 2, 3, 1, 5, 3, 2, 1, 2, 4, 2, 1, 8, 1, 2]
    graph_printer(populate_graph(data=data))