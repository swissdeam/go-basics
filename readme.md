# Basics of Golang
In this R. i will add some examples of basic technologues in golang. It will help you (and me) to understand them

## Base of pointers
In this commit is simpliest example how pointers work

### "New" function
There is a new function which allows you to create a var without name. We must create pointers-thing to do operations with the namless var. In this chapter we changing value of var with "x" pointer.

### Using pointers in funcs
There is a guide how to use pointers to handle with Pure functions args, as they pure, we need to give them pointers to the original vars, otherwise they cannot be changed in the main function, so there is wont be some impact from these pure-functions.

### Struct and pointers
There is how we can use pointers with structs and their fields

### Methods
Now we can create functions just for specal types, just adding (parameter type_name) between <func> and <func_name>.

### Methods for structs
***From this moment all of code will be in txt files , which i will add for each chapter to "code examples" folder***
there is how we can create some methods(func for certain struct or type) for "developer" struct

### Methods for Pointers
we cannot change some fields in struct when we write <name_of_struct.name_of_methods()> we have to create pointer for struct. With pointer methods can change value of fields in struct 

### Interfaces
Interface is show how some struct acts. There is methods in interface that each struct have for themself. Struct in interface should have all of methods that declarated in interface type.
#### struct can be a member of several interfaces
    If struct have all of methods from several interfaces its can be given to funcs which request different objects of thesee interfaces



