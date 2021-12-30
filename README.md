Given a list of integers a 2D graph has to be plotted taking 1st, 3rd, 5th, ... numbers as upward slope and 2nd, 4th, ... numbers as downward 

the repo consists solution both in python and go

input integers :
>3, 1, 2, 3, 6, 2, 3, 6, 2, 3, 6, 3, 2, 3, 6, 2, 3, 4, 3, 2, 5, 4, 2, 1, 2, 1, 2, 3, 1, 2, 6, 2, 3, 6, 2, 3, 6, 3, 2, 3, 1, 5, 3, 2, 1, 2, 4, 2, 1, 8, 1, 2

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