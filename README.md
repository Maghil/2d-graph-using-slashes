Given a list of integers a 2D graph has to be plotted taking 1st, 3rd, 5th, ... numbers as upward slope and 2nd, 4th, ... numbers as downward slope

I have used python list to do it. is there any further optimization possible with python?

input integers :
>3, 1, 2, 3, 6, 2, 3, 6, 2, 3, 6, 3, 2, 3, 6, 2, 3, 4, 3, 2, 5, 4, 2, 1, 2, 1, 2, 3, 1, 2, 6, 2, 3, 6, 2, 3, 6, 3, 2, 3, 1, 5, 3, 2, 1, 2, 4, 2, 1, 8, 1, 2

my code :
```
data = [3, 1, 2, 3, 6, 2, 3, 6, 2, 3, 6, 3, 2, 3, 6, 2, 3, 4, 3, 2, 5, 4, 2, 1, 2, 1, 2, 3, 1, 2, 6, 2, 3, 6, 2, 3, 6, 3, 2, 3, 1, 5, 3, 2, 1, 2, 4, 2, 1, 8, 1, 2]
length = sum(data)
print_list=[]

def populate_list():
    top = -1             #to keep track of max x-axis of list
    x=0
    y=0
    for i in range(0,len(data)):
        if i % 2==0:
            for _ in range(0,data[i]):
                if x>top:
                    print_list.append([])
                    top=x
                    for _ in range(0,length):               #to initialize a new top list
                        print_list[x].append(" ")
                print_list[x][y]="/"
                x+=1
                y+=1
        else:
            for _ in range(0,data[i]):
                x-=1
                print_list[x][y]="\\"
                y+=1

populate_list()
for i in range(len(print_list)-1,-1,-1):                    #printing from bottom up since generated graph is inverted
    print(''.join(map(str,print_list[i])))
```

OUTPUT :
```
                                                                                                  /\
                                                                                             /\  /  \              /\
                                                                                            /  \/    \            /  \  /\
                                                                     /\          /\        /          \          /    \/  \
                                                                    /  \      /\/  \      /            \  /\    /          \/\
                                                       /\          /    \  /\/      \/\  /              \/  \  /              \              /\
                                                  /\  /  \    /\  /      \/            \/                    \/                \    /\      /  \/\
                   /\                            /  \/    \  /  \/                                                              \  /  \/\  /      \
              /\  /  \              /\          /          \/                                                                    \/      \/        \
             /  \/    \            /  \  /\    /                                                                                                    \
            /          \          /    \/  \  /                                                                                                      \
     /\    /            \  /\    /          \/                                                                                                        \
  /\/  \  /              \/  \  /                                                                                                                      \
 /      \/                    \/                                                                                                                        \/\
/                                                                                                                                                          \
```